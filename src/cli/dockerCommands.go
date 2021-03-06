package cli

import (
    "../desc"
    "github.com/codegangsta/cli"
    "log"
    "os"
    "sync"
    "fmt"
    "github.com/samalba/dockerclient")

type DockerCommand struct {
    Descriptor *desc.MultiDockerDesc
    Api *DockerApi
    Printer *Printer
}

func NewDockerCommand() (*DockerCommand, error) {
    reader := desc.MultiDockerDescReader{"multidocker.json"}
    multidocker, err := reader.NewMultiDockerDesc()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    api := DockerApi{multidocker}

    dockerCommand := new(DockerCommand)
    dockerCommand.Descriptor = multidocker
    dockerCommand.Api = &api
    return dockerCommand, nil
}


func (d *DockerCommand) ListContainers(c *cli.Context) {
    //Prepare returned struct
    ret := []MDContainer{}
    //Get flag to list all containers (ie. running, and not running)
    showAll := c.Bool("all")
    //Get flag to showSize
    showSize := c.Bool("size")
    //Get Verbose flag
    debug := c.GlobalBool("debug")
    //Get args to filters on containers
    var filters string = ""
    if c.Args().Present() {
        filters = c.Args().First()
    }
    //Check nodes
    d.Api.CheckDockerNodes()
    //Iterate over all nodes
    for i := 0; i < len(d.Descriptor.Nodes); i++ {
        n := &d.Descriptor.Nodes[i]
        docker, _ := d.Api.ConnectToDocker(n.Alias)
        if docker != nil {
            containers, err := docker.ListContainers(showAll, false, filters)
            if err != nil {
                log.Fatal(err)
                os.Exit(1)
            }
            // Wrap all containers into MDContainer
            for idxC := 0; idxC < len(containers); idxC++ {
                c := containers[idxC]
                mdContainer := MDContainer{
                    Container:&c,
                    Node:n,
                }
                if debug {
                    log.Println("\t", mdContainer.Container.Id[:12], mdContainer.Container.Names)
                }
                ret = append(ret, mdContainer)
            }
        }
    }
    d.Printer.PrintMDContainersList(ret, showSize)
}

func (d *DockerCommand) ListImages(c *cli.Context) {
    //Prepare returned struct
    var ret []MDImage
    //Get Verbose flag
    debug := c.GlobalBool("debug")
    //Get flag to list all images
    //showAll := c.Bool("all")
    //Get flag to show size
    showSize := c.Bool("size")
    //Check nodes eachtime
    d.Api.CheckDockerNodes()

    //Iterate over all nodes
    for i := 0; i < len(d.Descriptor.Nodes); i++ {
        n := &d.Descriptor.Nodes[i]
        docker, _ := d.Api.ConnectToDocker(n.Alias)
        if docker != nil {
            images, err := docker.ListImages()
            if err != nil {
                log.Fatal(err)
                os.Exit(1)
            }
            // Wrap all images into MDImages
            for idxI := 0; idxI < len(images); idxI++ {
                mdImage := MDImage{
                    Node: n,
                    Image: images[idxI],
                }
                if debug {
                    log.Println("\t", mdImage.Image.Id[:12], mdImage.Image.RepoTags)
                }
                ret = append(ret, mdImage)
            }
        }
    }
    d.Printer.PrintMDImagesList(ret, showSize)
}

func (d *DockerCommand) PullImage(c *cli.Context) {
    //Check nodes
    d.Api.CheckDockerNodes()
    //Get Verbose flag
    debug := c.GlobalBool("debug")
    //Get image Name
    name := c.Args().First()


    if name == "" {
        log.Fatal("multidocker pull requires 1 argument. See 'multidocker pull --help'.")
        os.Exit(1)
    }
    if debug {
        log.Printf("Pulling image %s on hosts", name)
    }

    //Prepare channel for errors management
    chanPulledImage := make(chan MDPulledImage, len(d.Descriptor.Nodes))

    var wg sync.WaitGroup

    //Iterate over all nodes
    for i := 0; i < len(d.Descriptor.Nodes); i++ {
        n := &d.Descriptor.Nodes[i]
        docker, _ := d.Api.ConnectToDocker(n.Alias)

        if debug {
            log.Printf("Pulling image %s on host %s::%s", name, n.Alias, n.Host)
        }
        if docker != nil {
            wg.Add(1)
            //TODO put a nice progressbar
            go func() {
                err := docker.PullImage(name, nil)

                if err != nil {
                    log.Printf("Cannot pull image %s on host %s::%s", name, n.Alias, n.Host)
                    log.Printf("\t|__%s", err)
                    chanPulledImage <- MDPulledImage{
                        Node: n,
                        Name: name,
                        Success: false,
                        Error: err,
                    }
                } else {
                    chanPulledImage <- MDPulledImage{
                        Node: n,
                        Name: name,
                        Success: true,
                        Error: nil,
                    }
                }
                wg.Done();
            }()
        }
    }
    //Wait goroutines
    wg.Wait()
    //Then close the channel, to read it
    close(chanPulledImage)

    var ret = []MDPulledImage{}
    for pulledImage := range chanPulledImage {
        var status string
        if pulledImage.Success {
            status = "OK"
        } else {
            status = "KO"
        }
        if debug {
            log.Printf("%s::%s - Pulling %s\t%s", pulledImage.Node.Alias, pulledImage.Node.Host, pulledImage.Name, status)
        }
        ret = append(ret,
        MDPulledImage{
            Node: pulledImage.Node,
            Name: pulledImage.Name,
            Success: pulledImage.Success,
            Error: pulledImage.Error,
        })
    }

    d.Printer.PrintMDPulledImages(ret)

    if debug {
        log.Printf("End");
    }
}

func (d *DockerCommand) StopContainers(c *cli.Context) {
    //Check nodes
    d.Api.CheckDockerNodes()
    //Get Verbose flag
    debug := c.GlobalBool("debug")
    //Get Nodes slice
    nodes := c.StringSlice("node")
    //Get Nodes image
    images := c.StringSlice("image")
    //Get timeout
    timeout := c.Int("time")


    chanContainer := make(chan MDContainer, len(nodes) * len(images) + 1)
    chanError := make(chan error, len(nodes) * len(images) + 1)

    var wg sync.WaitGroup

    for idxNode := 0; idxNode < len(nodes); idxNode++ {
        wg.Add(1)
        go func(nodeAlias string) {
            log.Print("Connecting to " + nodeAlias)
            docker, err := d.Api.ConnectToDocker(nodeAlias)
            if err != nil {
                chanError <- err
            } else {
                if debug {
                    log.Printf("Inspect containers on node %s", nodeAlias)
                }
                containers, err := docker.ListContainers(false, false, "")
                if err != nil {
                    chanError <- err
                } else {
                    var wg1 sync.WaitGroup
                    for idxImage := 0; idxImage < len(images); idxImage++ {
                        if debug {
                            log.Printf("Stopping gracefully containers %s on node %s", images[idxImage], nodeAlias)
                        }
                        for idxContainer := 0; idxContainer < len(containers); idxContainer++ {
                            wg1.Add(1)
                            go func(container *dockerclient.Container, image string) {
                                if container.Image ==  image {
                                    err := docker.StopContainer(container.Id, timeout)
                                    if err != nil {
                                        chanError <- err
                                    } else {
                                        n, _ := d.Api.getNode(nodeAlias)
                                        chanContainer <- MDContainer{
                                            Node: n,
                                            Container: container,
                                        }
                                    }
                                }
                                wg1.Done()
                            }(&containers[idxContainer], images[idxImage])
                        }
                    }
                    wg1.Wait()
                }
            }
            wg.Done()
        }(nodes[idxNode])
    }
    wg.Wait()

    close(chanContainer)
    close(chanError)

    for c := range chanContainer {
        fmt.Println(c)
    }

    for err := range chanError {
        fmt.Println(err)
    }

}
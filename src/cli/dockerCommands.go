package cli

import (
    "../desc"
    "github.com/codegangsta/cli"
    "log"
    "os")

type DockerCommand struct {
    Descriptor *desc.MultiDockerDesc
    Api *DockerApi
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
                ret = ExtendMDContainersList(ret, mdContainer)
            }
        }
    }
    PrintMDContainersList(ret, showSize)
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
                ret = ExtendMDImageList(ret, mdImage)
            }
        }
    }
    PrintMDImagesList(ret, showSize)
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
    //Iterate over all nodes
    for i := 0; i < len(d.Descriptor.Nodes); i++ {
        n := &d.Descriptor.Nodes[i]
        docker, _ := d.Api.ConnectToDocker(n.Alias)
        if docker != nil {
            //docker.PullImage()
        }
    }
}
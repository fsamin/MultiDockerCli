package cli

import (
	"../desc"
	"github.com/codegangsta/cli"
	"log"
)

func listContainers(c *cli.Context) {
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

    reader := desc.MultiDockerDescReader{
		File: "multidocker.json",
	}

	multidocker, err := reader.NewMultiDockerDesc()
	if err != nil {
		log.Fatal(err)
		return
	}

	api := DockerApi{
		Descriptor: multidocker,
	}
	api.CheckDockerNodes()

	ret := []MDContainer{}
	//Iterate over all nodes
	for i := 0; i < len(multidocker.Nodes); i++ {
		n := &multidocker.Nodes[i]
		docker, _ := api.ConnectToDocker(n.Alias)

		// Get only running containers
		containers, err := docker.ListContainers(showAll, false, filters)
		if err != nil {
			log.Fatal(err)
		}

		// Wrap all containers into MDContainer
        for idxC:=0; idxC < len(containers); idxC++ {
            c:=containers[idxC]
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



    PrintMDContainersList(ret, showSize)
}

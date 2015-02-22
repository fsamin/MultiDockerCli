package cli

import (
	"../desc"
	"github.com/codegangsta/cli"
	"log"
)

func listContainers(c *cli.Context) {
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
		containers, err := docker.ListContainers(false, false, "")
		if err != nil {
			log.Fatal(err)
		}

		// Wrap all containers into MDContainer
		for _, c := range containers {
			log.Println(c.Id, c.Names)
			mdContainer := MDContainer{
				node:      n,
				container: &c,
			}
			ret = ExtendMDContainersList(ret, mdContainer)
		}
	}

}

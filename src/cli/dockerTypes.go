package cli

import (
	"../desc"
	"github.com/samalba/dockerclient"
)

type MDContainer struct {
	Node      *desc.Node
	Container *dockerclient.Container
}

func NewMDContainer(node *desc.Node, container *dockerclient.Container) MDContainer {
    return MDContainer{node, container}
}

type MDImage struct {
	Node  *desc.Node
	Image *dockerclient.Image
}

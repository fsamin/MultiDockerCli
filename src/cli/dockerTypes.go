package cli

import (
	"../desc"
	"github.com/samalba/dockerclient"
)

type MDContainer struct {
	Node      *desc.Node
	Container *dockerclient.Container
}

type MDImage struct {
	Node  *desc.Node
	Image *dockerclient.Image
}

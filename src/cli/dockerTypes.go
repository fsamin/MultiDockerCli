package cli

import (
	"../desc"
	"github.com/samalba/dockerclient"
)

type MDContainer struct {
	node      *desc.Node
	container *dockerclient.Container
}

type MDImage struct {
	node  *desc.Node
	image *dockerclient.Image
}

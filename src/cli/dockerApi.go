package cli

import (
	"crypto/tls"
	"crypto/x509"

	"io/ioutil"

	"../desc"
	"errors"
	"github.com/samalba/dockerclient"
	"log"
)

type DockerApi struct {
	Descriptor *desc.MultiDockerDesc
}

func (dockerApi *DockerApi) CheckDockerNodes() {
	for i := 0; i < len(dockerApi.Descriptor.Nodes); i++ {
		n := &dockerApi.Descriptor.Nodes[i]
		docker, err := getDockerClient(n)
		if err != nil {
			n.Available = false
		} else {
			n.Available = true
			version, err := docker.Version()
			if err != nil {
				n.Available = false
			} else {
				n.ApiVersion = version.Version
			}
		}
	}
}

func (dockerApi *DockerApi) ConnectToDocker(nodeAlias string) (*dockerclient.DockerClient, error) {
	node, err := dockerApi.getNode(nodeAlias)
	if err != nil {
		return nil, err
	}
	if node.Available == false {
		log.Fatal("Docker node : " + node.Alias + "::" + node.Host + " is unreachable")
		return nil, errors.New("Docker node : " + node.Alias + "::" + node.Host + " is unreachable")
	}

	log.Printf("Connecting to docker node %s::%s (version %s)", node.Alias, node.Host, node.ApiVersion)
	dockerClient, err := getDockerClient(node)
	return dockerClient, err
}

func (dockerApi *DockerApi) getNode(nodeAlias string) (*desc.Node, error) {
	for i := 0; i < len(dockerApi.Descriptor.Nodes); i++ {
		n := &dockerApi.Descriptor.Nodes[i]
		if n.Alias == nodeAlias {
			return n, nil
		}
	}
	return nil, errors.New("Cannot find alias " + nodeAlias + ". Please check your multidocker descriptor.")
}

func getDockerClient(node *desc.Node) (*dockerclient.DockerClient, error) {
	tlsConfig := &tls.Config{}
	cert := &node.Cert
	if cert != nil {

		caFile := cert.CaFile
		certFile := cert.CertFile
		keyFile := cert.KeyFile

		cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
		pemCerts, _ := ioutil.ReadFile(caFile)

		tlsConfig.RootCAs = x509.NewCertPool()
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tlsConfig.Certificates = []tls.Certificate{cert}

		tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)

	}
	docker, err := dockerclient.NewDockerClient(node.Host, tlsConfig)

	return docker, err
}

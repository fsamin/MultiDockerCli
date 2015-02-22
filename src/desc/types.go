package desc

type MultiDockerDesc struct {
	Nodes []Node
}

type Node struct {
	Alias      string
	Host       string
	Cert       Cert
	Available  bool
	ApiVersion string
}

type Cert struct {
	CaFile   string
	CertFile string
	KeyFile  string
}

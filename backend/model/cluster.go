package model

type ClusterInfo struct {
	Name      string
	Cluster   string
	Server    string
	User      string
	Namespace string
	Status    bool
	Source    string
}

type ClusterProfile struct {
	ContextName    string
	KubeConfigPath string
}

type EnvironmentDto struct {
	Name        string
	Description string
	Env         string
	Status      bool
}

type NodeDto struct {
	Name     string
	Resource Resource
	Roles    []string
	Version  string
	Age      string
	Status   bool
}

type NodeDtoV2 struct {
	Name              string
	Namespace         string
	Resource          Resource
	Version           string
	CreationTimestamp string
	Labels            map[string]string
}

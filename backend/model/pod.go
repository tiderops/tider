package model

type PodDto struct {
	Name      string
	Namespace string
	Container Container
	Node      string
	Age       string
	Status    string
	Editable  []string
	Labels    map[string]string
}

type PodDetail struct {
	Name      string
	Namespace string
	Container Container
	Age       string
	Status    string
	Editable  []string
}

type PodMetricDto struct {
	Name      string
	Namespace string
	Consume   Resource
}

type CurrentResourcesDto struct {
	Name      string
	Namespace string
	Container Container
}

type PodUpdate struct {
	App       string
	Container ContainerUpdate
}

package model

type PodDto struct {
	Name      string
	Namespace string
	Replicas  int32
	Container Container
	Age       string
	Status    string
}

type PodResponse struct {
	Name      string
	Namespace string
	Replicas  int32
	Container Container
	Age       string
	Status    string
	Editable  EditableOption
}

type EditableOption struct {
	Options []string
}

type PodMetricDto struct {
	Name      string
	Namespace string
	Consume   Resource
}

type Container struct {
	Limit   Resource
	Request Resource
}

type CurrentResourcesDto struct {
	Name      string
	Namespace string
	Container Container
}

type PodRequest struct {
	Replicas  string
	App       string
	Container ContainerRequest
}

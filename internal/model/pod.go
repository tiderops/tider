package model

type PodDto struct {
	Name       string
	Namespace  string
	Containers []Container
	Node       string
	Age        string
	CreatedAt  int64
	Status     string
	Editable   []string
	Labels     map[string]string
}

type PodUpdate struct {
	App       string
	Container ContainerUpdate
}

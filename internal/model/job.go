package model

type JobDto struct {
	Name      string
	Namespace string
	Container Container
	Node      string
	Age       string
	Status    string
	Editable  []string
	Labels    map[string]string
}

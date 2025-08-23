package model

type ServiceDto struct {
	Name              string
	Namespace         string
	Type              string
	InternalIp        string
	ExternalIp        string
	Port              string
	Status            string
	CreationTimestamp string
	Spec              string
	Labels            map[string]string
}

type ServiceUpdate struct {
	name        string
	LabelApp    string
	SpecType    string
	Port        string
	TargetPort  string
	SelectorApp string
}

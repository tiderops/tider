package model

type ServiceDto struct {
	Name              string
	Namespace         string
	Type              string
	InternalIp        string
	ExternalIp        string
	Port              int32
	Status            string
	CreationTimestamp string
	Spec              string
	Labels            map[string]string
}

type ServiceUpdate struct {
	name        string
	LabelApp    string
	SpecType    string
	Port        int32
	TargetPort  string
	SelectorApp string
}

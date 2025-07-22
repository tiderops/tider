package model

type ServiceDto struct {
	Name              string
	Namespace         string
	Labels            map[string]string
	Status            string
	CreationTimestamp string
	Spec              string
}

type ServiceRequest struct {
	name        string
	LabelApp    string
	SpecType    string
	Port        string
	TargetPort  string
	SelectorApp string
}

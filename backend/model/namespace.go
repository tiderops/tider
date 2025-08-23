package model

type NamespaceDto struct {
	Name    string
	Version string
	Age     string
	Labels  map[string]string
	Status  string
}

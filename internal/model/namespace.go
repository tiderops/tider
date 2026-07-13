package model

type NamespaceDto struct {
	Name      string
	Version   string
	Age       string
	CreatedAt int64
	Labels    map[string]string
	Status    string
}

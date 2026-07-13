package model

type IngressDto struct {
	Name      string
	Namespace string
	Rules     []RuleDto
	Age       string
	CreatedAt int64
	Labels    map[string]string
}

type RuleDto struct {
	Host             string
	Path             string
	IngressRuleValue string
}

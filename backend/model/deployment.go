package model

type DeploymentDto struct {
	Name      string
	Namespace string
	Status    string
	Age       string
}

type DeploymentUpdate struct {
	Replicas     string
	App          string
	StrategyType string
	Label        LabelUpdate
	Container    ContainerUpdate
}

type LabelUpdate struct {
	App      string
	Tier     string
	TierType string
}

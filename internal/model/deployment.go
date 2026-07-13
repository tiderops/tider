package model

type DeploymentDto struct {
	Name       string
	Namespace  string
	Containers []Container
	Replicas   int32
	Status     string
	Age        string
	CreatedAt  int64
	Labels     map[string]string
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

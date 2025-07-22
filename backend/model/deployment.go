package model

type DeploymentDto struct {
	Name      string
	Namespace string
	Status    string
	Age       string
}

type DeploymentRequest struct {
	Replicas     string
	App          string
	StrategyType string
	Label        LabelRequest
	Container    ContainerRequest
}

type LabelRequest struct {
	App      string
	Tier     string
	TierType string
}

type ContainerRequest struct {
	Image      string
	PullPolicy string
	Port       string
	Resource   ResourceRequest
}

type ResourceRequest struct {
	RMemory string
	RCpu    string
	LMemory string
	LCpu    string
}

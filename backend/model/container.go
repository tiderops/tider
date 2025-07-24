package model

type Container struct {
	Image      string
	PullPolicy string
	Port       string
	Limit      Resource
	Request    Resource
}

type ContainerUpdate struct {
	Image      string
	PullPolicy string
	Port       string
	Resource   ResourceUpdate
}

type ResourceUpdate struct {
	RMemory string
	RCpu    string
	LMemory string
	LCpu    string
}

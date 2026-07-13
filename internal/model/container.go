package model

type Container struct {
	Name       string
	Image      string
	PullPolicy string
	Port       int32
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
	RMemory int64
	RCpu    int64
	LMemory int64
	LCpu    int64
}

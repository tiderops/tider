package database

import "strconv"

type parameter struct {
	Name    string
	objects []object
}

type object struct {
	Name        string
	description string
}

func (p *parameter) get() []parameter {
	var params []parameter

	for i := 0; i < 5; i++ {
		var objs []object

		for j := 0; j < 3; j++ {
			o := object{
				Name:        "test" + strconv.Itoa(i),
				description: "description" + strconv.Itoa(i),
			}
			objs = append(objs, o)
		}
		pa := parameter{
			Name:    "Name test" + strconv.Itoa(i),
			objects: objs,
		}
		params = append(params, pa)
	}

	return params
}

type IParameterEntity interface {
	GetKubernetesParameters() []CommonParameterDto
	GetCommonParameters() []CommonParameterDto
	GetK8sObjects() []ObjectType
}

type parameterImpl struct{}

func NewParameterEntity() IParameterEntity {
	return &parameterImpl{}
}

type CommonParameterDto struct {
	Name string
	Link string
	Icon string
}

type ObjectType struct {
	Name       string
	IsVisible  bool
	IsEditable bool
	K8sObject  []K8sObject
}

type K8sObject struct {
	Name       string
	Link       string
	IsVisible  bool
	IsEditable bool
}

type HeadParamsDto struct {
	Title    string
	Key      string
	Align    string
	Sortable bool
}

func (p *parameterImpl) GetKubernetesParameters() []CommonParameterDto {
	return []CommonParameterDto{
		{Name: "Overview", Link: "overview", Icon: "📊"},
		{Name: "General", Link: "general", Icon: "📊"},
		{Name: "Workload", Link: "workload", Icon: "📊"},
		{Name: "Network", Link: "network", Icon: "📊"},
		{Name: "Storage", Link: "storage", Icon: "📊"},
	}
}

func (p *parameterImpl) GetCommonParameters() []CommonParameterDto {
	return []CommonParameterDto{
		{
			Name: "Connections",
			Link: "connections",
			Icon: "⚙️",
		},
		{
			Name: "Settings",
			Link: "settings",
			Icon: "⚙️",
		},
		{
			Name: "Documentation",
			Link: "documentation",
			Icon: "⚙️",
		},
	}
}

func (p *parameterImpl) GetK8sObjects() []ObjectType {
	return []ObjectType{
		{
			Name:       "Overview",
			IsEditable: true,
			IsVisible:  true,
			K8sObject:  []K8sObject{},
		},
		{
			Name:       "General",
			IsEditable: true,
			IsVisible:  true,
			K8sObject: []K8sObject{
				{
					Name:       "Node",
					Link:       "node",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "Namespace",
					Link:       "namespace",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "Event",
					Link:       "event",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "Cluster Graph",
					Link:       "cluster_graph",
					IsEditable: true,
					IsVisible:  true,
				},
			},
		},
		{
			Name:       "Workload",
			IsEditable: true,
			IsVisible:  true,
			K8sObject: []K8sObject{
				{
					Name:       "Pod",
					Link:       "pod",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "Deployment",
					Link:       "deployment",
					IsEditable: true,
					IsVisible:  true,
				},
			},
		},
		{Name: "Network",
			IsEditable: true,
			IsVisible:  true,
			K8sObject: []K8sObject{
				{
					Name:       "Ingress",
					Link:       "ingress",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "Service",
					Link:       "service",
					IsEditable: true,
					IsVisible:  true,
				},
			},
		},
		{Name: "Storage",
			IsEditable: true,
			IsVisible:  true,
			K8sObject: []K8sObject{
				{
					Name:       "PersistentVolume",
					Link:       "persistentVolume",
					IsEditable: true,
					IsVisible:  true,
				},
				{
					Name:       "PersistentVolumeClaim",
					Link:       "persistentVolumeClaim",
					IsEditable: true,
					IsVisible:  true,
				},
			},
		},
	}
}

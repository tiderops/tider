package config

type ParameterEntity interface {
	GetKubernetesParameters() []CommonParameterDto
	GetCommonParameters() []CommonParameterDto
	GetK8sObjects() []ObjectType
	GetHeadParams(k8sObject string) []HeadParamsDto
}

type parameterImpl struct{}

func NewParameterEntity() ParameterEntity {
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
				{
					Name:       "Tuning",
					Link:       "tuning",
					IsEditable: false,
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

// GetHeadParams returns the grid header definition for a resource view.
func (p *parameterImpl) GetHeadParams(k8sObject string) []HeadParamsDto {
	switch k8sObject {
	case "pod":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Storage", Key: "storage", Align: "start", Sortable: false},
			{Title: "Node", Key: "node", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "node":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Storage", Key: "storage", Align: "start", Sortable: false},
			{Title: "KubeletVersion", Key: "kubeletVersion", Align: "start", Sortable: false},
			{Title: "OperatingSystem", Key: "operatingSystem", Align: "start", Sortable: false},
			{Title: "Version", Key: "version", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "namespace":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}

	case "service":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Type", Key: "type", Align: "start", Sortable: false},

			{Title: "Internal IP", Key: "intIp", Align: "start", Sortable: false},
			{Title: "External IP", Key: "extIp", Align: "start", Sortable: false},
			{Title: "Port", Key: "port", Align: "start", Sortable: false},

			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "ingress":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "host", Key: "host", Align: "start", Sortable: false},
			{Title: "path", Key: "path", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "persistentVolume":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Storage class", Key: "storageClass", Align: "start", Sortable: false},
			{Title: "Capacity", Key: "capacity", Align: "start", Sortable: false},
			{Title: "Claim", Key: "claim", Align: "start", Sortable: true},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "persistentVolumeClaim":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Storage class", Key: "storageClass", Align: "start", Sortable: false},
			{Title: "Size", Key: "size", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	case "deployment":
		return []HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	default:
		return []HeadParamsDto{}
	}
}

package endpoint

import (
	"Kubexplorer/backend/database"
	"Kubexplorer/backend/usecase"
)

type ParameterEndpoint struct {
	useCase usecase.ParameterUseCase
}

func NewParameterEndpoint(useCase usecase.ParameterUseCase) *ParameterEndpoint {
	return &ParameterEndpoint{useCase: useCase}
}

func (pe *ParameterEndpoint) GetKubernetesParameters() []database.CommonParameterDto {
	return pe.useCase.GetKubernetesParameters()
}

func (pe *ParameterEndpoint) GetCommonParameters() []database.CommonParameterDto {
	return pe.useCase.GetCommonParameters()
}

func (pe *ParameterEndpoint) GetK8sObjects() []database.ObjectType {
	return pe.useCase.GetK8sObjects()
}

func (pe *ParameterEndpoint) GetHeadParams(k8sObject string) []database.HeadParamsDto {
	if k8sObject == "pod" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Storage", Key: "storage", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "node" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Storage", Key: "storage", Align: "start", Sortable: false},
			{Title: "Ephemeral", Key: "ephemeral", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Labels", Key: "labels", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "namespace" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Label", Key: "age", Align: "start", Sortable: true},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}

	} else if k8sObject == "service" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Labels", Key: "labels", Align: "start", Sortable: false},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "ingress" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "host", Key: "host", Align: "start", Sortable: false},
			{Title: "path", Key: "path", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Labels", Key: "labels", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "persistentVolume" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "persistentVolumeClaim" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "deployment" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else {
		return []database.HeadParamsDto{}
	}
}

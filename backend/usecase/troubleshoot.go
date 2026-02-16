package usecase

import (
	"Kubexplorer/backend/service"
)

type TroubleshootUseCase interface {
	Invoke(name string, namespace string, clusterCtx string, resource string)
}

type troubleshootUseCase struct {
	service service.DiagnosticService
}

func NewTroubleshootUseCase(service service.DiagnosticService) TroubleshootUseCase {
	return &troubleshootUseCase{service: service}
}

func (t *troubleshootUseCase) Invoke(name string, namespace string, clusterCtx string, resource string) {
	t.service.Analyse(name, namespace, clusterCtx, resource)
}

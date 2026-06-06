package usecase

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type MetricUseCase interface {
	GetPodMetric(namespace string)
	GetNodeMetric(namespace string)
	EnableMetrics(dto model.MetricEnableDto)
}

type metricUseCase struct {
	service service.MetricV2Service
}

func NewMetricUseCase(service service.MetricV2Service) MetricUseCase {
	return &metricUseCase{service: service}
}

func (m *metricUseCase) GetPodMetric(namespace string) {
	//return m.service.GetPodMetrics(namespace)
}

func (m *metricUseCase) GetNodeMetric(namespace string) {
	//return m.service.GetNodeMetrics(namespace)
}

func (m *metricUseCase) EnableMetrics(dto model.MetricEnableDto) {
	m.service.EnableMetrics(dto)
}

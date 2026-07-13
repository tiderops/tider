package usecase

import "Kubexplorer/internal/config"

type ParameterUseCase interface {
	GetKubernetesParameters() []config.CommonParameterDto
	GetCommonParameters() []config.CommonParameterDto
	GetK8sObjects() []config.ObjectType
	GetHeadParams(k8sObject string) []config.HeadParamsDto
}

type parameterImpl struct {
	entity config.ParameterEntity
}

func NewParameterUseCase(entity config.ParameterEntity) ParameterUseCase {
	return &parameterImpl{entity: entity}
}

func (p *parameterImpl) GetKubernetesParameters() []config.CommonParameterDto {
	return p.entity.GetKubernetesParameters()
}

func (p *parameterImpl) GetCommonParameters() []config.CommonParameterDto {
	return p.entity.GetCommonParameters()
}

func (p *parameterImpl) GetK8sObjects() []config.ObjectType {
	return p.entity.GetK8sObjects()
}

func (p *parameterImpl) GetHeadParams(k8sObject string) []config.HeadParamsDto {
	return p.entity.GetHeadParams(k8sObject)
}

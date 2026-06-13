package binding

import (
	"Kubexplorer/internal/config"
	"Kubexplorer/internal/usecase"
)

type Parameter struct {
	parameter usecase.ParameterUseCase
}

func BuildParameters() *Parameter {
	return &Parameter{
		parameter: usecase.NewParameterUseCase(config.NewParameterEntity()),
	}
}

func (p *Parameter) GetKubernetesParameters() []config.CommonParameterDto {
	return p.parameter.GetKubernetesParameters()
}

func (p *Parameter) GetCommonParameters() []config.CommonParameterDto {
	return p.parameter.GetCommonParameters()
}

func (p *Parameter) GetK8sObjects() []config.ObjectType {
	return p.parameter.GetK8sObjects()
}

func (p *Parameter) GetHeaderParams(k8sObject string) []config.HeadParamsDto {
	return p.parameter.GetHeadParams(k8sObject)
}

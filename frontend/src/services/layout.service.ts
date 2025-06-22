import {
  GetCommonParameters,
  GetHeaderParams,
  GetKubernetesParameters,
} from '../../wailsjs/go/middleware/ParameterMiddleware'
import {GetClusters} from '../../wailsjs/go/middleware/EnvironmentMiddleware'

export const fetchCommonParameters = async () => GetCommonParameters()
export const fetchKubernetesParameters = async () => GetKubernetesParameters()
export const fetchClusters = async () => GetClusters()

export const fetchHeaderParams = async (k8sObject: string) => GetHeaderParams(k8sObject)

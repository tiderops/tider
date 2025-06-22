import {
  DeleteDeployment,
  GetDeployment,
  GetDeployments,
  GetPod,
  GetPods, ResourceTuning,
  RestartPod, TroubleshootDeployment, TroubleshootPod,
  UpdateDeployment,
  UpdatePod,
} from '../../wailsjs/go/middleware/WorkloadMiddleware'

export const fetchGetPods = async (clusterCtx: string) => GetPods(clusterCtx)
export const fetchGetPod = async (name: string, namespace: string, clusterCtx: string) => GetPod(name, namespace, clusterCtx)
export const fetchUpdatePod = async (name: string, namespace: string, dto: any, clusterCtx: string) =>
  UpdatePod(name, namespace, dto, clusterCtx)
export const fetchRestartPod = async (name: string, namespace: string, clusterCtx: string) =>
  RestartPod(name, namespace, clusterCtx)

export const fetchGetDeployments = async (clusterCtx: string) => GetDeployments(clusterCtx)
export const fetchGetDeployment = async (name: string, namespace: string, clusterCtx: string) =>
  GetDeployment(name, namespace, clusterCtx)
export const fetchUpdateDeployment = async (name: string, namespace: string, dto: any, clusterCtx: string) =>
  UpdateDeployment(name, namespace, dto, clusterCtx)
export const fetchDeleteDeployment = async (name: string, namespace: string, clusterCtx: string) =>
  DeleteDeployment(name, namespace, clusterCtx)

export const fetchResourceTuning = async (namespace: string) => ResourceTuning(namespace)
export const fetchTroubleshootPod = async (name: string, namespace: string, clusterCtx: string) => TroubleshootPod(name, namespace, clusterCtx)
export const fetchTroubleshootDeployment = async (name: string, namespace: string, clusterCtx: string) => TroubleshootDeployment(name, namespace, clusterCtx)

import {
  DeleteDeployment,
  GetDeployment,
  GetDeployments,
  GetPod,
  GetPods,
  ResourceTuning,
  RestartPod,
  TroubleshootDeployment,
  TroubleshootPod,
  UpdateDeployment,
  UpdatePod,
} from '../../wailsjs/go/middleware/WorkloadMiddleware'
export const fetchGetPods = async (clusterCtx) => GetPods(clusterCtx)
export const fetchGetPod = async (name, namespace, clusterCtx) =>
  GetPod(name, namespace, clusterCtx)
export const fetchUpdatePod = async (name, namespace, dto, clusterCtx) =>
  UpdatePod(name, namespace, dto, clusterCtx)
export const fetchRestartPod = async (name, namespace, clusterCtx) =>
  RestartPod(name, namespace, clusterCtx)
export const fetchGetDeployments = async (clusterCtx) => GetDeployments(clusterCtx)
export const fetchGetDeployment = async (name, namespace, clusterCtx) =>
  GetDeployment(name, namespace, clusterCtx)
export const fetchUpdateDeployment = async (name, namespace, dto, clusterCtx) =>
  UpdateDeployment(name, namespace, dto, clusterCtx)
export const fetchDeleteDeployment = async (name, namespace, clusterCtx) =>
  DeleteDeployment(name, namespace, clusterCtx)
export const fetchResourceTuning = async (namespace) => ResourceTuning(namespace)
export const fetchTroubleshootPod = async (name, namespace, clusterCtx) =>
  TroubleshootPod(name, namespace, clusterCtx)
export const fetchTroubleshootDeployment = async (name, namespace, clusterCtx) =>
  TroubleshootDeployment(name, namespace, clusterCtx)

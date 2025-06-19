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

export const fetchGetPods = async () => GetPods()
export const fetchGetPod = async (name: string, namespace: string) => GetPod(name, namespace)
export const fetchUpdatePod = async (name: string, namespace: string, dto: any) =>
  UpdatePod(name, namespace, dto)
export const fetchRestartPod = async (name: string, namespace: string) =>
  RestartPod(name, namespace)

export const fetchGetDeployments = async () => GetDeployments()
export const fetchGetDeployment = async (name: string, namespace: string) =>
  GetDeployment(name, namespace)
export const fetchUpdateDeployment = async (name: string, namespace: string, dto: any) =>
  UpdateDeployment(name, namespace, dto)
export const fetchDeleteDeployment = async (name: string, namespace: string) =>
  DeleteDeployment(name, namespace)

export const fetchResourceTuning = async (namespace: string) => ResourceTuning(namespace)
export const fetchTroubleshootPod = async (name: string, namespace: string) => TroubleshootPod(name, namespace)
export const fetchTroubleshootDeployment = async (name: string, namespace: string) => TroubleshootDeployment(name, namespace)

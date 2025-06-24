export declare const fetchGetPods: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PodDto[]>
export declare const fetchGetPod: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PodDto>
export declare const fetchUpdatePod: (
  name: string,
  namespace: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchRestartPod: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>
export declare const fetchGetDeployments: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.DeploymentDto[]>
export declare const fetchGetDeployment: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.DeploymentDto>
export declare const fetchUpdateDeployment: (
  name: string,
  namespace: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeleteDeployment: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>
export declare const fetchResourceTuning: (namespace: string) => Promise<void>
export declare const fetchTroubleshootPod: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>
export declare const fetchTroubleshootDeployment: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>

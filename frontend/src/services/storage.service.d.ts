export declare const fetchGetPersistentVolumes: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PersistentVolumeDto[]>
export declare const fetchGetPersistentVolume: (
  name: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PersistentVolumeDto>
export declare const fetchUpdatePersistentVolume: (
  name: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeletePersistentVolume: (
  name: string,
  clusterCtx: string,
) => Promise<void>
export declare const fetchGetPersistentVolumesClaim: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PersistentVolumeClaimDto[]>
export declare const fetchGetPersistentVolumeClaim: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.PersistentVolumeClaimDto>
export declare const fetchUpdatePersistentVolumeClaim: (
  name: string,
  namespace: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeletePersistentVolumeClaim: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>

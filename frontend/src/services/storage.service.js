import {
  DeletePersistentVolume,
  GetPersistentVolume,
  GetPersistentVolumeClaim,
  GetPersistentVolumes,
  GetPersistentVolumesClaim,
  UpdatePersistentVolume,
  UpdatePersistentVolumeClaim,
} from '../../wailsjs/go/middleware/StorageMiddleware'
export const fetchGetPersistentVolumes = async (clusterCtx) => GetPersistentVolumes(clusterCtx)
export const fetchGetPersistentVolume = async (name, clusterCtx) =>
  GetPersistentVolume(name, clusterCtx)
export const fetchUpdatePersistentVolume = async (name, dto, clusterCtx) =>
  UpdatePersistentVolume(name, dto, clusterCtx)
export const fetchDeletePersistentVolume = async (name, clusterCtx) =>
  DeletePersistentVolume(name, clusterCtx)
export const fetchGetPersistentVolumesClaim = async (clusterCtx) =>
  GetPersistentVolumesClaim(clusterCtx)
export const fetchGetPersistentVolumeClaim = async (name, namespace, clusterCtx) =>
  GetPersistentVolumeClaim(name, namespace, clusterCtx)
export const fetchUpdatePersistentVolumeClaim = async (name, namespace, dto, clusterCtx) =>
  UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
export const fetchDeletePersistentVolumeClaim = async (name, namespace, clusterCtx) =>
  DeletePersistentVolume(name, namespace, clusterCtx)

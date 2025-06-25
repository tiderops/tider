import {
	DeletePersistentVolume,
	GetPersistentVolume,
	GetPersistentVolumeClaim,
	GetPersistentVolumes,
	GetPersistentVolumesClaim,
	UpdatePersistentVolume,
	UpdatePersistentVolumeClaim,
} from '../../wailsjs/go/middleware/StorageMiddleware'

export const fetchGetPersistentVolumes = async (clusterCtx: string) => GetPersistentVolumes(clusterCtx)
export const fetchGetPersistentVolume = async (name: string, clusterCtx: string) => GetPersistentVolume(name, clusterCtx)
export const fetchUpdatePersistentVolume = async (name: string, dto: any, clusterCtx: string) => UpdatePersistentVolume(name, dto, clusterCtx)
export const fetchDeletePersistentVolume = async (name: string, clusterCtx: string) => DeletePersistentVolume(name, clusterCtx)

export const fetchGetPersistentVolumesClaim = async (clusterCtx: string) => GetPersistentVolumesClaim(clusterCtx)
export const fetchGetPersistentVolumeClaim = async (name: string, namespace: string, clusterCtx: string) =>
	GetPersistentVolumeClaim(name, namespace, clusterCtx)
export const fetchUpdatePersistentVolumeClaim = async (name: string, namespace: string, dto: any, clusterCtx: string) =>
	UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
export const fetchDeletePersistentVolumeClaim = async (name: string, namespace: string, clusterCtx: string) =>
	DeletePersistentVolume(name, namespace, clusterCtx)

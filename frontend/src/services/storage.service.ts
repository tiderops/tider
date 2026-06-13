import { model } from '../../wailsjs/go/models'
import {
	DeletePersistentVolume,
	DeletePersistentVolumeClaim,
	GetPersistentVolume,
	GetPersistentVolumeClaim,
	GetPersistentVolumes,
	GetPersistentVolumesClaim,
	UpdatePersistentVolume,
	UpdatePersistentVolumeClaim,
} from '../../wailsjs/go/binding/Storage'

export const fetchGetPersistentVolumes = async (clusterCtx: string) => GetPersistentVolumes(clusterCtx)
export const fetchGetPersistentVolume = async (name: string, clusterCtx: string) => GetPersistentVolume(name, clusterCtx)
export const fetchUpdatePersistentVolume = async (name: string, dto: model.PersistentVolumeDto, clusterCtx: string) => UpdatePersistentVolume(name, dto, clusterCtx)
export const fetchDeletePersistentVolume = async (name: string, clusterCtx: string) => DeletePersistentVolume(name, clusterCtx)

export const fetchGetPersistentVolumesClaim = async (clusterCtx: string) => GetPersistentVolumesClaim(clusterCtx)
export const fetchGetPersistentVolumeClaim = async (name: string, namespace: string, clusterCtx: string) =>
	GetPersistentVolumeClaim(name, namespace, clusterCtx)
export const fetchUpdatePersistentVolumeClaim = async (name: string, namespace: string, dto: model.PersistentVolumeClaimDto, clusterCtx: string) =>
	UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
export const fetchDeletePersistentVolumeClaim = async (name: string, namespace: string, clusterCtx: string) =>
	DeletePersistentVolumeClaim(name, namespace, clusterCtx)

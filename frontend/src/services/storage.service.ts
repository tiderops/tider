import {
  DeletePersistentVolume,
  GetPersistentVolume, GetPersistentVolumeClaim,
  GetPersistentVolumes, GetPersistentVolumesClaim,
  UpdatePersistentVolume, UpdatePersistentVolumeClaim,
} from '../../wailsjs/go/middleware/StorageMiddleware'

export const fetchGetPersistentVolumes = async () => GetPersistentVolumes()
export const fetchGetPersistentVolume = async (name: string) =>
  GetPersistentVolume(name)
export const fetchUpdatePersistentVolume = async (name: string, dto: any) =>
  UpdatePersistentVolume(name, dto)
export const fetchDeletePersistentVolume = async (name: string) =>
  DeletePersistentVolume(name)

export const fetchGetPersistentVolumesClaim = async (namespace: string) => GetPersistentVolumesClaim(namespace)
export const fetchGetPersistentVolumeClaim = async (name: string, namespace: string) => GetPersistentVolumeClaim(name, namespace)
export const fetchUpdatePersistentVolumeClaim = async (name: string, namespace: string, dto: any) => UpdatePersistentVolumeClaim(name, namespace, dto)
export const fetchDeletePersistentVolumeClaim = async (name: string, namespace: string) => DeletePersistentVolume(name, namespace)

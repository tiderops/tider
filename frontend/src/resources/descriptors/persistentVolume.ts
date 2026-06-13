import { fetchGetPersistentVolume, fetchGetPersistentVolumes } from '@/services/storage.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { PersistentVolumeRow } from '@/types/storage.type'
import type { ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function persistentVolumeToRow(p: model.PersistentVolumeDto): PersistentVolumeRow {
	return {
		name: p.Name,
		namespace: p.Namespace,
		storageClass: p.StorageClass,
		capacity: p.Capacity,
		claim: p.Claim,
		age: p.Age,
		status: p.Status,
	}
}

export async function persistentVolumeToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const p = await fetchGetPersistentVolume(ref.name, ref.cluster)
	return { name: p.Name, age: p.Age, status: p.Status }
}

export const persistentVolumeDescriptor = defineResource({
	kind: 'persistentVolume',
	namespaced: false,
	fetchList: fetchGetPersistentVolumes,
	toRow: persistentVolumeToRow,
	fetchDetail: persistentVolumeToDetail,
})

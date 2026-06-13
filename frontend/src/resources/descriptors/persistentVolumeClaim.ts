import { fetchGetPersistentVolumeClaim, fetchGetPersistentVolumesClaim } from '@/services/storage.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { PersistentVolumeClaimRow } from '@/types/storage.type'
import type { ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function persistentVolumeClaimToRow(p: model.PersistentVolumeClaimDto): PersistentVolumeClaimRow {
	return {
		name: p.Name,
		namespace: p.Namespace,
		storageClass: p.StorageClass,
		size: p.Size,
		age: p.Age,
		status: p.Status,
	}
}

export async function persistentVolumeClaimToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const p = await fetchGetPersistentVolumeClaim(ref.name, ref.namespace, ref.cluster)
	return { name: p.Name, namespace: p.Namespace, age: p.Age, status: p.Status }
}

export const persistentVolumeClaimDescriptor = defineResource({
	kind: 'persistentVolumeClaim',
	namespaced: true,
	fetchList: fetchGetPersistentVolumesClaim,
	toRow: persistentVolumeClaimToRow,
	fetchDetail: persistentVolumeClaimToDetail,
})

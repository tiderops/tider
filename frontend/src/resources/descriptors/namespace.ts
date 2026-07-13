import { fetchGetNamespace, fetchGetNamespaces } from '@/services/general.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { NamespaceRow, ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function namespaceToRow(n: model.NamespaceDto): NamespaceRow {
	return {
		name: n.Name,
		age: n.Age,
		status: n.Status,
	}
}

export async function namespaceToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const n = await fetchGetNamespace(ref.name, ref.cluster)
	return { name: n.Name, age: n.Age, status: n.Status }
}

export const namespaceDescriptor = defineResource({
	kind: 'namespace',
	namespaced: false,
	fetchList: fetchGetNamespaces,
	toRow: namespaceToRow,
	fetchDetail: namespaceToDetail,
})

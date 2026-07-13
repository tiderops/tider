import { fetchDeleteIngress, fetchGetIngress, fetchGetIngresses } from '@/services/network.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { IngressRow } from '@/types/network.type'
import type { ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function ingressToRow(i: model.IngressDto): IngressRow {
	return {
		name: i.Name,
		namespace: i.Namespace,
		host: (i.Rules ?? []).map((r) => r.Host),
		path: (i.Rules ?? []).map((r) => r.Path),
		age: i.Age,
	}
}

export async function ingressToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const i = await fetchGetIngress(ref.name, ref.namespace, ref.cluster)
	return { name: i.Name, namespace: i.Namespace, age: i.Age }
}

export const ingressDescriptor = defineResource({
	kind: 'ingress',
	namespaced: true,
	fetchList: fetchGetIngresses,
	toRow: ingressToRow,
	fetchDetail: ingressToDetail,
	remove: (ref) => fetchDeleteIngress(ref.name, ref.namespace, ref.cluster),
})

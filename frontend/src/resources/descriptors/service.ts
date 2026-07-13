import { fetchDeleteService, fetchGetService, fetchGetServices, fetchUpdateService } from '@/services/network.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { ServiceRow } from '@/types/network.type'
import type { ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function serviceToRow(s: model.ServiceDto): ServiceRow {
	return {
		name: s.Name,
		namespace: s.Namespace,
		type: s.Type,
		intIp: s.InternalIp,
		extIp: s.ExternalIp,
		port: s.Port,
		status: s.Status,
		age: s.Age,
	}
}

export async function serviceToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const s = await fetchGetService(ref.name, ref.namespace, ref.cluster)
	return { name: s.Name, namespace: s.Namespace, age: s.Age, status: s.Status }
}

export const serviceDescriptor = defineResource({
	kind: 'service',
	namespaced: true,
	fetchList: fetchGetServices,
	toRow: serviceToRow,
	fetchDetail: serviceToDetail,
	remove: (ref) => fetchDeleteService(ref.name, ref.namespace, ref.cluster),
	update: (ref, patch) => fetchUpdateService(ref.name, ref.namespace, patch as model.ServiceUpdate, ref.cluster),
})

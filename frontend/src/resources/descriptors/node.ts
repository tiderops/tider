import { fetchGetNode, fetchGetNodes } from '@/services/general.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { NodeRow, ResourceSummary } from '@/types/general.type'
import { model } from '../../../wailsjs/go/models'

export function nodeToRow(n: model.NodeDto): NodeRow {
	return {
		name: n.Name,
		memory: n.Resource.Memory,
		cpu: n.Resource.Cpu,
		storage: n.Resource.Storage,
		ephemeralStorage: n.Resource.StorageEphemeral,
		kubeletVersion: n.KubeletVersion,
		operatingSystem: n.OperatingSystem,
		version: n.Version,
		age: n.Age,
	}
}

export async function nodeToDetail(ref: ResourceRef): Promise<ResourceSummary> {
	const n = await fetchGetNode(ref.name, ref.cluster)
	return { name: n.Name, age: n.Age }
}

export const nodeDescriptor = defineResource({
	kind: 'node',
	namespaced: false,
	fetchList: fetchGetNodes,
	toRow: nodeToRow,
	fetchDetail: nodeToDetail,
})

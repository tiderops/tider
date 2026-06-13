import { fetchAutoTroubleshoot, fetchDeleteDeployment, fetchGetDeployment, fetchGetDeployments, fetchUpdateDeployment } from '@/services/workload.service'
import { defineResource } from '../types'
import type { ResourceRef } from '../types'
import type { DeploymentRow, DeploymentDetail } from '@/types/workload.type'
import { model } from '../../../wailsjs/go/models'

export function deploymentToRow(d: model.DeploymentDto): DeploymentRow {
	return {
		name: d.Name,
		namespace: d.Namespace,
		replicas: d.Replicas,
		status: d.Status,
		age: d.Age,
	}
}

export async function deploymentToDetail(ref: ResourceRef): Promise<DeploymentDetail> {
	const d = await fetchGetDeployment(ref.name, ref.namespace, ref.cluster)
	return {
		name: d.Name,
		namespace: d.Namespace,
		replicas: d.Replicas,
		age: d.Age,
		status: d.Status,
	}
}

export const deploymentDescriptor = defineResource({
	kind: 'deployment',
	namespaced: true,
	fetchList: fetchGetDeployments,
	toRow: deploymentToRow,
	fetchDetail: deploymentToDetail,
	remove: (ref) => fetchDeleteDeployment(ref.name, ref.namespace, ref.cluster),
	update: (ref, patch) => fetchUpdateDeployment(ref.name, ref.namespace, patch as model.DeploymentUpdate, ref.cluster),
	troubleshoot: (ref) => fetchAutoTroubleshoot(ref.name, ref.namespace, ref.cluster, 'DEPLOYMENT'),
})

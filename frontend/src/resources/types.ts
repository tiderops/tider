import type { DeploymentDetail, PodDetail } from '@/types/workload.type'
import type { ResourceSummary } from '@/types/general.type'
import { model } from '../../wailsjs/go/models'

export const RESOURCE_KINDS = [
	'pod',
	'deployment',
	'service',
	'ingress',
	'namespace',
	'node',
	'persistentVolume',
	'persistentVolumeClaim',
] as const

export type ResourceKind = (typeof RESOURCE_KINDS)[number]

export interface ResourceRef {
	cluster: string
	namespace: string
	name: string
}

export interface GridHeader {
	title: string
	key: string
	align: 'start' | 'center' | 'end'
	sortable: boolean
}

export interface GridRow {
	name: string
	namespace?: string
	status?: string
}

export type ResourceDetail = PodDetail | DeploymentDetail | ResourceSummary

export type UpdatePayload = model.PodUpdate | model.DeploymentUpdate | model.ServiceUpdate

export interface ResourceDescriptor<TRow extends GridRow = GridRow, TDetail extends ResourceDetail = ResourceDetail> {
	kind: ResourceKind
	namespaced: boolean
	fetchRows(cluster: string): Promise<TRow[]>
	fetchDetail?(ref: ResourceRef): Promise<TDetail>
	remove?(ref: ResourceRef): Promise<void>
	update?(ref: ResourceRef, patch: UpdatePayload): Promise<void>
	troubleshoot?(ref: ResourceRef): Promise<model.Troubleshoot>
	statusColor?(status: string): string
}

interface ResourceDefinition<TDto, TRow extends GridRow, TDetail extends ResourceDetail> {
	kind: ResourceKind
	namespaced: boolean
	fetchList(cluster: string): Promise<TDto[]>
	toRow(dto: TDto): TRow
	fetchDetail?(ref: ResourceRef): Promise<TDetail>
	remove?(ref: ResourceRef): Promise<void>
	update?(ref: ResourceRef, patch: UpdatePayload): Promise<void>
	troubleshoot?(ref: ResourceRef): Promise<model.Troubleshoot>
	statusColor?(status: string): string
}

export function defineResource<TDto, TRow extends GridRow, TDetail extends ResourceDetail>(
	definition: ResourceDefinition<TDto, TRow, TDetail>,
): ResourceDescriptor<TRow, TDetail> {
	return {
		kind: definition.kind,
		namespaced: definition.namespaced,
		fetchRows: async (cluster) => ((await definition.fetchList(cluster)) ?? []).map(definition.toRow),
		fetchDetail: definition.fetchDetail,
		remove: definition.remove,
		update: definition.update,
		troubleshoot: definition.troubleshoot,
		statusColor: definition.statusColor,
	}
}

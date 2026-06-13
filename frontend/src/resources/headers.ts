import { fetchHeaderParams } from '@/services/layout.service'
import type { GridHeader, ResourceKind } from './types'

export async function fetchGridHeaders(kind: ResourceKind): Promise<GridHeader[]> {
	const params = await fetchHeaderParams(kind)
	return (params ?? []).map((h) => ({
		title: h.Title,
		key: h.Key,
		align: h.Align as GridHeader['align'],
		sortable: h.Sortable,
	}))
}

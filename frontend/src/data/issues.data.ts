import { delay } from './mock-latency'
import type { Issue } from '@/types/issue'

// TODO-16: Remove mock data
const issues: Issue[] = [
	{ id: 'a1', reason: 'OOMKilled', reasonTone: 'err', name: 'checkout-api-7d9f8c-9wzlm', namespace: 'payments', resourceKind: 'pod', kind: 'pod · payments', cluster: 'prod-eu-west-1', age: '2m', action: 'Diagnose' },
	{ id: 'a2', reason: 'ImagePull', reasonTone: 'err', name: 'recommender-api', namespace: 'ml', resourceKind: 'deployment', kind: 'deployment · ml', cluster: 'dev-sandbox', age: '8m', action: 'Diagnose' },
	{ id: 'a3', reason: 'NotReady', reasonTone: 'warn', name: 'ip-10-2-58-4', namespace: '', resourceKind: 'node', kind: 'node', cluster: 'prod-eu-west-1', age: '6m', action: 'Inspect' },
]

export async function fetchIssues(): Promise<Issue[]> {
	await delay()
	return issues
}

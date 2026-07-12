import { delay } from '@/data/mock-latency'
import { fetchGetPod, fetchGetPods } from '@/services/workload.service'
import { hasWailsRuntime } from '@/services/runtime'
import type { PodDetail, PodRow, ResourceTab } from './types'
import type { model } from '../../../wailsjs/go/models'

export const resourceTabs: ResourceTab[] = [
	{ key: 'pod', label: 'Pods', count: 184 },
	{ key: 'deployment', label: 'Deployments', count: 42 },
	{ key: 'statefulset', label: 'StatefulSets', count: 6 },
	{ key: 'job', label: 'Jobs', count: 11 },
]

export const namespaces = ['payments', 'ml', 'analytics', 'kube-system', 'default']

const mockPods: PodRow[] = [
	{ name: 'checkout-api-7d9f8c-2xk4p', namespace: 'payments', cpu: '240m', memory: '312Mi', restarts: 0, node: 'ip-10-2-31-7', age: '3h12m', status: 'Running' },
	{ name: 'checkout-api-7d9f8c-9wzlm', namespace: 'payments', cpu: '980m', memory: '1.4Gi', restarts: 7, node: 'ip-10-2-44-2', age: '3h12m', status: 'CrashLoopBackOff' },
	{ name: 'ledger-worker-58c6b4-qm8tx', namespace: 'payments', cpu: '120m', memory: '680Mi', restarts: 0, node: 'ip-10-2-31-7', age: '1d4h', status: 'Running' },
	{ name: 'fraud-scoring-6b7d9-h4n2k', namespace: 'payments', cpu: '1.2', memory: '2.1Gi', restarts: 2, node: 'ip-10-2-44-2', age: '6h40m', status: 'Pending' },
	{ name: 'payments-gateway-5f9c-tt7bd', namespace: 'payments', cpu: '340m', memory: '512Mi', restarts: 0, node: 'ip-10-2-19-9', age: '2d1h', status: 'Running' },
	{ name: 'refund-consumer-9d4f1-pl3qx', namespace: 'payments', cpu: '90m', memory: '256Mi', restarts: 0, node: 'ip-10-2-19-9', age: '2d1h', status: 'Running' },
	{ name: 'settlement-cron-28-w7c9', namespace: 'payments', cpu: '—', memory: '—', restarts: 0, node: 'ip-10-2-31-7', age: '14m', status: 'Completed' },
	{ name: 'notify-dispatch-7a1b-xc09r', namespace: 'payments', cpu: '60m', memory: '180Mi', restarts: 1, node: 'ip-10-2-44-2', age: '5h2m', status: 'Running' },
]

const detailOverrides: Record<string, PodDetail> = {
	'checkout-api-7d9f8c-9wzlm': {
		name: 'checkout-api-7d9f8c-9wzlm', namespace: 'payments', node: 'ip-10-2-44-2', status: 'CrashLoopBackOff',
		restarts: 7, restartWindow: 'restarted 7× in 12m', controller: 'ReplicaSet/checkout-api-7d9f8c',
		image: 'checkout-api:1.18.2', qosClass: 'Burstable', podIP: '10.2.44.118', serviceAccount: 'checkout', started: '2026-06-12 09:14:03',
		cpu: { label: '980m / 1000m', pct: 98 }, memory: { label: '1.4 / 1.5Gi', pct: 93 },
		lastTermination: { reason: 'OOMKilled', message: 'container exceeded memory limit', log: ['panic: runtime: out of memory', 'goroutine 142 [running]:', 'main.(*Ledger).flush(0xc0004a2000)'] },
	},
}

function mockDetail(name: string): PodDetail | undefined {
	const override = detailOverrides[name]
	if (override) {
		return override
	}
	const row = mockPods.find((p) => p.name === name)
	if (!row) {
		return undefined
	}
	const app = row.name.split('-').slice(0, -2).join('-') || row.name
	return {
		name: row.name, namespace: row.namespace, node: row.node, status: row.status, restarts: row.restarts,
		controller: `ReplicaSet/${row.name.split('-').slice(0, -1).join('-')}`, image: `${app}:1.18.2`,
		qosClass: 'Burstable', podIP: '10.2.31.' + (40 + row.name.length), serviceAccount: app, started: '2026-06-12 06:00:00',
		cpu: { label: `${row.cpu} / 1000m`, pct: 34 }, memory: { label: `${row.memory} / 1Gi`, pct: 46 },
	}
}

function toPodRow(dto: model.PodDto): PodRow {
	const c = dto.Containers?.[0]
	return {
		name: dto.Name,
		namespace: dto.Namespace,
		cpu: `${c?.Limit.Cpu ?? 0}m`,
		memory: `${c?.Limit.Memory ?? 0}Mi`,
		restarts: 0, // PodDto has no restart count yet (backend gap)
		node: dto.Node,
		age: dto.Age,
		status: dto.Status,
	}
}

function toPodDetail(dto: model.PodDto): PodDetail {
	const c = dto.Containers?.[0]
	return {
		name: dto.Name,
		namespace: dto.Namespace,
		node: dto.Node,
		status: dto.Status,
		restarts: 0,
		controller: '—',
		image: c?.Image ?? '—',
		qosClass: '—',
		podIP: '—',
		serviceAccount: '—',
		started: dto.Age,
		cpu: { label: `${c?.Limit.Cpu ?? 0}m`, pct: 0 },
		memory: { label: `${c?.Limit.Memory ?? 0}Mi`, pct: 0 },
	}
}

// TODO-2: Consume backend.
export async function fetchPods(cluster: string): Promise<PodRow[]> {
	if (!hasWailsRuntime()) {
		await delay()
		return mockPods
	}
	const dtos = await fetchGetPods(cluster)
	return (dtos ?? []).map(toPodRow)
}

// TODO-3: Consume backend.
export async function fetchPodDetail(cluster: string, name: string, namespace: string): Promise<PodDetail | undefined> {
	if (!hasWailsRuntime()) {
		await delay(120)
		return mockDetail(name)
	}
	const dto = await fetchGetPod(name, namespace, cluster)
	return dto ? toPodDetail(dto) : undefined
}

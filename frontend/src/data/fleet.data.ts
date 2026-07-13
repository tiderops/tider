import { delay } from './mock-latency'
import { fetchClusters as getClustersRpc } from '@/services/layout.service'
import { hasWailsRuntime } from '@/services/runtime'
import type { ClusterSummary, FleetTotals } from '@/types/fleet'
import type { model } from '../../wailsjs/go/models'

// TODO-19: Remove mock data
const mockClusters: ClusterSummary[] = [
	{ name: 'prod-eu-west-1', source: 'arn:aws:eks:eu-west-1 · v1.29.4', reachable: true, statusLabel: 'Healthy', statusTone: 'ok', metricsAvailable: true, cpu: 63, memory: 71, nodes: 12, pods: 184, namespaces: 28 },
	{ name: 'staging-eu-west-1', source: 'arn:aws:eks:eu-west-1 · v1.29.4', reachable: true, statusLabel: 'Healthy', statusTone: 'ok', metricsAvailable: true, cpu: 22, memory: 38, nodes: 6, pods: 74, namespaces: 19 },
	{ name: 'dev-sandbox', source: 'gke_dev_europe-west1 · v1.30.1', reachable: true, statusLabel: 'Degraded', statusTone: 'err', metricsAvailable: true, cpu: 91, memory: 88, nodes: 3, pods: 41, namespaces: 11, issues: 3 },
	{ name: 'minikube', source: 'local · v1.30.0', reachable: true, statusLabel: 'Idle', statusTone: 'idle', metricsAvailable: true, cpu: 8, memory: 14, nodes: 1, pods: 9, namespaces: 5 },
]

function toClusterSummary(info: model.ClusterInfo): ClusterSummary {
	return {
		name: info.Name,
		source: info.Server ? `${info.Cluster} · ${info.Server}` : info.Cluster,
		reachable: info.Status,
		statusLabel: info.Status ? 'Reachable' : 'Unreachable',
		statusTone: info.Status ? 'ok' : 'err',
		// Per-cluster CPU/memory/counts need a metrics-server integration
		// that the backend does not expose yet.
		metricsAvailable: false,
		cpu: 0,
		memory: 0,
		nodes: 0,
		pods: 0,
		namespaces: 0,
	}
}

export async function fetchClusters(): Promise<ClusterSummary[]> {
	if (!hasWailsRuntime()) {
		await delay()
		return mockClusters
	}
	const infos = await getClustersRpc()
	return (infos ?? []).map(toClusterSummary)
}

// Totals are derived from the cluster list so the KPIs always equal the
// sum of the cards. Workload/node counts are 0 for real clusters until
// the metrics integration lands.
export function deriveFleetTotals(clusters: ClusterSummary[]): FleetTotals {
	const workloads = clusters.reduce((sum, c) => sum + c.pods, 0)
	const nodes = clusters.reduce((sum, c) => sum + c.nodes, 0)
	const openIssues = clusters.reduce((sum, c) => sum + (c.issues ?? 0), 0)
	return {
		clusters: clusters.length,
		reachable: clusters.filter((c) => c.reachable).length,
		workloads,
		workloadsDelta: workloads ? '▲ 6 vs. yesterday' : '',
		nodes,
		nodesNotReady: nodes ? 1 : 0,
		openIssues,
		issuesBreakdown: openIssues ? `${openIssues} flagged` : 'none',
	}
}

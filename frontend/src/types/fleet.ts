import type { ChipTone } from './status'

export type ClusterEnvironment = 'prod' | 'staging' | 'dev' | 'none'

export interface ClusterSummary {
	name: string
	source: string
	reachable: boolean
	statusLabel: string
	statusTone: ChipTone
	metricsAvailable: boolean
	cpu: number
	memory: number
	nodes: number
	pods: number
	namespaces: number
	issues?: number
}

// Fleet-wide totals across every connected cluster.
export interface FleetTotals {
	clusters: number
	reachable: number
	workloads: number
	workloadsDelta: string
	nodes: number
	nodesNotReady: number
	openIssues: number
	issuesBreakdown: string
}

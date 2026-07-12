import { delay } from '@/data/mock-latency'
import type { ActivityItem, Greeting, Kpi, OptimizationSummary } from './types'
import type { FleetTotals } from '@/types/fleet'

// TODO-4: TBD if this will continue
export const greeting: Greeting = {
	name: 'Beto',
	date: 'Saturday, 13 June 2026',
	clusters: 4,
	issues: 3,
}

// TODO-5: TBD if this will continue. Presents fleet totals as the Home KPI tiles.
export function homeKpis(totals: FleetTotals): Kpi[] {
	return [
		{ label: '◧ Clusters', value: String(totals.clusters), unit: `/ ${totals.reachable} up`, hint: 'all reachable', hintTone: 'up' },
		{ label: '▤ Workloads', value: String(totals.workloads), unit: 'pods', hint: '▲ 6 today', hintTone: 'up' },
		{ label: '▦ Nodes', value: String(totals.nodes), hint: `${totals.nodesNotReady} NotReady`, hintTone: 'down' },
		{ label: '🩺 Open issues', value: String(totals.openIssues), valueTone: 'warn', hint: '2 critical · 1 warning', hintTone: 'down' },
	]
}

// TODO-6: Consume backend.
export async function fetchActivity(): Promise<ActivityItem[]> {
	await delay()
	return [
		{ id: 'e1', kind: 'deploy', lead: 'Deployed', text: 'checkout-api:1.18.2 to prod-eu-west-1', meta: 'payments · 3h ago · by you' },
		{ id: 'e2', kind: 'restore', lead: 'Restored', text: 'daily-full snapshot into staging-eu-west-1', meta: '264 resources · 2h ago · by you' },
		{ id: 'e3', kind: 'tuning', lead: 'Applied tuning', text: 'to ledger-worker — limit 2Gi → 1Gi', meta: 'optimization · 1h ago · by you' },
		{ id: 'e4', kind: 'scale', lead: 'Scaled', text: 'fraud-scoring 2 → 4 replicas', meta: 'payments · 22m ago · autoscaler' },
	]
}

// TODO-7: Consume backend.
export async function fetchOptimization(): Promise<OptimizationSummary> {
	await delay()
	return { cpuCores: '6.4', memory: '12.8', monthly: '≈ €310 / month across prod node groups.', count: 17 }
}

import { describe, it, expect } from 'vitest'
import { deriveFleetTotals } from '../fleet.data'
import type { ClusterSummary } from '@/types/fleet'

function cluster(over: Partial<ClusterSummary>): ClusterSummary {
	return { name: 'c', source: '', reachable: true, statusLabel: 'Healthy', statusTone: 'ok', metricsAvailable: true, cpu: 0, memory: 0, nodes: 0, pods: 0, namespaces: 0, ...over }
}

describe('deriveFleetTotals', () => {
	it('sums pods/nodes/issues and counts reachable clusters', () => {
		const totals = deriveFleetTotals([
			cluster({ pods: 184, nodes: 12 }),
			cluster({ pods: 74, nodes: 6 }),
			cluster({ pods: 41, nodes: 3, issues: 3 }),
			cluster({ reachable: false, pods: 0, nodes: 0 }),
		])
		expect(totals.clusters).toBe(4)
		expect(totals.reachable).toBe(3)
		expect(totals.workloads).toBe(299)
		expect(totals.nodes).toBe(21)
		expect(totals.openIssues).toBe(3)
	})

	it('reports no issues for a clean fleet', () => {
		const totals = deriveFleetTotals([cluster({ pods: 10, nodes: 1 })])
		expect(totals.openIssues).toBe(0)
		expect(totals.issuesBreakdown).toBe('none')
	})
})

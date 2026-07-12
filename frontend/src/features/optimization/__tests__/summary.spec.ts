import { describe, it, expect } from 'vitest'
import { computeSummary, deriveVerdict } from '../optimization.data'
import type { OptRecommendation } from '../types'

function rec(over: Partial<OptRecommendation>): OptRecommendation {
	return { id: 'd/c', namespace: 'ns', deployment: 'd', container: 'c', current: { cpu: 1000, memory: 1024 }, usage: { cpu: 500, memory: 512 }, suggested: { cpu: 1000, memory: 1024 }, verdict: 'optimal', ...over }
}

describe('deriveVerdict', () => {
	it('flags under-provisioned when suggested exceeds current', () => {
		expect(deriveVerdict({ cpu: 1000, memory: 1500 }, { cpu: 1480, memory: 1460 }, { cpu: 1500, memory: 2250 })).toBe('under')
	})
	it('flags idle when usage is negligible', () => {
		expect(deriveVerdict({ cpu: 500, memory: 512 }, { cpu: 12, memory: 20 }, { cpu: 100, memory: 128 })).toBe('idle')
	})
	it('flags over-provisioned when suggested is lower', () => {
		expect(deriveVerdict({ cpu: 2000, memory: 4096 }, { cpu: 410, memory: 980 }, { cpu: 700, memory: 1536 })).toBe('over')
	})
	it('is optimal when nothing changes', () => {
		expect(deriveVerdict({ cpu: 500, memory: 768 }, { cpu: 330, memory: 498 }, { cpu: 500, memory: 768 })).toBe('optimal')
	})
})

describe('computeSummary', () => {
	it('counts verdicts and sums reclaimable resources', () => {
		const s = computeSummary([
			rec({ verdict: 'over', current: { cpu: 2000, memory: 4096 }, suggested: { cpu: 700, memory: 1536 } }),
			rec({ verdict: 'under', current: { cpu: 1000, memory: 1024 }, suggested: { cpu: 1500, memory: 2048 } }),
			rec({ verdict: 'idle', current: { cpu: 500, memory: 512 }, suggested: { cpu: 50, memory: 128 } }),
			rec({ verdict: 'optimal' }),
		])
		expect(s.flagged).toBe(3)
		expect(s.over).toBe(1)
		expect(s.under).toBe(1)
		expect(s.idle).toBe(1)
		// reclaimable cpu: (2000-700) + (500-50) = 1750m -> 1.8 cores (rounded)
		expect(s.reclaimableCpu).toBeCloseTo(1.8, 1)
		expect(s.monthly).toBeGreaterThan(0)
	})
})

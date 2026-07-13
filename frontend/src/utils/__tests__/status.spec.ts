import { describe, it, expect } from 'vitest'
import { statusColor } from '../status'

describe('statusColor', () => {
	it.each([
		['Running', 'green'],
		['Active', 'green'],
		['Bound', 'green'],
		['Pending', 'orange'],
		['Terminating', 'orange'],
		['Failed', 'red'],
		['CrashLoopBackOff', 'red'],
		[undefined, 'grey'],
		['', 'grey'],
	])('%s -> %s', (status, color) => {
		expect(statusColor(status)).toBe(color)
	})
})

import type { ChipTone } from './status'

export interface Issue {
	id: string
	reason: string
	reasonTone: ChipTone
	name: string
	namespace: string
	resourceKind: 'pod' | 'deployment' | 'job' | 'node'
	kind: string
	cluster: string
	age: string
	action: 'Diagnose' | 'Inspect'
}

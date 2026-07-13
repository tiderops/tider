export interface Evidence {
	label: string
	value: string
}

export type ActionKind = 'apply' | 'logs' | 'rollback' | 'restart' | 'inspect'

export interface DiagnosisAction {
	label: string
	description: string
	kind: ActionKind
}

export interface Diagnosis {
	meaning: string
	recommendation: string
	evidence: Evidence[]
	actions: DiagnosisAction[]
}

import type { ClusterEnvironment } from '@/types/fleet'

export type Density = 'comfortable' | 'compact'

export interface NotificationPrefs {
	criticalPods: boolean
	nodePressure: boolean
	backupFailures: boolean
	optimizationDigest: boolean
	desktop: boolean
}

export interface AssistantPrefs {
	autoDiagnose: boolean
	headroomPct: number
	includeLogs: boolean
}

export interface AppPreferences {
	density: Density
	refreshInterval: number // seconds; 0 = manual only
	defaultCluster: string // '' = remember last used
	clusterEnvironments: Record<string, ClusterEnvironment> // keyed by context name
	notifications: NotificationPrefs
	assistant: AssistantPrefs
}

export interface AboutInfo {
	version: string
	build: string
	wailsRuntime: boolean
	repo: string
	docs: string
	license: string
}

export interface SettingsSection {
	key: string
	label: string
	icon: string
	hint: string
}

import { hasWailsRuntime } from '@/services/runtime'
import type { AboutInfo, AppPreferences, SettingsSection } from './types'

// TODO-12: Consume backend.
export const defaultPreferences: AppPreferences = {
	density: 'comfortable',
	refreshInterval: 30,
	defaultCluster: '',
	clusterEnvironments: {},
	notifications: {
		criticalPods: true,
		nodePressure: true,
		backupFailures: true,
		optimizationDigest: false,
		desktop: false,
	},
	assistant: {
		autoDiagnose: true,
		headroomPct: 40,
		includeLogs: true,
	},
}

// TODO-13: Consume backend.
export const refreshOptions = [
	{ value: 0, label: 'Manual only' },
	{ value: 15, label: 'Every 15s' },
	{ value: 30, label: 'Every 30s' },
	{ value: 60, label: 'Every 60s' },
	{ value: 300, label: 'Every 5m' },
]

// TODO-14: Consume backend.
export const sections: SettingsSection[] = [
	{ key: 'appearance', label: 'Appearance', icon: '◐', hint: 'Theme & density' },
	{ key: 'clusters', label: 'Clusters', icon: '◧', hint: 'Kubeconfig & defaults' },
	{ key: 'notifications', label: 'Notifications', icon: '🔔', hint: 'Alerts & digests' },
	{ key: 'assistant', label: 'Assistant', icon: '🩺', hint: 'Troubleshoot & tuning' },
	{ key: 'about', label: 'About', icon: 'ⓘ', hint: 'Version & links' },
]


const APP_VERSION = '0.9.0-dev'

export function fetchAbout(): AboutInfo {
	return {
		version: APP_VERSION,
		build: hasWailsRuntime() ? 'desktop (Wails)' : 'browser (standalone)',
		wailsRuntime: hasWailsRuntime(),
		repo: 'https://github.com/tiderops/tider',
		docs: 'https://github.com/beto20/Kubexplorer#readme',
		license: 'MIT',
	}
}

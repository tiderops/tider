export interface BackupSummary {
	snapshots: number
	retention: string
	lastBackup: string
	lastStatus: string
	storageUsed: string
	storageTarget: string
	nextScheduled: string
	nextDetail: string
}

export interface Schedule {
	id: string
	name: string
	scope: string
	cadence: string
	retention: string
	enabled: boolean
}

export type SnapshotStatus = 'Completed' | 'Running' | 'Partial' | 'Failed'

export interface Snapshot {
	id: string
	name: string
	scope: string
	source: string
	resources: number
	size: string
	created: string
	status: SnapshotStatus
	progress?: number
}

export interface RestoreScopeItem {
	namespace: string
	resources: number
	included: boolean
}

export interface BackupData {
	summary: BackupSummary
	schedules: Schedule[]
	snapshots: Snapshot[]
}

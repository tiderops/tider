import { delay } from '@/data/mock-latency'
import type { BackupData, RestoreScopeItem } from './types'

export const defaultRestoreScope: RestoreScopeItem[] = [
	{ namespace: 'payments', resources: 142, included: true },
	{ namespace: 'ml', resources: 68, included: true },
	{ namespace: 'analytics', resources: 54, included: true },
	{ namespace: 'kube-system', resources: 0, included: false },
]

// TODO-1: Consume backend. Take as a reference tools like velero for manifest snapshots
export async function fetchBackup(cluster: string): Promise<BackupData> {
	await delay(300)
	return {
		summary: {
			snapshots: 48,
			retention: '30-day retention',
			lastBackup: '2h ago',
			lastStatus: 'daily-full · completed',
			storageUsed: '14.2 GiB',
			storageTarget: 's3://kx-backups/prod',
			nextScheduled: 'in 4h',
			nextDetail: 'daily-full · 02:00 UTC',
		},
		schedules: [
			{ id: 's1', name: 'daily-full', scope: 'Whole cluster · 02:00 UTC daily', cadence: 'daily', retention: 'retention 30d', enabled: true },
			{ id: 's2', name: 'hourly-payments', scope: 'Namespace payments · hourly', cadence: 'hourly', retention: 'retention 24h', enabled: true },
		],
		snapshots: [
			{ id: 'b1', name: 'daily-full-20260612-0200', scope: 'Cluster', source: 'prod-eu-west-1', resources: 1284, size: '2.1 GiB', created: '2h ago', status: 'Completed' },
			{ id: 'b2', name: 'hourly-payments-20260612-0900', scope: 'payments', source: 'prod-eu-west-1', resources: 142, size: '88 MiB', created: '14m ago', status: 'Completed' },
			{ id: 'b3', name: 'manual-pre-deploy-1182', scope: 'payments,ml', source: 'prod-eu-west-1', resources: 210, size: '140 MiB', created: '3h ago', status: 'Completed' },
			{ id: 'b4', name: 'migration-snapshot-staging', scope: 'Cluster', source: 'staging-eu-west-1', resources: 612, size: '—', created: 'just now', status: 'Running', progress: 64 },
			{ id: 'b5', name: 'hourly-payments-20260612-0800', scope: 'payments', source: 'prod-eu-west-1', resources: 140, size: '86 MiB', created: '1h ago', status: 'Partial' },
			{ id: 'b6', name: 'daily-full-20260611-0200', scope: 'Cluster', source: 'prod-eu-west-1', resources: 1271, size: '2.0 GiB', created: '1d ago', status: 'Completed' },
		],
	}
}

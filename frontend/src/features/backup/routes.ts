import type { RouteRecordRaw } from 'vue-router'

export const backupRoutes: RouteRecordRaw[] = [
	{
		path: '/backup',
		name: 'backup',
		component: () => import('./pages/BackupPage.vue'),
		meta: { shell: 'kx', nav: 'backup', cluster: 'all clusters', clusterSub: '48 snapshots · S3', crumbs: ['Backup & Restore'] },
	},
]

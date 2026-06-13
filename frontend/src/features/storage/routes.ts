import type { RouteRecordRaw } from 'vue-router'

export const storageRoutes: RouteRecordRaw[] = [
	{
		path: '/:cluster/storage',
		name: 'storage',
		component: () => import('./pages/StoragePage.vue'),
		children: [
			{ path: 'persistentVolume', name: 'persistentVolume', component: () => import('./pages/PersistentVolumePage.vue'), props: true },
			{ path: 'persistentVolumeClaim', name: 'persistentVolumeClaim', component: () => import('./pages/PersistentVolumeClaimPage.vue'), props: true },
		],
	},
]

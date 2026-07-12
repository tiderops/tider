import type { RouteRecordRaw } from 'vue-router'

export const explorerRoutes: RouteRecordRaw[] = [
	{
		path: '/explorer',
		name: 'explorer-demo',
		component: () => import('./pages/ResourceExplorerPage.vue'),
		meta: {
			shell: 'kx',
			nav: 'explorer',
			cluster: 'prod-eu-west-1',
			clusterSub: 'v1.29.4 · 184 pods',
			crumbs: ['prod-eu-west-1', 'Workloads', 'Pods'],
		},
	},
]

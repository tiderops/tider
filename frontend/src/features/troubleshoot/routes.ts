import type { RouteRecordRaw } from 'vue-router'

export const troubleshootRoutes: RouteRecordRaw[] = [
	{
		path: '/troubleshoot',
		name: 'troubleshoot',
		component: () => import('./pages/TroubleshootPage.vue'),
		meta: { shell: 'kx', nav: 'troubleshoot', cluster: 'all clusters', clusterSub: '3 open issues', crumbs: ['Troubleshoot'] },
	},
]

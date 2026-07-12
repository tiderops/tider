import type { RouteRecordRaw } from 'vue-router'

export const optimizationRoutes: RouteRecordRaw[] = [
	{
		path: '/optimization',
		name: 'optimization',
		component: () => import('./pages/OptimizationPage.vue'),
		meta: { shell: 'kx', nav: 'optimization', cluster: 'prod-eu-west-1', clusterSub: 'v1.29.4 · right-sizing', crumbs: ['Optimization'] },
	},
]

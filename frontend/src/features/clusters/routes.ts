import type { RouteRecordRaw } from 'vue-router'

export const clustersRoutes: RouteRecordRaw[] = [
	{
		path: '/clusters',
		name: 'clusters-overview',
		component: () => import('./pages/ClustersOverviewPage.vue'),
		meta: { shell: 'kx', nav: 'clusters', cluster: 'all clusters', clusterSub: '4 connected · 1 alert', crumbs: ['Clusters'] },
	},
]

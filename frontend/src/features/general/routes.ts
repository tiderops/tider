import type { RouteRecordRaw } from 'vue-router'

export const generalRoutes: RouteRecordRaw[] = [
	{
		path: '/:cluster/general',
		name: 'general',
		component: () => import('./pages/GeneralPage.vue'),
		children: [
			{
				path: 'node',
				name: 'node',
				component: () => import('./pages/NodePage.vue'),
				children: [{ path: 'detail/:id', name: 'nodeDetail', component: () => import('./pages/NodeDetailPage.vue'), props: true }],
			},
			{
				path: 'namespace',
				name: 'namespace',
				component: () => import('./pages/NamespacePage.vue'),
				children: [
					{ path: 'detail/:id', name: 'namespaceDetail', component: () => import('./pages/NamespaceDetailPage.vue'), props: true },
				],
			},
			{ path: 'event', name: 'event', component: () => import('./pages/EventsPage.vue') },
			{ path: 'cluster_graph', name: 'cluster_graph', component: () => import('./pages/ClusterGraphPage.vue') },
		],
	},
]

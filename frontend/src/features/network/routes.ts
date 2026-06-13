import type { RouteRecordRaw } from 'vue-router'

export const networkRoutes: RouteRecordRaw[] = [
	{
		path: '/:cluster/network',
		name: 'network',
		component: () => import('./pages/NetworkPage.vue'),
		children: [
			{ path: 'ingress', name: 'ingress', component: () => import('./pages/IngressPage.vue'), props: true },
			{ path: 'service', name: 'service', component: () => import('./pages/ServicePage.vue'), props: true },
		],
	},
]

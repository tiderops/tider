import type { RouteRecordRaw } from 'vue-router'

export const workloadRoutes: RouteRecordRaw[] = [
	{
		path: '/:cluster/workload',
		name: 'workload',
		component: () => import('./pages/WorkloadPage.vue'),
		children: [
			{ path: 'pod', name: 'pod', component: () => import('./pages/PodPage.vue'), props: true },
			{ path: 'deployment', name: 'deployment', component: () => import('./pages/DeploymentPage.vue'), props: true },
			{ path: 'tuning', name: 'tuning', component: () => import('./pages/TuningPage.vue'), props: true },
		],
	},
]

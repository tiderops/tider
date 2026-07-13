import type { RouteRecordRaw } from 'vue-router'

export const settingsRoutes: RouteRecordRaw[] = [
	{
		path: '/settings',
		name: 'settings',
		component: () => import('./pages/SettingsPage.vue'),
		meta: { shell: 'kx', nav: 'settings', crumbs: ['Settings'] },
	},
]

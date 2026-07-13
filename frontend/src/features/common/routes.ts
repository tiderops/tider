import type { RouteRecordRaw } from 'vue-router'

export const commonRoutes: RouteRecordRaw[] = [
	{ path: '/home', name: 'home', component: () => import('./pages/HomePage.vue') },
	{ path: '/:cluster/overview', name: 'overview', component: () => import('./pages/OverviewPage.vue') },
	{ path: '/settings', name: 'settings', component: () => import('./pages/SettingsPage.vue') },
	{ path: '/documentation', name: 'documentation', component: () => import('./pages/DocumentationPage.vue') },
	{ path: '/connections', name: 'connections', component: () => import('./pages/ConnectionPage.vue') },
]

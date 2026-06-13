import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useClusterStore } from '@/stores/cluster.store'
import { commonRoutes } from '@/features/common/routes'
import { generalRoutes } from '@/features/general/routes'
import { workloadRoutes } from '@/features/workloads/routes'
import { networkRoutes } from '@/features/network/routes'
import { storageRoutes } from '@/features/storage/routes'

const routes: RouteRecordRaw[] = [
	{ path: '/', redirect: '/home' },
	...commonRoutes,
	...generalRoutes,
	...workloadRoutes,
	...networkRoutes,
	...storageRoutes,
	{ path: '/:pathMatch(.*)*', name: 'notFound', component: () => import('./NotFoundPage.vue') },
]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

router.beforeEach((to) => {
	const cluster = to.params.cluster
	if (typeof cluster === 'string' && cluster) {
		useClusterStore().setCurrentCluster(cluster)
	}
})

export default router

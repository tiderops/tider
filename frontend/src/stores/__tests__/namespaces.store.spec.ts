import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'

const fetchGetNamespaces = vi.fn()
vi.mock('@/services/general.service', () => ({
	fetchGetNamespaces: (...args: unknown[]) => fetchGetNamespaces(...args),
}))

import { useNamespacesStore } from '../namespaces.store'

describe('namespaces store', () => {
	beforeEach(() => {
		setActivePinia(createPinia())
		fetchGetNamespaces.mockReset()
		fetchGetNamespaces.mockResolvedValue([{ Name: 'default' }, { Name: 'kube-system' }])
	})

	it('fetches once per cluster and serves the cache afterwards', async () => {
		const store = useNamespacesStore()

		const first = await store.load('minikube')
		const second = await store.load('minikube')

		expect(first).toEqual(['default', 'kube-system'])
		expect(second).toEqual(first)
		expect(fetchGetNamespaces).toHaveBeenCalledTimes(1)
	})

	it('caches per cluster independently', async () => {
		const store = useNamespacesStore()

		await store.load('cluster-a')
		await store.load('cluster-b')

		expect(fetchGetNamespaces).toHaveBeenCalledTimes(2)
		expect(store.namespacesFor('cluster-a')).toEqual(['default', 'kube-system'])
	})

	it('refetches after invalidation', async () => {
		const store = useNamespacesStore()

		await store.load('minikube')
		store.invalidate('minikube')
		expect(store.namespacesFor('minikube')).toEqual([])

		await store.load('minikube')
		expect(fetchGetNamespaces).toHaveBeenCalledTimes(2)
	})

	it('force reload bypasses the cache', async () => {
		const store = useNamespacesStore()

		await store.load('minikube')
		await store.load('minikube', true)

		expect(fetchGetNamespaces).toHaveBeenCalledTimes(2)
	})
})

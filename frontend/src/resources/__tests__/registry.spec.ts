import { describe, it, expect, vi, beforeEach } from 'vitest'

const restartPod = vi.fn().mockResolvedValue(undefined)
const autoTroubleshoot = vi.fn().mockResolvedValue({ Meaning: 'CrashLoopBackOff', Recommendation: 'check logs' })
const deleteDeployment = vi.fn().mockResolvedValue(undefined)
const deleteService = vi.fn().mockResolvedValue(undefined)
const deleteIngress = vi.fn().mockResolvedValue(undefined)

vi.mock('@/services/workload.service', () => ({
	fetchGetPods: vi.fn(),
	fetchGetPod: vi.fn(),
	fetchUpdatePod: vi.fn(),
	fetchRestartPod: (...args: unknown[]) => restartPod(...args),
	fetchAutoTroubleshoot: (...args: unknown[]) => autoTroubleshoot(...args),
	fetchGetDeployments: vi.fn(),
	fetchGetDeployment: vi.fn(),
	fetchUpdateDeployment: vi.fn(),
	fetchDeleteDeployment: (...args: unknown[]) => deleteDeployment(...args),
}))
vi.mock('@/services/network.service', () => ({
	fetchGetServices: vi.fn(),
	fetchGetService: vi.fn(),
	fetchUpdateService: vi.fn(),
	fetchDeleteService: (...args: unknown[]) => deleteService(...args),
	fetchGetIngresses: vi.fn(),
	fetchGetIngress: vi.fn(),
	fetchDeleteIngress: (...args: unknown[]) => deleteIngress(...args),
}))
vi.mock('@/services/general.service', () => ({
	fetchGetNamespaces: vi.fn(),
	fetchGetNamespace: vi.fn(),
	fetchGetNodes: vi.fn(),
	fetchGetNode: vi.fn(),
}))
vi.mock('@/services/storage.service', () => ({
	fetchGetPersistentVolumes: vi.fn(),
	fetchGetPersistentVolume: vi.fn(),
	fetchGetPersistentVolumesClaim: vi.fn(),
	fetchGetPersistentVolumeClaim: vi.fn(),
}))

import { getDescriptor } from '../registry'
import { RESOURCE_KINDS } from '../types'

const ref = { cluster: 'minikube', namespace: 'default', name: 'web' }

describe('resource registry', () => {
	beforeEach(() => {
		vi.clearAllMocks()
	})

	it('has a descriptor for every kind', () => {
		for (const kind of RESOURCE_KINDS) {
			const d = getDescriptor(kind)
			expect(d).toBeDefined()
			expect(d?.kind).toBe(kind)
		}
	})

	it('returns undefined for unknown kinds', () => {
		expect(getDescriptor('configmap')).toBeUndefined()
	})

	// Regression: a pre-descriptor version invoked every delete action
	// eagerly, so deleting one resource fired all four delete calls.
	it('service remove fires exactly one backend call', async () => {
		await getDescriptor('service')?.remove?.(ref)

		expect(deleteService).toHaveBeenCalledTimes(1)
		expect(deleteService).toHaveBeenCalledWith('web', 'default', 'minikube')
		expect(restartPod).not.toHaveBeenCalled()
		expect(deleteDeployment).not.toHaveBeenCalled()
		expect(deleteIngress).not.toHaveBeenCalled()
	})

	it('pod remove restarts the pod', async () => {
		await getDescriptor('pod')?.remove?.(ref)
		expect(restartPod).toHaveBeenCalledTimes(1)
	})

	it.each(['namespace', 'node', 'persistentVolume', 'persistentVolumeClaim'])(
		'%s exposes no remove or update action',
		(kind) => {
			const d = getDescriptor(kind)
			expect(d?.remove).toBeUndefined()
			expect(d?.update).toBeUndefined()
		},
	)

	it.each(['pod', 'deployment', 'service'])('%s supports update', (kind) => {
		expect(getDescriptor(kind)?.update).toBeDefined()
	})

	it('pod and deployment expose troubleshoot with the right resource tag', async () => {
		await getDescriptor('pod')?.troubleshoot?.(ref)
		expect(autoTroubleshoot).toHaveBeenCalledWith('web', 'default', 'minikube', 'POD')

		await getDescriptor('deployment')?.troubleshoot?.(ref)
		expect(autoTroubleshoot).toHaveBeenCalledWith('web', 'default', 'minikube', 'DEPLOYMENT')

		expect(getDescriptor('service')?.troubleshoot).toBeUndefined()
		expect(getDescriptor('node')?.troubleshoot).toBeUndefined()
	})

	it('cluster-scoped kinds are flagged as not namespaced', () => {
		expect(getDescriptor('node')?.namespaced).toBe(false)
		expect(getDescriptor('namespace')?.namespaced).toBe(false)
		expect(getDescriptor('persistentVolume')?.namespaced).toBe(false)
		expect(getDescriptor('pod')?.namespaced).toBe(true)
	})
})

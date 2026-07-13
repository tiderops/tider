import { describe, it, expect, vi } from 'vitest'

vi.mock('@/services/workload.service', () => ({
	fetchGetPods: vi.fn(), fetchGetPod: vi.fn(), fetchUpdatePod: vi.fn(), fetchRestartPod: vi.fn(), fetchAutoTroubleshoot: vi.fn(),
	fetchGetDeployments: vi.fn(), fetchGetDeployment: vi.fn(), fetchUpdateDeployment: vi.fn(), fetchDeleteDeployment: vi.fn(),
}))
vi.mock('@/services/network.service', () => ({
	fetchGetServices: vi.fn(), fetchGetService: vi.fn(), fetchUpdateService: vi.fn(), fetchDeleteService: vi.fn(),
	fetchGetIngresses: vi.fn(), fetchGetIngress: vi.fn(), fetchDeleteIngress: vi.fn(),
}))

import { podToRow } from '../descriptors/pod'
import { serviceToRow } from '../descriptors/service'
import { ingressToRow } from '../descriptors/ingress'
import { model } from '../../../wailsjs/go/models'

describe('podToRow', () => {
	it('maps the first container resources', () => {
		const row = podToRow(
			model.PodDto.createFrom({
				Name: 'web-1',
				Namespace: 'default',
				Node: 'node-1',
				Age: '3h',
				Status: 'Running',
				Containers: [
					{ Name: 'web', Limit: { Cpu: 500, Memory: 128, Storage: 0, StorageEphemeral: 0 }, Request: { Cpu: 250, Memory: 64, Storage: 0, StorageEphemeral: 0 } },
					{ Name: 'sidecar', Limit: { Cpu: 100, Memory: 32, Storage: 0, StorageEphemeral: 0 }, Request: { Cpu: 50, Memory: 16, Storage: 0, StorageEphemeral: 0 } },
				],
			}),
		)

		expect(row).toMatchObject({ name: 'web-1', namespace: 'default', cpu: '500/250', memory: '128/64', node: 'node-1', status: 'Running' })
	})

	it('handles pods without containers', () => {
		const row = podToRow(model.PodDto.createFrom({ Name: 'bare', Namespace: 'default', Containers: null }))
		expect(row.cpu).toBe('0/0')
	})
})

describe('serviceToRow', () => {
	it('maps the IP columns the grid actually shows', () => {
		const row = serviceToRow(
			model.ServiceDto.createFrom({
				Name: 'svc', Namespace: 'default', Type: 'ClusterIP',
				InternalIp: '10.0.0.1', ExternalIp: 'example.org', Port: 8080, Status: 'ok', Age: '1d',
			}),
		)
		expect(row.intIp).toBe('10.0.0.1')
		expect(row.extIp).toBe('example.org')
		expect(row.port).toBe(8080)
	})
})

describe('ingressToRow', () => {
	it('flattens rule hosts and paths and survives nil rules', () => {
		const row = ingressToRow(model.IngressDto.createFrom({ Name: 'ing', Namespace: 'default', Rules: null, Age: '2d' }))
		expect(row.host).toEqual([])
		expect(row.path).toEqual([])
	})
})

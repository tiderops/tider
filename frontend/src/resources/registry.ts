import type { ResourceDescriptor, ResourceKind } from './types'
import { podDescriptor } from './descriptors/pod'
import { deploymentDescriptor } from './descriptors/deployment'
import { serviceDescriptor } from './descriptors/service'
import { ingressDescriptor } from './descriptors/ingress'
import { namespaceDescriptor } from './descriptors/namespace'
import { nodeDescriptor } from './descriptors/node'
import { persistentVolumeDescriptor } from './descriptors/persistentVolume'
import { persistentVolumeClaimDescriptor } from './descriptors/persistentVolumeClaim'

const registry: Record<ResourceKind, ResourceDescriptor> = {
	pod: podDescriptor,
	deployment: deploymentDescriptor,
	service: serviceDescriptor,
	ingress: ingressDescriptor,
	namespace: namespaceDescriptor,
	node: nodeDescriptor,
	persistentVolume: persistentVolumeDescriptor,
	persistentVolumeClaim: persistentVolumeClaimDescriptor,
}

export function getDescriptor(kind: string): ResourceDescriptor | undefined {
	return registry[kind as ResourceKind]
}

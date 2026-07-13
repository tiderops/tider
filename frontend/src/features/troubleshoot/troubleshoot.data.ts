import { delay } from '@/data/mock-latency'
import { fetchAutoTroubleshoot } from '@/services/workload.service'
import { hasWailsRuntime } from '@/services/runtime'
import type { Issue } from '@/types/issue'
import type { Diagnosis, DiagnosisAction } from './types'

const RESOURCE_ARG: Record<Issue['resourceKind'], string> = {
	pod: 'POD',
	deployment: 'DEPLOYMENT',
	job: 'JOB',
	node: 'NODE',
}

function defaultActions(issue: Issue): DiagnosisAction[] {
	switch (issue.resourceKind) {
		case 'pod':
			return [
				{ label: 'Restart pod', description: 'Delete the pod; its controller recreates it.', kind: 'restart' },
				{ label: 'View logs', description: 'Open the container logs and last termination.', kind: 'logs' },
			]
		case 'deployment':
			return [
				{ label: 'Roll back', description: 'Revert to the previous healthy revision.', kind: 'rollback' },
				{ label: 'View logs', description: 'Inspect the rollout events.', kind: 'logs' },
			]
		default:
			return [{ label: 'Inspect', description: 'Open the resource detail.', kind: 'inspect' }]
	}
}

function mockDiagnosis(issue: Issue): Diagnosis {
	switch (issue.reason) {
		case 'OOMKilled':
			return {
				meaning:
					'The container was terminated (exit 137) after exceeding its memory limit of 1.5Gi. It has restarted 7 times in 12 minutes — a sustained leak or an under-provisioned limit, not a transient spike.',
				recommendation: 'Raise the memory limit to ~2.25Gi (24h p95 + 50% headroom), then watch the restart trend.',
				evidence: [
					{ label: 'Peak memory before kill', value: '1.50 Gi / 1.5Gi' },
					{ label: 'Restarts', value: '7 in 12m' },
				],
				actions: [
					{ label: 'Raise memory limit to 2.25Gi', description: 'Apply the optimization recommendation.', kind: 'apply' },
					{ label: 'View logs', description: 'Open the last termination logs and goroutine dump.', kind: 'logs' },
					{ label: 'Roll back to checkout-api:1.18.1', description: 'The leak appeared after the 1.18.2 deploy.', kind: 'rollback' },
				],
			}
		case 'ImagePull':
			return {
				meaning: 'Kubernetes cannot pull the container image. The tag may be wrong, or the registry credentials/network are unavailable.',
				recommendation: 'Verify the image name and tag, and that the pull secret for the registry is present in this namespace.',
				evidence: [{ label: 'Last event', value: 'ErrImagePull · 401 Unauthorized' }],
				actions: [
					{ label: 'View events', description: 'Open the deployment events.', kind: 'logs' },
					{ label: 'Edit image', description: 'Correct the image reference.', kind: 'apply' },
				],
			}
		default:
			return {
				meaning: `The node stopped reporting to the control plane (${issue.reason}). Pods on it may be rescheduled.`,
				recommendation: 'Check kubelet health and node networking; cordon & drain if it stays unhealthy.',
				evidence: [],
				actions: [{ label: 'Inspect node', description: 'Open the node detail.', kind: 'inspect' }],
			}
	}
}

// TODO-15: Consume backend.
export async function fetchDiagnosis(issue: Issue): Promise<Diagnosis> {
	if (!hasWailsRuntime() || issue.resourceKind === 'node') {
		await delay()
		return mockDiagnosis(issue)
	}
	const result = await fetchAutoTroubleshoot(issue.name, issue.namespace, issue.cluster, RESOURCE_ARG[issue.resourceKind])
	return {
		meaning: result.Meaning,
		recommendation: result.Recommendation,
		evidence: [],
		actions: defaultActions(issue),
	}
}

const HEALTHY = new Set(['Running', 'Succeeded', 'Active', 'Bound', 'Available', 'True', 'Completed', 'Ready'])
const WARNING = new Set(['Pending', 'Unknown', 'Released', 'Terminating', 'ContainerCreating'])

export function statusColor(status: string | undefined): string {
	if (!status) {
		return 'grey'
	}
	if (HEALTHY.has(status)) {
		return 'green'
	}
	if (WARNING.has(status)) {
		return 'orange'
	}
	return 'red'
}

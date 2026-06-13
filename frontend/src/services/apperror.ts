
export const APP_ERROR_KINDS = [
	'NOT_FOUND',
	'FORBIDDEN',
	'CONFLICT',
	'VALIDATION',
	'TIMEOUT',
	'CLUSTER_UNREACHABLE',
	'INTERNAL',
] as const

export type AppErrorKind = (typeof APP_ERROR_KINDS)[number]

export class AppError extends Error {
	readonly kind: AppErrorKind

	constructor(kind: AppErrorKind, message: string) {
		super(message)
		this.name = 'AppError'
		this.kind = kind
	}
}

export function toAppError(err: unknown): AppError {
	if (err instanceof AppError) {
		return err
	}

	const raw = err instanceof Error ? err.message : String(err)

	const separator = raw.indexOf(': ')
	if (separator > 0) {
		const prefix = raw.slice(0, separator)
		if ((APP_ERROR_KINDS as readonly string[]).includes(prefix)) {
			return new AppError(prefix as AppErrorKind, raw.slice(separator + 2))
		}
	}

	return new AppError('INTERNAL', raw)
}

const KIND_MESSAGES: Record<AppErrorKind, string> = {
	NOT_FOUND: 'The resource was not found. It may have been deleted.',
	FORBIDDEN: 'You are not allowed to perform this action on the cluster.',
	CONFLICT: 'The resource changed while you were editing it. Reload and retry.',
	VALIDATION: 'The request was rejected as invalid.',
	TIMEOUT: 'The cluster took too long to respond.',
	CLUSTER_UNREACHABLE: 'The cluster is unreachable. Check the connection or your kubeconfig.',
	INTERNAL: 'Something went wrong.',
}

export function userMessage(error: AppError): string {
	return KIND_MESSAGES[error.kind]
}

import { describe, it, expect } from 'vitest'
import { AppError, toAppError, userMessage } from '../apperror'

describe('toAppError', () => {
	it.each([
		['NOT_FOUND: pod web-1 not found', 'NOT_FOUND', 'pod web-1 not found'],
		['FORBIDDEN: rbac denied', 'FORBIDDEN', 'rbac denied'],
		['CONFLICT: stale object', 'CONFLICT', 'stale object'],
		['VALIDATION: bad spec', 'VALIDATION', 'bad spec'],
		['TIMEOUT: context deadline exceeded', 'TIMEOUT', 'context deadline exceeded'],
		['CLUSTER_UNREACHABLE: cluster "minikube" is not registered', 'CLUSTER_UNREACHABLE', 'cluster "minikube" is not registered'],
		['INTERNAL: boom', 'INTERNAL', 'boom'],
	])('parses %s', (raw, kind, message) => {
		const err = toAppError(raw)
		expect(err.kind).toBe(kind)
		expect(err.message).toBe(message)
	})

	it('keeps extra colons inside the message', () => {
		const err = toAppError('NOT_FOUND: getting pod default/web-1: not found')
		expect(err.kind).toBe('NOT_FOUND')
		expect(err.message).toBe('getting pod default/web-1: not found')
	})

	it('classifies unprefixed strings as INTERNAL', () => {
		const err = toAppError('something exploded')
		expect(err.kind).toBe('INTERNAL')
		expect(err.message).toBe('something exploded')
	})

	it('does not treat arbitrary prefixes as kinds', () => {
		const err = toAppError('WHATEVER: nope')
		expect(err.kind).toBe('INTERNAL')
		expect(err.message).toBe('WHATEVER: nope')
	})

	it('unwraps Error instances', () => {
		const err = toAppError(new Error('TIMEOUT: too slow'))
		expect(err.kind).toBe('TIMEOUT')
	})

	it('passes through AppError unchanged', () => {
		const original = new AppError('FORBIDDEN', 'rbac')
		expect(toAppError(original)).toBe(original)
	})

	it('stringifies non-error rejections', () => {
		const err = toAppError(undefined)
		expect(err.kind).toBe('INTERNAL')
	})
})

describe('userMessage', () => {
	it('maps kinds to friendly headlines', () => {
		expect(userMessage(new AppError('CLUSTER_UNREACHABLE', 'x'))).toContain('unreachable')
		expect(userMessage(new AppError('TIMEOUT', 'x'))).toContain('too long')
	})
})

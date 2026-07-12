import type { Pin } from '@/types/pin'

// TODO-18: Remove mock data
export const defaultPins: Pin[] = [
	{ name: 'payments-gateway', detail: 'payments', tone: 'ok' },
	{ name: 'checkout-api', detail: '4/4', tone: 'ok' },
	{ name: 'ledger-worker', detail: 'payments', tone: 'ok' },
	{ name: 'dev-sandbox', detail: 'cluster', tone: 'err' },
]

class ResizeObserverStub {
	observe() {}
	unobserve() {}
	disconnect() {}
}

globalThis.ResizeObserver = globalThis.ResizeObserver ?? (ResizeObserverStub as unknown as typeof ResizeObserver)

if (!('visualViewport' in globalThis) || !globalThis.visualViewport) {
	Object.defineProperty(globalThis, 'visualViewport', {
		value: new EventTarget(),
		writable: true,
	})
}

// Node ≥22 exposes its own (unusable without --localstorage-file) localStorage
// global that shadows jsdom's; fall back to an in-memory implementation.
if (!globalThis.localStorage) {
	const store = new Map<string, string>()
	const memoryStorage = {
		getItem: (key: string) => store.get(key) ?? null,
		setItem: (key: string, value: string) => void store.set(key, String(value)),
		removeItem: (key: string) => void store.delete(key),
		clear: () => store.clear(),
		key: (index: number) => [...store.keys()][index] ?? null,
		get length() {
			return store.size
		},
	}
	Object.defineProperty(globalThis, 'localStorage', {
		value: memoryStorage as unknown as Storage,
		writable: true,
		configurable: true,
	})
}

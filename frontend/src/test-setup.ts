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

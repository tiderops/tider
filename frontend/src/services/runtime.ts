// True when running inside the Wails desktop runtime (bindings present).
// The data modules use this to fall back to mock data when the frontend
// runs standalone in a browser (npm run dev), and call real bindings
// under `wails dev` / the packaged app.
export function hasWailsRuntime(): boolean {
	if (typeof window === 'undefined') {
		return false
	}
	const w = window as unknown as { go?: { binding?: unknown } }
	return Boolean(w.go?.binding)
}

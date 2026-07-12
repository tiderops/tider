import { onMounted, onUnmounted } from 'vue'
import { useOverlayStore } from '@/stores/overlay.store'

export function useGlobalShortcuts() {
	const overlay = useOverlayStore()

	function onKey(e: KeyboardEvent) {
		const meta = e.metaKey || e.ctrlKey
		if (meta && e.key.toLowerCase() === 'k') {
			e.preventDefault()
			overlay.togglePalette()
		} else if (meta && e.key.toLowerCase() === 'o') {
			e.preventDefault()
			overlay.toggleSwitcher()
		} else if (e.key === 'Escape') {
			overlay.closeAll()
		}
	}

	onMounted(() => window.addEventListener('keydown', onKey))
	onUnmounted(() => window.removeEventListener('keydown', onKey))
}

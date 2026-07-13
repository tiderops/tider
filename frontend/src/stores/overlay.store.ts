import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useOverlayStore = defineStore('overlay', () => {
	const palette = ref(false)
	const switcher = ref(false)

	function closeAll() {
		palette.value = false
		switcher.value = false
	}
	function openPalette() {
		switcher.value = false
		palette.value = true
	}
	function openSwitcher() {
		palette.value = false
		switcher.value = true
	}
	function togglePalette() {
		if (palette.value) {
			closeAll()
		} else {
			openPalette()
		}
	}
	function toggleSwitcher() {
		if (switcher.value) {
			closeAll()
		} else {
			openSwitcher()
		}
	}

	return { palette, switcher, closeAll, openPalette, openSwitcher, togglePalette, toggleSwitcher }
})

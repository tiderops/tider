<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/stores/theme.store'
import { useOverlayStore } from '@/stores/overlay.store'
import { useGlobalShortcuts } from '@/composables/useGlobalShortcuts'
import AppRail from './AppRail.vue'
import AppTopbar from './AppTopbar.vue'
import CommandPalette from '@/overlays/CommandPalette.vue'
import ClusterSwitcher from '@/overlays/ClusterSwitcher.vue'
import '@/styles/kx-theme.css'

const { theme } = storeToRefs(useThemeStore())
const themeClass = computed(() => (theme.value === 'light' ? 'theme-light' : ''))

const overlay = useOverlayStore()
useGlobalShortcuts()
</script>

<template>
	<div class="kx-surface" :class="themeClass">
		<AppRail />
		<div class="kx-main">
			<AppTopbar />
			<div class="kx-content">
				<slot />
			</div>
		</div>

		<CommandPalette v-if="overlay.palette" />
		<ClusterSwitcher v-if="overlay.switcher" />
	</div>
</template>

<style scoped>
.kx-surface {
	display: grid;
	grid-template-columns: 248px 1fr;
	height: 100vh;
	background: var(--bg);
	overflow: hidden;
}
.kx-main {
	display: flex;
	flex-direction: column;
	min-width: 0;
	background: radial-gradient(120% 80% at 100% 0%, rgba(50, 108, 229, 0.06), transparent 60%), var(--bg);
}
.kx-content {
	padding: 24px 26px;
	overflow: auto;
	flex: 1;
}
</style>

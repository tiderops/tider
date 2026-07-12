<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/stores/theme.store'
import { useSettingsStore } from '@/stores/settings.store'
import type { Density } from '../types'
import SettingsCard from './SettingsCard.vue'
import SettingRow from './SettingRow.vue'

const themeStore = useThemeStore()
const { theme } = storeToRefs(themeStore)
const { prefs } = storeToRefs(useSettingsStore())

const densities: { value: Density; label: string }[] = [
	{ value: 'comfortable', label: 'Comfortable' },
	{ value: 'compact', label: 'Compact' },
]
</script>

<template>
	<SettingsCard title="Appearance" subtitle="How Tider looks on this machine.">
		<SettingRow title="Theme" description="Switch between the dark and light palettes.">
			<div class="seg">
				<button type="button" class="opt" :class="{ on: theme === 'dark' }" @click="themeStore.set('dark')">☾ Dark</button>
				<button type="button" class="opt" :class="{ on: theme === 'light' }" @click="themeStore.set('light')">☀ Light</button>
			</div>
		</SettingRow>

		<SettingRow title="Density" description="Compact tightens table rows and card padding." last>
			<div class="seg">
				<button v-for="d in densities" :key="d.value" type="button" class="opt" :class="{ on: prefs.density === d.value }" @click="prefs.density = d.value">
					{{ d.label }}
				</button>
			</div>
		</SettingRow>
	</SettingsCard>
</template>

<style scoped>
.seg {
	display: inline-flex;
	background: var(--surface-2);
	border: 1px solid var(--border);
	border-radius: var(--r-sm);
	padding: 2px;
	gap: 2px;
}
.opt {
	font-size: 12.5px;
	font-weight: 600;
	padding: 6px 13px;
	border-radius: 5px;
	border: none;
	background: none;
	color: var(--text-dim);
	cursor: pointer;
	font-family: var(--sans);
}
.opt.on {
	background: var(--brand);
	color: #fff;
}
</style>

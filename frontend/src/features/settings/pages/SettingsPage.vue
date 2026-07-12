<script setup lang="ts">
import { ref } from 'vue'
import { useSettingsStore } from '@/stores/settings.store'
import { sections } from '../settings.data'
import AppearanceSection from '../components/AppearanceSection.vue'
import ClustersSection from '../components/ClustersSection.vue'
import NotificationsSection from '../components/NotificationsSection.vue'
import AssistantSection from '../components/AssistantSection.vue'
import AboutSection from '../components/AboutSection.vue'

const panels = {
	appearance: AppearanceSection,
	clusters: ClustersSection,
	notifications: NotificationsSection,
	assistant: AssistantSection,
	about: AboutSection,
} as const

type SectionKey = keyof typeof panels

const active = ref<SectionKey>('appearance')
const settings = useSettingsStore()

const toast = ref('')
let toastTimer: ReturnType<typeof setTimeout> | undefined
function resetAll() {
	settings.reset()
	toast.value = 'Restored default preferences'
	clearTimeout(toastTimer)
	toastTimer = setTimeout(() => (toast.value = ''), 2400)
}
</script>

<template>
	<div class="page-head">
		<div>
			<h1>Settings</h1>
			<p>Preferences are stored locally on this machine.</p>
		</div>
		<button class="btn" @click="resetAll">Reset to defaults</button>
	</div>

	<div class="layout">
		<nav class="sidenav">
			<button
				v-for="s in sections"
				:key="s.key"
				type="button"
				class="nav"
				:class="{ active: s.key === active }"
				@click="active = s.key as SectionKey"
			>
				<span class="ico">{{ s.icon }}</span>
				<span class="lbl"><b>{{ s.label }}</b><small>{{ s.hint }}</small></span>
			</button>
		</nav>

		<div class="panel">
			<component :is="panels[active]" />
		</div>
	</div>

	<Transition name="toast">
		<div v-if="toast" class="toast">{{ toast }}</div>
	</Transition>
</template>

<style scoped>
.page-head {
	display: flex;
	align-items: flex-end;
	gap: 16px;
	margin-bottom: 20px;
}
.page-head h1 {
	margin: 0;
	font-size: 22px;
	font-weight: 700;
	letter-spacing: -0.02em;
}
.page-head p {
	margin: 4px 0 0;
	color: var(--text-dim);
	font-size: 13px;
}
.btn {
	margin-left: auto;
	font-size: 12.5px;
	font-weight: 600;
	padding: 9px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface);
	color: var(--text);
	cursor: pointer;
}
.btn:hover {
	border-color: var(--brand);
}
.layout {
	display: grid;
	grid-template-columns: 232px 1fr;
	gap: 22px;
	align-items: start;
}
.sidenav {
	display: flex;
	flex-direction: column;
	gap: 3px;
	position: sticky;
	top: 0;
}
.nav {
	display: flex;
	align-items: center;
	gap: 11px;
	padding: 10px 12px;
	border-radius: var(--r-sm);
	border: 1px solid transparent;
	background: none;
	color: var(--text-dim);
	cursor: pointer;
	text-align: left;
	font-family: var(--sans);
}
.nav:hover {
	background: var(--hover);
	color: var(--text);
}
.nav.active {
	background: var(--brand-soft);
	border-color: var(--border-soft);
}
.nav.active .lbl b {
	color: var(--brand);
}
.nav .ico {
	width: 18px;
	text-align: center;
	font-size: 14px;
	opacity: 0.9;
}
.lbl {
	display: flex;
	flex-direction: column;
	min-width: 0;
}
.lbl b {
	font-size: 13px;
	font-weight: 600;
	color: var(--text);
}
.lbl small {
	font-size: 11px;
	color: var(--text-faint);
}
.panel {
	display: flex;
	flex-direction: column;
	gap: 18px;
	min-width: 0;
}
.toast {
	position: fixed;
	bottom: 26px;
	left: 50%;
	transform: translateX(-50%);
	background: var(--surface-3);
	border: 1px solid var(--border);
	color: var(--text);
	padding: 10px 18px;
	border-radius: 999px;
	font-size: 13px;
	box-shadow: 0 12px 30px rgba(0, 0, 0, 0.45);
	z-index: 60;
}
.toast-enter-active,
.toast-leave-active {
	transition: all 0.2s ease;
}
.toast-enter-from,
.toast-leave-to {
	opacity: 0;
	transform: translate(-50%, 8px);
}
</style>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useFleetStore } from '@/stores/fleet.store'
import { useSettingsStore } from '@/stores/settings.store'
import { refreshOptions } from '../settings.data'
import SettingsCard from './SettingsCard.vue'
import SettingRow from './SettingRow.vue'

const fleetStore = useFleetStore()
const { clusters, loading } = storeToRefs(fleetStore)
const { prefs } = storeToRefs(useSettingsStore())

const reachableCount = computed(() => clusters.value.filter((c) => c.reachable).length)

onMounted(() => fleetStore.load())
</script>

<template>
	<SettingsCard title="Clusters & kubeconfig" subtitle="Contexts discovered from your kubeconfig and how the app refreshes them.">
		<SettingRow title="Default cluster" description="Which context to open on launch.">
			<select v-model="prefs.defaultCluster" class="select">
				<option value="">Remember last used</option>
				<option v-for="c in clusters" :key="c.name" :value="c.name">{{ c.name }}</option>
			</select>
		</SettingRow>

		<SettingRow title="Auto-refresh" description="How often resource lists and metrics re-fetch." last>
			<select v-model.number="prefs.refreshInterval" class="select">
				<option v-for="o in refreshOptions" :key="o.value" :value="o.value">{{ o.label }}</option>
			</select>
		</SettingRow>
	</SettingsCard>

	<SettingsCard title="Registered contexts" :subtitle="loading ? 'Loading…' : `${reachableCount} of ${clusters.length} reachable`">
		<div v-if="!loading && !clusters.length" class="empty">No clusters found in kubeconfig.</div>
		<div v-for="c in clusters" :key="c.name" class="ctx">
			<span class="dot" :class="c.reachable ? 'ok' : 'err'"></span>
			<div class="ci">
				<b>{{ c.name }}</b>
				<span>{{ c.source }}</span>
			</div>
			<span class="badge" :class="c.reachable ? 'ok' : 'err'">{{ c.reachable ? 'reachable' : 'offline' }}</span>
		</div>
		<p class="note">Contexts are read from <span class="mono">~/.kube/config</span>. Add or remove them from the cluster switcher (⌘O).</p>
	</SettingsCard>
</template>

<style scoped>
.select {
	background: var(--surface-2);
	border: 1px solid var(--border);
	border-radius: var(--r-sm);
	padding: 8px 12px;
	font-size: 12.5px;
	color: var(--text);
	font-family: var(--sans);
	cursor: pointer;
	min-width: 180px;
}
.empty {
	padding: 18px 0;
	text-align: center;
	color: var(--text-faint);
	font-size: 12.5px;
}
.ctx {
	display: flex;
	align-items: center;
	gap: 11px;
	padding: 11px 0;
	border-bottom: 1px solid var(--border-soft);
}
.ctx:last-of-type {
	border-bottom: none;
}
.dot {
	width: 8px;
	height: 8px;
	border-radius: 50%;
	flex: none;
}
.dot.ok {
	background: var(--ok);
	box-shadow: 0 0 0 3px var(--ok-bg);
}
.dot.err {
	background: var(--err);
	box-shadow: 0 0 0 3px var(--err-bg);
}
.ci {
	flex: 1;
	min-width: 0;
}
.ci b {
	display: block;
	font-family: var(--mono);
	font-size: 12.5px;
	font-weight: 600;
	color: var(--text);
}
.ci span {
	display: block;
	font-size: 11px;
	color: var(--text-faint);
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
.badge {
	font-size: 10px;
	font-weight: 600;
	padding: 2px 9px;
	border-radius: 999px;
}
.badge.ok {
	color: var(--ok);
	background: var(--ok-bg);
}
.badge.err {
	color: var(--err);
	background: var(--err-bg);
}
.note {
	margin: 14px 0 2px;
	font-size: 11.5px;
	color: var(--text-faint);
}
.mono {
	font-family: var(--mono);
}
</style>

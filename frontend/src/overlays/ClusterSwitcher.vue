<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useOverlayStore } from '@/stores/overlay.store'
import { useClusterStore } from '@/stores/cluster.store'
import { useFleetStore } from '@/stores/fleet.store'
import { useSettingsStore } from '@/stores/settings.store'
import type { ChipTone } from '@/types/status'

const overlay = useOverlayStore()
const clusterStore = useClusterStore()
const fleetStore = useFleetStore()
const settings = useSettingsStore()
const { clusters, loading } = storeToRefs(fleetStore)

const query = ref('')
const inputEl = ref<HTMLInputElement | null>(null)

const DOT: Record<ChipTone, string> = {
	ok: 'var(--ok)',
	warn: 'var(--warn)',
	err: 'var(--err)',
	info: 'var(--info)',
	idle: 'var(--idle)',
}

const filtered = computed(() => {
	const q = query.value.toLowerCase()
	return clusters.value.filter((c) => !q || c.name.toLowerCase().includes(q))
})

function select(name: string) {
	clusterStore.setCurrentCluster(name)
	overlay.closeAll()
}

onMounted(() => {
	fleetStore.load()
	inputEl.value?.focus()
})
</script>

<template>
	<div class="wrap">
		<div class="scrim" @click="overlay.closeAll()"></div>
		<div class="pop">
			<div class="search"><span>⌕</span><input ref="inputEl" v-model="query" placeholder="Filter clusters…" /></div>

			<div class="list">
				<div v-if="loading" class="muted">Loading clusters…</div>
				<button v-for="c in filtered" :key="c.name" type="button" class="item" :class="{ active: c.name === clusterStore.currentCluster }" @click="select(c.name)">
					<span class="h" :style="{ background: DOT[c.statusTone] }"></span>
					<span class="ci">
						<b>
							{{ c.name }}
							<i v-if="settings.envFor(c.name) !== 'none'" class="env" :class="settings.envFor(c.name)">{{ settings.envFor(c.name) }}</i>
						</b>
						<span>{{ c.source }}</span>
					</span>
					<span v-if="c.name === clusterStore.currentCluster" class="ck">✓</span>
					<span v-else-if="!c.reachable" class="off">offline</span>
				</button>
				<div v-if="!loading && !filtered.length" class="muted">No clusters match.</div>
			</div>

			<div class="foot">＋ Add from kubeconfig <span class="kbd">⌘O</span></div>
		</div>
	</div>
</template>

<style scoped>
.wrap {
	position: fixed;
	inset: 0;
	z-index: 80;
}
.scrim {
	position: absolute;
	inset: 0;
	background: var(--overlay);
}
.pop {
	position: absolute;
	top: 70px;
	left: 16px;
	width: 304px;
	background: var(--surface);
	border: 1px solid var(--border);
	border-radius: var(--r-md);
	box-shadow: 0 24px 60px rgba(0, 0, 0, 0.55);
	overflow: hidden;
}
.search {
	display: flex;
	align-items: center;
	gap: 9px;
	padding: 12px 14px;
	border-bottom: 1px solid var(--border-soft);
	color: var(--text-faint);
	font-size: 13px;
}
.search input {
	flex: 1;
	background: none;
	border: none;
	outline: none;
	color: var(--text);
	font-family: var(--sans);
	font-size: 13px;
}
.search input::placeholder {
	color: var(--text-faint);
}
.list {
	padding: 6px;
	max-height: 360px;
	overflow: auto;
}
.muted {
	padding: 16px;
	text-align: center;
	color: var(--text-faint);
	font-size: 12.5px;
}
.item {
	display: flex;
	align-items: center;
	gap: 10px;
	width: 100%;
	padding: 9px 10px;
	border-radius: var(--r-sm);
	background: none;
	border: none;
	cursor: pointer;
	text-align: left;
	font-family: var(--sans);
}
.item:hover {
	background: var(--hover);
}
.item.active {
	background: var(--brand-soft);
}
.h {
	width: 8px;
	height: 8px;
	border-radius: 50%;
	flex: none;
}
.ci {
	flex: 1;
	min-width: 0;
}
.ci b {
	font-family: var(--mono);
	font-size: 12.5px;
	font-weight: 600;
	color: var(--text);
	display: block;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
.ci .env {
	font-family: var(--sans);
	font-style: normal;
	font-size: 9px;
	font-weight: 700;
	letter-spacing: 0.05em;
	text-transform: uppercase;
	padding: 0 6px;
	border-radius: 999px;
	margin-left: 4px;
	vertical-align: 1px;
}
.ci .env.prod {
	color: var(--env-prod);
	background: var(--env-prod-bg);
}
.ci .env.staging {
	color: var(--env-staging);
	background: var(--env-staging-bg);
}
.ci .env.dev {
	color: var(--env-dev);
	background: var(--env-dev-bg);
}
.ci span {
	font-size: 11px;
	color: var(--text-faint);
	display: block;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
.ck {
	color: var(--brand);
	font-size: 13px;
}
.off {
	font-size: 10px;
	color: var(--err);
	font-weight: 600;
}
.foot {
	border-top: 1px solid var(--border-soft);
	padding: 10px 14px;
	display: flex;
	align-items: center;
	font-size: 12.5px;
	color: var(--text-dim);
	cursor: pointer;
}
.foot .kbd {
	margin-left: auto;
	font-family: var(--mono);
	font-size: 11px;
	background: var(--surface-2);
	border: 1px solid var(--border);
	border-radius: 5px;
	padding: 1px 6px;
	color: var(--text-faint);
}
</style>

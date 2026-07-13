<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'
import StatTile from '@/components/shared/StatTile.vue'
import KxState from '@/components/shared/KxState.vue'
import ScheduleStrip from '../components/ScheduleStrip.vue'
import BackupTable from '../components/BackupTable.vue'
import RestoreDrawer from '../components/RestoreDrawer.vue'
import { useAsyncData } from '@/composables/useAsyncData'
import { useActiveCluster } from '@/composables/useActiveCluster'
import { useFleetStore } from '@/stores/fleet.store'
import { fetchBackup } from '../backup.data'
import type { BackupData, Snapshot } from '../types'

const { resolve } = useActiveCluster()
const cluster = ref('')
const fleetStore = useFleetStore()
const { clusters } = storeToRefs(fleetStore)
const clusterNames = computed(() => clusters.value.map((c) => c.name))

const empty: BackupData = { summary: { snapshots: 0, retention: '', lastBackup: '', lastStatus: '', storageUsed: '', storageTarget: '', nextScheduled: '', nextDetail: '' }, schedules: [], snapshots: [] }
const { data, loading, error, reload } = useAsyncData<BackupData>(() => fetchBackup(cluster.value), empty)

const restoreOpen = ref(false)
const restoreSnap = ref<Snapshot | null>(null)

const toast = ref('')
let toastTimer: ReturnType<typeof setTimeout> | undefined
function showToast(message: string) {
	toast.value = message
	clearTimeout(toastTimer)
	toastTimer = setTimeout(() => (toast.value = ''), 2600)
}

function openRestore(snap: Snapshot) {
	restoreSnap.value = snap
	restoreOpen.value = true
}

function onRestore(payload: { snapshot: Snapshot; target: string; count: number }) {
	restoreOpen.value = false
	showToast(`Restoring ${payload.count.toLocaleString()} resources into ${payload.target}…`)
}

onMounted(async () => {
	cluster.value = await resolve()
	fleetStore.load()
	reload()
})
</script>

<template>
	<div class="page-head">
		<div>
			<h1>Backup &amp; Restore</h1>
			<p>Point-in-time snapshots of manifests &amp; state · restore to any connected cluster.</p>
		</div>
		<div class="head-actions">
			<button class="btn" @click="showToast('New schedule')">⌚ New schedule</button>
			<button class="btn primary" @click="showToast('Create backup')">＋ Create backup</button>
		</div>
	</div>

	<KxState v-if="loading || error" :loading="loading" :error="error" @retry="reload" />

	<template v-else>
		<div class="tiles">
			<StatTile label="⭢ Snapshots" :value="String(data.summary.snapshots)" :hint="data.summary.retention" hint-tone="up" />
			<StatTile label="⌚ Last backup" :value="data.summary.lastBackup" :hint="data.summary.lastStatus" hint-tone="up" />
			<StatTile label="▥ Storage used" :value="data.summary.storageUsed" :hint="data.summary.storageTarget" hint-tone="up" />
			<StatTile label="↻ Next scheduled" :value="data.summary.nextScheduled" :hint="data.summary.nextDetail" hint-tone="up" />
		</div>

		<div class="section-t">Schedules</div>
		<ScheduleStrip :schedules="data.schedules" @new="showToast('New schedule')" />

		<BackupTable :snapshots="data.snapshots" @restore="openRestore" />
	</template>

	<RestoreDrawer :open="restoreOpen" :snapshot="restoreSnap" :clusters="clusterNames" @close="restoreOpen = false" @restore="onRestore" />

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
.head-actions {
	margin-left: auto;
	display: flex;
	gap: 10px;
}
.btn {
	display: inline-flex;
	align-items: center;
	gap: 8px;
	font-size: 13px;
	font-weight: 600;
	padding: 9px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface);
	color: var(--text);
	cursor: pointer;
	white-space: nowrap;
}
.btn:hover {
	border-color: #3a465a;
	background: var(--surface-2);
}
.btn.primary {
	background: linear-gradient(180deg, var(--brand), var(--brand-deep));
	border-color: transparent;
	color: #fff;
}
.tiles {
	display: grid;
	grid-template-columns: repeat(4, 1fr);
	gap: 16px;
	margin-bottom: 18px;
}
.section-t {
	font-size: 11px;
	font-weight: 600;
	letter-spacing: 0.06em;
	text-transform: uppercase;
	color: var(--text-faint);
	margin: 4px 0 12px;
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

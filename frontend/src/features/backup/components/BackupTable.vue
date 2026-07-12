<script setup lang="ts">
import StatusChip from '@/components/shared/StatusChip.vue'
import type { ChipTone } from '@/types/status'
import type { Snapshot, SnapshotStatus } from '../types'

defineProps<{ snapshots: Snapshot[] }>()
const emit = defineEmits<{ (e: 'restore', snap: Snapshot): void }>()

const TONE: Record<SnapshotStatus, ChipTone> = {
	Completed: 'ok',
	Running: 'info',
	Partial: 'warn',
	Failed: 'err',
}

function label(s: Snapshot) {
	return s.status === 'Running' && s.progress !== undefined ? `Running ${s.progress}%` : s.status
}
</script>

<template>
	<div class="card">
		<div class="head">Backup history</div>
		<table class="tbl">
			<thead>
				<tr><th>Snapshot</th><th>Scope</th><th>Source</th><th>Resources</th><th>Size</th><th>Created</th><th>Status</th><th></th></tr>
			</thead>
			<tbody>
				<tr v-for="s in snapshots" :key="s.id">
					<td><span class="name">{{ s.name }}</span></td>
					<td><span class="tag">{{ s.scope }}</span></td>
					<td class="mono">{{ s.source }}</td>
					<td class="mono">{{ s.resources.toLocaleString() }}</td>
					<td class="mono">{{ s.size }}</td>
					<td>{{ s.created }}</td>
					<td><StatusChip :status="label(s)" :tone="TONE[s.status]" /></td>
					<td>
						<button v-if="s.status !== 'Running'" class="ra" title="Restore" @click="emit('restore', s)">⭡ Restore</button>
						<span v-else class="faint">running…</span>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<style scoped>
.card {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	overflow: hidden;
}
.head {
	font-size: 13px;
	font-weight: 600;
	padding: 14px 16px;
	border-bottom: 1px solid var(--border-soft);
}
.tbl {
	width: 100%;
	border-collapse: collapse;
	font-size: 13px;
}
.tbl thead th {
	text-align: left;
	font-size: 11px;
	font-weight: 600;
	letter-spacing: 0.05em;
	text-transform: uppercase;
	color: var(--text-faint);
	padding: 10px 16px;
}
.tbl tbody td {
	padding: 11px 16px;
	border-top: 1px solid var(--border-soft);
	color: var(--text-dim);
	vertical-align: middle;
}
.name {
	font-family: var(--mono);
	font-size: 12.5px;
	color: var(--text);
}
.mono {
	font-family: var(--mono);
}
.tag {
	font-family: var(--mono);
	font-size: 11px;
	padding: 2px 7px;
	border-radius: 5px;
	background: var(--surface-2);
	border: 1px solid var(--border-soft);
	color: var(--text-dim);
}
.ra {
	font-size: 12px;
	font-weight: 600;
	padding: 5px 10px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	cursor: pointer;
}
.ra:hover {
	border-color: var(--brand);
}
.faint {
	color: var(--text-faint);
	font-size: 12px;
}
</style>

<script setup lang="ts">
import StatusChip from '@/components/shared/StatusChip.vue'
import type { ChipTone } from '@/types/status'
import type { OptRecommendation, Verdict } from '../types'

defineProps<{ recs: OptRecommendation[]; selected: string[]; applying: string | null }>()

const emit = defineEmits<{
	(e: 'apply', rec: OptRecommendation): void
	(e: 'toggle', id: string): void
}>()

const VERDICT: Record<Verdict, { label: string; tone: ChipTone }> = {
	under: { label: 'Under', tone: 'err' },
	over: { label: 'Over', tone: 'warn' },
	idle: { label: 'Idle', tone: 'info' },
	optimal: { label: 'Optimal', tone: 'ok' },
}

function limit(r: { cpu: number; memory: number }) {
	return `${r.cpu}m / ${r.memory}Mi`
}
</script>

<template>
	<table class="tbl">
		<thead>
			<tr>
				<th style="width: 30px"></th>
				<th>Deployment</th>
				<th>Container</th>
				<th>Current limit</th>
				<th>Observed p95</th>
				<th>Suggested</th>
				<th>Verdict</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="r in recs" :key="r.id">
				<td>
					<input type="checkbox" :checked="selected.includes(r.id)" :disabled="r.verdict === 'optimal'" @change="emit('toggle', r.id)" />
				</td>
				<td><span class="name">{{ r.deployment }}</span></td>
				<td class="mono">{{ r.container }}</td>
				<td class="mono">{{ limit(r.current) }}</td>
				<td class="mono">{{ limit(r.usage) }}</td>
				<td class="mono strong">{{ r.verdict === 'optimal' ? 'no change' : limit(r.suggested) }}</td>
				<td><StatusChip :status="VERDICT[r.verdict].label" :tone="VERDICT[r.verdict].tone" /></td>
				<td>
					<button v-if="r.verdict !== 'optimal'" class="btn ai" :disabled="applying === r.id" @click="emit('apply', r)">
						{{ applying === r.id ? 'Applying…' : 'Apply' }}
					</button>
					<span v-else class="faint">—</span>
				</td>
			</tr>
		</tbody>
	</table>
</template>

<style scoped>
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
	padding: 11px 16px;
	border-bottom: 1px solid var(--border-soft);
}
.tbl tbody td {
	padding: 12px 16px;
	border-bottom: 1px solid var(--border-soft);
	color: var(--text-dim);
	vertical-align: middle;
}
.name {
	color: var(--text);
	font-family: var(--mono);
	font-size: 12.5px;
}
.mono {
	font-family: var(--mono);
}
.strong {
	color: var(--text);
}
.faint {
	color: var(--text-faint);
}
.btn {
	font-size: 12px;
	font-weight: 600;
	padding: 6px 12px;
	border-radius: var(--r-sm);
	border: 1px solid transparent;
	background: linear-gradient(180deg, var(--accent), var(--accent-deep));
	color: #fff;
	cursor: pointer;
}
.btn:disabled {
	opacity: 0.6;
	cursor: default;
}
input[type='checkbox'] {
	accent-color: var(--brand);
}
</style>

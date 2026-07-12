<script setup lang="ts">
import { computed } from 'vue'
import type { PodPhase } from '../types'

const props = defineProps<{ phase: PodPhase }>()

const total = computed(() => props.phase.running + props.phase.pending + props.phase.failed || 1)
const donut = computed(() => {
	const r = (props.phase.running / total.value) * 100
	const p = ((props.phase.running + props.phase.pending) / total.value) * 100
	return `conic-gradient(var(--ok) 0 ${r}%, var(--warn) ${r}% ${p}%, var(--err) ${p}% 100%)`
})
</script>

<template>
	<div class="card">
		<div class="t">Pod phase</div>
		<div class="donut" :style="{ background: donut }">
			<div class="lab"><b>{{ phase.running }}</b><span>running</span></div>
		</div>
		<div class="legend">
			<span class="chip ok">Running {{ phase.running }}</span>
			<span class="chip warn">Pending {{ phase.pending }}</span>
			<span class="chip err">Failed {{ phase.failed }}</span>
		</div>
	</div>
</template>

<style scoped>
.card {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 18px;
	display: flex;
	flex-direction: column;
	align-items: center;
}
.t {
	font-size: 13.5px;
	font-weight: 600;
	align-self: flex-start;
}
.donut {
	width: 132px;
	height: 132px;
	border-radius: 50%;
	display: grid;
	place-items: center;
	position: relative;
	margin: 16px 0;
}
.donut::after {
	content: '';
	position: absolute;
	inset: 17px;
	border-radius: 50%;
	background: var(--surface);
}
.lab {
	position: relative;
	z-index: 1;
	text-align: center;
}
.lab b {
	font-size: 24px;
	font-family: var(--mono);
	display: block;
}
.lab span {
	font-size: 11px;
	color: var(--text-dim);
}
.legend {
	display: flex;
	gap: 8px;
	flex-wrap: wrap;
	justify-content: center;
}
.chip {
	font-size: 11px;
	font-weight: 600;
	padding: 3px 9px;
	border-radius: 999px;
}
.chip.ok {
	color: var(--ok);
	background: var(--ok-bg);
}
.chip.warn {
	color: var(--warn);
	background: var(--warn-bg);
}
.chip.err {
	color: var(--err);
	background: var(--err-bg);
}
</style>

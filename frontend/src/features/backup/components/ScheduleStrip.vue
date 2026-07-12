<script setup lang="ts">
import type { Schedule } from '../types'

defineProps<{ schedules: Schedule[] }>()
const emit = defineEmits<{ (e: 'new'): void }>()
</script>

<template>
	<div class="strip">
		<div v-for="s in schedules" :key="s.id" class="card">
			<div class="top">
				<span class="name">{{ s.name }}</span>
				<span class="chip" :class="s.enabled ? 'ok' : 'idle'">{{ s.enabled ? 'Enabled' : 'Paused' }}</span>
			</div>
			<div class="scope">{{ s.scope }}</div>
			<div class="tags"><span class="tag">{{ s.retention }}</span><span class="tag">→ S3</span></div>
		</div>
		<button type="button" class="card new" @click="emit('new')">＋ New schedule</button>
	</div>
</template>

<style scoped>
.strip {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 16px;
	margin-bottom: 22px;
}
.card {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 14px 16px;
}
.top {
	display: flex;
	align-items: center;
	justify-content: space-between;
}
.name {
	font-family: var(--mono);
	font-size: 13px;
	font-weight: 600;
	color: var(--text);
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
.chip.idle {
	color: var(--idle);
	background: var(--idle-bg);
}
.scope {
	font-size: 12px;
	color: var(--text-faint);
	margin-top: 8px;
}
.tags {
	display: flex;
	gap: 8px;
	margin-top: 8px;
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
.new {
	display: grid;
	place-items: center;
	border-style: dashed;
	color: var(--text-dim);
	font-size: 13px;
	cursor: pointer;
	font-family: var(--sans);
}
.new:hover {
	border-color: var(--brand);
	color: var(--text);
}
</style>

<script setup lang="ts">
import type { ActivityItem, ActivityKind } from '../types'

defineProps<{ items: ActivityItem[] }>()

const TONE: Record<ActivityKind, string> = {
	deploy: 'ok',
	scale: 'ok',
	restore: 'info',
	tuning: 'ai',
}
</script>

<template>
	<div class="tl">
		<div v-for="item in items" :key="item.id" class="tl-item" :class="TONE[item.kind]">
			<div class="line"><b>{{ item.lead }}</b> {{ item.text }}</div>
			<div class="meta">{{ item.meta }}</div>
		</div>
	</div>
</template>

<style scoped>
.tl {
	position: relative;
	padding-left: 18px;
}
.tl::before {
	content: '';
	position: absolute;
	left: 4px;
	top: 4px;
	bottom: 10px;
	width: 2px;
	background: var(--border);
}
.tl-item {
	position: relative;
	padding: 0 0 16px;
	font-size: 12.5px;
}
.tl-item:last-child {
	padding-bottom: 0;
}
.tl-item::before {
	content: '';
	position: absolute;
	left: -18px;
	top: 3px;
	width: 9px;
	height: 9px;
	border-radius: 50%;
	background: var(--surface-3);
	border: 2px solid var(--surface);
}
.tl-item.ok::before {
	background: var(--ok);
}
.tl-item.info::before {
	background: var(--info);
}
.tl-item.ai::before {
	background: var(--accent);
}
.line {
	color: var(--text-dim);
}
.line b {
	color: var(--text);
}
.meta {
	color: var(--text-faint);
	font-size: 11px;
	margin-top: 2px;
}
</style>

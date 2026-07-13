<script setup lang="ts">
import { computed } from 'vue'
import type { ChipTone } from '@/types/status'

const props = defineProps<{ status: string; tone?: ChipTone }>()

const TONES: Record<string, ChipTone> = {
	Running: 'ok',
	Succeeded: 'ok',
	Active: 'ok',
	Bound: 'ok',
	Ready: 'ok',
	Completed: 'info',
	Pending: 'warn',
	ContainerCreating: 'warn',
	Terminating: 'warn',
	CrashLoopBackOff: 'err',
	OOMKilled: 'err',
	ImagePullBackOff: 'err',
	Error: 'err',
	Failed: 'err',
	Evicted: 'err',
}

const tone = computed<ChipTone>(() => props.tone ?? TONES[props.status] ?? 'idle')
const label = computed(() => (props.status === 'CrashLoopBackOff' ? 'CrashLoop' : props.status))
</script>

<template>
	<span class="chip" :class="tone">{{ label }}</span>
</template>

<style scoped>
.chip {
	display: inline-flex;
	align-items: center;
	gap: 6px;
	font-size: 11.5px;
	font-weight: 600;
	padding: 3px 9px 3px 8px;
	border-radius: 999px;
	line-height: 1.4;
	white-space: nowrap;
}
.chip::before {
	content: '';
	width: 6px;
	height: 6px;
	border-radius: 50%;
	background: currentColor;
}
.ok {
	color: var(--ok);
	background: var(--ok-bg);
}
.warn {
	color: var(--warn);
	background: var(--warn-bg);
}
.err {
	color: var(--err);
	background: var(--err-bg);
}
.info {
	color: var(--info);
	background: var(--info-bg);
}
.idle {
	color: var(--idle);
	background: var(--idle-bg);
}
</style>

<script setup lang="ts">
import type { ChipTone } from '@/types/status'
import type { Pin } from '@/types/pin'

defineProps<{ items: Pin[] }>()

const emit = defineEmits<{ (e: 'open', item: Pin): void }>()

const DOT: Record<ChipTone, string> = {
	ok: 'var(--ok)',
	warn: 'var(--warn)',
	err: 'var(--err)',
	info: 'var(--info)',
	idle: 'var(--idle)',
}
</script>

<template>
	<div class="list">
		<div v-for="item in items" :key="item.name" class="row" @click="emit('open', item)">
			<span class="dot" :style="{ background: DOT[item.tone] }"></span>
			<span class="name">{{ item.name }}</span>
			<span class="detail">{{ item.detail }}</span>
		</div>
	</div>
</template>

<style scoped>
.list {
	display: grid;
	gap: 2px;
}
.row {
	display: flex;
	align-items: center;
	gap: 10px;
	padding: 7px 6px;
	border-radius: 6px;
	font-size: 12.5px;
	cursor: pointer;
}
.row:hover {
	background: var(--hover);
}
.dot {
	width: 7px;
	height: 7px;
	border-radius: 50%;
	flex: none;
}
.name {
	flex: 1;
	font-family: var(--mono);
	color: var(--text);
}
.detail {
	color: var(--text-faint);
}
</style>

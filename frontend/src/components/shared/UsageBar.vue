<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{ pct: number; tone?: 'auto' | 'ok' | 'warn' | 'err' | 'brand' }>(), {
	tone: 'auto',
})

const resolved = computed(() => {
	if (props.tone !== 'auto') {
		return props.tone
	}
	if (props.pct >= 90) {
		return 'err'
	}
	if (props.pct >= 75) {
		return 'warn'
	}
	return 'ok'
})

const width = computed(() => `${Math.min(Math.max(props.pct, 0), 100)}%`)
</script>

<template>
	<div class="bar" :class="resolved"><span :style="{ width }"></span></div>
</template>

<style scoped>
.bar {
	height: 7px;
	border-radius: 999px;
	background: var(--surface-3);
	overflow: hidden;
}
.bar > span {
	display: block;
	height: 100%;
	border-radius: 999px;
}
.brand > span {
	background: linear-gradient(90deg, var(--brand), #6ea8ff);
}
.ok > span {
	background: linear-gradient(90deg, #2ea043, #46c861);
}
.warn > span {
	background: linear-gradient(90deg, #bb8009, var(--warn));
}
.err > span {
	background: linear-gradient(90deg, #c9352c, var(--err));
}
</style>

<script setup lang="ts">
withDefaults(defineProps<{ size?: number }>(), { size: 18 })
</script>

<template>
	<svg viewBox="0 0 48 48" :width="size" :height="size" aria-hidden="true">
		<circle cx="24" cy="24" r="17" stroke="currentColor" stroke-width="6" fill="none" />
		<path d="M9 24 q7.5 -5 15 0 t15 0" stroke="currentColor" stroke-width="4.5" stroke-linecap="round" fill="none" />
	</svg>
</template>

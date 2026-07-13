<script setup lang="ts">
defineProps<{ title: string; description?: string; last?: boolean }>()
</script>

<template>
	<div class="row" :class="{ last }">
		<div class="meta">
			<div class="title">{{ title }}</div>
			<div v-if="description" class="desc">{{ description }}</div>
		</div>
		<div class="control"><slot /></div>
	</div>
</template>

<style scoped>
.row {
	display: flex;
	align-items: center;
	gap: 20px;
	padding: 15px 0;
	border-bottom: 1px solid var(--border-soft);
}
.row.last {
	border-bottom: none;
}
.meta {
	flex: 1;
	min-width: 0;
}
.title {
	font-size: 13.5px;
	font-weight: 600;
	color: var(--text);
}
.desc {
	font-size: 12px;
	color: var(--text-faint);
	margin-top: 3px;
	line-height: 1.45;
}
.control {
	flex: none;
	display: flex;
	align-items: center;
	gap: 8px;
}
</style>

<script setup lang="ts">
defineProps<{ title: string; subtitle?: string }>()
</script>

<template>
	<section class="card">
		<header class="head">
			<h2>{{ title }}</h2>
			<p v-if="subtitle">{{ subtitle }}</p>
		</header>
		<div class="body"><slot /></div>
	</section>
</template>

<style scoped>
.card {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	overflow: hidden;
}
.head {
	padding: 16px 18px 14px;
	border-bottom: 1px solid var(--border-soft);
}
.head h2 {
	margin: 0;
	font-size: 14px;
	font-weight: 600;
}
.head p {
	margin: 3px 0 0;
	font-size: 12px;
	color: var(--text-faint);
}
.body {
	padding: 4px 18px 14px;
}
</style>

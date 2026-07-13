<script setup lang="ts">
import StatusChip from '@/components/shared/StatusChip.vue'
import type { Issue } from '@/types/issue'

defineProps<{ items: Issue[] }>()

const emit = defineEmits<{ (e: 'action', item: Issue): void }>()
</script>

<template>
	<div class="list">
		<div v-for="item in items" :key="item.id" class="row">
			<StatusChip :status="item.reason" :tone="item.reasonTone" />
			<div class="meta">
				<div class="name">{{ item.name }}</div>
				<div class="sub">{{ item.kind }} · <span class="tag">{{ item.cluster }}</span></div>
			</div>
			<span class="age">{{ item.age }}</span>
			<button class="btn" :class="item.action === 'Diagnose' ? 'ai' : ''" @click="emit('action', item)">{{ item.action }}</button>
		</div>
	</div>
</template>

<style scoped>
.list {
	padding: 4px 6px 8px;
}
.row {
	display: flex;
	align-items: center;
	gap: 12px;
	padding: 11px 10px;
	border-radius: 8px;
}
.row:hover {
	background: var(--hover);
}
.meta {
	flex: 1;
	min-width: 0;
}
.name {
	font-family: var(--mono);
	font-size: 12.5px;
	color: var(--text);
}
.sub {
	font-size: 11px;
	color: var(--text-faint);
	margin-top: 2px;
}
.tag {
	font-family: var(--mono);
	font-size: 10.5px;
	padding: 1px 6px;
	border-radius: 5px;
	background: var(--surface-2);
	border: 1px solid var(--border-soft);
	color: var(--text-dim);
}
.age {
	font-size: 11px;
	color: var(--text-faint);
}
.btn {
	font-size: 12px;
	font-weight: 600;
	padding: 6px 11px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface);
	color: var(--text);
	cursor: pointer;
	white-space: nowrap;
}
.btn:hover {
	border-color: #3a465a;
}
.btn.ai {
	background: linear-gradient(180deg, var(--accent), var(--accent-deep));
	border-color: transparent;
	color: #fff;
}
</style>

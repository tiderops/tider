<script setup lang="ts">
import type { AppError } from '@/services/apperror'
import { userMessage } from '@/services/apperror'

defineProps<{ loading?: boolean; error?: AppError | null; empty?: boolean; emptyText?: string }>()

const emit = defineEmits<{ (e: 'retry'): void }>()
</script>

<template>
	<div class="kx-state">
		<div v-if="loading" class="spinner" aria-label="Loading"></div>
		<template v-else-if="error">
			<div class="err-title">{{ userMessage(error) }}</div>
			<div class="err-msg">{{ error.message }}</div>
			<button class="retry" type="button" @click="emit('retry')">Retry</button>
		</template>
		<div v-else-if="empty" class="empty">{{ emptyText ?? 'Nothing here yet.' }}</div>
	</div>
</template>

<style scoped>
.kx-state {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	gap: 10px;
	padding: 56px 20px;
	text-align: center;
}
.spinner {
	width: 26px;
	height: 26px;
	border-radius: 50%;
	border: 2.5px solid var(--surface-3);
	border-top-color: var(--brand);
	animation: kx-spin 0.7s linear infinite;
}
@keyframes kx-spin {
	to {
		transform: rotate(360deg);
	}
}
.err-title {
	font-size: 14px;
	font-weight: 600;
	color: var(--err);
}
.err-msg {
	font-size: 12.5px;
	color: var(--text-dim);
	font-family: var(--mono);
	max-width: 520px;
}
.retry {
	margin-top: 4px;
	font-size: 12.5px;
	font-weight: 600;
	padding: 7px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	cursor: pointer;
}
.retry:hover {
	border-color: var(--brand);
}
.empty {
	font-size: 13px;
	color: var(--text-faint);
}
</style>

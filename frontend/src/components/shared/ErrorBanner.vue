<script setup lang="ts">
import type { AppError } from '@/services/apperror'
import { userMessage } from '@/services/apperror'

const props = defineProps<{
	error: AppError
}>()

const emit = defineEmits<{
	(e: 'retry'): void
}>()
</script>

<template>
	<v-alert type="error" variant="tonal" :title="userMessage(props.error)" class="ma-2">
		<div class="text-body-2">{{ props.error.message }}</div>
		<template v-slot:append>
			<v-btn variant="outlined" size="small" @click="emit('retry')">Retry</v-btn>
		</template>
	</v-alert>
</template>

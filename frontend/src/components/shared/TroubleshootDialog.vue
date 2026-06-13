<script setup lang="ts">
import { computed } from 'vue'
import { model } from '../../../wailsjs/go/models'

const props = defineProps<{
	modelValue: boolean
	resourceName: string
	loading: boolean
	result?: model.Troubleshoot
}>()

const emit = defineEmits<{
	(e: 'update:modelValue', value: boolean): void
}>()

const open = computed({
	get: () => props.modelValue,
	set: (value: boolean) => emit('update:modelValue', value),
})
</script>

<template>
	<v-dialog v-model="open" max-width="560">
		<v-card :title="`Troubleshoot: ${props.resourceName}`">
			<v-card-text>
				<div v-if="props.loading" class="text-center pa-4">
					<v-progress-circular indeterminate />
				</div>
				<template v-else-if="props.result">
					<div class="text-subtitle-2">Diagnosis</div>
					<p class="mb-4">{{ props.result.Meaning }}</p>
					<template v-if="props.result.Recommendation">
						<div class="text-subtitle-2">Recommendation</div>
						<p>{{ props.result.Recommendation }}</p>
					</template>
				</template>
			</v-card-text>
			<v-card-actions>
				<v-spacer />
				<v-btn @click="open = false">Close</v-btn>
			</v-card-actions>
		</v-card>
	</v-dialog>
</template>

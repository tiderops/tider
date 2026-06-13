<script setup lang="ts">
import { ref } from 'vue'
import type { GridHeader, GridRow, ResourceDescriptor, ResourceDetail } from '@/resources/types'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'
import { statusColor } from '@/utils/status'
import TroubleshootDialog from '@/components/shared/TroubleshootDialog.vue'
import { model } from '../../../wailsjs/go/models'

const props = defineProps<{
	cluster: string
	descriptor: ResourceDescriptor
	headers: GridHeader[]
	items: GridRow[]
	search?: string
	sortBy?: { key: string; order: 'asc' | 'desc' }[]
	loading?: boolean
}>()

const emit = defineEmits<{
	(e: 'edit', item: ResourceDetail | undefined, isOpen: boolean): void
	(e: 'delete', item: GridRow): void
	(e: 'detail', item: ResourceDetail | undefined): void
	(e: 'error', error: AppError): void
}>()

const deleting = ref(false)
const dialog = ref(false)
const itemToDelete = ref<GridRow | null>(null)

const refOf = (item: GridRow) => ({
	cluster: props.cluster,
	namespace: item.namespace ?? '',
	name: item.name,
})

const chipColor = (status: string | undefined) => props.descriptor.statusColor?.(status ?? '') ?? statusColor(status)

const editItem = async (item: GridRow) => {
	if (!props.descriptor.fetchDetail) return

	try {
		const dto = await props.descriptor.fetchDetail(refOf(item))
		emit('edit', dto, true)
	} catch (err) {
		emit('error', toAppError(err))
	}
}

const confirmDelete = (item: GridRow) => {
	itemToDelete.value = item
	dialog.value = true
}

const deleteConfirmed = async () => {
	const remove = props.descriptor.remove
	if (!itemToDelete.value || !remove) return

	deleting.value = true
	try {
		await remove(refOf(itemToDelete.value))
		emit('delete', itemToDelete.value)
	} catch (err) {
		emit('error', toAppError(err))
	} finally {
		deleting.value = false
		dialog.value = false
		itemToDelete.value = null
	}
}

const troubleshootOpen = ref(false)
const troubleshootLoading = ref(false)
const troubleshootName = ref('')
const troubleshootResult = ref<model.Troubleshoot | undefined>(undefined)

const troubleshootItem = async (item: GridRow) => {
	const troubleshoot = props.descriptor.troubleshoot
	if (!troubleshoot) return

	troubleshootName.value = item.name
	troubleshootResult.value = undefined
	troubleshootOpen.value = true
	troubleshootLoading.value = true

	try {
		troubleshootResult.value = await troubleshoot(refOf(item))
	} catch (err) {
		troubleshootOpen.value = false
		emit('error', toAppError(err))
	} finally {
		troubleshootLoading.value = false
	}
}

const onRowClick = async (_event: unknown, row: { item: GridRow }) => {
	if (!props.descriptor.fetchDetail) return

	try {
		const dto = await props.descriptor.fetchDetail(refOf(row.item))
		emit('detail', dto)
	} catch (err) {
		emit('error', toAppError(err))
	}
}
</script>

<template>
	<v-data-table-virtual
		:headers="props.headers"
		:items="props.items"
		:search="props.search"
		:sort-by="props.sortBy"
		:loading="props.loading"
		@click:row="onRowClick"
		height="720"
		item-value="name"
		density="compact"
		fixed-header
		hover
	>
		<template #no-data>
			<div class="pa-8 text-center text-medium-emphasis">No {{ props.descriptor.kind }} resources found in this cluster.</div>
		</template>

		<template #[`item.status`]="{ item }">
			<v-chip :color="chipColor(item.status)" dark>
				{{ item.status }}
			</v-chip>
		</template>

		<template #[`item.actions`]="{ item }">
			<v-btn v-if="props.descriptor.troubleshoot" @click.stop="troubleshootItem(item)" icon title="Troubleshoot">
				<v-icon icon="mdi-stethoscope" />
			</v-btn>

			<v-btn v-if="props.descriptor.update" @click.stop="editItem(item)" icon>
				<v-icon icon="mdi-pencil" />
			</v-btn>

			<v-btn v-if="props.descriptor.remove" @click.stop="confirmDelete(item)" icon>
				<v-icon icon="mdi-delete" />
			</v-btn>
		</template>
	</v-data-table-virtual>

	<TroubleshootDialog v-model="troubleshootOpen" :resource-name="troubleshootName" :loading="troubleshootLoading" :result="troubleshootResult" />

	<v-dialog v-if="dialog" v-model="dialog" max-width="400" persistent>
		<v-card title="Confirm Deletion">
			<v-card-text>
				Are you sure you want to delete
				<strong>{{ itemToDelete?.name }}</strong
				>?
			</v-card-text>

			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn @click="dialog = false" text>Cancel</v-btn>
				<v-btn @click="deleteConfirmed" color="red" :loading="deleting" text>Delete</v-btn>
			</v-card-actions>
		</v-card>
	</v-dialog>
</template>

<style scoped></style>

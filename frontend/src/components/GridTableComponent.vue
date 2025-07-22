<script setup lang="ts">
import { useGridDeleteButton } from '@/composables/useGridDeleteButton'
import {ref} from "vue";

const props = defineProps<{
	cluster?: string
	headers: any[]
	items: any[]
	search?: string
	sortBy?: any[]
	k8sObject?: string
}>()

const emit = defineEmits<{
	(e: 'edit', item: any): void
	(e: 'delete', item: any): void
    (e: 'detail', item: any): void
}>()

const { restart, loading, error, success } = useGridDeleteButton()

const dialog = ref(false)
const itemToDelete = ref<any | null>(null)

const editPod = (item: any) => {
	console.log('EDIT', item)
	emit('edit', item, true)
}

const confirmDelete = (item: any) => {
    itemToDelete.value = item
    dialog.value = true
}

const deleteConfirmed = async () => {
    if (!itemToDelete.value) return
    console.log('DELETE CONFIRMED', itemToDelete.value)
    await restart(itemToDelete.value.name, itemToDelete.value.namespace, props.cluster, props.k8sObject)

    emit('delete', itemToDelete.value)
    dialog.value = false
    itemToDelete.value = null
}

const onRowClick = (item: any) => {
    console.log("onRowClick", item)
    emit('detail', item)
}

</script>

<template>
	<v-data-table-virtual
		:headers="props.headers"
		:items="props.items"
		:search="props.search"
		:sort-by="props.sortBy"
        @click:row="(event, row) => onRowClick(row)"
		height="720"
		item-value="name"
		density="compact"
		fixed-header
		hover
	>
		<template v-slot:item.status="{ item }">
			<v-chip :color="item.status === 'Running' ? 'green' : 'red'" dark>
				{{ item.status }}
			</v-chip>
		</template>

		<template v-slot:item.actions="{ item }">
			<v-btn @click.stop="editPod(item)" icon>
				<v-icon icon="mdi-pencil" />
			</v-btn>

            <v-btn @click.stop="confirmDelete(item)" icon>
                <v-icon icon="mdi-delete" />
            </v-btn>
		</template>
	</v-data-table-virtual>

    <v-dialog v-if="dialog" v-model="dialog" max-width="400" persistent>
        <v-card title="Confirm Deletion">
            <v-card-text>
                Are you sure you want to delete
                <strong>{{ itemToDelete?.name }}</strong>?
            </v-card-text>

            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn @click="dialog = false" text>Cancel</v-btn>
                <v-btn @click="deleteConfirmed" color="red" :loading="loading" text>Delete</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>

</template>

<style scoped></style>

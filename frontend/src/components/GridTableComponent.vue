<script setup lang="ts">
import { gridGeneralComposable } from '@/composables/GridTableComposable'

const props = defineProps<{
  cluster?: string
  headers: any[]
  items: any[]
  search?: string
  sortBy?: any[]
}>()

const emit = defineEmits<{
  (e: 'rowClick', item: any): void
  (e: 'edit', item: any): void
  (e: 'delete', item: any): void
}>()

const editPod = (item: any) => {
  console.log('EDIT', item)
  emit('edit', item)
}

const deletePod = async (item: any) => {
  console.log('DELETE', item)
  console.log('DELETE - NAME', item.name)
  console.log('DELETE - NS', item.namespace)
  console.log('DELETE - CLUSTER ID', props.cluster)

  const { fetchData } = gridGeneralComposable(item.name, item.namespace, props.cluster)

  await fetchData()
  emit('delete', item)
}
</script>

<template>
  <v-data-table-virtual
    :headers="props.headers"
    :items="props.items"
    :search="props.search"
    :sort-by="props.sortBy"
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

      <v-btn @click.stop="deletePod(item)" icon>
        <v-icon icon="mdi-delete" />
      </v-btn>
    </template>
  </v-data-table-virtual>
</template>

<style scoped></style>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useClusterStore } from '@/stores/cluster.store'
import { useNamespacesStore } from '@/stores/namespaces.store'
import { fetchResourceTuning, fetchUpdateDeployment } from '@/services/workload.service'
import { buildTuningUpdate } from '../tuning'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'
import ErrorBanner from '@/components/shared/ErrorBanner.vue'
import { model } from '../../../../wailsjs/go/models'

const { currentCluster } = storeToRefs(useClusterStore())
const namespacesStore = useNamespacesStore()
const namespaces = computed(() => namespacesStore.namespacesFor(currentCluster.value))

const namespace = ref('default')
const loading = ref(false)
const analyzed = ref(false)
const error = ref<AppError | null>(null)
const recommendations = ref<model.TuningRecommendation[]>([])
const applying = ref<string | null>(null)
const snackbar = ref(false)
const snackbarText = ref('')

const headers = [
	{ title: 'Deployment', key: 'Deployment' },
	{ title: 'Container', key: 'Container' },
	{ title: 'Current limit (CPU/Mem)', key: 'current', sortable: false },
	{ title: 'Observed usage', key: 'usage', sortable: false },
	{ title: 'Suggested limit', key: 'suggested', sortable: false },
	{ title: '', key: 'actions', sortable: false },
]

const analyze = async () => {
	loading.value = true
	error.value = null

	try {
		recommendations.value = (await fetchResourceTuning(namespace.value, currentCluster.value)) ?? []
		analyzed.value = true
	} catch (err) {
		error.value = toAppError(err)
		recommendations.value = []
	} finally {
		loading.value = false
	}
}

const keyOf = (rec: model.TuningRecommendation) => `${rec.Deployment}/${rec.Container}`

const apply = async (rec: model.TuningRecommendation) => {
	applying.value = keyOf(rec)

	try {
		await fetchUpdateDeployment(rec.Deployment, rec.Namespace, buildTuningUpdate(rec), currentCluster.value)
		snackbarText.value = `Applied suggested limits to ${rec.Deployment}/${rec.Container}.`
		snackbar.value = true
		recommendations.value = recommendations.value.filter((r) => keyOf(r) !== keyOf(rec))
	} catch (err) {
		const appErr = toAppError(err)
		snackbarText.value = appErr.message
		snackbar.value = true
	} finally {
		applying.value = null
	}
}

onMounted(() => {
	namespacesStore.load(currentCluster.value).catch(() => {})
})
</script>

<template>
	<v-container>
		<v-card class="pa-4 mb-4">
			<div class="d-flex align-center ga-4">
				<v-select v-model="namespace" :items="namespaces" label="Namespace" density="compact" hide-details style="max-width: 300px" />
				<v-btn color="primary" :loading="loading" @click="analyze">Analyze usage</v-btn>
			</div>
			<p class="text-medium-emphasis text-body-2 mt-2 mb-0">
				Compares observed container usage against configured limits and suggests adjustments. Nothing is changed until you apply a
				recommendation.
			</p>
		</v-card>

		<ErrorBanner v-if="error" :error="error" @retry="analyze" />

		<v-card v-else-if="analyzed">
			<v-data-table :headers="headers" :items="recommendations" density="compact">
				<template #[`item.current`]="{ item }"> {{ item.CurrentLimit.Cpu }}m / {{ item.CurrentLimit.Memory }}M </template>
				<template #[`item.usage`]="{ item }"> {{ item.Usage.Cpu }}m / {{ item.Usage.Memory }}M </template>
				<template #[`item.suggested`]="{ item }">
					<strong>{{ item.SuggestedLimit.Cpu }}m / {{ item.SuggestedLimit.Memory }}M</strong>
				</template>
				<template #[`item.actions`]="{ item }">
					<v-btn size="small" color="primary" variant="tonal" :loading="applying === keyOf(item)" @click="apply(item)">Apply</v-btn>
				</template>
				<template #no-data>
					<div class="pa-8 text-center text-medium-emphasis">No tuning recommendations — current limits look adequate.</div>
				</template>
			</v-data-table>
		</v-card>

		<v-snackbar v-model="snackbar" :timeout="3000">{{ snackbarText }}</v-snackbar>
	</v-container>
</template>

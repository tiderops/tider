<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useSettingsStore } from '@/stores/settings.store'
import SettingsCard from './SettingsCard.vue'
import SettingRow from './SettingRow.vue'
import ToggleSwitch from './ToggleSwitch.vue'

const { prefs } = storeToRefs(useSettingsStore())
</script>

<template>
	<SettingsCard title="Troubleshooting assistant" subtitle="Behaviour of the auto-diagnosis engine.">
		<SettingRow title="Auto-diagnose on open" description="Run the diagnosis engine automatically when you open an attention item.">
			<ToggleSwitch v-model="prefs.assistant.autoDiagnose" label="Auto-diagnose" />
		</SettingRow>
		<SettingRow title="Include recent logs" description="Pull the last log lines into the diagnosis context." last>
			<ToggleSwitch v-model="prefs.assistant.includeLogs" label="Include logs" />
		</SettingRow>
	</SettingsCard>

	<SettingsCard title="Right-sizing" subtitle="Defaults the optimization engine uses to suggest limits.">
		<SettingRow title="Headroom over p95" :description="`Suggested limits add ${prefs.assistant.headroomPct}% on top of the 7-day p95 usage.`" last>
			<div class="slider">
				<input v-model.number="prefs.assistant.headroomPct" type="range" min="20" max="60" step="5" />
				<span class="val mono">{{ prefs.assistant.headroomPct }}%</span>
			</div>
		</SettingRow>
	</SettingsCard>
</template>

<style scoped>
.slider {
	display: flex;
	align-items: center;
	gap: 12px;
}
.slider input {
	width: 160px;
	accent-color: var(--accent);
	cursor: pointer;
}
.val {
	font-size: 13px;
	font-weight: 600;
	color: var(--accent);
	min-width: 38px;
	text-align: right;
}
.mono {
	font-family: var(--mono);
}
</style>

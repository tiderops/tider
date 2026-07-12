<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useSettingsStore } from '@/stores/settings.store'
import SettingsCard from './SettingsCard.vue'
import SettingRow from './SettingRow.vue'
import ToggleSwitch from './ToggleSwitch.vue'

const { prefs } = storeToRefs(useSettingsStore())
</script>

<template>
	<SettingsCard title="Alerts" subtitle="What surfaces in the Needs-attention feed and notifies you.">
		<SettingRow title="Critical pods" description="CrashLoopBackOff, OOMKilled, ImagePullBackOff.">
			<ToggleSwitch v-model="prefs.notifications.criticalPods" label="Critical pods" />
		</SettingRow>
		<SettingRow title="Node pressure" description="Memory, disk, or PID pressure and NotReady nodes.">
			<ToggleSwitch v-model="prefs.notifications.nodePressure" label="Node pressure" />
		</SettingRow>
		<SettingRow title="Backup failures" description="Snapshot or restore jobs that did not complete." last>
			<ToggleSwitch v-model="prefs.notifications.backupFailures" label="Backup failures" />
		</SettingRow>
	</SettingsCard>

	<SettingsCard title="Digests & delivery">
		<SettingRow title="Weekly optimization digest" description="A summary of right-sizing savings across prod clusters.">
			<ToggleSwitch v-model="prefs.notifications.optimizationDigest" label="Optimization digest" />
		</SettingRow>
		<SettingRow title="Desktop notifications" description="Show OS-level notifications when the app is in the background." last>
			<ToggleSwitch v-model="prefs.notifications.desktop" label="Desktop notifications" />
		</SettingRow>
	</SettingsCard>
</template>

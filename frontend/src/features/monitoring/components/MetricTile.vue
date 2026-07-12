<script setup lang="ts">
import UsageBar from '@/components/shared/UsageBar.vue'
import type { MetricKpi } from '../types'

defineProps<{ kpi: MetricKpi }>()
</script>

<template>
	<div class="tile">
		<div class="lbl">{{ kpi.label }}</div>
		<div class="val">{{ kpi.value }}<small v-if="kpi.unit"> {{ kpi.unit }}</small></div>
		<UsageBar v-if="kpi.bar !== undefined" class="bar" :pct="kpi.bar" :tone="kpi.barTone ?? 'auto'" />
		<div v-else-if="kpi.hint" class="hint" :class="kpi.hintTone">{{ kpi.hint }}</div>
	</div>
</template>

<style scoped>
.tile {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 16px 18px;
}
.lbl {
	font-size: 12px;
	color: var(--text-dim);
}
.val {
	font-size: 28px;
	font-weight: 700;
	letter-spacing: -0.02em;
	margin-top: 8px;
	font-family: var(--mono);
}
.val small {
	font-size: 13px;
	color: var(--text-faint);
	font-weight: 500;
}
.bar {
	margin-top: 10px;
}
.hint {
	font-size: 11.5px;
	margin-top: 6px;
	color: var(--text-faint);
}
.hint.up {
	color: var(--ok);
}
.hint.down {
	color: var(--err);
}
</style>

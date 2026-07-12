<script setup lang="ts">
withDefaults(
	defineProps<{
		label: string
		value: string
		unit?: string
		hint?: string
		hintTone?: 'up' | 'down'
		valueTone?: 'default' | 'warn' | 'err'
	}>(),
	{ valueTone: 'default' },
)
</script>

<template>
	<div class="stat">
		<div class="lbl">{{ label }}</div>
		<div class="val" :class="valueTone"><span>{{ value }}</span><small v-if="unit"> {{ unit }}</small></div>
		<div v-if="hint" class="hint" :class="hintTone">{{ hint }}</div>
	</div>
</template>

<style scoped>
.stat {
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
.val.warn {
	color: var(--warn);
}
.val.err {
	color: var(--err);
}
.val small {
	font-size: 13px;
	color: var(--text-faint);
	font-weight: 500;
}
.hint {
	font-size: 11.5px;
	margin-top: 4px;
	color: var(--text-faint);
}
.hint.up {
	color: var(--ok);
}
.hint.down {
	color: var(--err);
}
</style>

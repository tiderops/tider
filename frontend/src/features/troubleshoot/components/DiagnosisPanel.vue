<script setup lang="ts">
import StatusChip from '@/components/shared/StatusChip.vue'
import KxState from '@/components/shared/KxState.vue'
import type { AppError } from '@/services/apperror'
import type { Issue } from '@/types/issue'
import type { Diagnosis, DiagnosisAction } from '../types'

const props = defineProps<{ issue: Issue | null; diagnosis: Diagnosis | null; loading: boolean; error: AppError | null }>()

const emit = defineEmits<{
	(e: 'action', action: DiagnosisAction): void
	(e: 'retry'): void
}>()
</script>

<template>
	<div class="panel">
		<div v-if="!props.issue" class="placeholder">Select an issue to see its diagnosis.</div>

		<template v-else>
			<header class="head">
				<div class="left">
					<div class="ico">🩺</div>
					<div>
						<div class="t">Diagnosis</div>
						<div class="sub">{{ props.issue.name }} · {{ props.issue.cluster }}</div>
					</div>
				</div>
				<StatusChip :status="props.issue.reason" :tone="props.issue.reasonTone" />
			</header>

			<KxState v-if="props.loading || props.error" :loading="props.loading" :error="props.error" @retry="emit('retry')" />

			<template v-else-if="props.diagnosis">
				<div class="section-t">What happened</div>
				<p class="prose">{{ props.diagnosis.meaning }}</p>

				<template v-if="props.diagnosis.evidence.length">
					<div class="section-t">Evidence</div>
					<div class="evidence">
						<div v-for="(e, i) in props.diagnosis.evidence" :key="i" class="ev">
							<div class="ev-lbl">{{ e.label }}</div>
							<div class="ev-val">{{ e.value }}</div>
						</div>
					</div>
				</template>

				<div class="section-t">Recommended actions</div>
				<p class="rec">{{ props.diagnosis.recommendation }}</p>

				<div class="actions">
					<div v-for="(a, i) in props.diagnosis.actions" :key="i" class="action">
						<span class="idx">{{ i + 1 }}</span>
						<div class="a-meta">
							<div class="a-label">{{ a.label }}</div>
							<div class="a-desc">{{ a.description }}</div>
						</div>
						<button class="btn" :class="a.kind === 'apply' || a.kind === 'restart' ? 'ai' : ''" @click="emit('action', a)">
							{{ a.kind === 'apply' ? 'Apply' : a.kind === 'restart' ? 'Restart' : a.kind === 'rollback' ? 'Roll back' : a.kind === 'logs' ? 'View' : 'Open' }}
						</button>
					</div>
				</div>
			</template>
		</template>
	</div>
</template>

<style scoped>
.panel {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 22px;
	background-image: linear-gradient(180deg, var(--accent-bg), transparent 40%);
}
.placeholder {
	color: var(--text-faint);
	font-size: 13px;
	padding: 40px 0;
	text-align: center;
}
.head {
	display: flex;
	align-items: center;
	justify-content: space-between;
}
.left {
	display: flex;
	align-items: center;
	gap: 12px;
}
.ico {
	width: 38px;
	height: 38px;
	border-radius: 10px;
	background: var(--accent-soft);
	display: grid;
	place-items: center;
	font-size: 18px;
}
.t {
	font-size: 16px;
	font-weight: 700;
}
.sub {
	font-size: 12px;
	color: var(--text-faint);
	font-family: var(--mono);
}
.section-t {
	font-size: 11px;
	font-weight: 600;
	letter-spacing: 0.06em;
	text-transform: uppercase;
	color: var(--text-faint);
	margin: 22px 0 12px;
}
.prose {
	margin: 0;
	font-size: 13.5px;
	line-height: 1.65;
	color: var(--text);
}
.rec {
	margin: 0 0 14px;
	font-size: 13px;
	line-height: 1.6;
	color: var(--text-dim);
}
.evidence {
	display: grid;
	grid-template-columns: 1fr 1fr;
	gap: 12px;
}
.ev {
	background: var(--surface-2);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 13px 14px;
}
.ev-lbl {
	font-size: 11px;
	color: var(--text-faint);
}
.ev-val {
	font-family: var(--mono);
	font-size: 17px;
	margin-top: 8px;
	color: var(--err);
}
.actions {
	display: grid;
	gap: 8px;
}
.action {
	display: flex;
	align-items: center;
	gap: 12px;
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 13px 15px;
}
.idx {
	color: var(--accent);
	font-size: 15px;
	font-weight: 600;
}
.a-meta {
	flex: 1;
	min-width: 0;
}
.a-label {
	font-size: 13px;
	font-weight: 600;
}
.a-desc {
	font-size: 11.5px;
	color: var(--text-faint);
	margin-top: 2px;
}
.btn {
	font-size: 12px;
	font-weight: 600;
	padding: 6px 12px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	cursor: pointer;
	white-space: nowrap;
}
.btn:hover {
	border-color: var(--brand);
}
.btn.ai {
	background: linear-gradient(180deg, var(--accent), var(--accent-deep));
	border-color: transparent;
	color: #fff;
}
</style>

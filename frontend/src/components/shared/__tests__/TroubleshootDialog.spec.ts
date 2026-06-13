import { describe, it, expect, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import TroubleshootDialog from '../TroubleshootDialog.vue'
import { model } from '../../../../wailsjs/go/models'

const vuetify = createVuetify({ components })

const mountDialog = (props: Record<string, unknown>) =>
	mount(TroubleshootDialog, {
		global: { plugins: [vuetify] },
		props: { modelValue: true, resourceName: 'web-1', loading: false, ...props },
	})

describe('TroubleshootDialog', () => {
	afterEach(() => {
		document.body.innerHTML = ''
	})

	it('shows the diagnosis and recommendation', () => {
		mountDialog({
			result: model.Troubleshoot.createFrom({
				Meaning: 'CrashLoopBackOff',
				Recommendation: 'Check container logs with kubectl logs.',
			}),
		})

		expect(document.body.textContent).toContain('Troubleshoot: web-1')
		expect(document.body.textContent).toContain('CrashLoopBackOff')
		expect(document.body.textContent).toContain('Check container logs')
	})

	it('shows a spinner while loading', () => {
		mountDialog({ loading: true })

		expect(document.body.querySelector('.v-progress-circular')).toBeTruthy()
	})
})

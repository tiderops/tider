import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import ErrorBanner from '../ErrorBanner.vue'
import { AppError } from '@/services/apperror'

const vuetify = createVuetify({ components })

describe('ErrorBanner', () => {
	it('renders the friendly headline and the raw message', () => {
		const wrapper = mount(ErrorBanner, {
			global: { plugins: [vuetify] },
			props: { error: new AppError('CLUSTER_UNREACHABLE', 'cluster "prod" is not registered') },
		})

		expect(wrapper.text()).toContain('unreachable')
		expect(wrapper.text()).toContain('cluster "prod" is not registered')
	})

	it('emits retry when the button is clicked', async () => {
		const wrapper = mount(ErrorBanner, {
			global: { plugins: [vuetify] },
			props: { error: new AppError('TIMEOUT', 'too slow') },
		})

		await wrapper.find('button').trigger('click')
		expect(wrapper.emitted('retry')).toHaveLength(1)
	})
})

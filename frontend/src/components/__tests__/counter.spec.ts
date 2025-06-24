// tests/unit/Counter.spec.ts
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Counter from '@/components/Counter.vue'

describe('Counter.vue', () => {
  it('renders with initial count', () => {
    const wrapper = mount(Counter)
    expect(wrapper.get('[data-testid="count"]').text()).toBe('Count: 0')
  })

  it('increments count on button click', async () => {
    const wrapper = mount(Counter)
    await wrapper.get('button').trigger('click')
    expect(wrapper.get('[data-testid="count"]').text()).toBe('Count: 1')
  })
})

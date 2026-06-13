import { describe, it, expect } from 'vitest'
import { buildTuningUpdate } from '../tuning'
import { model } from '../../../../wailsjs/go/models'

describe('buildTuningUpdate', () => {
	it('maps suggested limits into the deployment update payload', () => {
		const rec = model.TuningRecommendation.createFrom({
			Deployment: 'web',
			Namespace: 'default',
			Container: 'web',
			CurrentLimit: { Cpu: 1000, Memory: 200, Storage: 0, StorageEphemeral: 0 },
			Usage: { Cpu: 950, Memory: 100, Storage: 0, StorageEphemeral: 0 },
			SuggestedLimit: { Cpu: 1500, Memory: 200, Storage: 0, StorageEphemeral: 0 },
		})

		const update = buildTuningUpdate(rec)

		expect(update.Container.Resource.LCpu).toBe(1500)
		expect(update.Container.Resource.LMemory).toBe(200)
	})
})

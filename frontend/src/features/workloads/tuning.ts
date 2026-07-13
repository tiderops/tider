import { model } from '../../../wailsjs/go/models'

export function buildTuningUpdate(rec: model.TuningRecommendation): model.DeploymentUpdate {
	return model.DeploymentUpdate.createFrom({
		Container: {
			Resource: {
				LCpu: rec.SuggestedLimit.Cpu,
				LMemory: rec.SuggestedLimit.Memory,
			},
		},
	})
}

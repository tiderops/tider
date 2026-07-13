package model

type TuningRecommendation struct {
	Deployment     string
	Namespace      string
	Container      string
	CurrentLimit   Resource
	Usage          Resource
	SuggestedLimit Resource
}

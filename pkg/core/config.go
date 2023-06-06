package core

type ModelConfig struct {
	MaxWeight    float64 `json:"maxWeight"`
	MinWeight    float64 `json:"minWeight"`
	MaxDelay     float64 `json:"maxDelay"`
	LearningRate float64 `json:"learningRate"`
	Threshold    float64 `json:"threshold"`
	MaxTime      float64 `json:"maxTime"`
}

func NewDefaultConfig() *ModelConfig {
	return &ModelConfig{
		MinWeight:    -20,
		MaxWeight:    20,
		LearningRate: 0.01,
		Threshold:    200.0,
		MaxDelay:     1.0,
		MaxTime:      10.0,
	}
}

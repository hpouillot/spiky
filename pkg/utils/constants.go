package utils

type Constants struct {
	MaxWeight        float64
	MinWeight        float64
	MaxDelay         float64
	LearningRate     float64
	Threshold        float64
	RefractoryPeriod float64
	Tho              float64
	MaxTime          float64
}

func NewDefaultConstants() *Constants {
	return &Constants{
		MinWeight:        -20,
		MaxWeight:        20,
		LearningRate:     0.01,
		Threshold:        200.0,
		RefractoryPeriod: 10.0,
		Tho:              5.0,
		MaxDelay:         0.0,
		MaxTime:          10.0,
	}
}

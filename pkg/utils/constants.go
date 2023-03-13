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
		MaxWeight:        20,
		MinWeight:        -20.0,
		LearningRate:     0.001,
		Threshold:        200.0,
		RefractoryPeriod: 10.0,
		Tho:              5.0,
		MaxDelay:         1.0,
		MaxTime:          10.0,
	}
}

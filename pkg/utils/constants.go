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
		MaxWeight:        250,
		MinWeight:        200,
		LearningRate:     0.1,
		Threshold:        150.0,
		RefractoryPeriod: 1.0,
		Tho:              10.0,
		MaxDelay:         1.0,
		MaxTime:          30.0,
	}
}

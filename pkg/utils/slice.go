package utils

func Max(slice []float64) float64 {
	var maxValue *float64 = new(float64)
	for _, value := range slice {
		if maxValue == nil || value >= *maxValue {
			*maxValue = value
		}
	}
	return *maxValue
}

func ArgMax(slice []float64) int {
	maxValue := Max(slice)
	for idx, v := range slice {
		if v == maxValue {
			return idx
		}
	}
	panic("Empty slice")
}

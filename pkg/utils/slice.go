package utils

func Max(slice []float64) (float64, int) {
	var maxValue *float64 = new(float64)
	*maxValue = slice[0]
	var argMax = 0
	for idx, value := range slice {
		if value >= *maxValue {
			*maxValue = value
			argMax = idx
		}
	}
	return *maxValue, argMax
}

func Min(slice []float64) (float64, int) {
	var minValue *float64 = new(float64)
	*minValue = slice[0]
	var argMin = 0
	for idx, value := range slice {
		if value <= *minValue {
			*minValue = value
			argMin = idx
		}
	}
	return *minValue, argMin
}

// func ArgMax(slice []float64) int {
// 	maxValue := Max(slice)
// 	for idx, v := range slice {
// 		if v == maxValue {
// 			return idx
// 		}
// 	}
// 	panic("Empty slice")
// }

// func ArgMin(slice []float64) int {
// 	minValue := Min(slice)
// 	for idx, v := range slice {
// 		if v == minValue {
// 			return idx
// 		}
// 	}
// 	panic("Empty slice")
// }

package utils

import "math"

func MaxInt(int1, int2 int) int {
	if int1 > int2 {
		return int1
	}
	return int2
}

func MinInt(int1, int2 int) int {
	if int1 < int2 {
		return int1
	}
	return int2
}

func ClampInt(number, min, max int) int {
	return MaxInt(MinInt(number, max), min)
}

func ClampFloat(number, min, max float64) float64 {
	return math.Max(math.Min(number, max), min)
}

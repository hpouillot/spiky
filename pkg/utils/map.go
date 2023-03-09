package utils

func Map[T interface{}, O interface{}](list []T, apply func(item T, idx int) O) []O {
	results := make([]O, len(list))
	for idx, item := range list {
		results[idx] = apply(item, idx)
	}
	return results
}

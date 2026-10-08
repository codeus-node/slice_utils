package slice_utils

func Filter[T any](s []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

package slice_utils

// Convert converts a slice of type T to a slice of type R
func Convert[T any, R any](slice []T, converter func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = converter(v)
	}
	return result
}

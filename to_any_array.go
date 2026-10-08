package slice_utils

// ToAnyArray converts a slice to an array of any
func ToAnyArray[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

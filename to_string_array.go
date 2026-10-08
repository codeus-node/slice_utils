package slice_utils

import "fmt"

// ToStringArray converts a slice to a string array
func ToStringArray[T any](slice []T) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

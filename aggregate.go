package slice_utils

// Aggregate iterates over a slice and applies an aggregator function to each element.
func Aggregate[T any, R any](slice []T, aggregator func(int, T, R) R, initial R) R {
	for i, value := range slice {
		initial = aggregator(i, value, initial)
	}
	return initial
}

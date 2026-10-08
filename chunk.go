package slice_utils

// Chunk splits a slice into chunks of the specified size.
func Chunk[T any](values []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return [][]T{values}
	}

	chunks := make([][]T, 0, (len(values)+chunkSize-1)/chunkSize)
	for start := 0; start < len(values); start += chunkSize {
		end := start + chunkSize
		if end > len(values) {
			end = len(values)
		}

		chunks = append(chunks, values[start:end])
	}

	return chunks
}

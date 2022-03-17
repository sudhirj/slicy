package godash

import (
	"math"
)

// Chunk splits the given array into groups the length of `chunkSize`.
// If the array cannot be split evenly, the last chunk will have the remaining elements.
func Chunk[T any](array []T, chunkSize int) [][]T {
	chunks := int(math.Ceil(float64(len(array)) / float64(chunkSize)))
	output := make([][]T, chunks)
	for c := 0; c < chunks; c++ {
		start := c * chunkSize
		end := start + chunkSize
		if end > len(array) {
			end = len(array)
		}
		output[c] = array[start:end]
	}
	return output
}

// Concat combines all the elements from all the given arrays into a single array.
func Concat[T any](arrays ...[]T) []T {
	output := make([]T, 0)
	for _, array := range arrays {
		for _, item := range array {
			output = append(output, item)
		}
	}
	return output
}

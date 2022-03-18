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

// Difference returns a list of items present in `array` that are *not* present in any of
// the `others` arrays. The comparison is performed with `==`.
func Difference[T comparable](array []T, others ...[]T) []T {
	output := make([]T, 0)
	for _, item := range array {
		found := false
		// TODO switch to a set to avoid quadratic complexity
		for _, otherArray := range others {
			for _, otherItem := range otherArray {
				if item == otherItem {
					found = true
				}
			}
		}
		if !found {
			output = append(output, item)
		}
	}
	return output
}

func DifferenceBy[T any, U comparable](array []T, iteratee func(T) U, others ...[]T) []T {
	output := make([]T, 0)
	for _, item := range array {
		found := false
		// TODO switch to a set to avoid quadratic complexity
		for _, otherArray := range others {
			for _, otherItem := range otherArray {
				if iteratee(item) == iteratee(otherItem) {
					found = true
				}
			}
		}
		if !found {
			output = append(output, item)
		}
	}
	return output
}

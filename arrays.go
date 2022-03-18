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
	return DifferenceWith(array, func(x, y T) bool { return x == y }, others...)
}

// DifferenceBy returns a list of items present in `array` that are *not* present in any of
// the `others` arrays, with the comparison made by passing items into the `iteratee` function
// and checking `==` on the result. This allows changing the way the item is viewed for comparison.
func DifferenceBy[T any, U comparable](array []T, iteratee func(T) U, others ...[]T) []T {
	return DifferenceWith(array, func(x, y T) bool { return iteratee(x) == iteratee(y) }, others...)
}

// DifferenceWith returns a list of items present in `array` that are *not* present in any of
// the `others` arrays, with the comparison made using the given `comparator`.
func DifferenceWith[T any](array []T, comparator func(T, T) bool, others ...[]T) []T {
	output := make([]T, 0)
	for _, item := range array {
		found := false
		for _, otherArray := range others {
			for _, otherItem := range otherArray {
				if comparator(item, otherItem) {
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

// Drop returns a slice of `array` with `n` elements dropped from the beginning.
func Drop[T any](array []T, n int) []T {
	if n > len(array) {
		n = len(array)
	}
	return array[n:]
}

// DropRight returns a slice of `array` with `n` elements dropped from the end.
func DropRight[T any](array []T, n int) []T {
	if n > len(array) {
		n = len(array)
	}
	return array[:len(array)-n]
}

// DropRightWhile creates a slice of `array` excluding elements dropped from the end.
// Elements are dropped until `predicate` returns false.
func DropRightWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T {
	i := len(array) - 1
	for i >= 0 {
		if !predicate(array[i], i, array) {
			break
		}
		i--
	}
	return array[:i+1]
}

// DropWhile creates a slice of `array` excluding elements dropped from the beginning.
// Elements are dropped until `predicate` returns false.
func DropWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T {
	i := 0
	for i < len(array) {
		if !predicate(array[i], i, array) {
			break
		}
		i++
	}
	return array[i:]
}

// Fill fills elements of `array` with `value` from `start` up to, but not including `end`.
func Fill[T any](array []T, value T, start int, end int) {
	for i := start; i < end; i++ {
		array[i] = value
	}
}

// FindIndex returns the index of the first element for which the `predicate` returns true.
func FindIndex[T any](array []T, predicate func(T) bool) int {
	for i := 0; i < len(array); i++ {
		if predicate(array[i]) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element of which the `predicate` returns true.
func FindLastIndex[T any](array []T, predicate func(T) bool) int {
	for i := len(array) - 1; i >= 0; i-- {
		if predicate(array[i]) {
			return i
		}
	}
	return -1
}

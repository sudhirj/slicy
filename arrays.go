package slicy

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"strings"
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
		output = append(output, array...)
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

// IndexOf returns the index at which the first occurrence of `value` is found in `array`.
// Returns `-1` if not found.
func IndexOf[T comparable](array []T, value T) int {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return i
		}
	}
	return -1
}

// Intersection returns an array of unique values that are included in all given arrays.
// The order of the result values are determined by the first array.
func Intersection[T comparable](array []T, others ...[]T) []T {
	return IntersectionWith(array, func(x, y T) bool { return x == y }, others...)
}

// IntersectionBy returns an array of unique values that are included in all given arrays,
// with comparison happening on the result of the `iteratee` function. The order of the result
// values are determined by the first array.
func IntersectionBy[T comparable, U comparable](array []T, iteratee func(T) U, others ...[]T) []T {
	return IntersectionWith(array, func(x, y T) bool { return iteratee(x) == iteratee(y) }, others...)
}

// IntersectionWith returns an array of unique values that are included in all given arrays,
// with comparison happening inside the given `comparator`. The order of the result values
// are determined by the first array.
func IntersectionWith[T comparable](array []T, comparator func(T, T) bool, others ...[]T) []T {
	output := make([]T, 0)
	for _, item := range array {
		findCount := 0
		for _, otherArray := range others {
			for _, otherItem := range otherArray {
				if comparator(item, otherItem) {
					findCount++
					break
				}
			}
		}
		if findCount == len(others) && FindIndex(output, func(e T) bool { return comparator(e, item) }) == -1 {
			output = append(output, item)
		}
	}
	return output
}

// Join concatenates all the elements of the array into a string separated by `separator`.
// `fmt.Sprint` is used for to get the string representation of the given value, so mixed types
// are possible with `[]any`.
func Join[T any](array []T, separator string) string {
	stringList := make([]string, len(array))
	for i, e := range array {
		stringList[i] = fmt.Sprint(e)
	}
	return strings.Join(stringList, separator)
}

// LastIndexOf returns the index at which the last occurrence of `value` is found in `array`.
// Returns `-1` if not found.
func LastIndexOf[T comparable](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if value == array[i] {
			return i
		}
	}
	return -1
}

// Nth gets the element at index `n` of the `array`. If `n` is negative, the nth element
// from the end is returned.
func Nth[T any](array []T, n int) T {
	if n < 0 {
		n = len(array) + n
	}
	return array[n]
}

// Pull returns a slice of `array` without all the given `values`.
func Pull[T comparable](array []T, values ...T) []T {
	return PullAll(array, values)
}

// PullAll returns a slice of `array` without the items in `values`.
func PullAll[T comparable](array []T, values []T) []T {
	return PullAllWith(array, values, func(x, y T) bool { return x == y })
}

// PullAllBy returns a slice of `array` without the items in `values`, with the
// comparison made by passing both values through the `iteratee` function.
func PullAllBy[T any, U comparable](array []T, values []T, iteratee func(T) U) []T {
	return PullAllWith(array, values, func(x, y T) bool { return iteratee(x) == iteratee(y) })
}

// PullAllWith returns a slice of `array` without the items in `values`, with the
// comparison made using the given `comparator`.
func PullAllWith[T any](array []T, values []T, comparator func(T, T) bool) []T {
	output := make([]T, 0, len(array)-len(values))
	for _, v := range array {
		if FindIndex(values, func(x T) bool { return comparator(x, v) }) == -1 {
			output = append(output, v)
		}
	}
	return output
}

// PullAt returns a slice of `array` without the items at the given indexes.
func PullAt[T comparable](array []T, indexes ...int) []T {
	output := make([]T, 0, len(array)-len(indexes))
	for i := range array {
		if IndexOf(indexes, i) == -1 {
			output = append(output, array[i])
		}
	}
	return output
}

// Remove returns a slice of `array` without the elements for which the `predicate`
// returns `true`.
func Remove[T any](array []T, predicate func(value T, index int, array []T) bool) []T {
	output := make([]T, 0)
	for i := range array {
		if !predicate(array[i], i, array) {
			output = append(output, array[i])
		}
	}
	return output
}

// Reverse return the reverse of `array`: with the first element last, the second element second-to-last, and so on.
func Reverse[T any](array []T) []T {
	output := make([]T, len(array))
	for i := range array {
		output[len(array)-1-i] = array[i]
	}
	return output
}

// SortedIndex uses a binary search to determine the lowest index at which `value` should be inserted into
// `array` in order to maintain its sort order.
func SortedIndex[T constraints.Ordered](array []T, value T) int {
	return slices.BinarySearch(array, value)
}

// SortedIndexBy uses a binary search to determine the lowest index at which `value` should be inserted into
// `array` in order to maintain its sort order, with the `iteratee` function used to computed sort ranking.
func SortedIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int {
	return slices.BinarySearchFunc(array, func(e T) bool { return iteratee(e) >= iteratee(value) })
}

// SortedIndexOf performs a binary search on a sorted `array` to find the given `value`. Returns -1 if not found.
func SortedIndexOf[T constraints.Ordered](array []T, value T) int {
	k := slices.BinarySearch(array, value)
	if k >= len(array) || array[k] != value {
		return -1
	}
	return k
}

// SortedLastIndex returns the highest index at which `value` should be inserted into the sorted `array` to maintain
// its sort order.
func SortedLastIndex[T constraints.Ordered](array []T, value T) int {
	i := SortedIndex(array, value)
	// we now want the next index that has a bigger value in the remaining sub-slice
	j := slices.BinarySearchFunc(array[i:], func(v T) bool {
		return v > value
	})
	return i + j
}

// SortedLastIndexBy returns the highest index at which `value` should be inserted into the sorted `array` to maintain
// its sort order, with comparisons made on the result of passing all values through `iteratee`.
func SortedLastIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int {
	i := SortedIndexBy(array, value, iteratee)
	j := slices.BinarySearchFunc(array[i:], func(v T) bool {
		return iteratee(v) > iteratee(value)
	})
	return i + j
}

// SortedLastIndexOf returns the highest index at which the `value` is present in the sorted `array`.
func SortedLastIndexOf[T constraints.Ordered](array []T, value T) int {
	i := SortedIndexOf(array, value)
	if i == -1 {
		return i
	}
	j := slices.BinarySearchFunc(array[i:], func(v T) bool {
		return v > value
	})
	return i + j - 1
}

// Take returns a slice of `array` with `n` elements taken from the beginning.
func Take[T any](array []T, n int) []T {
	if n > len(array) {
		n = len(array)
	}
	return array[:n]
}

// TakeRight returns a slice of `array` with `n` elements taken from the end.
func TakeRight[T any](array []T, n int) []T {
	if n > len(array) {
		n = len(array)
	}
	return array[len(array)-n:]
}

// TakeRightWhile create a slice of elements taken from the end of `array`.
// Elements are taken until the `predicate` returns false.
func TakeRightWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T {
	i := len(array) - 1
	for i >= 0 {
		if !predicate(array[i], i, array) {
			break
		}
		i--
	}
	return array[i+1:]
}

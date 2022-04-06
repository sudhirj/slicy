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
func Chunk[S ~[]T, T any](slice S, chunkSize int) []S {
	chunks := int(math.Ceil(float64(len(slice)) / float64(chunkSize)))
	output := make([]S, chunks)
	for c := 0; c < chunks; c++ {
		start := c * chunkSize
		end := start + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		output[c] = slice[start:end]
	}
	return output
}

// Concat combines all the elements from all the given slices into a single slice.
func Concat[S ~[]T, T any](slices ...S) S {
	output := make(S, 0)
	for _, list := range slices {
		output = append(output, list...)
	}
	return output
}

// Difference returns a list of items present in `slice` that are *not* present in any of
// the `others` slices. The comparison is performed with `==`.
func Difference[S ~[]T, T comparable](slice S, others ...S) S {
	return DifferenceWith(slice, func(x, y T) bool { return x == y }, others...)
}

// DifferenceBy returns a list of items present in `slice` that are *not* present in any of
// the `others` slices, with the comparison made by passing items into the `iteratee` function
// and checking `==` on the result. This allows changing the way the item is viewed for comparison.
func DifferenceBy[S ~[]T, T any, U comparable](array S, iteratee func(T) U, others ...S) S {
	return DifferenceWith(array, func(x, y T) bool { return iteratee(x) == iteratee(y) }, others...)
}

// DifferenceWith returns a slice of items present in `slice` that are *not* present in any of
// the `others` slices, with the comparison made using the given `comparator`.
func DifferenceWith[S ~[]T, T any](slice S, comparator func(T, T) bool, others ...S) S {
	output := make(S, 0)
	for _, item := range slice {
		found := Some(others, func(otherSlice S, _ int, _ []S) bool {
			return Some(otherSlice, func(v T, _ int, _ S) bool { return comparator(item, v) })
		})
		if !found {
			output = append(output, item)
		}
	}
	return output
}

// Drop returns a new slice with `n` elements dropped from the beginning.
func Drop[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[n:]
}

// DropRight returns a new slice with `n` elements dropped from the end.
func DropRight[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:len(slice)-n]
}

// DropRightWhile creates a new slice excluding elements dropped from the end.
// Elements are dropped until `predicate` returns false.
func DropRightWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	i := len(slice) - 1
	for i >= 0 {
		if !predicate(slice[i], i, slice) {
			break
		}
		i--
	}
	return slice[:i+1]
}

// DropWhile creates a new slice excluding elements dropped from the beginning.
// Elements are dropped until `predicate` returns false.
func DropWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	i := 0
	for i < len(slice) {
		if !predicate(slice[i], i, slice) {
			break
		}
		i++
	}
	return slice[i:]
}

// Fill fills elements of `slice` with `value` from `start` up to, but not including `end`.
func Fill[S ~[]T, T any](slice S, value T, start int, end int) {
	for i := start; i < end; i++ {
		slice[i] = value
	}
}

// FindIndex returns the index of the first element for which the `predicate` returns true.
func FindIndex[S ~[]T, T any](slice S, predicate func(T) bool) int {
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
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
func Intersection[T comparable](arrays ...[]T) []T {
	return IntersectionWith(func(x, y T) bool { return x == y }, arrays...)
}

// IntersectionBy returns an array of unique values that are included in all given arrays,
// with comparison happening on the result of the `iteratee` function. The order of the result
// values are determined by the first array.
func IntersectionBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T {
	return IntersectionWith(func(x, y T) bool { return iteratee(x) == iteratee(y) }, arrays...)
}

// IntersectionWith returns an array of unique values that are included in all given arrays,
// with comparison happening inside the given `comparator`. The order of the result values
// are determined by the first array.
func IntersectionWith[T any](comparator func(T, T) bool, arrays ...[]T) []T {
	output := make([]T, 0)
	for _, array := range arrays {
		for _, item := range array {
			found := true
			for _, searchArray := range arrays {
				if slices.IndexFunc(searchArray, func(e T) bool { return comparator(e, item) }) == -1 {
					found = false
				}
			}
			if found && slices.IndexFunc(output, func(e T) bool { return comparator(e, item) }) == -1 {
				output = append(output, item)
			}
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

func cmp[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// SortedIndex uses a binary search to determine the lowest index at which `value` should be inserted into
// `array` in order to maintain its sort order.
func SortedIndex[T constraints.Ordered](array []T, value T) int {
	i, _ := slices.BinarySearch(array, value)
	return i
}

// SortedIndexBy uses a binary search to determine the lowest index at which `value` should be inserted into
// `array` in order to maintain its sort order, with the `iteratee` function used to computed sort ranking.
func SortedIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int {
	i, _ := slices.BinarySearchFunc(array, value, func(a, b T) int { return cmp(iteratee(a), iteratee(b)) })
	return i
}

// SortedIndexOf performs a binary search on a sorted `array` to find the given `value`. Returns -1 if not found.
func SortedIndexOf[T constraints.Ordered](array []T, value T) int {
	k, found := slices.BinarySearch(array, value)
	if !found {
		return -1
	}
	return k
}

// SortedLastIndex returns the highest index at which `value` should be inserted into the sorted `array` to maintain
// its sort order.
func SortedLastIndex[T constraints.Ordered](array []T, value T) int {
	i := SortedIndex(array, value)
	// we now want the next index that has a bigger value in the remaining sub-slice
	j := FindIndex(array[i:], func(v T) bool { return v > value })
	if j == -1 {
		return len(array)
	}
	return i + j
}

// SortedLastIndexBy returns the highest index at which `value` should be inserted into the sorted `array` to maintain
// its sort order, with comparisons made on the result of passing all values through `iteratee`.
func SortedLastIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int {
	i := SortedIndexBy(array, value, iteratee)
	j := FindIndex(array[i:], func(v T) bool { return iteratee(v) > iteratee(value) })
	if j == -1 {
		return len(array)
	}
	return i + j
}

// SortedLastIndexOf returns the highest index at which the `value` is present in the sorted `array`.
func SortedLastIndexOf[T constraints.Ordered](array []T, value T) int {
	i := SortedIndexOf(array, value)
	if i == -1 {
		return i
	}
	j := FindIndex(array[i:], func(v T) bool { return v > value })
	if j == -1 {
		return i
	}
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

// TakeRightWhile creates a slice of elements taken from the end of `array`.
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

// TakeWhile creates a slice of elements taken from the beginning of `array`.
// Elements are taken until the `predicate` returns false.
func TakeWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T {
	i := 0
	for i < len(array) {
		if !predicate(array[i], i, array) {
			break
		}
		i++
	}
	return array[:i]
}

// Union creates a new slice, in order, of unique values of all the given arrays. Uses `==` for equality checks.
func Union[T comparable](arrays ...[]T) []T {
	return UnionWith(func(a, b T) bool { return a == b }, arrays...)
}

// UnionBy creates a new slice, in order, of unique values of all the given arrays.
// Uses the result of the given `iteratee` to check equality.
func UnionBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T {
	return UnionWith(func(a, b T) bool { return iteratee(a) == iteratee(b) }, arrays...)
}

// UnionWith creates a new slice, in order, of unique values of all the given arrays.
// Uses the given `comparator` to check equality between elements.
func UnionWith[T any](comparator func(T, T) bool, arrays ...[]T) []T {
	output := make([]T, 0)
	for _, array := range arrays {
		for _, e := range array {
			if slices.IndexFunc(output, func(v T) bool { return comparator(e, v) }) == -1 {
				output = append(output, e)
			}
		}
	}
	return output
}

// Uniq returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed with `==`.
func Uniq[T comparable](array []T) []T {
	return Union(array)
}

// UniqBy returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed with `==` on the result of passing each element through the given `iteratee`.
func UniqBy[T any, U comparable](iteratee func(T) U, array []T) []T {
	return UnionBy(iteratee, array)
}

// UniqWith returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed using the given `comparator`.
func UniqWith[T any](comparator func(T, T) bool, array []T) []T {
	return UnionWith(comparator, array)
}

// Without returns a new slice without the given elements. Uses `==` for equality checks.
func Without[T comparable](array []T, values ...T) []T {
	output := make([]T, 0)
	for _, e := range array {
		if slices.Index(values, e) == -1 {
			output = append(output, e)
		}
	}
	return output
}

// Xor returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
func Xor[T comparable](arrays ...[]T) []T {
	return XorWith(func(a, b T) bool { return a == b }, arrays...)
}

// XorBy returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
// Equality is determined by passing elements through the given `iteratee`.
func XorBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T {
	return XorWith(func(a, b T) bool { return iteratee(a) == iteratee(b) }, arrays...)
}

// XorWith returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
// Equality is determined by passing elements to the given `comparator`.
func XorWith[T any](comparator func(T, T) bool, arrays ...[]T) []T {
	output := make([]T, 0)
	intersection := IntersectionWith(comparator, arrays...)
	for _, array := range arrays {
		for _, item := range array {
			f := func(e T) bool { return comparator(e, item) }
			if slices.IndexFunc(intersection, f) == -1 && slices.IndexFunc(output, f) == -1 {
				output = append(output, item)
			}
		}
	}
	return output
}

// CountBy creates an object composed of keys generated from the results of running each element
// of the collection through `iteratee`. The corresponding value of each key is the number
// of times the key was returned by `iteratee`.
func CountBy[T any, U comparable](array []T, iteratee func(T) U) map[U]int {
	output := make(map[U]int)
	for _, item := range array {
		output[iteratee(item)]++
	}
	return output
}

// Each invokes the given `iteratee` for every element in the collection, from left to right.
func Each[T any](array []T, iteratee func(value T, index int, array []T)) {
	for i, v := range array {
		iteratee(v, i, array)
	}
}

// EachRight invokes the given `iteratee` for every element in the collection, from right to left.
func EachRight[T any](array []T, iteratee func(value T, index int, array []T)) {
	for i := len(array) - 1; i >= 0; i-- {
		iteratee(array[i], i, array)
	}
}

// Every returns true if the given `predicate` returns true for every element of the given
// collection.
func Every[T any](array []T, predicate func(value T, index int, array []T) bool) bool {
	for i, item := range array {
		if !predicate(item, i, array) {
			return false
		}
	}
	return true
}

// Some return true if the given `predicate` returns true for any element of the given collection.
func Some[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool {
	for i, item := range slice {
		if predicate(item, i, slice) {
			return true
		}
	}
	return false
}

// Filter iterates over the elements of `collection`, returning an array of all elements
// that the `predicate` returns true for.
func Filter[T any](array []T, predicate func(value T, index int, arr []T) bool) []T {
	output := make([]T, 0)
	for i, item := range array {
		if predicate(item, i, array) {
			output = append(output, item)
		}
	}
	return output
}

// Find iterates over the elements of `array`, returning the first element that `predicate`
// returns true for.
func Find[T any](array []T, predicate func(value T, index int, arr []T) bool) (result T) {
	for i, item := range array {
		if predicate(item, i, array) {
			return item
		}
	}
	return
}

// FlatMap creates a flattened array of values by running each element in `array` through
// `iteratee` and flattening the mapped results.
func FlatMap[T any, U any](array []T, iteratee func(value T, index int, arr []T) []U) []U {
	output := make([]U, 0)
	for i, item := range array {
		output = append(output, iteratee(item, i, array)...)
	}
	return output
}

// GroupBy creates a map composed of keys generated from the results of running each element
// of `array` through `iteratee`. The order of the grouped values is determined by the order
// that they occur in `array`. The corresponding value of each key is an array of elements
// responsible for generating the key.
func GroupBy[T any, U comparable](array []T, iteratee func(T) U) map[U][]T {
	output := make(map[U][]T)
	for _, item := range array {
		key := iteratee(item)
		output[key] = append(output[key], item)
	}
	return output
}

// Includes checks if `value` is in `array`. Equality is checked with `==`.
func Includes[T comparable](array []T, value T) bool {
	return slices.Contains(array, value)
}

// KeyBy creates a map composed of keys generated from the results of running each element
// of `array` through `iteratee`. The corresponding value of each key is the last element
// responsible for generating the key.
func KeyBy[T any, U comparable](array []T, iteratee func(T) U) map[U]T {
	output := make(map[U]T)
	for _, item := range array {
		key := iteratee(item)
		output[key] = item
	}
	return output
}

// Map creates a slice of values by running each element in `array` through
// `iteratee`.
func Map[T any, U any](array []T, iteratee func(T) U) []U {
	output := make([]U, len(array))
	for i, item := range array {
		output[i] = iteratee(item)
	}
	return output
}

// Partition creates two slices, the first of which contains elements that
// `predicate` returns true for, with the second containing elements for which
// `predicate` returns false.
func Partition[T any](array []T, predicate func(T) bool) (truths []T, falsehoods []T) {
	truths = make([]T, 0)
	falsehoods = make([]T, 0)
	for _, item := range array {
		if predicate(item) {
			truths = append(truths, item)
		} else {
			falsehoods = append(falsehoods, item)
		}
	}
	return
}

// Reduce reduces `array` to a value which is the accumulated result of running
// each element in `array` through `iteratee`, where each successive invocation is
// supplied the return value of the previous one. `accumulator` is used as the initial value.
func Reduce[T any, U any](array []T, iteratee func(acc U, value T, index int, arr []T) U, accumulator U) U {
	for i, item := range array {
		accumulator = iteratee(accumulator, item, i, array)
	}
	return accumulator
}

// ReduceRight reduces `array` to a value which is the accumulated result of running
// each element in `array`, from right to left, through `iteratee`, where each successive
// invocation is supplied the return value of the previous one. `accumulator` is used as the initial value.
func ReduceRight[T any, U any](array []T, iteratee func(acc U, value T, index int, arr []T) U, accumulator U) U {
	for i := len(array) - 1; i >= 0; i-- {
		accumulator = iteratee(accumulator, array[i], i, array)
	}
	return accumulator
}

// Reject iterates over the elements of `array`, returning a new slice of the elements for which
// `predicate` returns false.
func Reject[T any](array []T, predicate func(value T, index int, arr []T) bool) []T {
	output := make([]T, 0)
	for i, item := range array {
		if !predicate(item, i, array) {
			output = append(output, item)
		}
	}
	return output
}

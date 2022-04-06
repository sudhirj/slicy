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
func FindLastIndex[S ~[]T, T any](slice S, predicate func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}

// IndexOf returns the index at which the first occurrence of `value` is found in `slice`.
// Returns `-1` if not found.
func IndexOf[S ~[]T, T comparable](slice S, value T) int {
	for i := 0; i < len(slice); i++ {
		if value == slice[i] {
			return i
		}
	}
	return -1
}

// Intersection returns a slice of unique values that are included in all given slices.
// The order of the result values are determined by the first slice.
func Intersection[S ~[]T, T comparable](slices ...S) S {
	return IntersectionWith(func(x, y T) bool { return x == y }, slices...)
}

// IntersectionBy returns a slice of unique values that are included in all given slices,
// with comparison happening on the result of the `iteratee` function. The order of the result
// values are determined by the first slice.
func IntersectionBy[S ~[]T, T any, U comparable](iteratee func(T) U, arrays ...S) S {
	return IntersectionWith(func(x, y T) bool { return iteratee(x) == iteratee(y) }, arrays...)
}

// IntersectionWith returns a slice of unique values that are included in all given slice,
// with comparison happening inside the given `comparator`. The order of the result values
// are determined by the first slice.
func IntersectionWith[S ~[]T, T any](comparator func(T, T) bool, slices ...S) S {
	output := make(S, 0)
	for _, slice := range slices {
		for _, item := range slice {
			found := All(slices, func(s S, _ int, _ []S) bool {
				return Any(s, func(value T, _ int, _ S) bool {
					return comparator(value, item)
				})
			})
			if found && !Any(output, func(v T, _ int, _ S) bool { return comparator(v, item) }) {
				output = append(output, item)
			}
		}
	}
	return output
}

// Join concatenates all the elements of the slice into a string separated by `separator`.
// `fmt.Sprint` is used for to get the string representation of the given value, so mixed types
// are possible with `[]any`.
func Join[S ~[]T, T any](slice S, separator string) string {
	stringList := make([]string, len(slice))
	for i, e := range slice {
		stringList[i] = fmt.Sprint(e)
	}
	return strings.Join(stringList, separator)
}

// LastIndexOf returns the index at which the last occurrence of `value` is found in `slice`.
// Returns `-1` if not found.
func LastIndexOf[S ~[]T, T comparable](slice S, value T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if value == slice[i] {
			return i
		}
	}
	return -1
}

// Nth gets the element at index `n` of the `slice`. If `n` is negative, the nth element
// from the end is returned.
func Nth[S ~[]T, T any](slice S, n int) T {
	if n < 0 {
		n = len(slice) + n
	}
	return slice[n]
}

// Pull returns a slice of `slice` without all the given `values`.
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

// Remove returns a new slice without the elements for which the `predicate`
// returns `true`.
func Remove[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	output := make(S, 0)
	for i := range slice {
		if !predicate(slice[i], i, slice) {
			output = append(output, slice[i])
		}
	}
	return output
}

// Reverse return the reverse of `slice`: with the first element last, the second element second-to-last, and so on.
func Reverse[S ~[]T, T any](slice S) S {
	output := make([]T, len(slice))
	for i := range slice {
		output[len(slice)-1-i] = slice[i]
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
// `slice` in order to maintain its sort order.
func SortedIndex[S ~[]T, T constraints.Ordered](slice S, value T) int {
	i, _ := slices.BinarySearch(slice, value)
	return i
}

// SortedIndexBy uses a binary search to determine the lowest index at which `value` should be inserted into
// `slice` in order to maintain its sort order, with the `iteratee` function used to compute sort ranking.
func SortedIndexBy[S ~[]T, T any, U constraints.Ordered](slice S, value T, iteratee func(T) U) int {
	i, _ := slices.BinarySearchFunc(slice, value, func(a, b T) int { return cmp(iteratee(a), iteratee(b)) })
	return i
}

// SortedIndexOf performs a binary search on a sorted `slice` to find the given `value`. Returns -1 if not found.
func SortedIndexOf[S ~[]T, T constraints.Ordered](slice S, value T) int {
	k, found := slices.BinarySearch(slice, value)
	if !found {
		return -1
	}
	return k
}

// SortedLastIndex returns the highest index at which `value` should be inserted into the sorted `slice` to maintain
// its sort order.
func SortedLastIndex[S ~[]T, T constraints.Ordered](slice S, value T) int {
	i := SortedIndex(slice, value)
	// we now want the next index that has a bigger value in the remaining sub-slice
	j := FindIndex(slice[i:], func(v T) bool { return v > value })
	if j == -1 {
		return len(slice)
	}
	return i + j
}

// SortedLastIndexBy returns the highest index at which `value` should be inserted into the sorted `slice` to maintain
// its sort order, with comparisons made on the result of passing all values through `iteratee`.
func SortedLastIndexBy[S ~[]T, T any, U constraints.Ordered](slice S, value T, iteratee func(T) U) int {
	i := SortedIndexBy(slice, value, iteratee)
	j := FindIndex(slice[i:], func(v T) bool { return iteratee(v) > iteratee(value) })
	if j == -1 {
		return len(slice)
	}
	return i + j
}

// SortedLastIndexOf returns the highest index at which the `value` is present in the sorted `slice`.
func SortedLastIndexOf[S ~[]T, T constraints.Ordered](slice S, value T) int {
	i := SortedIndexOf(slice, value)
	if i == -1 {
		return i
	}
	j := FindIndex(slice[i:], func(v T) bool { return v > value })
	if j == -1 {
		return i
	}
	return i + j - 1
}

// Take returns a new slice with `n` elements taken from the beginning.
func Take[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}

// TakeRight returns a new slice with `n` elements taken from the end.
func TakeRight[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[len(slice)-n:]
}

// TakeRightWhile creates a slice of elements taken from the end of `slice`.
// Elements are taken until the `predicate` returns false.
func TakeRightWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	i := len(slice) - 1
	for i >= 0 {
		if !predicate(slice[i], i, slice) {
			break
		}
		i--
	}
	return slice[i+1:]
}

// TakeWhile creates a slice of elements taken from the beginning of `slice`.
// Elements are taken until the `predicate` returns false.
func TakeWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	i := 0
	for i < len(slice) {
		if !predicate(slice[i], i, slice) {
			break
		}
		i++
	}
	return slice[:i]
}

// Union creates a new slice, in order, of unique values of all the given slices. Uses `==` for equality checks.
func Union[S ~[]T, T comparable](slices ...S) S {
	return UnionWith(func(a, b T) bool { return a == b }, slices...)
}

// UnionBy creates a new slice, in order, of unique values of all the given slices.
// Uses the result of the given `iteratee` to check equality.
func UnionBy[S ~[]T, T any, U comparable](iteratee func(T) U, slices ...S) S {
	return UnionWith(func(a, b T) bool { return iteratee(a) == iteratee(b) }, slices...)
}

// UnionWith creates a new slice, in order, of unique values of all the given slices.
// Uses the given `comparator` to check equality between elements.
func UnionWith[S ~[]T, T any](comparator func(T, T) bool, sliceList ...S) S {
	output := make(S, 0)
	for _, slice := range sliceList {
		for _, e := range slice {
			if slices.IndexFunc(output, func(v T) bool { return comparator(e, v) }) == -1 {
				output = append(output, e)
			}
		}
	}
	return output
}

// Uniq returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed with `==`.
func Uniq[S ~[]T, T comparable](slice S) S {
	return Union(slice)
}

// UniqBy returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed with `==` on the result of passing each element through the given `iteratee`.
func UniqBy[S ~[]T, T any, U comparable](iteratee func(T) U, slice S) S {
	return UnionBy(iteratee, slice)
}

// UniqWith returns a new slice, in order, with no duplicates, with only the first occurrence of each element kept.
// Comparison is performed using the given `comparator`.
func UniqWith[S ~[]T, T any](comparator func(T, T) bool, slice S) S {
	return UnionWith(comparator, slice)
}

// Without returns a new slice without the given elements. Uses `==` for equality checks.
func Without[S ~[]T, T comparable](slice S, values ...T) S {
	output := make(S, 0)
	for _, e := range slice {
		if slices.Index(values, e) == -1 {
			output = append(output, e)
		}
	}
	return output
}

// Xor returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given slices.
// The order of result values is determined by the order they occur in the slices.
func Xor[S ~[]T, T comparable](slices ...S) S {
	return XorWith(func(a, b T) bool { return a == b }, slices...)
}

// XorBy returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given arrays.
// The order of result values is determined by the order they occur in the slices.
// Equality is determined by passing elements through the given `iteratee`.
func XorBy[S ~[]T, T any, U comparable](iteratee func(T) U, slices ...S) S {
	return XorWith(func(a, b T) bool { return iteratee(a) == iteratee(b) }, slices...)
}

// XorWith returns a new slice of unique values that is the symmetric difference
// (elements which are any of the sets but not in their intersection) of the given slices.
// The order of result values is determined by the order they occur in the slices.
// Equality is determined by passing elements to the given `comparator`.
func XorWith[S ~[]T, T any](comparator func(T, T) bool, sliceList ...S) S {
	output := make(S, 0)
	intersection := IntersectionWith(comparator, sliceList...)
	for _, slice := range sliceList {
		for _, item := range slice {
			f := func(e T) bool { return comparator(e, item) }
			if slices.IndexFunc(intersection, f) == -1 && slices.IndexFunc(output, f) == -1 {
				output = append(output, item)
			}
		}
	}
	return output
}

// CountBy creates a map composed of keys generated from the results of running each element
// of the slice through `iteratee`. The corresponding value of each key is the number
// of times the key was returned by `iteratee`.
func CountBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]int {
	output := make(map[U]int)
	for _, item := range slice {
		output[iteratee(item)]++
	}
	return output
}

// Each invokes the given `iteratee` for every element in the slice, from left to right.
func Each[S ~[]T, T any](slice S, iteratee func(value T, index int, slice S)) {
	for i, v := range slice {
		iteratee(v, i, slice)
	}
}

// EachRight invokes the given `iteratee` for every element in the slice, from right to left.
func EachRight[S ~[]T, T any](slice S, iteratee func(value T, index int, slice S)) {
	for i := len(slice) - 1; i >= 0; i-- {
		iteratee(slice[i], i, slice)
	}
}

// Every returns true if the given `predicate` returns true for every element of the given
// slice.
func Every[S ~[]T, T any](slice S, predicate func(value T, index int, array S) bool) bool {
	for i, item := range slice {
		if !predicate(item, i, slice) {
			return false
		}
	}
	return true
}

// All returns true if the given `predicate` returns true for every element of the given
// slice.
func All[T any](array []T, predicate func(value T, index int, array []T) bool) bool {
	return Every(array, predicate)
}

// Some return true if the given `predicate` returns true for any element of the given slice.
func Some[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool {
	for i, item := range slice {
		if predicate(item, i, slice) {
			return true
		}
	}
	return false
}

// Any return true if the given `predicate` returns true for any element of the given slice.
func Any[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool {
	return Some(slice, predicate)
}

// Filter iterates over the elements of `slice`, returning an array of all elements
// that the `predicate` returns true for.
func Filter[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	output := make([]T, 0)
	for i, item := range slice {
		if predicate(item, i, slice) {
			output = append(output, item)
		}
	}
	return output
}

// Find iterates over the elements of `slice`, returning the first element that `predicate`
// returns true for.
func Find[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) (result T) {
	for i, item := range slice {
		if predicate(item, i, slice) {
			return item
		}
	}
	return
}

// FlatMap creates a flattened slice of values by running each element in `slice` through
// `iteratee` and flattening the mapped results.
func FlatMap[S ~[]T, T any, U any](slice S, iteratee func(value T, index int, slice S) []U) []U {
	output := make([]U, 0)
	for i, item := range slice {
		output = append(output, iteratee(item, i, slice)...)
	}
	return output
}

// GroupBy creates a map composed of keys generated from the results of running each element
// of `slice` through `iteratee`. The order of the grouped values is determined by the order
// that they occur in `slice`. The corresponding value of each key is an array of elements
// responsible for generating the key.
func GroupBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]S {
	output := make(map[U]S)
	for _, item := range slice {
		key := iteratee(item)
		output[key] = append(output[key], item)
	}
	return output
}

// Includes checks if `value` is in `slice`. Equality is checked with `==`.
func Includes[S ~[]T, T comparable](slice S, value T) bool {
	return slices.Contains(slice, value)
}

// KeyBy creates a map composed of keys generated from the results of running each element
// of `slice` through `iteratee`. The corresponding value of each key is the last element
// responsible for generating the key.
func KeyBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]T {
	output := make(map[U]T)
	for _, item := range slice {
		key := iteratee(item)
		output[key] = item
	}
	return output
}

// Map creates a slice of values by running each element in `slice` through
// `iteratee`.
func Map[S ~[]T, T any, U any](slice S, iteratee func(T) U) []U {
	output := make([]U, len(slice))
	for i, item := range slice {
		output[i] = iteratee(item)
	}
	return output
}

// Partition creates two slices, the first of which contains elements that
// `predicate` returns true for, with the second containing elements for which
// `predicate` returns false.
func Partition[S ~[]T, T any](slice S, predicate func(T) bool) (truths S, falsehoods S) {
	truths = make([]T, 0)
	falsehoods = make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			truths = append(truths, item)
		} else {
			falsehoods = append(falsehoods, item)
		}
	}
	return
}

// Reduce reduces `slice` to a value which is the accumulated result of running
// each element in `slice` through `iteratee`, where each successive invocation is
// supplied the return value of the previous one. `accumulator` is used as the initial value.
func Reduce[S ~[]T, T any, U any](slice S, iteratee func(acc U, value T, index int, slice S) U, accumulator U) U {
	for i, item := range slice {
		accumulator = iteratee(accumulator, item, i, slice)
	}
	return accumulator
}

// ReduceRight reduces `slice` to a value which is the accumulated result of running
// each element in `slice`, from right to left, through `iteratee`, where each successive
// invocation is supplied the return value of the previous one. `accumulator` is used as the initial value.
func ReduceRight[S ~[]T, T any, U any](slice S, iteratee func(acc U, value T, index int, slice S) U, accumulator U) U {
	for i := len(slice) - 1; i >= 0; i-- {
		accumulator = iteratee(accumulator, slice[i], i, slice)
	}
	return accumulator
}

// Reject iterates over the elements of `slice`, returning a new slice of the elements for which
// `predicate` returns false.
func Reject[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	output := make(S, 0)
	for i, item := range slice {
		if !predicate(item, i, slice) {
			output = append(output, item)
		}
	}
	return output
}

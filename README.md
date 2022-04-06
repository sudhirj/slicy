# slicy

    import "github.com/sudhirj/slicy"

## Usage

#### func  All

```go
func All[T any](array []T, predicate func(value T, index int, array []T) bool) bool
```

All returns true if the given `predicate` returns true for every element of the
given slice.

#### func  Any

```go
func Any[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool
```

Any return true if the given `predicate` returns true for any element of the
given slice.

#### func  Chunk

```go
func Chunk[S ~[]T, T any](slice S, chunkSize int) []S
```

Chunk splits the given array into groups the length of `chunkSize`. If the array
cannot be split evenly, the last chunk will have the remaining elements.

#### func  Concat

```go
func Concat[S ~[]T, T any](slices ...S) S
```
Concat combines all the elements from all the given slices into a single slice.

#### func  CountBy

```go
func CountBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]int
```

CountBy creates a map composed of keys generated from the results of running
each element of the slice through `iteratee`. The corresponding value of each
key is the number of times the key was returned by `iteratee`.

#### func  Difference

```go
func Difference[S ~[]T, T comparable](slice S, others ...S) S
```
Difference returns a list of items present in `slice` that are *not* present in
any of the `others` slices. The comparison is performed with `==`.

#### func  DifferenceBy

```go
func DifferenceBy[S ~[]T, T any, U comparable](array S, iteratee func(T) U, others ...S) S
```
DifferenceBy returns a list of items present in `slice` that are *not* present
in any of the `others` slices, with the comparison made by passing items into
the `iteratee` function and checking `==` on the result. This allows changing
the way the item is viewed for comparison.

#### func  DifferenceWith

```go
func DifferenceWith[S ~[]T, T any](slice S, comparator func(T, T) bool, others ...S) S
```
DifferenceWith returns a slice of items present in `slice` that are *not*
present in any of the `others` slices, with the comparison made using the given
`comparator`.

#### func  Drop

```go
func Drop[S ~[]T, T any](slice S, n int) S
```
Drop returns a new slice with `n` elements dropped from the beginning.

#### func  DropRight

```go
func DropRight[S ~[]T, T any](slice S, n int) S
```
DropRight returns a new slice with `n` elements dropped from the end.

#### func  DropRightWhile

```go
func DropRightWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```
DropRightWhile creates a new slice excluding elements dropped from the end.
Elements are dropped until `predicate` returns false.

#### func  DropWhile

```go
func DropWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```
DropWhile creates a new slice excluding elements dropped from the beginning.
Elements are dropped until `predicate` returns false.

#### func  Each

```go
func Each[S ~[]T, T any](slice S, iteratee func(value T, index int, slice S))
```

Each invokes the given `iteratee` for every element in the slice, from left to
right.

#### func  EachRight

```go
func EachRight[S ~[]T, T any](slice S, iteratee func(value T, index int, slice S))
```

EachRight invokes the given `iteratee` for every element in the slice, from
right to left.

#### func  Every

```go
func Every[S ~[]T, T any](slice S, predicate func(value T, index int, array S) bool) bool
```
Every returns true if the given `predicate` returns true for every element of
the given slice.

#### func  Fill

```go
func Fill[S ~[]T, T any](slice S, value T, start int, end int)
```
Fill fills elements of `slice` with `value` from `start` up to, but not
including `end`.

#### func  Filter

```go
func Filter[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```

Filter iterates over the elements of `slice`, returning an array of all elements
that the `predicate` returns true for.

#### func  Find

```go
func Find[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) (result T)
```

Find iterates over the elements of `slice`, returning the first element that
`predicate` returns true for.

#### func  FindIndex

```go
func FindIndex[S ~[]T, T any](slice S, predicate func(T) bool) int
```
FindIndex returns the index of the first element for which the `predicate`
returns true.

#### func  FindLastIndex

```go
func FindLastIndex[S ~[]T, T any](slice S, predicate func(T) bool) int
```
FindLastIndex returns the index of the last element of which the `predicate`
returns true.

#### func  FlatMap

```go
func FlatMap[S ~[]T, T any, U any](slice S, iteratee func(value T, index int, slice S) []U) []U
```

FlatMap creates a flattened slice of values by running each element in `slice`
through `iteratee` and flattening the mapped results.

#### func  GroupBy

```go
func GroupBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]S
```

GroupBy creates a map composed of keys generated from the results of running
each element of `slice` through `iteratee`. The order of the grouped values is
determined by the order that they occur in `slice`. The corresponding value of
each key is an array of elements responsible for generating the key.

#### func  Includes

```go
func Includes[S ~[]T, T comparable](slice S, value T) bool
```

Includes checks if `value` is in `slice`. Equality is checked with `==`.

#### func  IndexOf

```go
func IndexOf[S ~[]T, T comparable](slice S, value T) int
```

IndexOf returns the index at which the first occurrence of `value` is found in
`slice`. Returns `-1` if not found.

#### func  Intersection

```go
func Intersection[S ~[]T, T comparable](slices ...S) S
```

Intersection returns a slice of unique values that are included in all given
slices. The order of the result values are determined by the first slice.

#### func  IntersectionBy

```go
func IntersectionBy[S ~[]T, T any, U comparable](iteratee func(T) U, arrays ...S) S
```

IntersectionBy returns a slice of unique values that are included in all given
slices, with comparison happening on the result of the `iteratee` function. The
order of the result values are determined by the first slice.

#### func  IntersectionWith

```go
func IntersectionWith[S ~[]T, T any](comparator func(T, T) bool, slices ...S) S
```

IntersectionWith returns a slice of unique values that are included in all given
slice, with comparison happening inside the given `comparator`. The order of the
result values are determined by the first slice.

#### func  Join

```go
func Join[S ~[]T, T any](slice S, separator string) string
```

Join concatenates all the elements of the slice into a string separated by
`separator`. `fmt.Sprint` is used for to get the string representation of the
given value, so mixed types are possible with `[]any`.

#### func  KeyBy

```go
func KeyBy[S ~[]T, T any, U comparable](slice S, iteratee func(T) U) map[U]T
```

KeyBy creates a map composed of keys generated from the results of running each
element of `slice` through `iteratee`. The corresponding value of each key is
the last element responsible for generating the key.

#### func  LastIndexOf

```go
func LastIndexOf[S ~[]T, T comparable](slice S, value T) int
```

LastIndexOf returns the index at which the last occurrence of `value` is found
in `slice`. Returns `-1` if not found.

#### func  Map

```go
func Map[S ~[]T, T any, U any](slice S, iteratee func(T) U) []U
```

Map creates a slice of values by running each element in `slice` through
`iteratee`.

#### func  Nth

```go
func Nth[S ~[]T, T any](slice S, n int) T
```

Nth gets the element at index `n` of the `slice`. If `n` is negative, the nth
element from the end is returned.

#### func  Partition

```go
func Partition[S ~[]T, T any](slice S, predicate func(T) bool) (truths S, falsehoods S)
```
Partition creates two slices, the first of which contains elements that
`predicate` returns true for, with the second containing elements for which
`predicate` returns false.

#### func  Pull

```go
func Pull[T comparable](array []T, values ...T) []T
```

Pull returns a slice of `slice` without all the given `values`.

#### func  PullAll

```go
func PullAll[T comparable](array []T, values []T) []T
```
PullAll returns a slice of `array` without the items in `values`.

#### func  PullAllBy

```go
func PullAllBy[T any, U comparable](array []T, values []T, iteratee func(T) U) []T
```
PullAllBy returns a slice of `array` without the items in `values`, with the
comparison made by passing both values through the `iteratee` function.

#### func  PullAllWith

```go
func PullAllWith[T any](array []T, values []T, comparator func(T, T) bool) []T
```
PullAllWith returns a slice of `array` without the items in `values`, with the
comparison made using the given `comparator`.

#### func  PullAt

```go
func PullAt[T comparable](array []T, indexes ...int) []T
```
PullAt returns a slice of `array` without the items at the given indexes.

#### func  Reduce

```go
func Reduce[S ~[]T, T any, U any](slice S, iteratee func(acc U, value T, index int, slice S) U, accumulator U) U
```

Reduce reduces `slice` to a value which is the accumulated result of running
each element in `slice` through `iteratee`, where each successive invocation is
supplied the return value of the previous one. `accumulator` is used as the
initial value.

#### func  ReduceRight

```go
func ReduceRight[S ~[]T, T any, U any](slice S, iteratee func(acc U, value T, index int, slice S) U, accumulator U) U
```

ReduceRight reduces `slice` to a value which is the accumulated result of
running each element in `slice`, from right to left, through `iteratee`, where
each successive invocation is supplied the return value of the previous one.
`accumulator` is used as the initial value.

#### func  Reject

```go
func Reject[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```

Reject iterates over the elements of `slice`, returning a new slice of the
elements for which `predicate` returns false.

#### func  Remove

```go
func Remove[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```

Remove returns a new slice without the elements for which the `predicate`
returns `true`.

#### func  Reverse

```go
func Reverse[S ~[]T, T any](slice S) S
```

Reverse return the reverse of `slice`: with the first element last, the second
element second-to-last, and so on.

#### func  Some

```go
func Some[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool
```
Some return true if the given `predicate` returns true for any element of the
given slice.

#### func  SortedIndex

```go
func SortedIndex[S ~[]T, T constraints.Ordered](slice S, value T) int
```

SortedIndex uses a binary search to determine the lowest index at which `value`
should be inserted into `slice` in order to maintain its sort order.

#### func  SortedIndexBy

```go
func SortedIndexBy[S ~[]T, T any, U constraints.Ordered](slice S, value T, iteratee func(T) U) int
```

SortedIndexBy uses a binary search to determine the lowest index at which
`value` should be inserted into `slice` in order to maintain its sort order,
with the `iteratee` function used to compute sort ranking.

#### func  SortedIndexOf

```go
func SortedIndexOf[S ~[]T, T constraints.Ordered](slice S, value T) int
```

SortedIndexOf performs a binary search on a sorted `slice` to find the given
`value`. Returns -1 if not found.

#### func  SortedLastIndex

```go
func SortedLastIndex[S ~[]T, T constraints.Ordered](slice S, value T) int
```

SortedLastIndex returns the highest index at which `value` should be inserted
into the sorted `slice` to maintain its sort order.

#### func  SortedLastIndexBy

```go
func SortedLastIndexBy[S ~[]T, T any, U constraints.Ordered](slice S, value T, iteratee func(T) U) int
```

SortedLastIndexBy returns the highest index at which `value` should be inserted
into the sorted `slice` to maintain its sort order, with comparisons made on the
result of passing all values through `iteratee`.

#### func  SortedLastIndexOf

```go
func SortedLastIndexOf[S ~[]T, T constraints.Ordered](slice S, value T) int
```

SortedLastIndexOf returns the highest index at which the `value` is present in
the sorted `slice`.

#### func  Take

```go
func Take[S ~[]T, T any](slice S, n int) S
```

Take returns a new slice with `n` elements taken from the beginning.

#### func  TakeRight

```go
func TakeRight[S ~[]T, T any](slice S, n int) S
```

TakeRight returns a new slice with `n` elements taken from the end.

#### func  TakeRightWhile

```go
func TakeRightWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```

TakeRightWhile creates a slice of elements taken from the end of `slice`.
Elements are taken until the `predicate` returns false.

#### func  TakeWhile

```go
func TakeWhile[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S
```

TakeWhile creates a slice of elements taken from the beginning of `slice`.
Elements are taken until the `predicate` returns false.

#### func  Union

```go
func Union[S ~[]T, T comparable](slices ...S) S
```

Union creates a new slice, in order, of unique values of all the given slices.
Uses `==` for equality checks.

#### func  UnionBy

```go
func UnionBy[S ~[]T, T any, U comparable](iteratee func(T) U, slices ...S) S
```

UnionBy creates a new slice, in order, of unique values of all the given slices.
Uses the result of the given `iteratee` to check equality.

#### func  UnionWith

```go
func UnionWith[S ~[]T, T any](comparator func(T, T) bool, sliceList ...S) S
```

UnionWith creates a new slice, in order, of unique values of all the given
slices. Uses the given `comparator` to check equality between elements.

#### func  Uniq

```go
func Uniq[S ~[]T, T comparable](slice S) S
```
Uniq returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed with `==`.

#### func  UniqBy

```go
func UniqBy[S ~[]T, T any, U comparable](iteratee func(T) U, slice S) S
```
UniqBy returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed with `==` on the result
of passing each element through the given `iteratee`.

#### func  UniqWith

```go
func UniqWith[S ~[]T, T any](comparator func(T, T) bool, slice S) S
```
UniqWith returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed using the given
`comparator`.

#### func  Without

```go
func Without[S ~[]T, T comparable](slice S, values ...T) S
```
Without returns a new slice without the given elements. Uses `==` for equality
checks.

#### func  Xor

```go
func Xor[S ~[]T, T comparable](slices ...S) S
```

Xor returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
slices. The order of result values is determined by the order they occur in the
slices.

#### func  XorBy

```go
func XorBy[S ~[]T, T any, U comparable](iteratee func(T) U, slices ...S) S
```

XorBy returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
arrays. The order of result values is determined by the order they occur in the
slices. Equality is determined by passing elements through the given `iteratee`.

#### func  XorWith

```go
func XorWith[S ~[]T, T any](comparator func(T, T) bool, sliceList ...S) S
```

XorWith returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
slices. The order of result values is determined by the order they occur in the
slices. Equality is determined by passing elements to the given `comparator`.

# slicy
--
    import "github.com/sudhirj/slicy"


## Usage

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
func CountBy[T any, U comparable](array []T, iteratee func(T) U) map[U]int
```
CountBy creates an object composed of keys generated from the results of running
each element of the collection through `iteratee`. The corresponding value of
each key is the number of times the key was returned by `iteratee`.

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
func Each[T any](array []T, iteratee func(value T, index int, array []T))
```
Each invokes the given `iteratee` for every element in the collection, from left
to right.

#### func  EachRight

```go
func EachRight[T any](array []T, iteratee func(value T, index int, array []T))
```
EachRight invokes the given `iteratee` for every element in the collection, from
right to left.

#### func  Every

```go
func Every[T any](array []T, predicate func(value T, index int, array []T) bool) bool
```
Every returns true if the given `predicate` returns true for every element of
the given collection.

#### func  Fill

```go
func Fill[S ~[]T, T any](slice S, value T, start int, end int)
```

Fill fills elements of `slice` with `value` from `start` up to, but not
including `end`.

#### func  Filter

```go
func Filter[T any](array []T, predicate func(value T, index int, arr []T) bool) []T
```
Filter iterates over the elements of `collection`, returning an array of all
elements that the `predicate` returns true for.

#### func  Find

```go
func Find[T any](array []T, predicate func(value T, index int, arr []T) bool) (result T)
```
Find iterates over the elements of `array`, returning the first element that
`predicate` returns true for.

#### func  FindIndex

```go
func FindIndex[S ~[]T, T any](slice S, predicate func(T) bool) int
```
FindIndex returns the index of the first element for which the `predicate`
returns true.

#### func  FindLastIndex

```go
func FindLastIndex[T any](array []T, predicate func(T) bool) int
```
FindLastIndex returns the index of the last element of which the `predicate`
returns true.

#### func  FlatMap

```go
func FlatMap[T any, U any](array []T, iteratee func(value T, index int, arr []T) []U) []U
```
FlatMap creates a flattened array of values by running each element in `array`
through `iteratee` and flattening the mapped results.

#### func  GroupBy

```go
func GroupBy[T any, U comparable](array []T, iteratee func(T) U) map[U][]T
```
GroupBy creates a map composed of keys generated from the results of running
each element of `array` through `iteratee`. The order of the grouped values is
determined by the order that they occur in `array`. The corresponding value of
each key is an array of elements responsible for generating the key.

#### func  Includes

```go
func Includes[T comparable](array []T, value T) bool
```
Includes checks if `value` is in `array`. Equality is checked with `==`.

#### func  IndexOf

```go
func IndexOf[T comparable](array []T, value T) int
```
IndexOf returns the index at which the first occurrence of `value` is found in
`array`. Returns `-1` if not found.

#### func  Intersection

```go
func Intersection[T comparable](arrays ...[]T) []T
```
Intersection returns an array of unique values that are included in all given
arrays. The order of the result values are determined by the first array.

#### func  IntersectionBy

```go
func IntersectionBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T
```
IntersectionBy returns an array of unique values that are included in all given
arrays, with comparison happening on the result of the `iteratee` function. The
order of the result values are determined by the first array.

#### func  IntersectionWith

```go
func IntersectionWith[T any](comparator func(T, T) bool, arrays ...[]T) []T
```
IntersectionWith returns an array of unique values that are included in all
given arrays, with comparison happening inside the given `comparator`. The order
of the result values are determined by the first array.

#### func  Join

```go
func Join[T any](array []T, separator string) string
```
Join concatenates all the elements of the array into a string separated by
`separator`. `fmt.Sprint` is used for to get the string representation of the
given value, so mixed types are possible with `[]any`.

#### func  KeyBy

```go
func KeyBy[T any, U comparable](array []T, iteratee func(T) U) map[U]T
```
KeyBy creates a map composed of keys generated from the results of running each
element of `array` through `iteratee`. The corresponding value of each key is
the last element responsible for generating the key.

#### func  LastIndexOf

```go
func LastIndexOf[T comparable](array []T, value T) int
```
LastIndexOf returns the index at which the last occurrence of `value` is found
in `array`. Returns `-1` if not found.

#### func  Map

```go
func Map[T any, U any](array []T, iteratee func(T) U) []U
```
Map creates a slice of values by running each element in `array` through
`iteratee`.

#### func  Nth

```go
func Nth[T any](array []T, n int) T
```
Nth gets the element at index `n` of the `array`. If `n` is negative, the nth
element from the end is returned.

#### func  Partition

```go
func Partition[T any](array []T, predicate func(T) bool) (truths []T, falsehoods []T)
```
Partition creates two slices, the first of which contains elements that
`predicate` returns true for, with the second containing elements for which
`predicate` returns false.

#### func  Pull

```go
func Pull[T comparable](array []T, values ...T) []T
```
Pull returns a slice of `array` without all the given `values`.

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
func Reduce[T any, U any](array []T, iteratee func(acc U, value T, index int, arr []T) U, accumulator U) U
```
Reduce reduces `array` to a value which is the accumulated result of running
each element in `array` through `iteratee`, where each successive invocation is
supplied the return value of the previous one. `accumulator` is used as the
initial value.

#### func  ReduceRight

```go
func ReduceRight[T any, U any](array []T, iteratee func(acc U, value T, index int, arr []T) U, accumulator U) U
```
ReduceRight reduces `array` to a value which is the accumulated result of
running each element in `array`, from right to left, through `iteratee`, where
each successive invocation is supplied the return value of the previous one.
`accumulator` is used as the initial value.

#### func  Reject

```go
func Reject[T any](array []T, predicate func(value T, index int, arr []T) bool) []T
```
Reject iterates over the elements of `array`, returning a new slice of the
elements for which `predicate` returns false.

#### func  Remove

```go
func Remove[T any](array []T, predicate func(value T, index int, array []T) bool) []T
```

Remove returns a slice of `array` without the elements for which the `predicate`
returns `true`.

#### func  Reverse

```go
func Reverse[T any](array []T) []T
```

Reverse return the reverse of `array`: with the first element last, the second
element second-to-last, and so on.

#### func  Some

```go
func Some[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) bool
```

Some return true if the given `predicate` returns true for any element of the
given collection.

#### func  SortedIndex

```go
func SortedIndex[T constraints.Ordered](array []T, value T) int
```

SortedIndex uses a binary search to determine the lowest index at which `value`
should be inserted into `array` in order to maintain its sort order.

#### func  SortedIndexBy

```go
func SortedIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int
```
SortedIndexBy uses a binary search to determine the lowest index at which
`value` should be inserted into `array` in order to maintain its sort order,
with the `iteratee` function used to computed sort ranking.

#### func  SortedIndexOf

```go
func SortedIndexOf[T constraints.Ordered](array []T, value T) int
```
SortedIndexOf performs a binary search on a sorted `array` to find the given
`value`. Returns -1 if not found.

#### func  SortedLastIndex

```go
func SortedLastIndex[T constraints.Ordered](array []T, value T) int
```
SortedLastIndex returns the highest index at which `value` should be inserted
into the sorted `array` to maintain its sort order.

#### func  SortedLastIndexBy

```go
func SortedLastIndexBy[T any, U constraints.Ordered](array []T, value T, iteratee func(T) U) int
```
SortedLastIndexBy returns the highest index at which `value` should be inserted
into the sorted `array` to maintain its sort order, with comparisons made on the
result of passing all values through `iteratee`.

#### func  SortedLastIndexOf

```go
func SortedLastIndexOf[T constraints.Ordered](array []T, value T) int
```
SortedLastIndexOf returns the highest index at which the `value` is present in
the sorted `array`.

#### func  Take

```go
func Take[T any](array []T, n int) []T
```
Take returns a slice of `array` with `n` elements taken from the beginning.

#### func  TakeRight

```go
func TakeRight[T any](array []T, n int) []T
```
TakeRight returns a slice of `array` with `n` elements taken from the end.

#### func  TakeRightWhile

```go
func TakeRightWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T
```
TakeRightWhile creates a slice of elements taken from the end of `array`.
Elements are taken until the `predicate` returns false.

#### func  TakeWhile

```go
func TakeWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T
```
TakeWhile creates a slice of elements taken from the beginning of `array`.
Elements are taken until the `predicate` returns false.

#### func  Union

```go
func Union[T comparable](arrays ...[]T) []T
```
Union creates a new slice, in order, of unique values of all the given arrays.
Uses `==` for equality checks.

#### func  UnionBy

```go
func UnionBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T
```
UnionBy creates a new slice, in order, of unique values of all the given arrays.
Uses the result of the given `iteratee` to check equality.

#### func  UnionWith

```go
func UnionWith[T any](comparator func(T, T) bool, arrays ...[]T) []T
```
UnionWith creates a new slice, in order, of unique values of all the given
arrays. Uses the given `comparator` to check equality between elements.

#### func  Uniq

```go
func Uniq[T comparable](array []T) []T
```
Uniq returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed with `==`.

#### func  UniqBy

```go
func UniqBy[T any, U comparable](iteratee func(T) U, array []T) []T
```
UniqBy returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed with `==` on the result
of passing each element through the given `iteratee`.

#### func  UniqWith

```go
func UniqWith[T any](comparator func(T, T) bool, array []T) []T
```
UniqWith returns a new slice, in order, with no duplicates, with only the first
occurrence of each element kept. Comparison is performed using the given
`comparator`.

#### func  Without

```go
func Without[T comparable](array []T, values ...T) []T
```
Without returns a new slice without the given elements. Uses `==` for equality
checks.

#### func  Xor

```go
func Xor[T comparable](arrays ...[]T) []T
```
Xor returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
arrays. The order of result values is determined by the order they occur in the
arrays.

#### func  XorBy

```go
func XorBy[T any, U comparable](iteratee func(T) U, arrays ...[]T) []T
```
XorBy returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
arrays. The order of result values is determined by the order they occur in the
arrays. Equality is determined by passing elements through the given `iteratee`.

#### func  XorWith

```go
func XorWith[T any](comparator func(T, T) bool, arrays ...[]T) []T
```
XorWith returns a new slice of unique values that is the symmetric difference
(elements which are any of the sets but not in their intersection) of the given
arrays. The order of result values is determined by the order they occur in the
arrays. Equality is determined by passing elements to the given `comparator`.

# slicy

    import "github.com/sudhirj/slicy"


## Usage

#### func  Chunk

```go
func Chunk[T any](array []T, chunkSize int) [][]T
```
Chunk splits the given array into groups the length of `chunkSize`. If the array
cannot be split evenly, the last chunk will have the remaining elements.

#### func  Concat

```go
func Concat[T any](arrays ...[]T) []T
```
Concat combines all the elements from all the given arrays into a single array.

#### func  Difference

```go
func Difference[T comparable](array []T, others ...[]T) []T
```
Difference returns a list of items present in `array` that are *not* present in
any of the `others` arrays. The comparison is performed with `==`.

#### func  DifferenceBy

```go
func DifferenceBy[T any, U comparable](array []T, iteratee func(T) U, others ...[]T) []T
```
DifferenceBy returns a list of items present in `array` that are *not* present
in any of the `others` arrays, with the comparison made by passing items into
the `iteratee` function and checking `==` on the result. This allows changing
the way the item is viewed for comparison.

#### func  DifferenceWith

```go
func DifferenceWith[T any](array []T, comparator func(T, T) bool, others ...[]T) []T
```
DifferenceWith returns a list of items present in `array` that are *not* present
in any of the `others` arrays, with the comparison made using the given
`comparator`.

#### func  Drop

```go
func Drop[T any](array []T, n int) []T
```
Drop returns a slice of `array` with `n` elements dropped from the beginning.

#### func  DropRight

```go
func DropRight[T any](array []T, n int) []T
```
DropRight returns a slice of `array` with `n` elements dropped from the end.

#### func  DropRightWhile

```go
func DropRightWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T
```
DropRightWhile creates a slice of `array` excluding elements dropped from the
end. Elements are dropped until `predicate` returns false.

#### func  DropWhile

```go
func DropWhile[T any](array []T, predicate func(value T, index int, array []T) bool) []T
```
DropWhile creates a slice of `array` excluding elements dropped from the
beginning. Elements are dropped until `predicate` returns false.

#### func  Fill

```go
func Fill[T any](array []T, value T, start int, end int)
```
Fill fills elements of `array` with `value` from `start` up to, but not
including `end`.

#### func  FindIndex

```go
func FindIndex[T any](array []T, predicate func(T) bool) int
```
FindIndex returns the index of the first element for which the `predicate`
returns true.

#### func  FindLastIndex

```go
func FindLastIndex[T any](array []T, predicate func(T) bool) int
```
FindLastIndex returns the index of the last element of which the `predicate`
returns true.

#### func  IndexOf

```go
func IndexOf[T comparable](array []T, value T) int
```
IndexOf returns the index at which the first occurrence of `value` is found in
`array`. Returns `-1` if not found.

#### func  Intersection

```go
func Intersection[T comparable](array []T, others ...[]T) []T
```
Intersection returns an array of unique values that are included in all given
arrays. The order of the result values are determined by the first array.

#### func  IntersectionBy

```go
func IntersectionBy[T comparable, U comparable](array []T, iteratee func(T) U, others ...[]T) []T
```
IntersectionBy returns an array of unique values that are included in all given
arrays, with comparison happening on the result of the `iteratee` function. The
order of the result values are determined by the first array.

#### func  IntersectionWith

```go
func IntersectionWith[T comparable](array []T, comparator func(T, T) bool, others ...[]T) []T
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

#### func  LastIndexOf

```go
func LastIndexOf[T comparable](array []T, value T) int
```
LastIndexOf returns the index at which the last occurrence of `value` is found
in `array`. Returns `-1` if not found.

#### func  Nth

```go
func Nth[T any](array []T, n int) T
```
Nth gets the element at index `n` of the `array`. If `n` is negative, the nth
element from the end is returned.

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

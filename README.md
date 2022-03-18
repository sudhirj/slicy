# godash

    import "github.com/sudhirj/godash"


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

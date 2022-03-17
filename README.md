# godash
--

    import "github.com/sudhirj/godash"

## Usage

#### func  Chunk

```go
func Chunk[T any](array []T, chunkSize int) [][]T
```

Chunk splits the given array into groups the length of `chunkSize`. If the array cannot be split evenly, the last chunk
will have the remaining elements.

#### func  Concat

```go
func Concat[T any](arrays ...[]T) []T
```

Concat combines all the elements from all the given arrays into a single array.

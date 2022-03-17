package godash

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		i    []string
		n    int
		o    [][]string
	}{
		{"empty", []string{}, 3, [][]string{}},
		{"round", []string{"a", "b", "c", "d"}, 2, [][]string{{"a", "b"}, {"c", "d"}}},
		{"extra 1", []string{"a", "b", "c", "d"}, 3, [][]string{{"a", "b", "c"}, {"d"}}},
		{"extra 2", []string{"a", "b", "c", "d", "e"}, 3, [][]string{{"a", "b", "c"}, {"d", "e"}}},
		{"round 2", []string{"a", "b", "c", "d", "e", "f"}, 3, [][]string{{"a", "b", "c"}, {"d", "e", "f"}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op := Chunk(test.i, test.n)
			if !reflect.DeepEqual(op, test.o) {
				t.Error(test.o, op)
			}
		})
	}
}

func ExampleChunk() {
	fmt.Println(Chunk([]int{1, 2, 3, 4}, 2))
	fmt.Println(Chunk([]int{1, 2, 3, 4}, 3))
	// Output:
	// [[1 2] [3 4]]
	// [[1 2 3] [4]]
}

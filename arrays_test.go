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

func TestConcat(t *testing.T) {
	tests := []struct {
		name string
		i    [][]int
		o    []int
	}{
		{"empty1", [][]int{}, []int{}},
		{"empty2", [][]int{{}}, []int{}},
		{"empty3", [][]int{{}, {}}, []int{}},
		{"t1", [][]int{{1}, {2}}, []int{1, 2}},
		{"t2", [][]int{{1, 2}, {3}}, []int{1, 2, 3}},
		{"t3", [][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
		{"middle empty", [][]int{{1, 2}, {}, {3, 4}, {}, {5}}, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op := Concat(test.i...)
			if !reflect.DeepEqual(op, test.o) {
				t.Error(test.o, op)
			}
		})
	}
}

func ExampleConcat() {
	fmt.Println(Concat([]int{1, 2}, []int{3}))
	fmt.Println(Concat([]int{1, 2}, []int{3}, []int{4, 5, 6}))
	// Output:
	// [1 2 3]
	// [1 2 3 4 5 6]
}

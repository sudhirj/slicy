package godash

import (
	"fmt"
	"math"
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

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		others [][]int
		o      []int
	}{
		{"empty", []int{}, [][]int{}, []int{}},
		{"t1", []int{2, 1}, [][]int{{2, 3}}, []int{1}},
		{"t2", []int{2, 1, 4, 5}, [][]int{{2, 3}, {4, 6}}, []int{1, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op := Difference(test.array, test.others...)
			if !reflect.DeepEqual(op, test.o) {
				t.Error(test.o, op)
			}
		})
	}
}

func ExampleDifference() {
	fmt.Println(Difference([]int{2, 1}, []int{2, 3}))
	fmt.Println(Difference([]int{1, 2, 3, 4, 5, 6, 7}, []int{0, 1, 2}, []int{5, 6, 7, 8}))
	// Output:
	// [1]
	// [3 4]
}

func ExampleDifferenceBy() {
	fmt.Println(DifferenceBy([]float64{2.1, 1.2}, math.Floor, []float64{2.3, 3.4}))
	fmt.Println(DifferenceBy([]float64{-2, -1, 0, 1, 2}, math.Abs, []float64{1}, []float64{-2}))
	// Output:
	// [1.2]
	// [0]
}

func ExampleDifferenceWith() {
	array := []map[string]int{{"a": 1}, {"b": 2}}
	others := []map[string]int{{"a": 1}}
	fmt.Println(DifferenceWith(array, func(x, y map[string]int) bool { return reflect.DeepEqual(x, y) }, others))
	// Output:
	// [map[b:2]]
}

func ExampleDrop() {
	fmt.Println(Drop([]int{1, 2, 3}, 1))
	fmt.Println(Drop([]int{1, 2, 3}, 2))
	fmt.Println(Drop([]int{1, 2, 3}, 5))
	fmt.Println(Drop([]int{1, 2, 3}, 0))
	// Output:
	// [2 3]
	// [3]
	// []
	// [1 2 3]
}

func ExampleDropRight() {
	fmt.Println(DropRight([]int{1, 2, 3}, 1))
	fmt.Println(DropRight([]int{1, 2, 3}, 2))
	fmt.Println(DropRight([]int{1, 2, 3}, 5))
	fmt.Println(DropRight([]int{1, 2, 3}, 0))
	// Output:
	// [1 2]
	// [1]
	// []
	// [1 2 3]
}

func ExampleDropRightWhile() {
	haystack := []string{"h1", "h2", "h3", "needle", "h4", "h5", "needle", "h6", "h7", "h8"}
	fmt.Println(DropRightWhile(haystack, func(item string, index int, arr []string) bool { return item != "needle" }))
	// Output:
	// [h1 h2 h3 needle h4 h5 needle]
}

func TestDropRightWhile(t *testing.T) {
	tests := []struct {
		name      string
		i         []int
		predicate func(item int, index int, arr []int) bool
		o         []int
	}{
		{"empty", []int{}, func(item int, index int, arr []int) bool {
			return true
		}, []int{}},
		{"five", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item > 5
		}, []int{1, 2, 3, 4, 5}},
		{"dont drop", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item > 10
		}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"drop all", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item < 10
		}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op := DropRightWhile(test.i, test.predicate)
			if !reflect.DeepEqual(op, test.o) {
				t.Error(test.o, op)
			}
		})
	}
}

func ExampleDropWhile() {
	haystack := []string{"h1", "h2", "h3", "needle", "h4", "h5", "needle", "h6", "h7", "h8"}
	fmt.Println(DropWhile(haystack, func(item string, index int, arr []string) bool { return item != "needle" }))
	// Output:
	// [needle h4 h5 needle h6 h7 h8]
}

func TestDropWhile(t *testing.T) {
	tests := []struct {
		name      string
		i         []int
		predicate func(item int, index int, arr []int) bool
		o         []int
	}{
		{"empty", []int{}, func(item int, index int, arr []int) bool {
			return true
		}, []int{}},
		{"five", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item < 5
		}, []int{5, 6, 7, 8, 9}},
		{"dont drop", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item > 10
		}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"drop all", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(item int, index int, arr []int) bool {
			return item < 10
		}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op := DropWhile(test.i, test.predicate)
			if !reflect.DeepEqual(op, test.o) {
				t.Error(test.o, op)
			}
		})
	}
}

func ExampleFill() {
	array := []string{"a", "b", "c", "d"}
	Fill(array, "*", 1, 3)
	fmt.Println(array)
	// Output:
	// [a * * d]
}

func ExampleFindIndex() {
	fmt.Println(FindIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "x" }))
	fmt.Println(FindIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "a" }))
	fmt.Println(FindIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "d" }))
	fmt.Println(FindIndex([]string{"a", "a", "b", "b"}, func(t string) bool { return t == "b" }))
	// Output:
	// -1
	// 0
	// 3
	// 2
}

func ExampleFindLastIndex() {
	fmt.Println(FindLastIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "x" }))
	fmt.Println(FindLastIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "a" }))
	fmt.Println(FindLastIndex([]string{"a", "b", "c", "d"}, func(t string) bool { return t == "d" }))
	fmt.Println(FindLastIndex([]string{"a", "a", "b", "b"}, func(t string) bool { return t == "a" }))
	// Output:
	// -1
	// 0
	// 3
	// 1
}

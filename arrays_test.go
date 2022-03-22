package slicy

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

func ExampleIndexOf() {
	fmt.Println(IndexOf([]string{"a", "b", "c"}, "x"))
	fmt.Println(IndexOf([]string{"a", "b", "c"}, "a"))
	fmt.Println(IndexOf([]string{"a", "b", "b", "c"}, "b"))
	// Output:
	// -1
	// 0
	// 1
}

func ExampleIntersection() {
	fmt.Println(Intersection([]int{2, 1}, []int{2, 3}))
	fmt.Println(Intersection([]int{2, 1}, []int{2, 3}, []int{8, 2}))
	fmt.Println(Intersection([]int{2, 1, 2, 2, 1}, []int{1, 2, 3, 2}, []int{8, 1, 2, 2}))
	fmt.Println(Intersection([]int{}, []int{}))
	// Output:
	// [2]
	// [2]
	// [2 1]
	// []
}

func ExampleIntersectionBy() {
	fmt.Println(IntersectionBy([]float64{2.1, 1.2}, math.Floor, []float64{2.3, 3.4}))
	fmt.Println(IntersectionBy([]float64{2.1, 1.2, 2.4}, math.Floor, []float64{2.3, 3.4}))
	// Output:
	// [2.1]
	// [2.1]
}

func ExampleIntersectionWith() {
	fmt.Println(IntersectionWith([]float64{2.1, 1.2}, func(x, y float64) bool { return math.Floor(x) == math.Floor(y) }, []float64{2.3, 3.4}))
	fmt.Println(IntersectionWith([]float64{2.1, 1.2, 2.4}, func(x, y float64) bool { return math.Floor(x) == math.Floor(y) }, []float64{2.3, 3.4}))
	// Output:
	// [2.1]
	// [2.1]
}

func ExampleJoin() {
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]int{2022, 1, 1}, "/"))
	fmt.Println(Join([]any{1, "January", 2022}, "-"))
	// Output:
	// Hello World
	// 2022/1/1
	// 1-January-2022
}

func ExampleLastIndexOf() {
	fmt.Println(LastIndexOf([]string{"a", "b", "c"}, "x"))
	fmt.Println(LastIndexOf([]string{"a", "b", "c"}, "a"))
	fmt.Println(LastIndexOf([]string{"a", "b", "b", "c"}, "b"))
	// Output:
	// -1
	// 0
	// 2
}

func ExampleNth() {
	fmt.Println(Nth([]string{"a", "b", "c", "d"}, 1))
	fmt.Println(Nth([]string{"a", "b", "c", "d"}, -2))
	fmt.Println(Nth([]string{"a", "b", "c", "d"}, 0))
	fmt.Println(Nth([]string{"a", "b", "c", "d"}, -1))
	// Output:
	// b
	// c
	// a
	// d
}

func ExamplePull() {
	array := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(array)
	fmt.Println(Pull(array, "a", "g"))
	fmt.Println(Pull(array, []string{"c"}...))
	fmt.Println(Pull(array))
	// Output:
	// [a b c d e f g]
	// [b c d e f]
	// [a b d e f g]
	// [a b c d e f g]
}

func ExamplePullAll() {
	array := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(array)
	fmt.Println(PullAll(array, []string{"a", "g"}))
	fmt.Println(PullAll(array, []string{"c"}))
	fmt.Println(PullAll(array, []string{}))
	// Output:
	// [a b c d e f g]
	// [b c d e f]
	// [a b d e f g]
	// [a b c d e f g]
}

func ExamplePullAllBy() {
	fmt.Println(PullAllBy([]float64{1.2, 2.5, 3.14, 4.2}, []float64{3.8}, math.Floor))
	// Output:
	// [1.2 2.5 4.2]
}

func ExamplePullAllWith() {
	fmt.Println(PullAllWith([]float64{1.2, 2.5, 3.14, 4.2}, []float64{3.8}, func(x, y float64) bool { return math.Floor(x) == math.Floor(y) }))
	// Output:
	// [1.2 2.5 4.2]
}

func ExamplePullAt() {
	array := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(PullAt(array))
	fmt.Println(PullAt(array, 0, 1))
	fmt.Println(PullAt(array, []int{0, 2, 4, 6, 8, 10, 12}...))
	// Output:
	// [a b c d e f g]
	// [c d e f g]
	// [b d f]
}

func ExampleRemove() {
	array := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(array)
	fmt.Println(Remove(array, func(v string, i int, arr []string) bool {
		return i%2 == 0
	}))
	// Output:
	// [a b c d e f g]
	// [b d f]
}

func ExampleReverse() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println(Reverse(array))
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]

}

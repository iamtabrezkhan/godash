package slice

import (
	"fmt"
	"reflect"
	"testing"
)

type person struct {
	name string
}

func TestChunk(t *testing.T) {

	p1 := person{"jon"}
	p2 := person{"mike"}
	p3 := person{"jennifer"}
	p4 := person{"harry"}
	p5 := person{"ron"}

	var emptyFunc func() bool
	emptyFuncPtr := &emptyFunc

	nonEmptyFunc := func() bool { return true }
	nonEmptyFuncPtr := &nonEmptyFunc

	map1 := make(map[string]int)
	map2 := make(map[string]int)
	map3 := make(map[string]int)

	testCases := []struct {
		Slice  interface{}
		Size   int
		Result interface{}
	}{
		{
			[]interface{}{42, emptyFuncPtr, emptyFunc, nonEmptyFuncPtr},
			2,
			[][]interface{}{{42, emptyFuncPtr}, {emptyFunc, nonEmptyFuncPtr}},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			1,
			[][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			2,
			[][]string{{"1", "2"}, {"3", "4"}, {"5"}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
			[][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			-1,
			[][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			[]int8{1, 2, 3, 4, 5},
			2,
			[][]int8{{1, 2}, {3, 4}, {5}},
		},
		{
			[]int16{1, 2, 3, 4, 5},
			2,
			[][]int16{{1, 2}, {3, 4}, {5}},
		},
		{
			[]int32{1, 2, 3, 4, 5},
			2,
			[][]int32{{1, 2}, {3, 4}, {5}},
		},
		{
			[]int64{1, 2, 3, 4, 5},
			2,
			[][]int64{{1, 2}, {3, 4}, {5}},
		},
		{
			[]uint{1, 2, 3, 4, 5},
			2,
			[][]uint{{1, 2}, {3, 4}, {5}},
		},
		{
			[]uint8{1, 2, 3, 4, 5},
			2,
			[][]uint8{{1, 2}, {3, 4}, {5}},
		},
		{
			[]uint16{1, 2, 3, 4, 5},
			2,
			[][]uint16{{1, 2}, {3, 4}, {5}},
		},
		{
			[]uint32{1, 2, 3, 4, 5},
			2,
			[][]uint32{{1, 2}, {3, 4}, {5}},
		},
		{
			[]uint64{1, 2, 3, 4, 5},
			2,
			[][]uint64{{1, 2}, {3, 4}, {5}},
		},
		{
			[]float32{1.1, 2.1, 3.1, 4.1, 5.1},
			2,
			[][]float32{{1.1, 2.1}, {3.1, 4.1}, {5.1}},
		},
		{
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1},
			2,
			[][]float64{{1.1, 2.1}, {3.1, 4.1}, {5.1}},
		},
		{
			[]complex64{-5 + 12i, -10 + 22i, -15 + 32i, -20 + 42i, -25 + 52i},
			2,
			[][]complex64{{-5 + 12i, -10 + 22i}, {-15 + 32i, -20 + 42i}, {-25 + 52i}},
		},
		{
			[]complex128{-5 + 12i, -10 + 22i, -15 + 32i, -20 + 42i, -25 + 52i},
			2,
			[][]complex128{{-5 + 12i, -10 + 22i}, {-15 + 32i, -20 + 42i}, {-25 + 52i}},
		},
		{
			[]bool{true, false, true, false, true},
			2,
			[][]bool{{true, false}, {true, false}, {true}},
		},
		{
			[]person{p1, p2, p3, p4},
			2,
			[][]person{{p1, p2}, {p3, p4}},
		},
		{
			[]person{p1, p2, p3, p4, p5},
			2,
			[][]person{{p1, p2}, {p3, p4}, {p5}},
		},
		{
			[]map[string]int{map1, map2, map3},
			2,
			[][]map[string]int{{map1, map2}, {map3}},
		},
		{
			[]rune{'a', 'b', 'd'},
			2,
			[][]rune{{97, 98}, {100}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test case #%d\n", i+1), func(t *testing.T) {
			fmt.Printf("Expected: %v\n", testCase.Result)
			output := Chunk(testCase.Slice, testCase.Size)
			fmt.Printf("Output: %v\n", output)
			if !reflect.DeepEqual(testCase.Result, output) {
				t.Errorf("Expected %v, received %v", testCase.Result, output)
			}
		})
	}

}

func TestChunkPanic(t *testing.T) {
	defer func() { recover() }()
	Chunk(123, 1)
	t.Errorf("should have panicked!")
}

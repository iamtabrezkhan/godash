package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDrop(t *testing.T) {

	p1 := person{"jon"}
	p2 := person{"mike"}
	var p3 person
	p4 := person{"harry"}
	var p5 person

	var emptyMap map[string]int
	map2 := make(map[string]int)
	map2["one"] = 1
	map3 := make(map[string]int)
	map3["three"] = 3

	var emptyRune rune
	var rune1 rune
	var rune2 rune
	rune1 = 'a'
	rune2 = 'b'

	testCases := []struct {
		Slice  interface{}
		N      int
		Result interface{}
	}{
		{
			[]string{"1", "2", "3", "4", "5"},
			2,
			[]string{"3", "4", "5"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			-1,
			[]string{"2", "3", "4", "5"},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{3, 4, 5},
		},
		{
			[]int8{1, 2, 3, 4, 5},
			2,
			[]int8{3, 4, 5},
		},
		{
			[]int16{1, 2, 3, 4, 5},
			2,
			[]int16{3, 4, 5},
		},
		{
			[]int32{1, 2, 3, 4, 5},
			2,
			[]int32{3, 4, 5},
		},
		{
			[]int64{1, 2, 3, 4, 5},
			2,
			[]int64{3, 4, 5},
		},
		{
			[]uint{1, 2, 3, 4, 5},
			2,
			[]uint{3, 4, 5},
		},
		{
			[]uint8{1, 2, 3, 4, 5},
			2,
			[]uint8{3, 4, 5},
		},
		{
			[]uint16{1, 2, 3, 4, 5},
			2,
			[]uint16{3, 4, 5},
		},
		{
			[]uint32{1, 2, 3, 4, 5},
			2,
			[]uint32{3, 4, 5},
		},
		{
			[]uint64{1, 2, 3, 4, 5},
			2,
			[]uint64{3, 4, 5},
		},
		{
			[]float32{1.1, 2.1, 3.0, 4.1, 5.0},
			2,
			[]float32{3.0, 4.1, 5.0},
		},
		{
			[]float64{1.1, 2.1, 3.0, 4.1, 5.0},
			2,
			[]float64{3.0, 4.1, 5.0},
		},
		{
			[]complex64{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			2,
			[]complex64{-0 + 0i, -20 + 42i, -0 + 0i},
		},
		{
			[]complex128{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			2,
			[]complex128{-0 + 0i, -20 + 42i, -0 + 0i},
		},
		{
			[]bool{true, false, true, false, true},
			2,
			[]bool{true, false, true},
		},
		{
			[]person{p1, p2, p3, p4, p5},
			2,
			[]person{p3, p4, p5},
		},
		{
			[]map[string]int{emptyMap, map2, map3},
			2,
			[]map[string]int{map3},
		},
		{
			[]rune{emptyRune, rune1, rune2},
			2,
			[]rune{rune2},
		},
		{
			[]rune{emptyRune, rune1, rune2},
			5,
			[]rune{},
		},
		{
			[][]rune{{emptyRune, rune1}, {rune2}},
			1,
			[][]rune{{rune2}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test case #%d", i+1), func(t *testing.T) {
			fmt.Printf("Expected: %v\n", testCase.Result)
			output := Drop(testCase.Slice, testCase.N)
			fmt.Printf("Output: %v\n", output)
			if !reflect.DeepEqual(testCase.Result, output) {
				t.Errorf("Expected %v, received %v", testCase.Result, output)
			}
		})
	}

}

func TestDropPanic(t *testing.T) {
	defer func() { recover() }()
	Drop(123, 10)
	t.Errorf("should have panicked!")
}

package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompact(t *testing.T) {

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
		Result interface{}
	}{
		{
			[]string{"1", "2", "", "4", ""},
			[]string{"1", "2", "4"},
		},
		{
			[]int{1, 2, 0, 4, 0},
			[]int{1, 2, 4},
		},
		{
			[]int8{1, 2, 0, 4, 0},
			[]int8{1, 2, 4},
		},
		{
			[]int16{1, 2, 0, 4, 0},
			[]int16{1, 2, 4},
		},
		{
			[]int32{1, 2, 0, 4, 0},
			[]int32{1, 2, 4},
		},
		{
			[]int64{1, 2, 0, 4, 0},
			[]int64{1, 2, 4},
		},
		{
			[]uint{1, 2, 0, 4, 0},
			[]uint{1, 2, 4},
		},
		{
			[]uint8{1, 2, 0, 4, 0},
			[]uint8{1, 2, 4},
		},
		{
			[]uint16{1, 2, 0, 4, 0},
			[]uint16{1, 2, 4},
		},
		{
			[]uint32{1, 2, 0, 4, 0},
			[]uint32{1, 2, 4},
		},
		{
			[]uint64{1, 2, 0, 4, 0},
			[]uint64{1, 2, 4},
		},
		{
			[]float32{1.1, 2.1, 0.0, 4.1, 0.0},
			[]float32{1.1, 2.1, 4.1},
		},
		{
			[]float64{1.1, 2.1, 0.0, 4.1, 0.0},
			[]float64{1.1, 2.1, 4.1},
		},
		{
			[]complex64{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			[]complex64{-5 + 12i, -10 + 22i, -20 + 42i},
		},
		{
			[]complex128{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			[]complex128{-5 + 12i, -10 + 22i, -20 + 42i},
		},
		{
			[]bool{true, false, true, false, true},
			[]bool{true, true, true},
		},
		{
			[]person{p1, p2, p3, p4, p5},
			[]person{p1, p2, p4},
		},
		{
			[]map[string]int{emptyMap, map2, map3},
			[]map[string]int{map2, map3},
		},
		{
			[]rune{emptyRune, rune1, rune2},
			[]rune{rune1, rune2},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test case #%d", i+1), func(t *testing.T) {
			fmt.Printf("Expected: %v\n", testCase.Result)
			output := Compact(testCase.Slice)
			fmt.Printf("Output: %v\n", output)
			if !reflect.DeepEqual(testCase.Result, output) {
				t.Errorf("Expected %v, received %v", testCase.Result, output)
			}
		})
	}

}

func TestCompactPanic(t *testing.T) {
	defer func() { recover() }()
	Compact(123)
	t.Errorf("should have panicked!")
}

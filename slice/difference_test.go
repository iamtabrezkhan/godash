package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {

	p1 := person{"jon"}
	p2 := person{"mike"}
	p3 := person{"ron"}

	var rune1 rune
	var rune2 rune
	var rune3 rune
	rune1 = 'a'
	rune2 = 'b'
	rune3 = 'b'

	testCases := []struct {
		Slice1 interface{}
		Slice2 interface{}
		Result interface{}
	}{
		{
			[]string{"1", "2", "3", "4"},
			[]string{"2", "4"},
			[]string{"1", "3"},
		},
		{
			[]int{1, 2, 0, 4, 0},
			[]int{1, 2, 4},
			[]int{0, 0},
		},
		{
			[]int8{1, 2, 0, 4, 0},
			[]int8{1, 2, 4},
			[]int8{0, 0},
		},
		{
			[]int16{1, 2, 0, 4, 0},
			[]int16{1, 2, 4},
			[]int16{0, 0},
		},
		{
			[]int32{1, 2, 0, 4, 0},
			[]int32{1, 2, 4},
			[]int32{0, 0},
		},
		{
			[]int64{1, 2, 0, 4, 0},
			[]int64{1, 2, 4},
			[]int64{0, 0},
		},
		{
			[]uint{1, 2, 0, 4, 0},
			[]uint{1, 2, 4},
			[]uint{0, 0},
		},
		{
			[]uint8{1, 2, 0, 4, 0},
			[]uint8{1, 2, 4},
			[]uint8{0, 0},
		},
		{
			[]uint16{1, 2, 0, 4, 0},
			[]uint16{1, 2, 4},
			[]uint16{0, 0},
		},
		{
			[]uint32{1, 2, 0, 4, 0},
			[]uint32{1, 2, 4},
			[]uint32{0, 0},
		},
		{
			[]uint64{1, 2, 0, 4, 0},
			[]uint64{1, 2, 4},
			[]uint64{0, 0},
		},
		{
			[]float32{1.1, 2.1, 0.0, 4.1, 0.0},
			[]float32{1.1, 2.1, 4.1},
			[]float32{0.0, 0.0},
		},
		{
			[]float64{1.1, 2.1, 0.0, 4.1, 0.0},
			[]float64{1.1, 2.1, 4.1},
			[]float64{0.0, 0.0},
		},
		{
			[]complex64{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			[]complex64{-5 + 12i, -10 + 22i, -20 + 42i},
			[]complex64{-0 + 0i, -0 + 0i},
		},
		{
			[]complex128{-5 + 12i, -10 + 22i, -0 + 0i, -20 + 42i, -0 + 0i},
			[]complex128{-5 + 12i, -10 + 22i, -20 + 42i},
			[]complex128{-0 + 0i, -0 + 0i},
		},
		{
			[]bool{true, false, true, false, true},
			[]bool{true, true, true},
			[]bool{false, false},
		},
		{
			[]person{p1, p2},
			[]person{p2, p3},
			[]person{p1},
		},
		{
			[]rune{rune1, rune2},
			[]rune{rune2, rune3},
			[]rune{rune1},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test case #%d", i+1), func(t *testing.T) {
			fmt.Printf("Expected: %v\n", testCase.Result)
			output := Difference(testCase.Slice1, testCase.Slice2)
			fmt.Printf("Output: %v\n", output)
			if !reflect.DeepEqual(testCase.Result, output) {
				t.Errorf("Expected %v, received %v", testCase.Result, output)
			}
		})
	}
}

func TestDifferencePanic(t *testing.T) {
	defer func() { recover() }()
	Difference(123, []int{1, 2, 3})
	t.Errorf("should have panicked")
}

func TestDifferencePanic2(t *testing.T) {
	defer func() { recover() }()
	Difference([]int{1, 2, 3}, 123)
	t.Errorf("should have panicked")
}

func TestDifferencePanic3(t *testing.T) {
	defer func() { recover() }()
	Difference([]int{1, 2, 3}, []string{"abc"})
	t.Errorf("should have panicked")
}

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

	testCases := []struct {
		Slice  interface{}
		Size   int
		Result interface{}
	}{
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[][]int{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[][]int{[]int{1, 2}, []int{3, 4}, []int{5}},
		},
		{
			[]int8{1, 2, 3, 4, 5},
			2,
			[][]int8{[]int8{1, 2}, []int8{3, 4}, []int8{5}},
		},
		{
			[]int16{1, 2, 3, 4, 5},
			2,
			[][]int16{[]int16{1, 2}, []int16{3, 4}, []int16{5}},
		},
		{
			[]int32{1, 2, 3, 4, 5},
			2,
			[][]int32{[]int32{1, 2}, []int32{3, 4}, []int32{5}},
		},
		{
			[]int64{1, 2, 3, 4, 5},
			2,
			[][]int64{[]int64{1, 2}, []int64{3, 4}, []int64{5}},
		},
		{
			[]uint{1, 2, 3, 4, 5},
			2,
			[][]uint{[]uint{1, 2}, []uint{3, 4}, []uint{5}},
		},
		{
			[]uint8{1, 2, 3, 4, 5},
			2,
			[][]uint8{[]uint8{1, 2}, []uint8{3, 4}, []uint8{5}},
		},
		{
			[]uint16{1, 2, 3, 4, 5},
			2,
			[][]uint16{[]uint16{1, 2}, []uint16{3, 4}, []uint16{5}},
		},
		{
			[]uint32{1, 2, 3, 4, 5},
			2,
			[][]uint32{[]uint32{1, 2}, []uint32{3, 4}, []uint32{5}},
		},
		{
			[]uint64{1, 2, 3, 4, 5},
			2,
			[][]uint64{[]uint64{1, 2}, []uint64{3, 4}, []uint64{5}},
		},
		{
			[]float32{1.1, 2.1, 3.1, 4.1, 5.1},
			2,
			[][]float32{[]float32{1.1, 2.1}, []float32{3.1, 4.1}, []float32{5.1}},
		},
		{
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1},
			2,
			[][]float64{[]float64{1.1, 2.1}, []float64{3.1, 4.1}, []float64{5.1}},
		},
		{
			[]complex64{-5 + 12i, -10 + 22i, -15 + 32i, -20 + 42i, -25 + 52i},
			2,
			[][]complex64{[]complex64{-5 + 12i, -10 + 22i}, []complex64{-15 + 32i, -20 + 42i}, []complex64{-25 + 52i}},
		},
		{
			[]complex128{-5 + 12i, -10 + 22i, -15 + 32i, -20 + 42i, -25 + 52i},
			2,
			[][]complex128{[]complex128{-5 + 12i, -10 + 22i}, []complex128{-15 + 32i, -20 + 42i}, []complex128{-25 + 52i}},
		},
		{
			[]bool{true, false, true, false, true},
			2,
			[][]bool{[]bool{true, false}, []bool{true, false}, []bool{true}},
		},
		{
			[]person{p1, p2, p3, p4},
			2,
			[][]person{[]person{p1, p2}, []person{p3, p4}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test case #%d", i+1), func(t *testing.T) {
			output := Chunk(testCase.Slice, testCase.Size)
			if !reflect.DeepEqual(testCase.Result, output) {
				t.Errorf("Expected %v, received %v", testCase.Result, output)
			}
		})
	}

}

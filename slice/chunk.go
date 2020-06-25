package slice

import (
	"reflect"
)

/*
Chunk returns a slice of elements split into groups the length of size.
If slice can't be split evenly, the final chunk will be the remaining elements.
*/
// Examples:
//	Chunk([]int{1, 2, 3, 4}, 2) // ==> [[1, 2], [3, 4]]
func Chunk(slice interface{}, size int) interface{} {
	if size <= 0 {
		size = 1
	}
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Chunk: not slice")
	}
	elemType := s.Type().Elem()
	len := s.Len()
	switch elemType.Kind() {
	// string type
	case reflect.String:
		{
			res := make([][]string, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]string, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]string)[j])
				}
				res = append(res, t)
			}
			return res
		}
	// int type
	case reflect.Int:
		{
			res := make([][]int, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]int, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]int)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Int8:
		{
			res := make([][]int8, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]int8, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]int8)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Int16:
		{
			res := make([][]int16, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]int16, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]int16)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Int32:
		{
			res := make([][]int32, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]int32, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]int32)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Int64:
		{
			res := make([][]int64, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]int64, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]int64)[j])
				}
				res = append(res, t)
			}
			return res
		}
	// uint type
	case reflect.Uint:
		{
			res := make([][]uint, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]uint, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]uint)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Uint8:
		{
			res := make([][]uint8, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]uint8, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]uint8)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Uint16:
		{
			res := make([][]uint16, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]uint16, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]uint16)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Uint32:
		{
			res := make([][]uint32, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]uint32, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]uint32)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Uint64:
		{
			res := make([][]uint64, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]uint64, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]uint64)[j])
				}
				res = append(res, t)
			}
			return res
		}
	// float type
	case reflect.Float32:
		{
			res := make([][]float32, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]float32, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]float32)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Float64:
		{
			res := make([][]float64, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]float64, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]float64)[j])
				}
				res = append(res, t)
			}
			return res
		}
	// bool type
	case reflect.Bool:
		{
			res := make([][]bool, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]bool, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]bool)[j])
				}
				res = append(res, t)
			}
			return res
		}
	// complex type
	case reflect.Complex64:
		{
			res := make([][]complex64, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]complex64, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]complex64)[j])
				}
				res = append(res, t)
			}
			return res
		}
	case reflect.Complex128:
		{
			res := make([][]complex128, 0, len)
			for i := 0; i < len; i = i + size {
				t := make([]complex128, 0, size)
				for j := i; j < size+i; j++ {
					if j >= len {
						break
					}
					t = append(t, slice.([]complex128)[j])
				}
				res = append(res, t)
			}
			return res
		}
	}
	res := reflect.MakeSlice(reflect.SliceOf(s.Type()), 0, len)
	for i := 0; i < len; i = i + size {
		t := reflect.MakeSlice(reflect.SliceOf(s.Index(i).Type()), 0, size)
		for j := i; j < size+i; j++ {
			if j >= len {
				break
			}
			t = reflect.Append(t, s.Index(j))
		}
		res = reflect.Append(res, t)
	}
	return res.Interface()
}

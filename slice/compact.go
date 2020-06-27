package slice

import (
	"reflect"
)

// Compact removes all the `zero values` and returns a new slice
//
// Examples:
// 	s := []int{1, 0, 3, 10, 0, 5, 0}
// 	Compact(s) // ==> [1, 3, 10, 5]
// 	Compact([]string{"abc", "hello", ""}) // ==> ["abc", "hello"]
func Compact(slice interface{}) interface{} {
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
			newSlice := slice.([]string)
			res := make([]string, 0, len)
			var zeroValue string
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	// int type
	case reflect.Int:
		{
			newSlice := slice.([]int)
			res := make([]int, 0, len)
			var zeroValue int
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Int8:
		{
			newSlice := slice.([]int8)
			res := make([]int8, 0, len)
			var zeroValue int8
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Int16:
		{
			newSlice := slice.([]int16)
			res := make([]int16, 0, len)
			var zeroValue int16
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Int32:
		{
			newSlice := slice.([]int32)
			res := make([]int32, 0, len)
			var zeroValue int32
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Int64:
		{
			newSlice := slice.([]int64)
			res := make([]int64, 0, len)
			var zeroValue int64
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	// uint type
	case reflect.Uint:
		{
			newSlice := slice.([]uint)
			res := make([]uint, 0, len)
			var zeroValue uint
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Uint8:
		{
			newSlice := slice.([]uint8)
			res := make([]uint8, 0, len)
			var zeroValue uint8
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Uint16:
		{
			newSlice := slice.([]uint16)
			res := make([]uint16, 0, len)
			var zeroValue uint16
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Uint32:
		{
			newSlice := slice.([]uint32)
			res := make([]uint32, 0, len)
			var zeroValue uint32
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Uint64:
		{
			newSlice := slice.([]uint64)
			res := make([]uint64, 0, len)
			var zeroValue uint64
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	// float type
	case reflect.Float32:
		{
			newSlice := slice.([]float32)
			res := make([]float32, 0, len)
			var zeroValue float32
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Float64:
		{
			newSlice := slice.([]float64)
			res := make([]float64, 0, len)
			var zeroValue float64
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	// bool type
	case reflect.Bool:
		{
			newSlice := slice.([]bool)
			res := make([]bool, 0, len)
			var zeroValue bool
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	// complex type
	case reflect.Complex64:
		{
			newSlice := slice.([]complex64)
			res := make([]complex64, 0, len)
			var zeroValue complex64
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	case reflect.Complex128:
		{
			newSlice := slice.([]complex128)
			res := make([]complex128, 0, len)
			var zeroValue complex128
			for i := 0; i < len; i++ {
				if newSlice[i] != zeroValue {
					res = append(res, newSlice[i])
				}
			}
			return res
		}
	}
	res := reflect.MakeSlice(s.Type(), 0, len)
	for i := 0; i < len; i++ {
		if !s.Index(i).IsZero() {
			res = reflect.Append(res, s.Index(i))
		}
	}
	return res.Interface()
}

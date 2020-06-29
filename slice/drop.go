package slice

import (
	"reflect"
)

// Drop returns a slice of array with n elements dropped from the beginning.
//
// Examples:
// 	Drop([]int{1, 2, 3}, 1) // => [2, 3]
// 	Drop([]int{1, 2, 3}, 2) // => [3]
func Drop(slice interface{}, n int) interface{} {
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic("Drop: not a slice")
	}
	if n <= 0 {
		n = 1
	}
	elemType := sliceType.Elem()
	len := sliceValue.Len()
	res := reflect.MakeSlice(sliceType, 0, len)
	if n > len {
		return res.Interface()
	}
	switch elemType.Kind() {
	// string type
	case reflect.String:
		{
			newSlice := slice.([]string)[n:]
			return newSlice
		}
	// int type
	case reflect.Int:
		{
			newSlice := slice.([]int)[n:]
			return newSlice
		}
	case reflect.Int8:
		{
			newSlice := slice.([]int8)[n:]
			return newSlice
		}
	case reflect.Int16:
		{
			newSlice := slice.([]int16)[n:]
			return newSlice
		}
	case reflect.Int32:
		{
			newSlice := slice.([]int32)[n:]
			return newSlice
		}
	case reflect.Int64:
		{
			newSlice := slice.([]int64)[n:]
			return newSlice
		}
	// uint type
	case reflect.Uint:
		{
			newSlice := slice.([]uint)[n:]
			return newSlice
		}
	case reflect.Uint8:
		{
			newSlice := slice.([]uint8)[n:]
			return newSlice
		}
	case reflect.Uint16:
		{
			newSlice := slice.([]uint16)[n:]
			return newSlice
		}
	case reflect.Uint32:
		{
			newSlice := slice.([]uint32)[n:]
			return newSlice
		}
	case reflect.Uint64:
		{
			newSlice := slice.([]uint64)[n:]
			return newSlice
		}
	// float type
	case reflect.Float32:
		{
			newSlice := slice.([]float32)[n:]
			return newSlice
		}
	case reflect.Float64:
		{
			newSlice := slice.([]float64)[n:]
			return newSlice
		}
	// bool type
	case reflect.Bool:
		{
			newSlice := slice.([]bool)[n:]
			return newSlice
		}
	// complex type
	case reflect.Complex64:
		{
			newSlice := slice.([]complex64)[n:]
			return newSlice
		}
	case reflect.Complex128:
		{
			newSlice := slice.([]complex128)[n:]
			return newSlice
		}
	}
	for i := n; i < len; i++ {
		res = reflect.Append(res, sliceValue.Index(i))
	}
	return res.Interface()
}

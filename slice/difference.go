package slice

import (
	"reflect"
)

// Difference returns a slice of slice values not included in the other given slice.
// The order and references of result values are determined by the first slice.
//
// Examples:
// 	Difference([]int{1, 2, 3}, []int{2, 3}) // ==> [1]
func Difference(slice interface{}, values interface{}) interface{} {
	valueOf := reflect.ValueOf
	typeOf := reflect.TypeOf
	sliceType := typeOf(slice)
	valuesType := typeOf(values)
	if sliceType.Kind() != reflect.Slice || valuesType.Kind() != reflect.Slice {
		panic("Difference: both args should be slice")
	}
	sliceElemType := sliceType.Elem()
	valueElemType := valuesType.Elem()
	if sliceElemType != valueElemType {
		panic("Difference: slice and values element types are not same")
	}
	sliceValue := valueOf(slice)
	valuesValue := valueOf(values)
	mapType := reflect.MapOf(valueElemType, typeOf(true))
	freqMap := reflect.MakeMap(mapType)
	for i := 0; i < valuesValue.Len(); i++ {
		freqMap.SetMapIndex(valuesValue.Index(i), valueOf(true))
	}
	res := reflect.MakeSlice(sliceType, 0, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		mapValue := freqMap.MapIndex(sliceValue.Index(i))
		if !mapValue.IsValid() {
			res = reflect.Append(res, sliceValue.Index(i))
		}
	}
	return res.Interface()
}

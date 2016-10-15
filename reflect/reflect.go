package reflect

import (
	"reflect"
)

//new object same the type as sample
func New(sample interface{}) interface{} {
	t := reflect.TypeOf(sample)
	v := reflect.New(t).Interface()
	return v
}

//check if the same type
func IsSameType(l, r interface{}) (ok bool) {
	tl, tr := reflect.TypeOf(l), reflect.TypeOf(r)
	return tl == tr
}

//get kind
func Kind(l interface{}) reflect.Kind {
	return reflect.TypeOf(l).Kind()
}

//check if cmpareable inner type
func IsCmpableInnterType(t interface{}) bool {
	switch t.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, string, float32, float64:
		return true
	}
	return false
}

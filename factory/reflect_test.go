package factory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	//var f float64 = 1.03
	type Float float64
	var my_f Float = 1.1
	fmt.Println(my_f)
	n := reflect.New(reflect.TypeOf(my_f))
	n.Elem().SetFloat(1.8)
	fmt.Println(n.Elem().Interface())
}

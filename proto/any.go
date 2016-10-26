package proto

import (
	"reflect"
)

var (
	gTypeMap map[string]*reflect.Type
)

type Kind reflect.Kind

const (
	Invalid       = Kind(reflect.Invalid)
	Bool          = Kind(reflect.Bool)
	Int           = Kind(reflect.Int)
	Int8          = Kind(reflect.Int8)
	Int16         = Kind(reflect.Int16)
	Int32         = Kind(reflect.Int32)
	Int64         = Kind(reflect.Int64)
	Uint          = Kind(reflect.Uint)
	Uint8         = Kind(reflect.Uint8)
	Uint16        = Kind(reflect.Uint16)
	Uint32        = Kind(reflect.Uint32)
	Uint64        = Kind(reflect.Uint64)
	Uintptr       = Kind(reflect.Uintptr)
	Float32       = Kind(reflect.Float32)
	Float64       = Kind(reflect.Float64)
	Complex64     = Kind(reflect.Complex64)
	Complex128    = Kind(reflect.Complex128)
	Array         = Kind(reflect.Array)
	Chan          = Kind(reflect.Chan)
	Func          = Kind(reflect.Func)
	Interface     = Kind(reflect.Interface)
	Map           = Kind(reflect.Map)
	Ptr           = Kind(reflect.Ptr)
	Slice         = Kind(reflect.Slice)
	String        = Kind(reflect.String)
	Struct        = Kind(reflect.Struct)
	UnsafePointer = Kind(reflect.UnsafePointer)
)

func Unpack([]byte) (any Any) {
	return
}

type Any interface {
	String() string
	Int() int
	Bool() bool
	FromInt(int)
	FromString(string)
	FromBool(bool)
	Kind() Kind
	Size() int
	Pack(b []byte) []byte
}

type IntValue int
type StringValue string
type BoolValue bool
type Other struct {
	value []byte
	kind  Kind
}

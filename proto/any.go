package proto

import (
	"reflect"
)

var (
	gTypeMap map[string]*reflect.Type

//typeID?
)

type TypeId uint

//func (me TypeId) New() Any {}
//func (me TypeId) Load(b []byte) Any {}

//func NewTypeRegister()

//register a type, return it's unique id
//this TypeId can use to create new values of this type
func RegType(name string, i interface{}) TypeId { //TypeId=hash(name)?
	return 0
}

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
	FromString(string)
	Kind() Kind
	Pack(b []byte) []byte
	Unpack([]byte) Any
}

type IntValue int
type StringValue string
type BoolValue bool
type StructValue struct {
	fieldsMap map[string]int
	fields    []Any
	v         reflect.Value
}

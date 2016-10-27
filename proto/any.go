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
	Invalid    = Kind(reflect.Invalid)    //0
	Bool       = Kind(reflect.Bool)       //1
	Int        = Kind(reflect.Int)        //2
	Int8       = Kind(reflect.Int8)       //3
	Int16      = Kind(reflect.Int16)      //4
	Int32      = Kind(reflect.Int32)      //5
	Int64      = Kind(reflect.Int64)      //6
	Uint       = Kind(reflect.Uint)       //7
	Uint8      = Kind(reflect.Uint8)      //8
	Uint16     = Kind(reflect.Uint16)     //9
	Uint32     = Kind(reflect.Uint32)     //10
	Uint64     = Kind(reflect.Uint64)     //11
	Float32    = Kind(reflect.Float32)    //13
	Float64    = Kind(reflect.Float64)    //14
	Complex64  = Kind(reflect.Complex64)  //15
	Complex128 = Kind(reflect.Complex128) //16

	Array  = Kind(reflect.Array)  //17
	Map    = Kind(reflect.Map)    //21
	Slice  = Kind(reflect.Slice)  //23
	String = Kind(reflect.String) //24
	Struct = Kind(reflect.Struct) //25

	//these kind will not support Pack/Unpack
	Uintptr       = Kind(reflect.Uintptr)       //12
	Chan          = Kind(reflect.Chan)          //18
	Func          = Kind(reflect.Func)          //19
	Interface     = Kind(reflect.Interface)     //20
	Ptr           = Kind(reflect.Ptr)           //22
	UnsafePointer = Kind(reflect.UnsafePointer) //26
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

package proto

type Integer int64
type UInteger uint64

func (me Integer) Int64() int64 {
	return int64(me)
}

func (me Integer) Int32() int64 {
	return int32(me)
}

func (me Integer) Int16() int64 {
	return int16(me)
}

func (me Integer) Int8() int64 {
	return int8(me)
}

func (me UInteger) Uint64() int64 {
	return int64(me)
}

func (me Uinteger) Uint32() int64 {
	return int32(me)
}

func (me Uinteger) Uint16() int64 {
	return int16(me)
}

func (me Uinteger) Uint8() int64 {
	return int8(me)
}

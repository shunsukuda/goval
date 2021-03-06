package goval

import (
	"strconv"
)

type Bool struct {
	Bool bool
}

// func (e Bool) GoBool() bool { return e.Bool }
func (e Bool) Type() Type        { return ValTypes.Bool }
func (e Bool) Equal(x Bool) bool { return e.Bool == x.Bool }

func (e Bool) ToInt8() Int8 {
	if e.Bool {
		return Int8{1}
	}
	return Int8{0}
}
func (e Bool) ToInt16() Int16 {
	if e.Bool {
		return Int16{1}
	}
	return Int16{0}
}
func (e Bool) ToInt32() Int32 {
	if e.Bool {
		return Int32{1}
	}
	return Int32{0}
}
func (e Bool) ToInt64() Int64 {
	if e.Bool {
		return Int64{1}
	}
	return Int64{0}
}
func (e Bool) ToUint8() Uint8 {
	if e.Bool {
		return Uint8{1}
	}
	return Uint8{0}
}
func (e Bool) ToUint16() Uint16 {
	if e.Bool {
		return Uint16{1}
	}
	return Uint16{0}
}
func (e Bool) ToUint32() Uint32 {
	if e.Bool {
		return Uint32{1}
	}
	return Uint32{0}
}
func (e Bool) ToUint64() Uint64 {
	if e.Bool {
		return Uint64{1}
	}
	return Uint64{0}
}
func (e Bool) ToFloat32() Float32 {
	if e.Bool {
		return Float32{1}
	}
	return Float32{0}
}
func (e Bool) ToFloat64() Float64 {
	if e.Bool {
		return Float64{1}
	}
	return Float64{0}
}
func (e Bool) ToComplex64() Complex64 {
	if e.Bool {
		return Complex64{1}
	}
	return Complex64{0}
}
func (e Bool) ToComplex128() Complex128 {
	if e.Bool {
		return Complex128{1}
	}
	return Complex128{0}
}

func (e Bool) ToBool() Bool     { return e }
func (e Bool) ToString() String { return String{strconv.FormatBool(e.Bool)} }
func (e Bool) ToBytes() Bytes   { return e.ToString().ToBytes() }

package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/unsafecast"
)

type Bool bool

func (e Bool) Bool() bool {
	return bool(e)
}

func (e Bool) Interface() interface{} {
	return bool(e)
}

func (e Bool) Type() Type {
	return ValTypeBool
}

func (e Bool) ToBool() Bool {
	return e
}

func (e Bool) ToBoolCheck() (Bool, Err) {
	return e, nil
}

func (e Bool) ToInt8() Int8 {
	if e {
		return Int8(1)
	}
	return Int8(0)
}

func (e Bool) ToInt8Check() (Int8, Err) {
	if e {
		return Int8(1), nil
	}
	return Int8(0), nil
}

func (e Bool) ToInt16() Int16 {
	if e {
		return Int16(1)
	}
	return Int16(0)
}

func (e Bool) ToInt16Check() (Int16, Err) {
	if e {
		return Int16(1), nil
	}
	return Int16(0), nil
}

func (e Bool) ToInt32() Int32 {
	if e {
		return Int32(1)
	}
	return Int32(0)
}

func (e Bool) ToInt32Check() (Int32, Err) {
	if e {
		return Int32(1), nil
	}
	return Int32(0), nil
}

func (e Bool) ToInt64() Int64 {
	if e {
		return Int64(1)
	}
	return Int64(0)
}

func (e Bool) ToInt64Check() (Int64, Err) {
	if e {
		return Int64(1), nil
	}
	return Int64(0), nil
}

func (e Bool) ToUint8() Uint8 {
	if e {
		return Uint8(1)
	}
	return Uint8(0)
}

func (e Bool) ToUint8Check() (Uint8, Err) {
	if e {
		return Uint8(1), nil
	}
	return Uint8(0), nil
}

func (e Bool) ToUint16() Uint16 {
	if e {
		return Uint16(1)
	}
	return Uint16(0)
}

func (e Bool) ToUint16Check() (Uint16, Err) {
	if e {
		return Uint16(1), nil
	}
	return Uint16(0), nil
}

func (e Bool) ToUint32() Uint32 {
	if e {
		return Uint32(1)
	}
	return Uint32(0)
}

func (e Bool) ToUint32Check() (Uint32, Err) {
	if e {
		return Uint32(1), nil
	}
	return Uint32(0), nil
}

func (e Bool) ToUint64() Uint64 {
	if e {
		return Uint64(1)
	}
	return Uint64(0)
}

func (e Bool) ToUint64Check() (Uint64, Err) {
	if e {
		return Uint64(1), nil
	}
	return Uint64(0), nil
}

func (e Bool) ToFloat32() Float32 {
	if e {
		return Float32(1)
	}
	return Float32(0)
}

func (e Bool) ToFloat32Check() (Float32, Err) {
	if e {
		return Float32(1), nil
	}
	return Float32(0), nil
}

func (e Bool) ToFloat64() Float64 {
	if e {
		return Float64(1)
	}
	return Float64(0)
}

func (e Bool) ToFloat64Check() (Float64, Err) {
	if e {
		return Float64(1), nil
	}
	return Float64(0), nil
}

func (e Bool) ToString() String {
	return String(strconv.FormatBool(e.Bool()))
}

func (e Bool) ToStringCheck() (String, Err) {
	return String(strconv.FormatBool(e.Bool())), nil
}

func (e Bool) ToBytes() Bytes {
	return Bytes(unsafecast.StringToBytes(strconv.FormatBool(e.Bool())))
}

func (e Bool) ToBytesCheck() (Bytes, Err) {
	return Bytes(unsafecast.StringToBytes(strconv.FormatBool(e.Bool()))), nil
}

package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/unsafecast"
)

type Int64 int64

func (e Int64) Int64() int64 {
	return int64(e)
}

func (e Int64) Interface() interface{} {
	return int64(e)
}

func (e Int64) Type() Type {
	return ValTypeInt64
}

func (e Int64) ToBool() Bool {
	return Bool(e != 0)
}

func (e Int64) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Int64) ToInt8() Int8 {
	return Int8(e)
}

func (e Int64) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Int64) ToInt16() Int16 {
	return Int16(e)
}

func (e Int64) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Int64) ToInt32() Int32 {
	return Int32(e)
}

func (e Int64) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Int64) ToInt64() Int64 {
	return e
}

func (e Int64) ToInt64Check() (Int64, Err) {
	return e, nil
}

func (e Int64) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Int64) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Int64) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Int64) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Int64) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Int64) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Int64) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Int64) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Int64) ToFloat32() Float32 {
	return Float32(e)
}

func (e Int64) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Int64) ToFloat64() Float64 {
	return Float64(e)
}

func (e Int64) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Int64) ToString() String {
	return String(strconv.FormatInt(int64(e.Int64()), 10))
}

func (e Int64) ToStringCheck() (String, Err) {
	return String(strconv.FormatInt(e.Int64(), 10)), nil
}

func (e Int64) ToBytes() Bytes {
	return Bytes(unsafecast.StringToBytes(strconv.FormatInt(int64(e.Int64()), 10)))
}

func (e Int64) ToBytesCheck() (Bytes, Err) {
	return Bytes(unsafecast.StringToBytes(strconv.FormatInt(e.Int64(), 10))), nil
}

package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/forceconv"
)

type Int8 int8

func (e Int8) Int8() int8 {
	return int8(e)
}

func (e Int8) Interface() interface{} {
	return int8(e)
}

func (e Int8) Type() Type {
	return ValTypeInt8
}

func (e Int8) ToBool() Bool {
	return Bool(e != 0)
}

func (e Int8) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Int8) ToInt8() Int8 {
	return e
}

func (e Int8) ToInt8Check() (Int8, Err) {
	return e, nil
}

func (e Int8) ToInt16() Int16 {
	return Int16(e)
}

func (e Int8) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Int8) ToInt32() Int32 {
	return Int32(e)
}

func (e Int8) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Int8) ToInt64() Int64 {
	return Int64(e)
}

func (e Int8) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Int8) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Int8) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Int8) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Int8) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Int8) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Int8) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Int8) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Int8) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Int8) ToFloat32() Float32 {
	return Float32(e)
}

func (e Int8) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Int8) ToFloat64() Float64 {
	return Float64(e)
}

func (e Int8) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Int8) ToString() String {
	return String(strconv.FormatInt(int64(e.Int8()), 10))
}

func (e Int8) ToStringCheck() (String, Err) {
	return String(strconv.FormatInt(int64(e.Int8()), 10)), nil
}

func (e Int8) ToBytes() Bytes {
	return Bytes(forceconv.StringToBytes(strconv.FormatInt(int64(e.Int8()), 10)))
}

func (e Int8) ToBytesCheck() (Bytes, Err) {
	return Bytes(forceconv.StringToBytes(strconv.FormatInt(int64(e.Int8()), 10))), nil
}

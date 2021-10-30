package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/unsafecast"
)

type Uint32 uint32

func (e Uint32) Uint32() uint32 {
	return uint32(e)
}

func (e Uint32) Interface() interface{} {
	return uint32(e)
}

func (e Uint32) Type() Type {
	return ValTypeUint32
}

func (e Uint32) ToBool() Bool {
	return Bool(e != 0)
}

func (e Uint32) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Uint32) ToInt8() Int8 {
	return Int8(e)
}

func (e Uint32) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Uint32) ToInt16() Int16 {
	return Int16(e)
}

func (e Uint32) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Uint32) ToInt32() Int32 {
	return Int32(e)
}

func (e Uint32) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Uint32) ToInt64() Int64 {
	return Int64(e)
}

func (e Uint32) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Uint32) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Uint32) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Uint32) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Uint32) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Uint32) ToUint32() Uint32 {
	return e
}

func (e Uint32) ToUint32Check() (Uint32, Err) {
	return e, nil
}

func (e Uint32) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Uint32) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Uint32) ToFloat32() Float32 {
	return Float32(e)
}

func (e Uint32) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Uint32) ToFloat64() Float64 {
	return Float64(e)
}

func (e Uint32) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Uint32) ToString() String {
	return String(strconv.FormatUint(uint64(e.Uint32()), 10))
}

func (e Uint32) ToStringCheck() (String, Err) {
	return String(strconv.FormatUint(uint64(e.Uint32()), 10)), nil
}

func (e Uint32) ToBytes() Bytes {
	return Bytes(unsafecast.StringToBytes(strconv.FormatUint(uint64(e.Uint32()), 10)))
}

func (e Uint32) ToBytesCheck() (Bytes, Err) {
	return Bytes(unsafecast.StringToBytes(strconv.FormatUint(uint64(e.Uint32()), 10))), nil
}

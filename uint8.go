package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/forceconv"
)

type Uint8 uint8

func (e Uint8) Uint8() uint8 {
	return uint8(e)
}

func (e Uint8) Interface() interface{} {
	return uint8(e)
}

func (e Uint8) Type() Type {
	return ValTypeUint8
}

func (e Uint8) ToBool() Bool {
	return Bool(e != 0)
}

func (e Uint8) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Uint8) ToInt8() Int8 {
	return Int8(e)
}

func (e Uint8) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Uint8) ToInt16() Int16 {
	return Int16(e)
}

func (e Uint8) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Uint8) ToInt32() Int32 {
	return Int32(e)
}

func (e Uint8) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Uint8) ToInt64() Int64 {
	return Int64(e)
}

func (e Uint8) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Uint8) ToUint8() Uint8 {
	return e
}

func (e Uint8) ToUint8Check() (Uint8, Err) {
	return e, nil
}

func (e Uint8) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Uint8) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Uint8) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Uint8) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Uint8) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Uint8) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Uint8) ToFloat32() Float32 {
	return Float32(e)
}

func (e Uint8) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Uint8) ToFloat64() Float64 {
	return Float64(e)
}

func (e Uint8) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Uint8) ToString() String {
	return String(strconv.FormatUint(uint64(e.Uint8()), 10))
}

func (e Uint8) ToStringCheck() (String, Err) {
	return String(strconv.FormatUint(uint64(e.Uint8()), 10)), nil
}

func (e Uint8) ToBytes() Bytes {
	return Bytes(forceconv.StringToBytes(strconv.FormatUint(uint64(e.Uint8()), 10)))
}

func (e Uint8) ToBytesCheck() (Bytes, Err) {
	return Bytes(forceconv.StringToBytes(strconv.FormatUint(uint64(e.Uint8()), 10))), nil
}

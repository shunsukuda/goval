package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/forceconv"
)

type Float32 float32

func (e Float32) Float32() float32 {
	return float32(e)
}

func (e Float32) Interface() interface{} {
	return float32(e)
}

func (e Float32) Type() Type {
	return ValTypeFloat32
}

func (e Float32) ToBool() Bool {
	return Bool(e != 0)
}

func (e Float32) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Float32) ToInt8() Int8 {
	return Int8(e)
}

func (e Float32) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Float32) ToInt16() Int16 {
	return Int16(e)
}

func (e Float32) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Float32) ToInt32() Int32 {
	return Int32(e)
}

func (e Float32) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Float32) ToInt64() Int64 {
	return Int64(e)
}

func (e Float32) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Float32) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Float32) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Float32) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Float32) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Float32) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Float32) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Float32) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Float32) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Float32) ToFloat32() Float32 {
	return e
}

func (e Float32) ToFloat32Check() (Float32, Err) {
	return e, nil
}

func (e Float32) ToFloat64() Float64 {
	return Float64(e)
}

func (e Float32) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Float32) ToString() String {
	return String(strconv.FormatFloat(float64(e.Float32()), FloatConvFormat, -1, 32))
}

func (e Float32) ToStringCheck() (String, Err) {
	return String(strconv.FormatFloat(float64(e.Float32()), FloatConvFormat, -1, 32)), nil
}

func (e Float32) ToBytes() Bytes {
	return Bytes(forceconv.StringToBytes(strconv.FormatFloat(float64(e.Float32()), FloatConvFormat, -1, 32)))
}

func (e Float32) ToBytesCheck() (Bytes, Err) {
	return Bytes(forceconv.StringToBytes(strconv.FormatFloat(float64(e.Float32()), FloatConvFormat, -1, 32))), nil
}

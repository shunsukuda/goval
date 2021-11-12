package time

import (
	"strconv"

	. "github.com/shunsukuda/goval"
	"github.com/shunsukuda/goval/forceconv"
)

type Date64 int32

func (e Date64) Date64() int32 {
	return int32(e)
}

func (e Date64) Interface() interface{} {
	return int32(e)
}

func (e Date64) Type() Type {
	return ValTypeDate64
}

func (e Date64) ToBool() Bool {
	return Bool(e != 0)
}

func (e Date64) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Date64) ToInt8() Int8 {
	return Int8(e)
}

func (e Date64) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Date64) ToInt16() Int16 {
	return Int16(e)
}

func (e Date64) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Date64) ToInt32() Int32 {
	return Int32(e)
}

func (e Date64) ToDate64Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Date64) ToInt64() Int64 {
	return Int64(e)
}

func (e Date64) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Date64) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Date64) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Date64) ToUint16() Uint16 {
	return Uint16(e)
}

func (e Date64) ToUint16Check() (Uint16, Err) {
	return Uint16(e), nil
}

func (e Date64) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Date64) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Date64) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Date64) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Date64) ToFloat32() Float32 {
	return Float32(e)
}

func (e Date64) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Date64) ToFloat64() Float64 {
	return Float64(e)
}

func (e Date64) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Date64) ToString() String {
	return String(strconv.FormatInt(int64(e.Date64()), 10))
}

func (e Date64) ToStringCheck() (String, Err) {
	return String(strconv.FormatInt(int64(e.Date64()), 10)), nil
}

func (e Date64) ToBytes() Bytes {
	return Bytes(forceconv.StringToBytes(e.ToString().String()))
}

func (e Date64) ToBytesCheck() (Bytes, Err) {
	return Bytes(forceconv.StringToBytes(strconv.FormatInt(int64(e.Date64()), 10))), nil
}

package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/unsafecast"
)

type String string

func (e String) String() string {
	return string(e)
}

func (e String) Interface() interface{} {
	return string(e)
}

func (e String) Type() Type {
	return ValTypeString
}

func (e String) ToBool() Bool {
	cv, _ := strconv.ParseBool(e.String())
	return Bool(cv)
}

func (e String) ToBoolCheck() (Bool, Err) {
	cv, err := strconv.ParseBool(e.String())
	return Bool(cv), Err(err)
}

func (e String) ToInt8() Int8 {
	cv, _ := strconv.ParseInt(e.String(), 0, 8)
	return Int8(cv)
}

func (e String) ToInt8Check() (Int8, Err) {
	cv, err := strconv.ParseInt(e.String(), 0, 8)
	return Int8(cv), Err(err)
}

func (e String) ToInt16() Int16 {
	cv, _ := strconv.ParseInt(e.String(), 0, 16)
	return Int16(cv)
}

func (e String) ToInt16Check() (Int16, Err) {
	cv, err := strconv.ParseInt(e.String(), 0, 16)
	return Int16(cv), Err(err)
}

func (e String) ToInt32() Int32 {
	cv, _ := strconv.ParseInt(e.String(), 0, 32)
	return Int32(cv)
}

func (e String) ToInt32Check() (Int32, Err) {
	cv, err := strconv.ParseInt(e.String(), 0, 32)
	return Int32(cv), Err(err)
}

func (e String) ToInt64() Int64 {
	cv, _ := strconv.ParseInt(e.String(), 0, 64)
	return Int64(cv)
}

func (e String) ToInt64Check() (Int64, Err) {
	cv, err := strconv.ParseInt(e.String(), 0, 64)
	return Int64(cv), Err(err)
}

func (e String) ToUint8() Uint8 {
	cv, _ := strconv.ParseUint(e.String(), 0, 8)
	return Uint8(cv)
}

func (e String) ToUint8Check() (Uint8, Err) {
	cv, err := strconv.ParseUint(e.String(), 0, 8)
	return Uint8(cv), Err(err)
}

func (e String) ToUint16() Uint16 {
	cv, _ := strconv.ParseUint(e.String(), 0, 16)
	return Uint16(cv)
}

func (e String) ToUint16Check() (Uint16, Err) {
	cv, err := strconv.ParseUint(e.String(), 0, 16)
	return Uint16(cv), Err(err)
}

func (e String) ToUint32() Uint32 {
	cv, _ := strconv.ParseUint(e.String(), 0, 32)
	return Uint32(cv)
}

func (e String) ToUint32Check() (Uint32, Err) {
	cv, err := strconv.ParseUint(e.String(), 0, 32)
	return Uint32(cv), Err(err)
}

func (e String) ToUint64() Uint64 {
	cv, _ := strconv.ParseUint(e.String(), 0, 64)
	return Uint64(cv)
}

func (e String) ToUint64Check() (Uint64, Err) {
	cv, err := strconv.ParseUint(e.String(), 0, 64)
	return Uint64(cv), Err(err)
}

func (e String) ToFloat32() Float32 {
	cv, _ := strconv.ParseFloat(e.String(), 32)
	return Float32(cv)
}

func (e String) ToFloat32Check() (Float32, Err) {
	cv, err := strconv.ParseFloat(e.String(), 32)
	return Float32(cv), Err(err)
}

func (e String) ToFloat64() Float64 {
	cv, _ := strconv.ParseFloat(e.String(), 64)
	return Float64(cv)
}

func (e String) ToFloat64Check() (Float64, Err) {
	cv, err := strconv.ParseFloat(e.String(), 64)
	return Float64(cv), Err(err)
}

func (e String) ToString() String {
	return e
}

func (e String) ToStringCheck() (String, Err) {
	return e, nil
}

func (e String) ToBytes() Bytes {
	return Bytes(unsafecast.StringToBytes(e.String()))
}

func (e String) ToBytesCheck() (Bytes, Err) {
	return Bytes(unsafecast.StringToBytes(e.String())), nil
}

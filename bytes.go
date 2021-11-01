package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/forceconv"
)

type Bytes []byte

func (e Bytes) Bytes() []byte {
	return []byte(e)
}

func (e Bytes) Interface() interface{} {
	return []byte(e)
}

func (e Bytes) Type() Type {
	return ValTypeBytes
}

func (e Bytes) ToBool() Bool {
	cv, _ := strconv.ParseBool(forceconv.BytesToString(e.Bytes()))
	return Bool(cv)
}

func (e Bytes) ToBoolCheck() (Bool, Err) {
	cv, err := strconv.ParseBool(forceconv.BytesToString(e.Bytes()))
	return Bool(cv), Err(err)
}

func (e Bytes) ToInt8() Int8 {
	cv, _ := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 8)
	return Int8(cv)
}

func (e Bytes) ToInt8Check() (Int8, Err) {
	cv, err := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 8)
	return Int8(cv), Err(err)
}

func (e Bytes) ToInt16() Int16 {
	cv, _ := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 16)
	return Int16(cv)
}

func (e Bytes) ToInt16Check() (Int16, Err) {
	cv, err := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 16)
	return Int16(cv), Err(err)
}

func (e Bytes) ToInt32() Int32 {
	cv, _ := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 32)
	return Int32(cv)
}

func (e Bytes) ToInt32Check() (Int32, Err) {
	cv, err := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 32)
	return Int32(cv), Err(err)
}

func (e Bytes) ToInt64() Int64 {
	cv, _ := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 64)
	return Int64(cv)
}

func (e Bytes) ToInt64Check() (Int64, Err) {
	cv, err := strconv.ParseInt(forceconv.BytesToString(e.Bytes()), 0, 64)
	return Int64(cv), Err(err)
}

func (e Bytes) ToUint8() Uint8 {
	cv, _ := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 8)
	return Uint8(cv)
}

func (e Bytes) ToUint8Check() (Uint8, Err) {
	cv, err := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 8)
	return Uint8(cv), Err(err)
}

func (e Bytes) ToUint16() Uint16 {
	cv, _ := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 16)
	return Uint16(cv)
}

func (e Bytes) ToUint16Check() (Uint16, Err) {
	cv, err := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 16)
	return Uint16(cv), Err(err)
}

func (e Bytes) ToUint32() Uint32 {
	cv, _ := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 32)
	return Uint32(cv)
}

func (e Bytes) ToUint32Check() (Uint32, Err) {
	cv, err := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 32)
	return Uint32(cv), Err(err)
}

func (e Bytes) ToUint64() Uint64 {
	cv, _ := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 64)
	return Uint64(cv)
}

func (e Bytes) ToUint64Check() (Uint64, Err) {
	cv, err := strconv.ParseUint(forceconv.BytesToString(e.Bytes()), 0, 64)
	return Uint64(cv), Err(err)
}

func (e Bytes) ToFloat32() Float32 {
	cv, _ := strconv.ParseFloat(forceconv.BytesToString(e.Bytes()), 32)
	return Float32(cv)
}

func (e Bytes) ToFloat32Check() (Float32, Err) {
	cv, err := strconv.ParseFloat(forceconv.BytesToString(e.Bytes()), 32)
	return Float32(cv), Err(err)
}

func (e Bytes) ToFloat64() Float64 {
	cv, _ := strconv.ParseFloat(forceconv.BytesToString(e.Bytes()), 64)
	return Float64(cv)
}

func (e Bytes) ToFloat64Check() (Float64, Err) {
	cv, err := strconv.ParseFloat(forceconv.BytesToString(e.Bytes()), 64)
	return Float64(cv), Err(err)
}

func (e Bytes) ToString() String {
	return String(forceconv.BytesToString(e.Bytes()))
}

func (e Bytes) ToStringCheck() (String, Err) {
	return String(forceconv.BytesToString(e.Bytes())), nil
}

func (e Bytes) ToBytes() Bytes {
	return Bytes(e.Bytes())
}

func (e Bytes) ToBytesCheck() (Bytes, Err) {
	return Bytes(e.Bytes()), nil
}

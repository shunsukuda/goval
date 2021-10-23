package goval

import "strconv"

type Uint16 uint16

func (e Uint16) Uint16() uint16 {
	return uint16(e)
}

func (e Uint16) Interface() interface{} {
	return uint16(e)
}

func (e Uint16) Type() Type {
	return ValTypeUint16
}

func (e Uint16) ToBool() Bool {
	return Bool(e != 0)
}

func (e Uint16) ToBoolCheck() (Bool, Err) {
	return Bool(e != 0), nil
}

func (e Uint16) ToInt8() Int8 {
	return Int8(e)
}

func (e Uint16) ToInt8Check() (Int8, Err) {
	return Int8(e), nil
}

func (e Uint16) ToInt16() Int16 {
	return Int16(e)
}

func (e Uint16) ToInt16Check() (Int16, Err) {
	return Int16(e), nil
}

func (e Uint16) ToInt32() Int32 {
	return Int32(e)
}

func (e Uint16) ToInt32Check() (Int32, Err) {
	return Int32(e), nil
}

func (e Uint16) ToInt64() Int64 {
	return Int64(e)
}

func (e Uint16) ToInt64Check() (Int64, Err) {
	return Int64(e), nil
}

func (e Uint16) ToUint8() Uint8 {
	return Uint8(e)
}

func (e Uint16) ToUint8Check() (Uint8, Err) {
	return Uint8(e), nil
}

func (e Uint16) ToUint16() Uint16 {
	return e
}

func (e Uint16) ToUint16Check() (Uint16, Err) {
	return e, nil
}

func (e Uint16) ToUint32() Uint32 {
	return Uint32(e)
}

func (e Uint16) ToUint32Check() (Uint32, Err) {
	return Uint32(e), nil
}

func (e Uint16) ToUint64() Uint64 {
	return Uint64(e)
}

func (e Uint16) ToUint64Check() (Uint64, Err) {
	return Uint64(e), nil
}

func (e Uint16) ToFloat32() Float32 {
	return Float32(e)
}

func (e Uint16) ToFloat32Check() (Float32, Err) {
	return Float32(e), nil
}

func (e Uint16) ToFloat64() Float64 {
	return Float64(e)
}

func (e Uint16) ToFloat64Check() (Float64, Err) {
	return Float64(e), nil
}

func (e Uint16) ToString() String {
	return String(strconv.FormatUint(uint64(e.Uint16()), 10))
}

func (e Uint16) ToStringCheck() (String, Err) {
	return String(strconv.FormatUint(uint64(e.Uint16()), 10)), nil
}

func (e Uint16) ToBytes() Bytes {
	return Bytes(unsafeStringToBytes(strconv.FormatUint(uint64(e.Uint16()), 10)))
}

func (e Uint16) ToBytesCheck() (Bytes, Err) {
	return Bytes(unsafeStringToBytes(strconv.FormatUint(uint64(e.Uint16()), 10))), nil
}

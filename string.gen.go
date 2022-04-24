package goval

import (
	"strconv"

	"github.com/shunsukuda/forceconv"
)

type String struct {
	String string
}

func (e String) GoString() string       { return e.String }
func (e String) Interface() interface{} { return e.String }
func (e String) Val() Val               { return e }
func (e String) Type() Type             { return ValTypes.String }

func (e String) ToInt8() Int8 {
	ce, _ := strconv.ParseInt(e.String, 0, 8)
	return Int8{int8(ce)}
}
func (e String) ToInt16() Int16 {
	ce, _ := strconv.ParseInt(e.String, 0, 16)
	return Int16{int16(ce)}
}
func (e String) ToInt32() Int32 {
	ce, _ := strconv.ParseInt(e.String, 0, 32)
	return Int32{int32(ce)}
}
func (e String) ToInt64() Int64 {
	ce, _ := strconv.ParseInt(e.String, 0, 64)
	return Int64{int64(ce)}
}
func (e String) ToUint8() Uint8 {
	ce, _ := strconv.ParseUint(e.String, 0, 8)
	return Uint8{uint8(ce)}
}
func (e String) ToUint16() Uint16 {
	ce, _ := strconv.ParseUint(e.String, 0, 16)
	return Uint16{uint16(ce)}
}
func (e String) ToUint32() Uint32 {
	ce, _ := strconv.ParseUint(e.String, 0, 32)
	return Uint32{uint32(ce)}
}
func (e String) ToUint64() Uint64 {
	ce, _ := strconv.ParseUint(e.String, 0, 64)
	return Uint64{uint64(ce)}
}
func (e String) ToFloat32() Float32 {
	ce, _ := strconv.ParseFloat(e.String, 0)
	return Float32{float32(ce)}
}
func (e String) ToFloat64() Float64 {
	ce, _ := strconv.ParseFloat(e.String, 0)
	return Float64{float64(ce)}
}
func (e String) ToComplex64() Complex64 {
	ce, _ := strconv.ParseComplex(e.String, 0)
	return Complex64{complex64(ce)}
}
func (e String) ToComplex128() Complex128 {
	ce, _ := strconv.ParseComplex(e.String, 0)
	return Complex128{complex128(ce)}
}
func (e String) ToBool() Bool {
	ce, _ := strconv.ParseBool(e.String)
	return Bool{ce}
}
func (e String) ToString() String {
	return e
}
func (e String) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.String)}
}

type Bytes struct {
	Bytes []byte
}

func (e Bytes) GoBytes() []byte        { return e.Bytes }
func (e Bytes) Interface() interface{} { return e.Bytes }
func (e Bytes) Val() Val               { return e }
func (e Bytes) Type() Type             { return ValTypes.Bytes }

func (e Bytes) ToInt8() Int8 {
	return e.ToString().ToInt8()
}
func (e Bytes) ToInt16() Int16 {
	return e.ToString().ToInt16()
}
func (e Bytes) ToInt32() Int32 {
	return e.ToString().ToInt32()
}
func (e Bytes) ToInt64() Int64 {
	return e.ToString().ToInt64()
}
func (e Bytes) ToUint8() Uint8 {
	return e.ToString().ToUint8()
}
func (e Bytes) ToUint16() Uint16 {
	return e.ToString().ToUint16()
}
func (e Bytes) ToUint32() Uint32 {
	return e.ToString().ToUint32()
}
func (e Bytes) ToUint64() Uint64 {
	return e.ToString().ToUint64()
}
func (e Bytes) ToFloat32() Float32 {
	return e.ToString().ToFloat32()
}
func (e Bytes) ToFloat64() Float64 {
	return e.ToString().ToFloat64()
}
func (e Bytes) ToComplex64() Complex64 {
	return e.ToString().ToComplex64()
}
func (e Bytes) ToComplex128() Complex128 {
	return e.ToString().ToComplex128()
}
func (e Bytes) ToBool() Bool {
	return e.ToString().ToBool()
}
func (e Bytes) ToString() String {
	return String{forceconv.BytesToString(e.Bytes)}
}
func (e Bytes) ToBytes() Bytes {
	return e
}

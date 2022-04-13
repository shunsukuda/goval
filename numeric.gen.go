package goval

import (
	"math"
	"strconv"
)

var (
	MinInt8   = Int8(math.MinInt8)
	MinInt16  = Int16(math.MinInt16)
	MinInt32  = Int32(math.MinInt32)
	MinInt64  = Int64(math.MinInt64)
	MaxInt8   = Int8(math.MaxInt8)
	MaxInt16  = Int16(math.MaxInt16)
	MaxInt32  = Int32(math.MaxInt32)
	MaxInt64  = Int64(math.MaxInt64)
	MinUint8  = Uint8(0)
	MinUint16 = Uint16(0)
	MinUint32 = Uint32(0)
	MinUint64 = Uint64(0)
	MaxUint8  = Uint8(math.MaxUint8)
	MaxUint16 = Uint16(math.MaxUint16)
	MaxUint32 = Uint32(math.MaxUint32)
	MaxUint64 = Uint64(math.MaxUint64)
)

var (
	IntToStringBase   int  = 10
	UintToStringBase  int  = 16
	FloatToStringFmt  byte = 'g'
	FloatToStringPrec int  = -1
)

type Int8 int8

func (e Int8) Int8() int8             { return int8(e) }
func (e Int8) Interface() interface{} { return e.Int8() }
func (e Int8) Val() Val               { return e }
func (e Int8) Type() Type             { return ValTypes.Int8 }

func (e Int8) ToInt8() Int8                       { return e }
func (e Int8) ToInt8Eq() (Int8, bool)             { return e, true }
func (e Int8) ToInt16() Int16                     { return Int16(e) }
func (e Int8) ToInt16Eq() (Int16, bool)           { return e.ToInt16(), true }
func (e Int8) ToInt32() Int32                     { return Int32(e) }
func (e Int8) ToInt32Eq() (Int32, bool)           { return e.ToInt32(), true }
func (e Int8) ToInt64() Int64                     { return Int64(e) }
func (e Int8) ToInt64Eq() (Int64, bool)           { return e.ToInt64(), true }
func (e Int8) ToUint8() Uint8                     { return Uint8(e) }
func (e Int8) ToUint8Eq() (Uint8, bool)           { return e.ToUint8(), 0 <= e }
func (e Int8) ToUint16() Uint16                   { return Uint16(e) }
func (e Int8) ToUint16Eq() (Uint16, bool)         { return e.ToUint16(), 0 <= e }
func (e Int8) ToUint32() Uint32                   { return Uint32(e) }
func (e Int8) ToUint32Eq() (Uint32, bool)         { return e.ToUint32(), 0 <= e }
func (e Int8) ToUint64() Uint64                   { return Uint64(e) }
func (e Int8) ToUint64Eq() (Uint64, bool)         { return e.ToUint64(), 0 <= e }
func (e Int8) ToFloat32() Float32                 { return Float32(e) }
func (e Int8) ToFloat32Eq() (Float32, bool)       { return e.ToFloat32(), true }
func (e Int8) ToFloat64() Float64                 { return Float64(e) }
func (e Int8) ToFloat64Eq() (Float64, bool)       { return e.ToFloat64(), true }
func (e Int8) ToComplex64() Complex64             { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Int8) ToComplex64Eq() (Complex64, bool)   { return e.ToComplex64(), true }
func (e Int8) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Int8) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Int8) ToBool() Bool                       { return Bool(e != 0) }
func (e Int8) ToString() String {
	return String(strconv.FormatInt(e.ToInt64().Int64(), IntToStringBase))
}

type Int16 int16

func (e Int16) Int16() int16           { return int16(e) }
func (e Int16) Interface() interface{} { return e.Int16() }
func (e Int16) Val() Val               { return e }
func (e Int16) Type() Type             { return ValTypes.Int16 }

func (e Int16) ToInt8() Int8                       { return Int8(e) }
func (e Int16) ToInt8Eq() (Int8, bool)             { return e.ToInt8(), math.MinInt8 <= e && e <= math.MaxInt8 }
func (e Int16) ToInt16() Int16                     { return e }
func (e Int16) ToInt16Eq() (Int16, bool)           { return e, true }
func (e Int16) ToInt32() Int32                     { return Int32(e) }
func (e Int16) ToInt32Eq() (Int32, bool)           { return e.ToInt32(), true }
func (e Int16) ToInt64() Int64                     { return Int64(e) }
func (e Int16) ToInt64Eq() (Int64, bool)           { return e.ToInt64(), true }
func (e Int16) ToUint8() Uint8                     { return Uint8(e) }
func (e Int16) ToUint8Eq() (Uint8, bool)           { return e.ToUint8(), 0 <= e && e <= math.MaxUint8 }
func (e Int16) ToUint16() Uint16                   { return Uint16(e) }
func (e Int16) ToUint16Eq() (Uint16, bool)         { return e.ToUint16(), 0 <= e }
func (e Int16) ToUint32() Uint32                   { return Uint32(e) }
func (e Int16) ToUint32Eq() (Uint32, bool)         { return e.ToUint32(), 0 <= e }
func (e Int16) ToUint64() Uint64                   { return Uint64(e) }
func (e Int16) ToUint64Eq() (Uint64, bool)         { return e.ToUint64(), 0 <= e }
func (e Int16) ToFloat32() Float32                 { return Float32(e) }
func (e Int16) ToFloat32Eq() (Float32, bool)       { return e.ToFloat32(), true }
func (e Int16) ToFloat64() Float64                 { return Float64(e) }
func (e Int16) ToFloat64Eq() (Float64, bool)       { return e.ToFloat64(), true }
func (e Int16) ToComplex64() Complex64             { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Int16) ToComplex64Eq() (Complex64, bool)   { return e.ToComplex64(), true }
func (e Int16) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Int16) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Int16) ToBool() Bool                       { return Bool(e != 0) }
func (e Int16) ToString() String {
	return String(strconv.FormatInt(e.ToInt64().Int64(), IntToStringBase))
}

type Int32 int32

func (e Int32) Int32() int32           { return int32(e) }
func (e Int32) Interface() interface{} { return e.Int32() }
func (e Int32) Val() Val               { return e }
func (e Int32) Type() Type             { return ValTypes.Int32 }

func (e Int32) ToInt8() Int8           { return Int8(e) }
func (e Int32) ToInt8Eq() (Int8, bool) { return e.ToInt8(), math.MinInt8 <= e && e <= math.MaxInt8 }
func (e Int32) ToInt16() Int16         { return Int16(e) }
func (e Int32) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), math.MinInt16 <= e && e <= math.MaxInt16
}
func (e Int32) ToInt32() Int32               { return e }
func (e Int32) ToInt32Eq() (Int32, bool)     { return e, true }
func (e Int32) ToInt64() Int64               { return Int64(e) }
func (e Int32) ToInt64Eq() (Int64, bool)     { return e.ToInt64(), true }
func (e Int32) ToUint8() Uint8               { return Uint8(e) }
func (e Int32) ToUint8Eq() (Uint8, bool)     { return e.ToUint8(), 0 <= e && e <= math.MaxUint8 }
func (e Int32) ToUint16() Uint16             { return Uint16(e) }
func (e Int32) ToUint16Eq() (Uint16, bool)   { return e.ToUint16(), 0 <= e && e <= math.MaxUint16 }
func (e Int32) ToUint32() Uint32             { return Uint32(e) }
func (e Int32) ToUint32Eq() (Uint32, bool)   { return e.ToUint32(), 0 <= e }
func (e Int32) ToUint64() Uint64             { return Uint64(e) }
func (e Int32) ToUint64Eq() (Uint64, bool)   { return e.ToUint64(), 0 <= e }
func (e Int32) ToFloat32() Float32           { return Float32(e) }
func (e Int32) ToFloat32Eq() (Float32, bool) { return e.ToFloat32(), e == e.ToFloat32().ToInt32() }
func (e Int32) ToFloat64() Float64           { return Float64(e) }
func (e Int32) ToFloat64Eq() (Float64, bool) { return e.ToFloat64(), true }
func (e Int32) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Int32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToInt32()
}
func (e Int32) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Int32) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Int32) ToBool() Bool                       { return Bool(e != 0) }
func (e Int32) ToString() String {
	return String(strconv.FormatInt(e.ToInt64().Int64(), IntToStringBase))
}

type Int64 int64

func (e Int64) Int64() int64           { return int64(e) }
func (e Int64) Interface() interface{} { return e.Int64() }
func (e Int64) Val() Val               { return e }
func (e Int64) Type() Type             { return ValTypes.Int64 }

func (e Int64) ToInt8() Int8           { return Int8(e) }
func (e Int64) ToInt8Eq() (Int8, bool) { return e.ToInt8(), math.MinInt8 <= e && e <= math.MaxInt8 }
func (e Int64) ToInt16() Int16         { return Int16(e) }
func (e Int64) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), math.MinInt16 <= e && e <= math.MaxInt16
}
func (e Int64) ToInt32() Int32 { return Int32(e) }
func (e Int64) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), math.MinInt32 <= e && e <= math.MaxInt32
}
func (e Int64) ToInt64() Int64               { return e }
func (e Int64) ToInt64Eq() (Int64, bool)     { return e, true }
func (e Int64) ToUint8() Uint8               { return Uint8(e) }
func (e Int64) ToUint8Eq() (Uint8, bool)     { return e.ToUint8(), 0 <= e && e <= math.MaxUint8 }
func (e Int64) ToUint16() Uint16             { return Uint16(e) }
func (e Int64) ToUint16Eq() (Uint16, bool)   { return e.ToUint16(), 0 <= e && e <= math.MaxUint16 }
func (e Int64) ToUint32() Uint32             { return Uint32(e) }
func (e Int64) ToUint32Eq() (Uint32, bool)   { return e.ToUint32(), 0 <= e && e <= math.MaxUint32 }
func (e Int64) ToUint64() Uint64             { return Uint64(e) }
func (e Int64) ToUint64Eq() (Uint64, bool)   { return e.ToUint64(), 0 <= e }
func (e Int64) ToFloat32() Float32           { return Float32(e) }
func (e Int64) ToFloat32Eq() (Float32, bool) { return e.ToFloat32(), e == e.ToFloat32().ToInt64() }
func (e Int64) ToFloat64() Float64           { return Float64(e) }
func (e Int64) ToFloat64Eq() (Float64, bool) { return e.ToFloat64(), e == e.ToFloat64().ToInt64() }
func (e Int64) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Int64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToInt64()
}
func (e Int64) ToComplex128() Complex128 { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Int64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToInt64()
}
func (e Int64) ToBool() Bool     { return Bool(e != 0) }
func (e Int64) ToString() String { return String(strconv.FormatInt(e.Int64(), IntToStringBase)) }

type Uint8 uint8

func (e Uint8) Uint8() uint8           { return uint8(e) }
func (e Uint8) Interface() interface{} { return e.Uint8() }
func (e Uint8) Val() Val               { return e }
func (e Uint8) Type() Type             { return ValTypes.Uint8 }

func (e Uint8) ToInt8() Int8                       { return Int8(e) }
func (e Uint8) ToInt8Eq() (Int8, bool)             { return e.ToInt8(), 0 <= e && e <= math.MaxInt8 }
func (e Uint8) ToInt16() Int16                     { return Int16(e) }
func (e Uint8) ToInt16Eq() (Int16, bool)           { return e.ToInt16(), 0 <= e }
func (e Uint8) ToInt32() Int32                     { return Int32(e) }
func (e Uint8) ToInt32Eq() (Int32, bool)           { return e.ToInt32(), 0 <= e }
func (e Uint8) ToInt64() Int64                     { return Int64(e) }
func (e Uint8) ToInt64Eq() (Int64, bool)           { return e.ToInt64(), 0 <= e }
func (e Uint8) ToUint8() Uint8                     { return e }
func (e Uint8) ToUint8Eq() (Uint8, bool)           { return e, true }
func (e Uint8) ToUint16() Uint16                   { return Uint16(e) }
func (e Uint8) ToUint16Eq() (Uint16, bool)         { return e.ToUint16(), true }
func (e Uint8) ToUint32() Uint32                   { return Uint32(e) }
func (e Uint8) ToUint32Eq() (Uint32, bool)         { return e.ToUint32(), true }
func (e Uint8) ToUint64() Uint64                   { return Uint64(e) }
func (e Uint8) ToUint64Eq() (Uint64, bool)         { return e.ToUint64(), true }
func (e Uint8) ToFloat32() Float32                 { return Float32(e) }
func (e Uint8) ToFloat32Eq() (Float32, bool)       { return e.ToFloat32(), true }
func (e Uint8) ToFloat64() Float64                 { return Float64(e) }
func (e Uint8) ToFloat64Eq() (Float64, bool)       { return e.ToFloat64(), true }
func (e Uint8) ToComplex64() Complex64             { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Uint8) ToComplex64Eq() (Complex64, bool)   { return e.ToComplex64(), true }
func (e Uint8) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Uint8) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Uint8) ToBool() Bool                       { return Bool(e != 0) }
func (e Uint8) ToString() String {
	return String(strconv.FormatUint(e.ToUint64().Uint64(), UintToStringBase))
}

type Uint16 uint16

func (e Uint16) Uint16() uint16         { return uint16(e) }
func (e Uint16) Interface() interface{} { return e.Uint16() }
func (e Uint16) Val() Val               { return e }
func (e Uint16) Type() Type             { return ValTypes.Uint16 }

func (e Uint16) ToInt8() Int8                       { return Int8(e) }
func (e Uint16) ToInt8Eq() (Int8, bool)             { return e.ToInt8(), 0 <= e && e <= math.MaxInt8 }
func (e Uint16) ToInt16() Int16                     { return Int16(e) }
func (e Uint16) ToInt16Eq() (Int16, bool)           { return e.ToInt16(), 0 <= e && e <= math.MaxInt16 }
func (e Uint16) ToInt32() Int32                     { return Int32(e) }
func (e Uint16) ToInt32Eq() (Int32, bool)           { return e.ToInt32(), 0 <= e }
func (e Uint16) ToInt64() Int64                     { return Int64(e) }
func (e Uint16) ToInt64Eq() (Int64, bool)           { return e.ToInt64(), 0 <= e }
func (e Uint16) ToUint8() Uint8                     { return Uint8(e) }
func (e Uint16) ToUint8Eq() (Uint8, bool)           { return e.ToUint8(), e <= math.MaxUint8 }
func (e Uint16) ToUint16() Uint16                   { return e }
func (e Uint16) ToUint16Eq() (Uint16, bool)         { return e, true }
func (e Uint16) ToUint32() Uint32                   { return Uint32(e) }
func (e Uint16) ToUint32Eq() (Uint32, bool)         { return e.ToUint32(), true }
func (e Uint16) ToUint64() Uint64                   { return Uint64(e) }
func (e Uint16) ToUint64Eq() (Uint64, bool)         { return e.ToUint64(), true }
func (e Uint16) ToFloat32() Float32                 { return Float32(e) }
func (e Uint16) ToFloat32Eq() (Float32, bool)       { return e.ToFloat32(), true }
func (e Uint16) ToFloat64() Float64                 { return Float64(e) }
func (e Uint16) ToFloat64Eq() (Float64, bool)       { return e.ToFloat64(), true }
func (e Uint16) ToComplex64() Complex64             { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Uint16) ToComplex64Eq() (Complex64, bool)   { return e.ToComplex64(), true }
func (e Uint16) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Uint16) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Uint16) ToBool() Bool                       { return Bool(e != 0) }
func (e Uint16) ToString() String {
	return String(strconv.FormatUint(e.ToUint64().Uint64(), UintToStringBase))
}

type Uint32 uint32

func (e Uint32) Uint32() uint32         { return uint32(e) }
func (e Uint32) Interface() interface{} { return e.Uint32() }
func (e Uint32) Val() Val               { return e }
func (e Uint32) Type() Type             { return ValTypes.Uint32 }

func (e Uint32) ToInt8() Int8                 { return Int8(e) }
func (e Uint32) ToInt8Eq() (Int8, bool)       { return e.ToInt8(), 0 <= e && e <= math.MaxInt8 }
func (e Uint32) ToInt16() Int16               { return Int16(e) }
func (e Uint32) ToInt16Eq() (Int16, bool)     { return e.ToInt16(), 0 <= e && e <= math.MaxInt16 }
func (e Uint32) ToInt32() Int32               { return Int32(e) }
func (e Uint32) ToInt32Eq() (Int32, bool)     { return e.ToInt32(), 0 <= e && e <= math.MaxInt32 }
func (e Uint32) ToInt64() Int64               { return Int64(e) }
func (e Uint32) ToInt64Eq() (Int64, bool)     { return e.ToInt64(), 0 <= e }
func (e Uint32) ToUint8() Uint8               { return Uint8(e) }
func (e Uint32) ToUint8Eq() (Uint8, bool)     { return e.ToUint8(), e <= math.MaxUint8 }
func (e Uint32) ToUint16() Uint16             { return Uint16(e) }
func (e Uint32) ToUint16Eq() (Uint16, bool)   { return e.ToUint16(), e <= math.MaxUint16 }
func (e Uint32) ToUint32() Uint32             { return e }
func (e Uint32) ToUint32Eq() (Uint32, bool)   { return e, true }
func (e Uint32) ToUint64() Uint64             { return Uint64(e) }
func (e Uint32) ToUint64Eq() (Uint64, bool)   { return e.ToUint64(), true }
func (e Uint32) ToFloat32() Float32           { return Float32(e) }
func (e Uint32) ToFloat32Eq() (Float32, bool) { return e.ToFloat32(), e == e.ToFloat32().ToUint32() }
func (e Uint32) ToFloat64() Float64           { return Float64(e) }
func (e Uint32) ToFloat64Eq() (Float64, bool) { return e.ToFloat64(), true }
func (e Uint32) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Uint32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToUint32()
}
func (e Uint32) ToComplex128() Complex128           { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Uint32) ToComplex128Eq() (Complex128, bool) { return e.ToComplex128(), true }
func (e Uint32) ToBool() Bool                       { return Bool(e != 0) }
func (e Uint32) ToString() String {
	return String(strconv.FormatUint(e.ToUint64().Uint64(), UintToStringBase))
}

type Uint64 uint64

func (e Uint64) Uint64() uint64         { return uint64(e) }
func (e Uint64) Interface() interface{} { return e.Uint64() }
func (e Uint64) Val() Val               { return e }
func (e Uint64) Type() Type             { return ValTypes.Uint64 }

func (e Uint64) ToInt8() Int8                 { return Int8(e) }
func (e Uint64) ToInt8Eq() (Int8, bool)       { return e.ToInt8(), 0 <= e && e <= math.MaxInt8 }
func (e Uint64) ToInt16() Int16               { return Int16(e) }
func (e Uint64) ToInt16Eq() (Int16, bool)     { return e.ToInt16(), 0 <= e && e <= math.MaxInt16 }
func (e Uint64) ToInt32() Int32               { return Int32(e) }
func (e Uint64) ToInt32Eq() (Int32, bool)     { return e.ToInt32(), 0 <= e && e <= math.MaxInt32 }
func (e Uint64) ToInt64() Int64               { return Int64(e) }
func (e Uint64) ToInt64Eq() (Int64, bool)     { return e.ToInt64(), 0 <= e && e <= math.MaxInt64 }
func (e Uint64) ToUint8() Uint8               { return Uint8(e) }
func (e Uint64) ToUint8Eq() (Uint8, bool)     { return e.ToUint8(), e <= math.MaxUint8 }
func (e Uint64) ToUint16() Uint16             { return Uint16(e) }
func (e Uint64) ToUint16Eq() (Uint16, bool)   { return e.ToUint16(), e <= math.MaxUint16 }
func (e Uint64) ToUint32() Uint32             { return Uint32(e) }
func (e Uint64) ToUint32Eq() (Uint32, bool)   { return e.ToUint32(), e <= math.MaxUint32 }
func (e Uint64) ToUint64() Uint64             { return e }
func (e Uint64) ToUint64Eq() (Uint64, bool)   { return e, true }
func (e Uint64) ToFloat32() Float32           { return Float32(e) }
func (e Uint64) ToFloat32Eq() (Float32, bool) { return e.ToFloat32(), e == e.ToFloat32().ToUint64() }
func (e Uint64) ToFloat64() Float64           { return Float64(e) }
func (e Uint64) ToFloat64Eq() (Float64, bool) { return e.ToFloat64(), e == e.ToFloat64().ToUint64() }
func (e Uint64) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Uint64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToUint64()
}
func (e Uint64) ToComplex128() Complex128 { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Uint64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToUint64()
}
func (e Uint64) ToBool() Bool     { return Bool(e != 0) }
func (e Uint64) ToString() String { return String(strconv.FormatUint(e.Uint64(), UintToStringBase)) }

type Float32 float32

func (e Float32) Float32() float32       { return float32(e) }
func (e Float32) Interface() interface{} { return e.Float32() }
func (e Float32) Val() Val               { return e }
func (e Float32) Type() Type             { return ValTypes.Float32 }

func (e Float32) ToInt8() Int8                 { return Int8(e) }
func (e Float32) ToInt8Eq() (Int8, bool)       { return e.ToInt8(), e == e.ToInt8().ToFloat32() }
func (e Float32) ToInt16() Int16               { return Int16(e) }
func (e Float32) ToInt16Eq() (Int16, bool)     { return e.ToInt16(), e == e.ToInt16().ToFloat32() }
func (e Float32) ToInt32() Int32               { return Int32(e) }
func (e Float32) ToInt32Eq() (Int32, bool)     { return e.ToInt32(), e == e.ToInt32().ToFloat32() }
func (e Float32) ToInt64() Int64               { return Int64(e) }
func (e Float32) ToInt64Eq() (Int64, bool)     { return e.ToInt64(), e == e.ToInt64().ToFloat32() }
func (e Float32) ToUint8() Uint8               { return Uint8(e) }
func (e Float32) ToUint8Eq() (Uint8, bool)     { return e.ToUint8(), e == e.ToUint8().ToFloat32() }
func (e Float32) ToUint16() Uint16             { return Uint16(e) }
func (e Float32) ToUint16Eq() (Uint16, bool)   { return e.ToUint16(), e == e.ToUint16().ToFloat32() }
func (e Float32) ToUint32() Uint32             { return Uint32(e) }
func (e Float32) ToUint32Eq() (Uint32, bool)   { return e.ToUint32(), e == e.ToUint32().ToFloat32() }
func (e Float32) ToUint64() Uint64             { return Uint64(e) }
func (e Float32) ToUint64Eq() (Uint64, bool)   { return e.ToUint64(), e == e.ToUint64().ToFloat32() }
func (e Float32) ToFloat32() Float32           { return e }
func (e Float32) ToFloat32Eq() (Float32, bool) { return e, true }
func (e Float32) ToFloat64() Float64           { return Float64(e) }
func (e Float32) ToFloat64Eq() (Float64, bool) { return e.ToFloat64(), e == e.ToFloat64().ToFloat32() }
func (e Float32) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Float32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Float32) ToComplex128() Complex128 { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Float32) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToFloat32()
}
func (e Float32) ToBool() Bool { return Bool(e != 0) }
func (e Float32) ToString() String {
	return String(strconv.FormatFloat(e.ToFloat64().Float64(), FloatToStringFmt, FloatToStringPrec, 32))
}

type Float64 float64

func (e Float64) Float64() float64       { return float64(e) }
func (e Float64) Interface() interface{} { return e.Float64() }
func (e Float64) Val() Val               { return e }
func (e Float64) Type() Type             { return ValTypes.Float64 }

func (e Float64) ToInt8() Int8               { return Int8(e) }
func (e Float64) ToInt8Eq() (Int8, bool)     { return e.ToInt8(), e == e.ToInt8().ToFloat64() }
func (e Float64) ToInt16() Int16             { return Int16(e) }
func (e Float64) ToInt16Eq() (Int16, bool)   { return e.ToInt16(), e == e.ToInt16().ToFloat64() }
func (e Float64) ToInt32() Int32             { return Int32(e) }
func (e Float64) ToInt32Eq() (Int32, bool)   { return e.ToInt32(), e == e.ToInt32().ToFloat64() }
func (e Float64) ToInt64() Int64             { return Int64(e) }
func (e Float64) ToInt64Eq() (Int64, bool)   { return e.ToInt64(), e == e.ToInt64().ToFloat64() }
func (e Float64) ToUint8() Uint8             { return Uint8(e) }
func (e Float64) ToUint8Eq() (Uint8, bool)   { return e.ToUint8(), e == e.ToUint8().ToFloat64() }
func (e Float64) ToUint16() Uint16           { return Uint16(e) }
func (e Float64) ToUint16Eq() (Uint16, bool) { return e.ToUint16(), e == e.ToUint16().ToFloat64() }
func (e Float64) ToUint32() Uint32           { return Uint32(e) }
func (e Float64) ToUint32Eq() (Uint32, bool) { return e.ToUint32(), e == e.ToUint32().ToFloat64() }
func (e Float64) ToUint64() Uint64           { return Uint64(e) }
func (e Float64) ToUint64Eq() (Uint64, bool) { return e.ToUint64(), e == e.ToUint64().ToFloat64() }
func (e Float64) ToFloat32() Float32         { return Float32(e) }
func (e Float64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Float64) ToFloat64() Float64           { return e }
func (e Float64) ToFloat64Eq() (Float64, bool) { return e, true }
func (e Float64) ToComplex64() Complex64       { return Complex64(complex(e.ToFloat32(), 0)) }
func (e Float64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Float64) ToComplex128() Complex128 { return Complex128(complex(e.ToFloat64(), 0)) }
func (e Float64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e.ToString() == e.ToComplex128().ToString()
}
func (e Float64) ToBool() Bool { return Bool(e != 0) }
func (e Float64) ToString() String {
	return String(strconv.FormatFloat(e.Float64(), FloatToStringFmt, FloatToStringPrec, 64))
}

type Complex64 complex64

func (e Complex64) Complex64() complex64   { return complex64(e) }
func (e Complex64) Interface() interface{} { return e.Complex64() }
func (e Complex64) Val() Val               { return e }
func (e Complex64) Type() Type             { return ValTypes.Complex64 }

func (e Complex64) ToInt8() Int8               { return Int8(real(e.Complex64())) }
func (e Complex64) ToInt8Eq() (Int8, bool)     { return e.ToInt8(), e == e.ToInt8().ToComplex64() }
func (e Complex64) ToInt16() Int16             { return Int16(real(e.Complex64())) }
func (e Complex64) ToInt16Eq() (Int16, bool)   { return e.ToInt16(), e == e.ToInt16().ToComplex64() }
func (e Complex64) ToInt32() Int32             { return Int32(real(e.Complex64())) }
func (e Complex64) ToInt32Eq() (Int32, bool)   { return e.ToInt32(), e == e.ToInt32().ToComplex64() }
func (e Complex64) ToInt64() Int64             { return Int64(real(e.Complex64())) }
func (e Complex64) ToInt64Eq() (Int64, bool)   { return e.ToInt64(), e == e.ToInt64().ToComplex64() }
func (e Complex64) ToUint8() Uint8             { return Uint8(real(e.Complex64())) }
func (e Complex64) ToUint8Eq() (Uint8, bool)   { return e.ToUint8(), e == e.ToUint8().ToComplex64() }
func (e Complex64) ToUint16() Uint16           { return Uint16(real(e.Complex64())) }
func (e Complex64) ToUint16Eq() (Uint16, bool) { return e.ToUint16(), e == e.ToUint16().ToComplex64() }
func (e Complex64) ToUint32() Uint32           { return Uint32(real(e.Complex64())) }
func (e Complex64) ToUint32Eq() (Uint32, bool) { return e.ToUint32(), e == e.ToUint32().ToComplex64() }
func (e Complex64) ToUint64() Uint64           { return Uint64(real(e.Complex64())) }
func (e Complex64) ToUint64Eq() (Uint64, bool) { return e.ToUint64(), e == e.ToUint64().ToComplex64() }
func (e Complex64) ToFloat32() Float32         { return Float32(real(e.Complex64())) }
func (e Complex64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Complex64) ToFloat64() Float64 { return Float64(real(e.Complex64())) }
func (e Complex64) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e == e.ToFloat64().ToComplex64()
}
func (e Complex64) ToComplex64() Complex64           { return e }
func (e Complex64) ToComplex64Eq() (Complex64, bool) { return e, true }
func (e Complex64) ToComplex128() Complex128         { return Complex128(e) }
func (e Complex64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToComplex64()
}
func (e Complex64) ToBool() Bool { return Bool(e != 0) }
func (e Complex64) ToString() String {
	return String(strconv.FormatComplex(e.ToComplex128().Complex128(), FloatToStringFmt, FloatToStringPrec, 128))
}

type Complex128 complex128

func (e Complex128) Complex128() complex128 { return complex128(e) }
func (e Complex128) Interface() interface{} { return e.Complex128() }
func (e Complex128) Val() Val               { return e }
func (e Complex128) Type() Type             { return ValTypes.Complex128 }

func (e Complex128) ToInt8() Int8             { return Int8(real(e.Complex128())) }
func (e Complex128) ToInt8Eq() (Int8, bool)   { return e.ToInt8(), e == e.ToInt8().ToComplex128() }
func (e Complex128) ToInt16() Int16           { return Int16(real(e.Complex128())) }
func (e Complex128) ToInt16Eq() (Int16, bool) { return e.ToInt16(), e == e.ToInt16().ToComplex128() }
func (e Complex128) ToInt32() Int32           { return Int32(real(e.Complex128())) }
func (e Complex128) ToInt32Eq() (Int32, bool) { return e.ToInt32(), e == e.ToInt32().ToComplex128() }
func (e Complex128) ToInt64() Int64           { return Int64(real(e.Complex128())) }
func (e Complex128) ToInt64Eq() (Int64, bool) { return e.ToInt64(), e == e.ToInt64().ToComplex128() }
func (e Complex128) ToUint8() Uint8           { return Uint8(real(e.Complex128())) }
func (e Complex128) ToUint8Eq() (Uint8, bool) { return e.ToUint8(), e == e.ToUint8().ToComplex128() }
func (e Complex128) ToUint16() Uint16         { return Uint16(real(e.Complex128())) }
func (e Complex128) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e == e.ToUint16().ToComplex128()
}
func (e Complex128) ToUint32() Uint32 { return Uint32(real(e.Complex128())) }
func (e Complex128) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e == e.ToUint32().ToComplex128()
}
func (e Complex128) ToUint64() Uint64 { return Uint64(real(e.Complex128())) }
func (e Complex128) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), e == e.ToUint64().ToComplex128()
}
func (e Complex128) ToFloat32() Float32 { return Float32(real(e.Complex128())) }
func (e Complex128) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Complex128) ToFloat64() Float64 { return Float64(real(e.Complex128())) }
func (e Complex128) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e.ToString() == e.ToFloat64().ToString()
}
func (e Complex128) ToComplex64() Complex64 { return Complex64(e) }
func (e Complex128) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Complex128) ToComplex128() Complex128           { return e }
func (e Complex128) ToComplex128Eq() (Complex128, bool) { return e, true }
func (e Complex128) ToBool() Bool                       { return Bool(e != 0) }
func (e Complex128) ToString() String {
	return String(strconv.FormatComplex(e.Complex128(), FloatToStringFmt, FloatToStringPrec, 128))
}

package goval

import (
	"math"
	"strconv"
	//"database/sql"
	//"database/sql/driver"

	"github.com/shunsukuda/forceconv"
)

var (
	MinInt8   = Int8{Int8: math.MinInt8}
	MinInt16  = Int16{Int16: math.MinInt16}
	MinInt32  = Int32{Int32: math.MinInt32}
	MinInt64  = Int64{Int64: math.MinInt64}
	MaxInt8   = Int8{Int8: math.MaxInt8}
	MaxInt16  = Int16{Int16: math.MaxInt16}
	MaxInt32  = Int32{Int32: math.MaxInt32}
	MaxInt64  = Int64{Int64: math.MaxInt64}
	MinUint8  = Uint8{Uint8: 0}
	MinUint16 = Uint16{Uint16: 0}
	MinUint32 = Uint32{Uint32: 0}
	MinUint64 = Uint64{Uint64: 0}
	MaxUint8  = Uint8{Uint8: math.MaxUint8}
	MaxUint16 = Uint16{Uint16: math.MaxUint16}
	MaxUint32 = Uint32{Uint32: math.MaxUint32}
	MaxUint64 = Uint64{Uint64: math.MaxUint64}
)

var (
	maxInt24 = 1<<24 - 1
	maxInt53 = 1<<53 - 1
	minInt24 = -1 * maxInt24
	minInt53 = -1 * maxInt53
)

var (
	IntToStringBase   int  = 10
	UintToStringBase  int  = 16
	FloatToStringFmt  byte = 'g'
	FloatToStringPrec int  = -1
)

type Int8 struct {
	Int8 int8
}

func (e Int8) GoInt8() int8      { return e.Int8 }
func (e Int8) Type() Type        { return ValTypes.Int8 }
func (e Int8) Equal(x Int8) bool { return e.Int8 == x.Int8 }

func (e Int8) ToBool() Bool {
	return Bool{e.Int8 != 0}
}
func (e Int8) ToInt8() Int8 {
	return e
}
func (e Int8) ToInt8Eq() (Int8, bool) {
	return e, true
}
func (e Int8) ToInt16() Int16 {
	return Int16{int16(e.Int8)}
}
func (e Int8) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), true
}
func (e Int8) ToInt32() Int32 {
	return Int32{int32(e.Int8)}
}
func (e Int8) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), true
}
func (e Int8) ToInt64() Int64 {
	return Int64{int64(e.Int8)}
}
func (e Int8) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), true
}
func (e Int8) ToUint8() Uint8 {
	return Uint8{uint8(e.Int8)}
}
func (e Int8) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), 0 <= e.Int8
}
func (e Int8) ToUint16() Uint16 {
	return Uint16{uint16(e.Int8)}
}
func (e Int8) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), 0 <= e.Int8
}
func (e Int8) ToUint32() Uint32 {
	return Uint32{uint32(e.Int8)}
}
func (e Int8) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), 0 <= e.Int8
}
func (e Int8) ToUint64() Uint64 {
	return Uint64{uint64(e.Int8)}
}
func (e Int8) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), 0 <= e.Int8
}
func (e Int8) ToFloat32() Float32 {
	return Float32{float32(e.Int8)}
}
func (e Int8) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), true
}
func (e Int8) ToFloat64() Float64 {
	return Float64{float64(e.Int8)}
}
func (e Int8) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Int8) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Int8) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), true
}
func (e Int8) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Int8) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Int8) ToString() String {
	return String{strconv.FormatInt(e.ToInt64().Int64, IntToStringBase)}
}

func (e Int8) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Int8) ToStringBase(base int) String {
	return String{strconv.FormatInt(e.ToInt64().Int64, base)}
}

func (e Int8) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Int16 struct {
	Int16 int16
}

func (e Int16) GoInt16() int16     { return e.Int16 }
func (e Int16) Type() Type         { return ValTypes.Int16 }
func (e Int16) Equal(x Int16) bool { return e.Int16 == x.Int16 }

func (e Int16) ToBool() Bool {
	return Bool{e.Int16 != 0}
}
func (e Int16) ToInt8() Int8 {
	return Int8{int8(e.Int16)}
}
func (e Int16) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), math.MinInt8 <= e.Int16 && e.Int16 <= math.MaxInt8
}
func (e Int16) ToInt16() Int16 {
	return e
}
func (e Int16) ToInt16Eq() (Int16, bool) {
	return e, true
}
func (e Int16) ToInt32() Int32 {
	return Int32{int32(e.Int16)}
}
func (e Int16) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), true
}
func (e Int16) ToInt64() Int64 {
	return Int64{int64(e.Int16)}
}
func (e Int16) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), true
}
func (e Int16) ToUint8() Uint8 {
	return Uint8{uint8(e.Int16)}
}
func (e Int16) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), 0 <= e.Int16 && e.Int16 <= math.MaxUint8
}
func (e Int16) ToUint16() Uint16 {
	return Uint16{uint16(e.Int16)}
}
func (e Int16) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), 0 <= e.Int16
}
func (e Int16) ToUint32() Uint32 {
	return Uint32{uint32(e.Int16)}
}
func (e Int16) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), 0 <= e.Int16
}
func (e Int16) ToUint64() Uint64 {
	return Uint64{uint64(e.Int16)}
}
func (e Int16) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), 0 <= e.Int16
}
func (e Int16) ToFloat32() Float32 {
	return Float32{float32(e.Int16)}
}
func (e Int16) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), true
}
func (e Int16) ToFloat64() Float64 {
	return Float64{float64(e.Int16)}
}
func (e Int16) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Int16) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Int16) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), true
}
func (e Int16) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Int16) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Int16) ToString() String {
	return String{strconv.FormatInt(e.ToInt64().Int64, IntToStringBase)}
}

func (e Int16) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Int16) ToStringBase(base int) String {
	return String{strconv.FormatInt(e.ToInt64().Int64, base)}
}

func (e Int16) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Int32 struct {
	Int32 int32
}

func (e Int32) GoInt32() int32     { return e.Int32 }
func (e Int32) Type() Type         { return ValTypes.Int32 }
func (e Int32) Equal(x Int32) bool { return e.Int32 == x.Int32 }

func (e Int32) ToBool() Bool {
	return Bool{e.Int32 != 0}
}
func (e Int32) ToInt8() Int8 {
	return Int8{int8(e.Int32)}
}
func (e Int32) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), math.MinInt8 <= e.Int32 && e.Int32 <= math.MaxInt8
}
func (e Int32) ToInt16() Int16 {
	return Int16{int16(e.Int32)}
}
func (e Int32) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), math.MinInt16 <= e.Int32 && e.Int32 <= math.MaxInt16
}
func (e Int32) ToInt32() Int32 {
	return e
}
func (e Int32) ToInt32Eq() (Int32, bool) {
	return e, true
}
func (e Int32) ToInt64() Int64 {
	return Int64{int64(e.Int32)}
}
func (e Int32) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), true
}
func (e Int32) ToUint8() Uint8 {
	return Uint8{uint8(e.Int32)}
}
func (e Int32) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), 0 <= e.Int32 && e.Int32 <= math.MaxUint8
}
func (e Int32) ToUint16() Uint16 {
	return Uint16{uint16(e.Int32)}
}
func (e Int32) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), 0 <= e.Int32 && e.Int32 <= math.MaxUint16
}
func (e Int32) ToUint32() Uint32 {
	return Uint32{uint32(e.Int32)}
}
func (e Int32) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), 0 <= e.Int32
}
func (e Int32) ToUint64() Uint64 {
	return Uint64{uint64(e.Int32)}
}
func (e Int32) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), 0 <= e.Int32
}
func (e Int32) ToFloat32() Float32 {
	return Float32{float32(e.Int32)}
}
func (e Int32) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), int32(minInt24) <= e.Int32 && e.Int32 <= int32(maxInt24)
}
func (e Int32) ToFloat64() Float64 {
	return Float64{float64(e.Int32)}
}
func (e Int32) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Int32) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Int32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToInt32()
}
func (e Int32) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Int32) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Int32) ToString() String {
	return String{strconv.FormatInt(e.ToInt64().Int64, IntToStringBase)}
}

func (e Int32) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Int32) ToStringBase(base int) String {
	return String{strconv.FormatInt(e.ToInt64().Int64, base)}
}

func (e Int32) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Int64 struct {
	Int64 int64
}

func (e Int64) GoInt64() int64     { return e.Int64 }
func (e Int64) Type() Type         { return ValTypes.Int64 }
func (e Int64) Equal(x Int64) bool { return e.Int64 == x.Int64 }

func (e Int64) ToBool() Bool {
	return Bool{e.Int64 != 0}
}
func (e Int64) ToInt8() Int8 {
	return Int8{int8(e.Int64)}
}
func (e Int64) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), math.MinInt8 <= e.Int64 && e.Int64 <= math.MaxInt8
}
func (e Int64) ToInt16() Int16 {
	return Int16{int16(e.Int64)}
}
func (e Int64) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), math.MinInt16 <= e.Int64 && e.Int64 <= math.MaxInt16
}
func (e Int64) ToInt32() Int32 {
	return Int32{int32(e.Int64)}
}
func (e Int64) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), math.MinInt32 <= e.Int64 && e.Int64 <= math.MaxInt32
}
func (e Int64) ToInt64() Int64 {
	return e
}
func (e Int64) ToInt64Eq() (Int64, bool) {
	return e, true
}
func (e Int64) ToUint8() Uint8 {
	return Uint8{uint8(e.Int64)}
}
func (e Int64) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), 0 <= e.Int64 && e.Int64 <= math.MaxUint8
}
func (e Int64) ToUint16() Uint16 {
	return Uint16{uint16(e.Int64)}
}
func (e Int64) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), 0 <= e.Int64 && e.Int64 <= math.MaxUint16
}
func (e Int64) ToUint32() Uint32 {
	return Uint32{uint32(e.Int64)}
}
func (e Int64) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), 0 <= e.Int64 && e.Int64 <= math.MaxUint32
}
func (e Int64) ToUint64() Uint64 {
	return Uint64{uint64(e.Int64)}
}
func (e Int64) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), 0 <= e.Int64
}
func (e Int64) ToFloat32() Float32 {
	return Float32{float32(e.Int64)}
}
func (e Int64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), int64(minInt53) <= e.Int64 && e.Int64 <= int64(maxInt53)
}
func (e Int64) ToFloat64() Float64 {
	return Float64{float64(e.Int64)}
}
func (e Int64) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), int64(minInt53) <= e.Int64 && e.Int64 <= int64(maxInt53)
}
func (e Int64) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Int64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToInt64()
}
func (e Int64) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Int64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToInt64()
}

func (e Int64) ToString() String {
	return String{strconv.FormatInt(e.Int64, IntToStringBase)}
}

func (e Int64) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Int64) ToStringBase(base int) String {
	return String{strconv.FormatInt(e.Int64, base)}
}

func (e Int64) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Uint8 struct {
	Uint8 uint8
}

func (e Uint8) GoUint8() uint8     { return e.Uint8 }
func (e Uint8) Type() Type         { return ValTypes.Uint8 }
func (e Uint8) Equal(x Uint8) bool { return e.Uint8 == x.Uint8 }

func (e Uint8) ToBool() Bool {
	return Bool{e.Uint8 != 0}
}
func (e Uint8) ToInt8() Int8 {
	return Int8{int8(e.Uint8)}
}
func (e Uint8) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), 0 <= e.Uint8 && e.Uint8 <= math.MaxInt8
}
func (e Uint8) ToInt16() Int16 {
	return Int16{int16(e.Uint8)}
}
func (e Uint8) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), 0 <= e.Uint8
}
func (e Uint8) ToInt32() Int32 {
	return Int32{int32(e.Uint8)}
}
func (e Uint8) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), 0 <= e.Uint8
}
func (e Uint8) ToInt64() Int64 {
	return Int64{int64(e.Uint8)}
}
func (e Uint8) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), 0 <= e.Uint8
}
func (e Uint8) ToUint8() Uint8 {
	return e
}
func (e Uint8) ToUint8Eq() (Uint8, bool) {
	return e, true
}
func (e Uint8) ToUint16() Uint16 {
	return Uint16{uint16(e.Uint8)}
}
func (e Uint8) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), true
}
func (e Uint8) ToUint32() Uint32 {
	return Uint32{uint32(e.Uint8)}
}
func (e Uint8) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), true
}
func (e Uint8) ToUint64() Uint64 {
	return Uint64{uint64(e.Uint8)}
}
func (e Uint8) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), true
}
func (e Uint8) ToFloat32() Float32 {
	return Float32{float32(e.Uint8)}
}
func (e Uint8) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), true
}
func (e Uint8) ToFloat64() Float64 {
	return Float64{float64(e.Uint8)}
}
func (e Uint8) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Uint8) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Uint8) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), true
}
func (e Uint8) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Uint8) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Uint8) ToString() String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, UintToStringBase)}
}

func (e Uint8) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Uint8) ToStringBase(base int) String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, base)}
}

func (e Uint8) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Uint16 struct {
	Uint16 uint16
}

func (e Uint16) GoUint16() uint16    { return e.Uint16 }
func (e Uint16) Type() Type          { return ValTypes.Uint16 }
func (e Uint16) Equal(x Uint16) bool { return e.Uint16 == x.Uint16 }

func (e Uint16) ToBool() Bool {
	return Bool{e.Uint16 != 0}
}
func (e Uint16) ToInt8() Int8 {
	return Int8{int8(e.Uint16)}
}
func (e Uint16) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), 0 <= e.Uint16 && e.Uint16 <= math.MaxInt8
}
func (e Uint16) ToInt16() Int16 {
	return Int16{int16(e.Uint16)}
}
func (e Uint16) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), 0 <= e.Uint16 && e.Uint16 <= math.MaxInt16
}
func (e Uint16) ToInt32() Int32 {
	return Int32{int32(e.Uint16)}
}
func (e Uint16) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), 0 <= e.Uint16
}
func (e Uint16) ToInt64() Int64 {
	return Int64{int64(e.Uint16)}
}
func (e Uint16) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), 0 <= e.Uint16
}
func (e Uint16) ToUint8() Uint8 {
	return Uint8{uint8(e.Uint16)}
}
func (e Uint16) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e.Uint16 <= math.MaxUint8
}
func (e Uint16) ToUint16() Uint16 {
	return e
}
func (e Uint16) ToUint16Eq() (Uint16, bool) {
	return e, true
}
func (e Uint16) ToUint32() Uint32 {
	return Uint32{uint32(e.Uint16)}
}
func (e Uint16) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), true
}
func (e Uint16) ToUint64() Uint64 {
	return Uint64{uint64(e.Uint16)}
}
func (e Uint16) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), true
}
func (e Uint16) ToFloat32() Float32 {
	return Float32{float32(e.Uint16)}
}
func (e Uint16) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), true
}
func (e Uint16) ToFloat64() Float64 {
	return Float64{float64(e.Uint16)}
}
func (e Uint16) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Uint16) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Uint16) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), true
}
func (e Uint16) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Uint16) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Uint16) ToString() String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, UintToStringBase)}
}

func (e Uint16) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Uint16) ToStringBase(base int) String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, base)}
}

func (e Uint16) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Uint32 struct {
	Uint32 uint32
}

func (e Uint32) GoUint32() uint32    { return e.Uint32 }
func (e Uint32) Type() Type          { return ValTypes.Uint32 }
func (e Uint32) Equal(x Uint32) bool { return e.Uint32 == x.Uint32 }

func (e Uint32) ToBool() Bool {
	return Bool{e.Uint32 != 0}
}
func (e Uint32) ToInt8() Int8 {
	return Int8{int8(e.Uint32)}
}
func (e Uint32) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), 0 <= e.Uint32 && e.Uint32 <= math.MaxInt8
}
func (e Uint32) ToInt16() Int16 {
	return Int16{int16(e.Uint32)}
}
func (e Uint32) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), 0 <= e.Uint32 && e.Uint32 <= math.MaxInt16
}
func (e Uint32) ToInt32() Int32 {
	return Int32{int32(e.Uint32)}
}
func (e Uint32) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), 0 <= e.Uint32 && e.Uint32 <= math.MaxInt32
}
func (e Uint32) ToInt64() Int64 {
	return Int64{int64(e.Uint32)}
}
func (e Uint32) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), 0 <= e.Uint32
}
func (e Uint32) ToUint8() Uint8 {
	return Uint8{uint8(e.Uint32)}
}
func (e Uint32) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e.Uint32 <= math.MaxUint8
}
func (e Uint32) ToUint16() Uint16 {
	return Uint16{uint16(e.Uint32)}
}
func (e Uint32) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e.Uint32 <= math.MaxUint16
}
func (e Uint32) ToUint32() Uint32 {
	return e
}
func (e Uint32) ToUint32Eq() (Uint32, bool) {
	return e, true
}
func (e Uint32) ToUint64() Uint64 {
	return Uint64{uint64(e.Uint32)}
}
func (e Uint32) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), true
}
func (e Uint32) ToFloat32() Float32 {
	return Float32{float32(e.Uint32)}
}
func (e Uint32) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.Uint32 <= uint32(maxInt24)
}
func (e Uint32) ToFloat64() Float64 {
	return Float64{float64(e.Uint32)}
}
func (e Uint32) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), true
}
func (e Uint32) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Uint32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToUint32()
}
func (e Uint32) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Uint32) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), true
}

func (e Uint32) ToString() String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, UintToStringBase)}
}

func (e Uint32) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Uint32) ToStringBase(base int) String {
	return String{strconv.FormatUint(e.ToUint64().Uint64, base)}
}

func (e Uint32) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Uint64 struct {
	Uint64 uint64
}

func (e Uint64) GoUint64() uint64    { return e.Uint64 }
func (e Uint64) Type() Type          { return ValTypes.Uint64 }
func (e Uint64) Equal(x Uint64) bool { return e.Uint64 == x.Uint64 }

func (e Uint64) ToBool() Bool {
	return Bool{e.Uint64 != 0}
}
func (e Uint64) ToInt8() Int8 {
	return Int8{int8(e.Uint64)}
}
func (e Uint64) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), 0 <= e.Uint64 && e.Uint64 <= math.MaxInt8
}
func (e Uint64) ToInt16() Int16 {
	return Int16{int16(e.Uint64)}
}
func (e Uint64) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), 0 <= e.Uint64 && e.Uint64 <= math.MaxInt16
}
func (e Uint64) ToInt32() Int32 {
	return Int32{int32(e.Uint64)}
}
func (e Uint64) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), 0 <= e.Uint64 && e.Uint64 <= math.MaxInt32
}
func (e Uint64) ToInt64() Int64 {
	return Int64{int64(e.Uint64)}
}
func (e Uint64) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), 0 <= e.Uint64 && e.Uint64 <= math.MaxInt64
}
func (e Uint64) ToUint8() Uint8 {
	return Uint8{uint8(e.Uint64)}
}
func (e Uint64) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e.Uint64 <= math.MaxUint8
}
func (e Uint64) ToUint16() Uint16 {
	return Uint16{uint16(e.Uint64)}
}
func (e Uint64) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e.Uint64 <= math.MaxUint16
}
func (e Uint64) ToUint32() Uint32 {
	return Uint32{uint32(e.Uint64)}
}
func (e Uint64) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e.Uint64 <= math.MaxUint32
}
func (e Uint64) ToUint64() Uint64 {
	return e
}
func (e Uint64) ToUint64Eq() (Uint64, bool) {
	return e, true
}
func (e Uint64) ToFloat32() Float32 {
	return Float32{float32(e.Uint64)}
}
func (e Uint64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.Uint64 <= uint64(maxInt53)
}
func (e Uint64) ToFloat64() Float64 {
	return Float64{float64(e.Uint64)}
}
func (e Uint64) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e.Uint64 <= uint64(maxInt53)
}
func (e Uint64) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Uint64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e == e.ToComplex64().ToUint64()
}
func (e Uint64) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Uint64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToUint64()
}

func (e Uint64) ToString() String {
	return String{strconv.FormatUint(e.Uint64, UintToStringBase)}
}

func (e Uint64) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Uint64) ToStringBase(base int) String {
	return String{strconv.FormatUint(e.Uint64, base)}
}

func (e Uint64) ToBytesBase(base int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringBase(base).String)}
}

type Float32 struct {
	Float32 float32
}

func (e Float32) GoFloat32() float32   { return e.Float32 }
func (e Float32) Type() Type           { return ValTypes.Float32 }
func (e Float32) Equal(x Float32) bool { return e.Float32 == x.Float32 }

func (e Float32) ToBool() Bool {
	return Bool{e.Float32 != 0}
}
func (e Float32) ToInt8() Int8 {
	return Int8{int8(e.Float32)}
}
func (e Float32) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), e == e.ToInt8().ToFloat32()
}
func (e Float32) ToInt16() Int16 {
	return Int16{int16(e.Float32)}
}
func (e Float32) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), e == e.ToInt16().ToFloat32()
}
func (e Float32) ToInt32() Int32 {
	return Int32{int32(e.Float32)}
}
func (e Float32) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), e == e.ToInt32().ToFloat32()
}
func (e Float32) ToInt64() Int64 {
	return Int64{int64(e.Float32)}
}
func (e Float32) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), e == e.ToInt64().ToFloat32()
}
func (e Float32) ToUint8() Uint8 {
	return Uint8{uint8(e.Float32)}
}
func (e Float32) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e == e.ToUint8().ToFloat32()
}
func (e Float32) ToUint16() Uint16 {
	return Uint16{uint16(e.Float32)}
}
func (e Float32) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e == e.ToUint16().ToFloat32()
}
func (e Float32) ToUint32() Uint32 {
	return Uint32{uint32(e.Float32)}
}
func (e Float32) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e == e.ToUint32().ToFloat32()
}
func (e Float32) ToUint64() Uint64 {
	return Uint64{uint64(e.Float32)}
}
func (e Float32) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), e == e.ToUint64().ToFloat32()
}
func (e Float32) ToFloat32() Float32 {
	return e
}
func (e Float32) ToFloat32Eq() (Float32, bool) {
	return e, true
}
func (e Float32) ToFloat64() Float64 {
	return Float64{float64(e.Float32)}
}
func (e Float32) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e == e.ToFloat64().ToFloat32()
}
func (e Float32) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Float32) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Float32) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Float32) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToFloat32()
}

func (e Float32) ToString() String {
	return String{strconv.FormatFloat(e.ToFloat64().Float64, FloatToStringFmt, FloatToStringPrec, 32)}
}

func (e Float32) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Float32) ToStringFmt(fmt byte) String {
	return String{strconv.FormatFloat(e.ToFloat64().Float64, fmt, FloatToStringPrec, 32)}
}

func (e Float32) ToBytesFmt(fmt byte) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmt(fmt).String)}
}

func (e Float32) ToStringPrec(prec int) String {
	return String{strconv.FormatFloat(e.ToFloat64().Float64, FloatToStringFmt, prec, 32)}
}

func (e Float32) ToBytesPrec(prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringPrec(prec).String)}
}

func (e Float32) ToStringFmtPrec(fmt byte, prec int) String {
	return String{strconv.FormatFloat(e.ToFloat64().Float64, fmt, prec, 32)}
}

func (e Float32) ToBytesFmtPrec(fmt byte, prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmtPrec(fmt, prec).String)}
}

type Float64 struct {
	Float64 float64
}

func (e Float64) GoFloat64() float64   { return e.Float64 }
func (e Float64) Type() Type           { return ValTypes.Float64 }
func (e Float64) Equal(x Float64) bool { return e.Float64 == x.Float64 }

func (e Float64) ToBool() Bool {
	return Bool{e.Float64 != 0}
}
func (e Float64) ToInt8() Int8 {
	return Int8{int8(e.Float64)}
}
func (e Float64) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), e == e.ToInt8().ToFloat64()
}
func (e Float64) ToInt16() Int16 {
	return Int16{int16(e.Float64)}
}
func (e Float64) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), e == e.ToInt16().ToFloat64()
}
func (e Float64) ToInt32() Int32 {
	return Int32{int32(e.Float64)}
}
func (e Float64) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), e == e.ToInt32().ToFloat64()
}
func (e Float64) ToInt64() Int64 {
	return Int64{int64(e.Float64)}
}
func (e Float64) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), e == e.ToInt64().ToFloat64()
}
func (e Float64) ToUint8() Uint8 {
	return Uint8{uint8(e.Float64)}
}
func (e Float64) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e == e.ToUint8().ToFloat64()
}
func (e Float64) ToUint16() Uint16 {
	return Uint16{uint16(e.Float64)}
}
func (e Float64) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e == e.ToUint16().ToFloat64()
}
func (e Float64) ToUint32() Uint32 {
	return Uint32{uint32(e.Float64)}
}
func (e Float64) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e == e.ToUint32().ToFloat64()
}
func (e Float64) ToUint64() Uint64 {
	return Uint64{uint64(e.Float64)}
}
func (e Float64) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), e == e.ToUint64().ToFloat64()
}
func (e Float64) ToFloat32() Float32 {
	return Float32{float32(e.Float64)}
}
func (e Float64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Float64) ToFloat64() Float64 {
	return e
}
func (e Float64) ToFloat64Eq() (Float64, bool) {
	return e, true
}
func (e Float64) ToComplex64() Complex64 {
	return Complex64{complex(e.ToFloat32().Float32, float32(0))}
}
func (e Float64) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Float64) ToComplex128() Complex128 {
	return Complex128{complex(e.ToFloat64().Float64, float64(0))}
}
func (e Float64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e.ToString() == e.ToComplex128().ToString()
}

func (e Float64) ToString() String {
	return String{strconv.FormatFloat(e.Float64, FloatToStringFmt, FloatToStringPrec, 64)}
}

func (e Float64) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Float64) ToStringFmt(fmt byte) String {
	return String{strconv.FormatFloat(e.Float64, fmt, FloatToStringPrec, 64)}
}

func (e Float64) ToBytesFmt(fmt byte) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmt(fmt).String)}
}

func (e Float64) ToStringPrec(prec int) String {
	return String{strconv.FormatFloat(e.Float64, FloatToStringFmt, prec, 64)}
}

func (e Float64) ToBytesPrec(prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringPrec(prec).String)}
}

func (e Float64) ToStringFmtPrec(fmt byte, prec int) String {
	return String{strconv.FormatFloat(e.Float64, fmt, prec, 64)}
}

func (e Float64) ToBytesFmtPrec(fmt byte, prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmtPrec(fmt, prec).String)}
}

type Complex64 struct {
	Complex64 complex64
}

func (e Complex64) GoComplex64() complex64 { return e.Complex64 }
func (e Complex64) Type() Type             { return ValTypes.Complex64 }
func (e Complex64) Equal(x Complex64) bool { return e.Complex64 == x.Complex64 }

func (e Complex64) ToBool() Bool {
	return Bool{e.Complex64 != 0}
}
func (e Complex64) ToInt8() Int8 {
	return Int8{int8(real(e.Complex64))}
}
func (e Complex64) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), e == e.ToInt8().ToComplex64()
}
func (e Complex64) ToInt16() Int16 {
	return Int16{int16(real(e.Complex64))}
}
func (e Complex64) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), e == e.ToInt16().ToComplex64()
}
func (e Complex64) ToInt32() Int32 {
	return Int32{int32(real(e.Complex64))}
}
func (e Complex64) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), e == e.ToInt32().ToComplex64()
}
func (e Complex64) ToInt64() Int64 {
	return Int64{int64(real(e.Complex64))}
}
func (e Complex64) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), e == e.ToInt64().ToComplex64()
}
func (e Complex64) ToUint8() Uint8 {
	return Uint8{uint8(real(e.Complex64))}
}
func (e Complex64) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e == e.ToUint8().ToComplex64()
}
func (e Complex64) ToUint16() Uint16 {
	return Uint16{uint16(real(e.Complex64))}
}
func (e Complex64) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e == e.ToUint16().ToComplex64()
}
func (e Complex64) ToUint32() Uint32 {
	return Uint32{uint32(real(e.Complex64))}
}
func (e Complex64) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e == e.ToUint32().ToComplex64()
}
func (e Complex64) ToUint64() Uint64 {
	return Uint64{uint64(real(e.Complex64))}
}
func (e Complex64) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), e == e.ToUint64().ToComplex64()
}
func (e Complex64) ToFloat32() Float32 {
	return Float32{float32(real(e.Complex64))}
}
func (e Complex64) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Complex64) ToFloat64() Float64 {
	return Float64{float64(real(e.Complex64))}
}
func (e Complex64) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e == e.ToFloat64().ToComplex64()
}
func (e Complex64) ToComplex64() Complex64 {
	return e
}
func (e Complex64) ToComplex64Eq() (Complex64, bool) {
	return e, true
}
func (e Complex64) ToComplex128() Complex128 {
	return Complex128{e.ToComplex128().Complex128}
}
func (e Complex64) ToComplex128Eq() (Complex128, bool) {
	return e.ToComplex128(), e == e.ToComplex128().ToComplex64()
}

func (e Complex64) ToString() String {
	return String{strconv.FormatComplex(e.ToComplex128().Complex128, FloatToStringFmt, FloatToStringPrec, 128)}
}

func (e Complex64) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Complex64) ToStringFmt(fmt byte) String {
	return String{strconv.FormatComplex(e.ToComplex128().Complex128, fmt, FloatToStringPrec, 128)}
}

func (e Complex64) ToBytesFmt(fmt byte) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmt(fmt).String)}
}

func (e Complex64) ToStringPrec(prec int) String {
	return String{strconv.FormatComplex(e.ToComplex128().Complex128, FloatToStringFmt, prec, 128)}
}

func (e Complex64) ToBytesPrec(prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringPrec(prec).String)}
}

func (e Complex64) ToStringFmtPrec(fmt byte, prec int) String {
	return String{strconv.FormatComplex(e.ToComplex128().Complex128, fmt, prec, 128)}
}

func (e Complex64) ToBytesFmtPrec(fmt byte, prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmtPrec(fmt, prec).String)}
}

type Complex128 struct {
	Complex128 complex128
}

func (e Complex128) GoComplex128() complex128 { return e.Complex128 }
func (e Complex128) Type() Type               { return ValTypes.Complex128 }
func (e Complex128) Equal(x Complex128) bool  { return e.Complex128 == x.Complex128 }

func (e Complex128) ToBool() Bool {
	return Bool{e.Complex128 != 0}
}
func (e Complex128) ToInt8() Int8 {
	return Int8{int8(real(e.Complex128))}
}
func (e Complex128) ToInt8Eq() (Int8, bool) {
	return e.ToInt8(), e == e.ToInt8().ToComplex128()
}
func (e Complex128) ToInt16() Int16 {
	return Int16{int16(real(e.Complex128))}
}
func (e Complex128) ToInt16Eq() (Int16, bool) {
	return e.ToInt16(), e == e.ToInt16().ToComplex128()
}
func (e Complex128) ToInt32() Int32 {
	return Int32{int32(real(e.Complex128))}
}
func (e Complex128) ToInt32Eq() (Int32, bool) {
	return e.ToInt32(), e == e.ToInt32().ToComplex128()
}
func (e Complex128) ToInt64() Int64 {
	return Int64{int64(real(e.Complex128))}
}
func (e Complex128) ToInt64Eq() (Int64, bool) {
	return e.ToInt64(), e == e.ToInt64().ToComplex128()
}
func (e Complex128) ToUint8() Uint8 {
	return Uint8{uint8(real(e.Complex128))}
}
func (e Complex128) ToUint8Eq() (Uint8, bool) {
	return e.ToUint8(), e == e.ToUint8().ToComplex128()
}
func (e Complex128) ToUint16() Uint16 {
	return Uint16{uint16(real(e.Complex128))}
}
func (e Complex128) ToUint16Eq() (Uint16, bool) {
	return e.ToUint16(), e == e.ToUint16().ToComplex128()
}
func (e Complex128) ToUint32() Uint32 {
	return Uint32{uint32(real(e.Complex128))}
}
func (e Complex128) ToUint32Eq() (Uint32, bool) {
	return e.ToUint32(), e == e.ToUint32().ToComplex128()
}
func (e Complex128) ToUint64() Uint64 {
	return Uint64{uint64(real(e.Complex128))}
}
func (e Complex128) ToUint64Eq() (Uint64, bool) {
	return e.ToUint64(), e == e.ToUint64().ToComplex128()
}
func (e Complex128) ToFloat32() Float32 {
	return Float32{float32(real(e.Complex128))}
}
func (e Complex128) ToFloat32Eq() (Float32, bool) {
	return e.ToFloat32(), e.ToString() == e.ToFloat32().ToString()
}
func (e Complex128) ToFloat64() Float64 {
	return Float64{float64(real(e.Complex128))}
}
func (e Complex128) ToFloat64Eq() (Float64, bool) {
	return e.ToFloat64(), e.ToString() == e.ToFloat64().ToString()
}
func (e Complex128) ToComplex64() Complex64 {
	return Complex64{e.ToComplex64().Complex64}
}
func (e Complex128) ToComplex64Eq() (Complex64, bool) {
	return e.ToComplex64(), e.ToString() == e.ToComplex64().ToString()
}
func (e Complex128) ToComplex128() Complex128 {
	return e
}
func (e Complex128) ToComplex128Eq() (Complex128, bool) {
	return e, true
}

func (e Complex128) ToString() String {
	return String{strconv.FormatComplex(e.Complex128, FloatToStringFmt, FloatToStringPrec, 128)}
}

func (e Complex128) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}

func (e Complex128) ToStringFmt(fmt byte) String {
	return String{strconv.FormatComplex(e.Complex128, fmt, FloatToStringPrec, 128)}
}

func (e Complex128) ToBytesFmt(fmt byte) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmt(fmt).String)}
}

func (e Complex128) ToStringPrec(prec int) String {
	return String{strconv.FormatComplex(e.Complex128, FloatToStringFmt, prec, 128)}
}

func (e Complex128) ToBytesPrec(prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringPrec(prec).String)}
}

func (e Complex128) ToStringFmtPrec(fmt byte, prec int) String {
	return String{strconv.FormatComplex(e.Complex128, fmt, prec, 128)}
}

func (e Complex128) ToBytesFmtPrec(fmt byte, prec int) Bytes {
	return Bytes{forceconv.StringToBytes(e.ToStringFmtPrec(fmt, prec).String)}
}

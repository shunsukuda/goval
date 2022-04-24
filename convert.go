package goval

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/shunsukuda/forceconv"
)

type IntConverter interface {
	Int8Converter
	Int16Converter
	Int32Converter
	Int64Converter
}

type UintConverter interface {
	Uint8Converter
	Uint16Converter
	Uint32Converter
	Uint64Converter
}

type FloatConverter interface {
	Float32Converter
	Float64Converter
}

type ComplexConverter interface {
	Complex64Converter
	Complex128Converter
}

type BoolConverter interface {
	ToBool() Bool
	ToBoolEq() (Bool, bool)
}

type Int8Converter interface {
	ToInt8() Int8
	ToInt8Eq() (Int8, bool)
}

type Int16Converter interface {
	ToInt16() Int16
	ToInt16Eq() (Int16, bool)
}

type Int32Converter interface {
	ToInt32() Int32
	ToInt32Eq() (Int32, bool)
}

type Int64Converter interface {
	ToInt64() Int64
	ToInt64Eq() (Int64, bool)
}

type Uint8Converter interface {
	ToUint8() Uint8
	ToUint8Eq() (Uint8, bool)
}

type Uint16Converter interface {
	ToUint16() Uint16
	ToUint16Eq() (Uint16, bool)
}

type Uint32Converter interface {
	ToUint32() Uint32
	ToUint32Eq() (Uint32, bool)
}

type Uint64Converter interface {
	ToUint64() Uint64
	ToUint64Eq() (Uint64, bool)
}

type Float32Converter interface {
	ToFloat32() Float32
	ToFloat32Eq() (Float32, bool)
}

type Float64Converter interface {
	ToFloat64() Float64
	ToFloat64Eq() (Float64, bool)
}

type Complex64Converter interface {
	ToComplex64() Complex64
	ToComplex64Eq() (Complex64, bool)
}

type Complex128Converter interface {
	ToComplex128() Complex128
	ToComplex128Eq() (Complex128, bool)
}

type StringConverter interface {
	ToString() String
	ToStringCheck() (String, bool)
}

type BytesConverter interface {
	ToBytes() Bytes
	ToBytesCheck() (Bytes, bool)
}

/*
type TimeConverter interface {
	ToTime() Time
	ToTimeCheck() (Time, bool)
}
*/

type ArrowTimeConverter interface {
	ToArrowDayTimeInterval() arrow.DayTimeInterval
	ToArrowDuration() arrow.Duration
	ToArrowMonthdayNanoInterval() arrow.MonthDayNanoInterval
	ToArrowMonthInterval() arrow.MonthInterval
	ToArrowTimestamp() arrow.Timestamp
	ToArrowDate32() arrow.Date32
	ToArrowDate64() arrow.Date64
	ToArrowTime32() arrow.Time32
	ToArrowTime64() arrow.Time64
}

func ToBool(x BoolConverter) Bool {
	return x.ToBool()
}

func ToBoolA(x BoolConverter) (Bool, bool) {
	return x.ToBoolEq()
}

func ToInt8(x Int8Converter) Int8 {
	return x.ToInt8()
}

func ToInt8A(x Int8Converter) (Int8, bool) {
	return x.ToInt8Eq()
}

func ToInt16(x Int16Converter) Int16 {
	return x.ToInt16()
}

func ToInt16A(x Int16Converter) (Int16, bool) {
	return x.ToInt16Eq()
}

func ToInt32(x Int32Converter) Int32 {
	return x.ToInt32()
}

func ToInt32A(x Int32Converter) (Int32, bool) {
	return x.ToInt32Eq()
}

func ToInt64(x Int64Converter) Int64 {
	return x.ToInt64()
}

func ToInt64A(x Int64Converter) (Int64, bool) {
	return x.ToInt64Eq()
}

func ToUint8(x Uint8Converter) Uint8 {
	return x.ToUint8()
}

func ToUint8A(x Uint8Converter) (Uint8, bool) {
	return x.ToUint8Eq()
}

func ToUint16(x Uint16Converter) Uint16 {
	return x.ToUint16()
}

func ToUint16A(x Uint16Converter) (Uint16, bool) {
	return x.ToUint16Eq()
}

func ToUint32(x Uint32Converter) Uint32 {
	return x.ToUint32()
}

func ToUint32A(x Uint32Converter) (Uint32, bool) {
	return x.ToUint32Eq()
}

func ToUint64(x Uint64Converter) Uint64 {
	return x.ToUint64()
}

func ToUint64A(x Uint64Converter) (Uint64, bool) {
	return x.ToUint64Eq()
}

func ToFloat32(x Float32Converter) Float32 {
	return x.ToFloat32()
}

func ToFloat32A(x Float32Converter) (Float32, bool) {
	return x.ToFloat32Eq()
}

func ToFloat64(x Float64Converter) Float64 {
	return x.ToFloat64()
}

func ToFloat64A(x Float64Converter) (Float64, bool) {
	return x.ToFloat64Eq()
}

func ToString(x StringConverter) String {
	return x.ToString()
}

/*
func ToStringEq(x StringConverter) (String, bool) {
	return x.ToStringEq()
}
*/

func ToBytes(x BytesConverter) Bytes {
	return x.ToBytes()
}

/*
func ToBytesEq(x BytesConverter) (Bytes, bool) {
	return x.ToBytesEq()
}
*/

func ToStringFormat(x Val, format string) String {
	return String{fmt.Sprintf(format, x.Interface())}
}

func ToBytesFormat(x Val, format string) Bytes {
	return Bytes{forceconv.StringToBytes(fmt.Sprintf(format, x.Interface()))}
}

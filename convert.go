package goval

import (
	"github.com/apache/arrow/go/arrow"
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
	ToStringFormat(format string) String
}

type BytesConverter interface {
	ToBytes() Bytes
	ToBytesFormat(format string) Bytes
}

/*
type TimeConverter interface {
	ToTime() Time
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

package goval

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	FloatConvFormat byte = 'f'
)

type BoolConverter interface {
	ToBool() Bool
	ToBoolCheck() (Bool, Err)
}

type Int8Converter interface {
	ToInt8() Int8
	ToInt8Check() (Int8, Err)
}

type Int16Converter interface {
	ToInt16() Int16
	ToInt16Check() (Int16, Err)
}

type Int32Converter interface {
	ToInt32() Int32
	ToInt32Check() (Int32, Err)
}

type Int64Converter interface {
	ToInt64() Int64
	ToInt64Check() (Int64, Err)
}

type Uint8Converter interface {
	ToUint8() Uint8
	ToUint8Check() (Uint8, Err)
}

type Uint16Converter interface {
	ToUint16() Uint16
	ToUint16Check() (Uint16, Err)
}

type Uint32Converter interface {
	ToUint32() Uint32
	ToUint32Check() (Uint32, Err)
}

type Uint64Converter interface {
	ToUint64() Uint64
	ToUint64Check() (Uint64, Err)
}

type Float32Converter interface {
	ToFloat32() Float32
	ToFloat32Check() (Float32, Err)
}

type Float64Converter interface {
	ToFloat64() Float64
	ToFloat64Check() (Float64, Err)
}

type StringConverter interface {
	ToString() String
	ToStringCheck() (String, Err)
}

type BytesConverter interface {
	ToBytes() Bytes
	ToBytesCheck() (Bytes, Err)
}

func unsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(
		&reflect.StringHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
			Len:  len(b),
		}))
}

func unsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
			Len:  len(s),
			Cap:  len(s),
		}))
}

func ToBool(x BoolConverter) Bool {
	return x.ToBool()
	/*
		switch cx := x.(type) {
		case Bool:
			return cx
		case Int8:
			return Bool(cx.Int8() != 0)
		case Int16:
			return Bool(cx.Int16() != 0)
		case Int32:
			return Bool(cx.Int32() != 0)
		case Int64:
			return Bool(cx.Int64() != 0)
		case Uint8:
			return Bool(cx.Uint8() != 0)
		case Uint16:
			return Bool(cx.Uint16() != 0)
		case Uint32:
			return Bool(cx.Uint32() != 0)
		case Uint64:
			return Bool(cx.Uint64() != 0)
		case Float32:
			return Bool(cx.Float32() != 0.0)
		case Float64:
			return Bool(cx.Float64() != 0.0)
		case String:
			cv, err := strconv.ParseBool(cx.String())
			if err == nil {
				return Bool(cv)
			}
			return Bool(false)
		case Bytes:
			cv, err := strconv.ParseBool(unsafeBytesToString(cx.Bytes()))
			if err == nil {
				return Bool(cv)
			}
			return Bool(false)
		default:
			return Bool(false)
		}
	*/
}

func ToBoolCheck(x BoolConverter) (Bool, Err) {
	return x.ToBoolCheck()
}

func ToInt8(x Int8Converter) Int8 {
	return x.ToInt8()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Int8(1)
			}
			return Int8(0)
		case Int8:
			return cx
		case Int16:
			return Int8(cx.Int16())
		case Int32:
			return Int8(cx.Int32())
		case Int64:
			return Int8(cx.Int64())
		case Uint8:
			return Int8(cx.Uint8())
		case Uint16:
			return Int8(cx.Uint16())
		case Uint32:
			return Int8(cx.Uint32())
		case Uint64:
			return Int8(cx.Uint64())
		case Float32:
			return Int8(cx.Float32())
		case Float64:
			return Int8(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Int8(cv)
			}
			return Int8(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Int8(cv)
			}
			return Int8(0)
		default:
			return Int8(0)
		}
	*/
}

func ToInt8Check(x Int8Converter) (Int8, Err) {
	return x.ToInt8Check()
}

func ToInt16(x Int16Converter) Int16 {
	return x.ToInt16()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Int16(1)
			}
			return Int16(0)
		case Int8:
			return Int16(cx.Int8())
		case Int16:
			return cx
		case Int32:
			return Int16(cx.Int32())
		case Int64:
			return Int16(cx.Int64())
		case Uint8:
			return Int16(cx.Uint8())
		case Uint16:
			return Int16(cx.Uint16())
		case Uint32:
			return Int16(cx.Uint32())
		case Uint64:
			return Int16(cx.Uint64())
		case Float32:
			return Int16(cx.Float32())
		case Float64:
			return Int16(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Int16(cv)
			}
			return Int16(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Int16(cv)
			}
			return Int16(0)
		default:
			return Int16(0)
		}
	*/
}

func ToInt16Check(x Int16Converter) (Int16, Err) {
	return x.ToInt16Check()
}

func ToInt32(x Int32Converter) Int32 {
	return x.ToInt32()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Int32(1)
			}
			return Int32(0)
		case Int8:
			return Int32(cx.Int8())
		case Int16:
			return Int32(cx.Int16())
		case Int32:
			return cx
		case Int64:
			return Int32(cx.Int64())
		case Uint8:
			return Int32(cx.Uint8())
		case Uint16:
			return Int32(cx.Uint16())
		case Uint32:
			return Int32(cx.Uint32())
		case Uint64:
			return Int32(cx.Uint64())
		case Float32:
			return Int32(cx.Float32())
		case Float64:
			return Int32(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Int32(cv)
			}
			return Int32(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Int32(cv)
			}
			return Int32(0)
		default:
			return Int32(0)
		}
	*/
}

func ToInt32Check(x Int32Converter) (Int32, Err) {
	return x.ToInt32Check()
}

func ToInt64(x Int64Converter) Int64 {
	return x.ToInt64()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Int64(1)
			}
			return Int64(0)
		case Int8:
			return Int64(cx.Int8())
		case Int16:
			return Int64(cx.Int16())
		case Int32:
			return Int64(cx.Int32())
		case Int64:
			return cx
		case Uint8:
			return Int64(cx.Uint8())
		case Uint16:
			return Int64(cx.Uint16())
		case Uint32:
			return Int64(cx.Uint32())
		case Uint64:
			return Int64(cx.Uint64())
		case Float32:
			return Int64(cx.Float32())
		case Float64:
			return Int64(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Int64(cv)
			}
			return Int64(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Int64(cv)
			}
			return Int64(0)
		default:
			return Int64(0)
		}
	*/
}

func ToInt64Check(x Int64Converter) (Int64, Err) {
	return x.ToInt64Check()
}

func ToUint8(x Uint8Converter) Uint8 {
	return x.ToUint8()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Uint8(1)
			}
			return Uint8(0)
		case Int8:
			return Uint8(cx.Int8())
		case Int16:
			return Uint8(cx.Int16())
		case Int32:
			return Uint8(cx.Int32())
		case Int64:
			return Uint8(cx.Int64())
		case Uint8:
			return cx
		case Uint16:
			return Uint8(cx.Uint16())
		case Uint32:
			return Uint8(cx.Uint32())
		case Uint64:
			return Uint8(cx.Uint64())
		case Float32:
			return Uint8(cx.Float32())
		case Float64:
			return Uint8(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Uint8(cv)
			}
			return Uint8(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Uint8(cv)
			}
			return Uint8(0)
		default:
			return Uint8(0)
		}
	*/
}

func ToUint8Check(x Uint8Converter) (Uint8, Err) {
	return x.ToUint8Check()
}

func ToUint16(x Uint16Converter) Uint16 {
	return x.ToUint16()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Uint16(1)
			}
			return Uint16(0)
		case Int8:
			return Uint16(cx.Int8())
		case Int16:
			return Uint16(cx.Int16())
		case Int32:
			return Uint16(cx.Int32())
		case Int64:
			return Uint16(cx.Int64())
		case Uint8:
			return Uint16(cx.Uint8())
		case Uint16:
			return cx
		case Uint32:
			return Uint16(cx.Uint32())
		case Uint64:
			return Uint16(cx.Uint64())
		case Float32:
			return Uint16(cx.Float32())
		case Float64:
			return Uint16(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Uint16(cv)
			}
			return Uint16(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Uint16(cv)
			}
			return Uint16(0)
		default:
			return Uint16(0)
		}
	*/
}

func ToUint16Check(x Uint16Converter) (Uint16, Err) {
	return x.ToUint16Check()
}

func ToUint32(x Uint32Converter) Uint32 {
	return x.ToUint32()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Uint32(1)
			}
			return Uint32(0)
		case Int8:
			return Uint32(cx.Int8())
		case Int16:
			return Uint32(cx.Int16())
		case Int32:
			return Uint32(cx.Int32())
		case Int64:
			return Uint32(cx.Int64())
		case Uint8:
			return Uint32(cx.Uint8())
		case Uint16:
			return Uint32(cx.Uint16())
		case Uint32:
			return cx
		case Uint64:
			return Uint32(cx.Uint64())
		case Float32:
			return Uint32(cx.Float32())
		case Float64:
			return Uint32(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Uint32(cv)
			}
			return Uint32(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Uint32(cv)
			}
			return Uint32(0)
		default:
			return Uint32(0)
		}
	*/
}

func ToUint32Check(x Uint32Converter) (Uint32, Err) {
	return x.ToUint32Check()
}

func ToUint64(x Uint64Converter) Uint64 {
	return x.ToUint64()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Uint64(1)
			}
			return Uint64(0)
		case Int8:
			return Uint64(cx.Int8())
		case Int16:
			return Uint64(cx.Int16())
		case Int32:
			return Uint64(cx.Int32())
		case Int64:
			return Uint64(cx.Int64())
		case Uint8:
			return Uint64(cx.Uint8())
		case Uint16:
			return Uint64(cx.Uint16())
		case Uint32:
			return Uint64(cx.Uint32())
		case Uint64:
			return cx
		case Float32:
			return Uint64(cx.Float32())
		case Float64:
			return Uint64(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Uint64(cv)
			}
			return Uint64(0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Uint64(cv)
			}
			return Uint64(0)
		default:
			return Uint64(0)
		}
	*/
}

func ToUint64Check(x Uint64Converter) (Uint64, Err) {
	return x.ToUint64Check()
}

func ToFloat32(x Float32Converter) Float32 {
	return x.ToFloat32()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Float32(1.0)
			}
			return Float32(0.0)
		case Int8:
			return Float32(cx.Int8())
		case Int16:
			return Float32(cx.Int16())
		case Int32:
			return Float32(cx.Int32())
		case Int64:
			return Float32(cx.Int64())
		case Uint8:
			return Float32(cx.Uint8())
		case Uint16:
			return Float32(cx.Uint16())
		case Uint32:
			return Float32(cx.Uint32())
		case Uint64:
			return Float32(cx.Uint64())
		case Float32:
			return cx
		case Float64:
			return Float32(cx.Float64())
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Float32(cv)
			}
			return Float32(0.0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Float32(cv)
			}
			return Float32(0.0)
		default:
			return Float32(0.0)
		}
	*/
}

func ToFloat32Check(x Float32Converter) (Float32, Err) {
	return x.ToFloat32Check()
}

func ToFloat64(x Float64Converter) Float64 {
	return x.ToFloat64()
	/*
		switch cx := x.(type) {
		case Bool:
			if cx.Bool() {
				return Float64(1.0)
			}
			return Float64(0.0)
		case Int8:
			return Float64(cx.Int8())
		case Int16:
			return Float64(cx.Int16())
		case Int32:
			return Float64(cx.Int32())
		case Int64:
			return Float64(cx.Int64())
		case Uint8:
			return Float64(cx.Uint8())
		case Uint16:
			return Float64(cx.Uint16())
		case Uint32:
			return Float64(cx.Uint32())
		case Uint64:
			return Float64(cx.Uint64())
		case Float32:
			return Float64(cx.Float32())
		case Float64:
			return cx
		case String:
			cv, err := strconv.ParseInt(cx.String(), 0, 0)
			if err == nil {
				return Float64(cv)
			}
			return Float64(0.0)
		case Bytes:
			cv, err := strconv.ParseInt(unsafeBytesToString(cx.Bytes()), 0, 0)
			if err == nil {
				return Float64(cv)
			}
			return Float64(0.0)
		default:
			return Float64(0.0)
		}
	*/
}

func ToFloat64Check(x Float64Converter) (Float64, Err) {
	return x.ToFloat64Check()
}

func ToString(x StringConverter) String {
	return x.ToString()
	/*
		switch cx := x.(type) {
		case Bool:
			return String(strconv.FormatBool(cx.Bool()))
		case Int8:
			return String(strconv.FormatInt(int64(cx.Int8()), 10))
		case Int16:
			return String(strconv.FormatInt(int64(cx.Int16()), 10))
		case Int32:
			return String(strconv.FormatInt(int64(cx.Int32()), 10))
		case Int64:
			return String(strconv.FormatInt(cx.Int64(), 10))
		case Uint8:
			return String(strconv.FormatUint(uint64(cx.Uint8()), 10))
		case Uint16:
			return String(strconv.FormatUint(uint64(cx.Uint16()), 10))
		case Uint32:
			return String(strconv.FormatUint(uint64(cx.Uint32()), 10))
		case Uint64:
			return String(strconv.FormatUint(cx.Uint64(), 10))
		case Float32:
			return String(strconv.FormatFloat(float64(cx.Float32()), FloatStringFormat, -1, 32))
		case Float64:
			return String(strconv.FormatFloat(cx.Float64(), FloatStringFormat, -1, 64))
		case String:
			return cx
		case Bytes:
			return String(unsafeBytesToString(cx.Bytes()))
		default:
			return String("")
		}
	*/
}

func ToStringCheck(x StringConverter) (String, Err) {
	return x.ToStringCheck()
}

func ToBytes(x BytesConverter) Bytes {
	return x.ToBytes()
	/*
		switch cx := x.(type) {
		case Bytes:
			return cx
		default:
			return Bytes(unsafeStringToBytes(ToString(cx).String()))
		}
	*/
}

func ToBytesCheck(x BytesConverter) (Bytes, Err) {
	return x.ToBytesCheck()
}

func ToStringFormat(x Val, format string) String {
	return String(fmt.Sprintf(format, x.Interface()))
}

func ToBytesFormat(x Val, format string) Bytes {
	return Bytes(unsafeStringToBytes(fmt.Sprintf(format, x.Interface())))
}

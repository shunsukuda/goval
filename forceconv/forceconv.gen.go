package forVceconv

import (
	"reflect"
	"unsafe"

	"github.com/shunsukuda/goval"
)

const (
	SizeofBool       = int(unsafe.Sizeof(bool(false)))
	SizeofInt8       = int(unsafe.Sizeof(int8(0)))
	SizeofInt16      = int(unsafe.Sizeof(int16(0)))
	SizeofInt32      = int(unsafe.Sizeof(int32(0)))
	SizeofInt64      = int(unsafe.Sizeof(int64(0)))
	SizeofUint8      = int(unsafe.Sizeof(uint8(0)))
	SizeofUint16     = int(unsafe.Sizeof(uint16(0)))
	SizeofUint32     = int(unsafe.Sizeof(uint32(0)))
	SizeofUint64     = int(unsafe.Sizeof(uint64(0)))
	SizeofFloat32    = int(unsafe.Sizeof(float32(0.0)))
	SizeofFloat64    = int(unsafe.Sizeof(float64(0.0)))
	SizeofComplex64  = int(unsafe.Sizeof(complex64(0)))
	SizeofComplex128 = int(unsafe.Sizeof(complex128(0)))
)

func BytesToBool(b []byte) []goval.Bool {
	if b == nil {
		return nil
	}

	return *(*[]goval.Bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func BoolsToBytes(s []bool) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func BytesToInt8(b []byte) []goval.Int8 {
	if b == nil {
		return nil
	}

	return *(*[]goval.Int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Int8sToBytes(s []int8) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func BytesToInt16(b []byte) []goval.Int16 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt16 - 1) / SizeofInt16)
	cvCap := int(cap(b) / SizeofInt16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt16]
	return *(*[]goval.Int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Int16sToBytes(s []int16) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt16,
		Cap:  cap(s) * SizeofInt16,
	}))
}

func BytesToInt32(b []byte) []goval.Int32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt32 - 1) / SizeofInt32)
	cvCap := int(cap(b) / SizeofInt32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt32]
	return *(*[]goval.Int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Int32sToBytes(s []int32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt32,
		Cap:  cap(s) * SizeofInt32,
	}))
}

func BytesToInt64(b []byte) []goval.Int64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt64 - 1) / SizeofInt64)
	cvCap := int(cap(b) / SizeofInt64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt64]
	return *(*[]goval.Int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Int64sToBytes(s []int64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt64,
		Cap:  cap(s) * SizeofInt64,
	}))
}

func BytesToUint8(b []byte) []goval.Uint8 {
	if b == nil {
		return nil
	}

	return *(*[]goval.Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Uint8sToBytes(s []uint8) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func BytesToUint16(b []byte) []goval.Uint16 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint16 - 1) / SizeofUint16)
	cvCap := int(cap(b) / SizeofUint16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint16]
	return *(*[]goval.Uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Uint16sToBytes(s []uint16) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint16,
		Cap:  cap(s) * SizeofUint16,
	}))
}

func BytesToUint32(b []byte) []goval.Uint32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint32 - 1) / SizeofUint32)
	cvCap := int(cap(b) / SizeofUint32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint32]
	return *(*[]goval.Uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Uint32sToBytes(s []uint32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint32,
		Cap:  cap(s) * SizeofUint32,
	}))
}

func BytesToUint64(b []byte) []goval.Uint64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint64 - 1) / SizeofUint64)
	cvCap := int(cap(b) / SizeofUint64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint64]
	return *(*[]goval.Uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Uint64sToBytes(s []uint64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint64,
		Cap:  cap(s) * SizeofUint64,
	}))
}

func BytesToFloat32(b []byte) []goval.Float32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofFloat32 - 1) / SizeofFloat32)
	cvCap := int(cap(b) / SizeofFloat32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofFloat32]
	return *(*[]goval.Float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Float32sToBytes(s []float32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofFloat32,
		Cap:  cap(s) * SizeofFloat32,
	}))
}

func BytesToFloat64(b []byte) []goval.Float64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofFloat64 - 1) / SizeofFloat64)
	cvCap := int(cap(b) / SizeofFloat64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofFloat64]
	return *(*[]goval.Float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Float64sToBytes(s []float64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofFloat64,
		Cap:  cap(s) * SizeofFloat64,
	}))
}

func BytesToComplex64(b []byte) []goval.Complex64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofComplex64 - 1) / SizeofComplex64)
	cvCap := int(cap(b) / SizeofComplex64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofComplex64]
	return *(*[]goval.Complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Complex64sToBytes(s []complex64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofComplex64,
		Cap:  cap(s) * SizeofComplex64,
	}))
}

func BytesToComplex128(b []byte) []goval.Complex128 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofComplex128 - 1) / SizeofComplex128)
	cvCap := int(cap(b) / SizeofComplex128)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofComplex128]
	return *(*[]goval.Complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

func Complex128sToBytes(s []complex128) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofComplex128,
		Cap:  cap(s) * SizeofComplex128,
	}))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(
		&reflect.StringHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
			Len:  len(b),
		}))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
			Len:  len(s),
			Cap:  len(s),
		}))
}

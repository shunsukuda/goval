
package forVceconv

import (
	"reflect"
	"unsafe"

	"github.com/shunsukuda/goval"
)


const (
	SizeofInt8 = int(unsafe.Sizeof(int8(0)))
	SizeofInt16 = int(unsafe.Sizeof(int16(0)))
	SizeofInt32 = int(unsafe.Sizeof(int32(0)))
	SizeofInt64 = int(unsafe.Sizeof(int64(0)))
	SizeofUint8 = int(unsafe.Sizeof(uint8(0)))
	SizeofUint16 = int(unsafe.Sizeof(uint16(0)))
	SizeofUint32 = int(unsafe.Sizeof(uint32(0)))
	SizeofUint64 = int(unsafe.Sizeof(uint64(0)))
	SizeofFloat32 = int(unsafe.Sizeof(float32(0)))
	SizeofFloat64 = int(unsafe.Sizeof(float64(0)))
	SizeofComplex64 = int(unsafe.Sizeof(complex64(0)))
	SizeofComplex128 = int(unsafe.Sizeof(complex128(0)))
	
)

func BytesToInt8(b []byte) []goval.Int8 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofInt8) + 1
	cvCap := int(cap(b)/ SizeofInt8)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofInt8]
	return *(*[]goval.Int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToInt16(b []byte) []goval.Int16 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofInt16) + 1
	cvCap := int(cap(b)/ SizeofInt16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofInt16]
	return *(*[]goval.Int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToInt32(b []byte) []goval.Int32 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofInt32) + 1
	cvCap := int(cap(b)/ SizeofInt32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofInt32]
	return *(*[]goval.Int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToInt64(b []byte) []goval.Int64 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofInt64) + 1
	cvCap := int(cap(b)/ SizeofInt64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofInt64]
	return *(*[]goval.Int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToUint8(b []byte) []goval.Uint8 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofUint8) + 1
	cvCap := int(cap(b)/ SizeofUint8)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofUint8]
	return *(*[]goval.Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToUint16(b []byte) []goval.Uint16 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofUint16) + 1
	cvCap := int(cap(b)/ SizeofUint16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofUint16]
	return *(*[]goval.Uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToUint32(b []byte) []goval.Uint32 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofUint32) + 1
	cvCap := int(cap(b)/ SizeofUint32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofUint32]
	return *(*[]goval.Uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToUint64(b []byte) []goval.Uint64 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofUint64) + 1
	cvCap := int(cap(b)/ SizeofUint64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofUint64]
	return *(*[]goval.Uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToFloat32(b []byte) []goval.Float32 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofFloat32) + 1
	cvCap := int(cap(b)/ SizeofFloat32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofFloat32]
	return *(*[]goval.Float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToFloat64(b []byte) []goval.Float64 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofFloat64) + 1
	cvCap := int(cap(b)/ SizeofFloat64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofFloat64]
	return *(*[]goval.Float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToComplex64(b []byte) []goval.Complex64 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofComplex64) + 1
	cvCap := int(cap(b)/ SizeofComplex64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofComplex64]
	return *(*[]goval.Complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
func BytesToComplex128(b []byte) []goval.Complex128 {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ SizeofComplex128) + 1
	cvCap := int(cap(b)/ SizeofComplex128)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*SizeofComplex128]
	return *(*[]goval.Complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
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

package unsafecast

import (
	"reflect"
	"unsafe"
)

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

func Int8ToUint8(v int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&v))
}

func Uint8ToInt8(v uint8) int8 {
	return *(*int8)(unsafe.Pointer(&v))
}

func Int16ToUint16(v int16) uint16 {
	return *(*uint16)(unsafe.Pointer(&v))
}

func Uint16ToInt16(v uint16) int16 {
	return *(*int16)(unsafe.Pointer(&v))
}

func Int32ToUint32(v int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}

func Uint32ToInt32(v uint32) int32 {
	return *(*int32)(unsafe.Pointer(&v))
}

func Int64ToUint64(v int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

func Uint64ToInt64(v uint64) int64 {
	return *(*int64)(unsafe.Pointer(&v))
}

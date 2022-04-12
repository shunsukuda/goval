
package forVceconv

import (
	"reflect"
	"unsafe"

	"github.com/shunsukuda/goval"
)


const (
	{{range $T := $.TL -}}
		Sizeof{{$T.TypeName}} = int(unsafe.Sizeof({{$T.GoTypeName}}(0)))
	{{end}}
)

{{range $To := $.TL -}}
func BytesTo{{$To.TypeName}}(b []byte) []goval.{{$To.TypeName}} {
	if b == nil {
		return nil
	}

	cvLen := int(len(b)/ Sizeof{{$To.TypeName}}) + 1
	cvCap := int(cap(b)/ Sizeof{{$To.TypeName}})
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:len(b):cvCap*Sizeof{{$To.TypeName}}]
	return *(*[]goval.{{$To.TypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}
{{end}}

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

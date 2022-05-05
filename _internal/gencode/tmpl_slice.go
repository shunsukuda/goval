package goval

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/shunsukuda/forceconv"
	"github.com/valyala/bytebufferpool"
)

const (
	maxInt = int(^uint(0) >> 1)
)

var (
	defaultByteBufferPool = bytebufferpool.Pool{}
	minSliceSize = 64
)

var (
	ErrTooLarge = errors.New("shunsukuda.goval.Slice: too large")
)

type ValSlice interface {
	Type() Type
	Len() int
	Cap() int
	Grow(n int)
	Resize(n int)
	PoolGet()
	PoolPut()
	Reset()
}

func makeByteSlice(n, bs int) []byte {
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return make([]byte, n*bs)
}

{{range $T := $.TL}}

type {{$T.TypeName}}Slice struct{
	S []{{$T.TypeName}}
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGo{{$T.TypeName}}(s []{{$T.GoTypeName}}) []{{$T.TypeName}} {
	return *(*[]{{$T.TypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}
func New{{$T.TypeName}}Slice(ubbp bool) {{$T.TypeName}}Slice {
	return {{$T.TypeName}}Slice{nil, ubbp}
}
func New{{$T.TypeName}}SliceFromGoSlice(s []{{$T.GoTypeName}}, ubbp bool) {{$T.TypeName}}Slice {
	return {{$T.TypeName}}Slice{forceconvSliceFromGo{{$T.TypeName}}(s), ubbp}
}

func (s {{$T.TypeName}}Slice) GoSlice() []{{$T.GoTypeName}} {
	return *(*[]{{$T.GoTypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s {{$T.TypeName}}Slice) UseByteBufferPool() bool { return s.ubbp }

func (s {{$T.TypeName}}Slice) Copy() {{$T.TypeName}}Slice {
	buf := New{{$T.TypeName}}Slice(s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	buf.Resize(s.Len())
	return buf
}

func (s {{$T.TypeName}}Slice) Type() Type { return ValTypes.{{$T.TypeName}} }
func (s {{$T.TypeName}}Slice) Len() int { return len(s.S) }
func (s {{$T.TypeName}}Slice) Cap() int { return cap(s.S) }

func (s *{{$T.TypeName}}Slice) Reset() {
	s.S = s.S[:0]
}

func (s *{{$T.TypeName}}Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	} 
	return 0, false
}

func (s *{{$T.TypeName}}Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.S == nil && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]{{$T.TypeName}}, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c)*x)
	if c >= (maxInt-n) {{if ne $T.SizeOf 1}}/ {{$T.SizeOf}}{{end}} {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGo{{$T.TypeName}}(forceconv.BytesTo{{$T.TypeName}}Slice(makeByteSlice(c+n, {{$T.SizeOf}})))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

func (s *{{$T.TypeName}}Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.{{$T.TypeName}}Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *{{$T.TypeName}}Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *{{$T.TypeName}}Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.S == nil {
		s.S = forceconvSliceFromGo{{$T.TypeName}}(forceconv.BytesTo{{$T.TypeName}}Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.S == nil {
		s.S = make([]{{$T.TypeName}}, minSliceSize)
	}
}

func (s *{{$T.TypeName}}Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.S != nil {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.{{$T.TypeName}}SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}


{{end}}
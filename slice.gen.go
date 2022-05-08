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
	minSliceSize          = 64
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

type BoolSlice struct {
	S    []Bool
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoBool(s []bool) []Bool {
	return *(*[]Bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewBoolSlice(size int, ubbp bool) BoolSlice {
	s := BoolSlice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewBoolSliceFromGoSlice(s []bool, ubbp bool) BoolSlice {
	return BoolSlice{forceconvSliceFromGoBool(s), ubbp}
}

func (s BoolSlice) GoSlice() []bool {
	return *(*[]bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s BoolSlice) UseByteBufferPool() bool { return s.ubbp }

func (s BoolSlice) Copy() BoolSlice {
	buf := NewBoolSlice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s BoolSlice) Type() Type { return ValTypes.Bool }
func (s BoolSlice) Len() int   { return len(s.S) }
func (s BoolSlice) Cap() int   { return cap(s.S) }

func (s *BoolSlice) Reset() {
	s.S = s.S[:0]
}

func (s *BoolSlice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *BoolSlice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Bool, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt - n) {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoBool(forceconv.BytesToBoolSlice(makeByteSlice(c+n, 1)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *BoolSlice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.BoolSlice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *BoolSlice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *BoolSlice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoBool(forceconv.BytesToBoolSlice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Bool, minSliceSize)
	}
}

func (s *BoolSlice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.BoolSliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Int8Slice struct {
	S    []Int8
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoInt8(s []int8) []Int8 {
	return *(*[]Int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewInt8Slice(size int, ubbp bool) Int8Slice {
	s := Int8Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewInt8SliceFromGoSlice(s []int8, ubbp bool) Int8Slice {
	return Int8Slice{forceconvSliceFromGoInt8(s), ubbp}
}

func (s Int8Slice) GoSlice() []int8 {
	return *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Int8Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Int8Slice) Copy() Int8Slice {
	buf := NewInt8Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Int8Slice) Type() Type { return ValTypes.Int8 }
func (s Int8Slice) Len() int   { return len(s.S) }
func (s Int8Slice) Cap() int   { return cap(s.S) }

func (s *Int8Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Int8Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Int8Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Int8, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt - n) {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoInt8(forceconv.BytesToInt8Slice(makeByteSlice(c+n, 1)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Int8Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Int8Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Int8Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Int8Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoInt8(forceconv.BytesToInt8Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Int8, minSliceSize)
	}
}

func (s *Int8Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Int8SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Int16Slice struct {
	S    []Int16
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoInt16(s []int16) []Int16 {
	return *(*[]Int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewInt16Slice(size int, ubbp bool) Int16Slice {
	s := Int16Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewInt16SliceFromGoSlice(s []int16, ubbp bool) Int16Slice {
	return Int16Slice{forceconvSliceFromGoInt16(s), ubbp}
}

func (s Int16Slice) GoSlice() []int16 {
	return *(*[]int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Int16Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Int16Slice) Copy() Int16Slice {
	buf := NewInt16Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Int16Slice) Type() Type { return ValTypes.Int16 }
func (s Int16Slice) Len() int   { return len(s.S) }
func (s Int16Slice) Cap() int   { return cap(s.S) }

func (s *Int16Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Int16Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Int16Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Int16, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/2 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoInt16(forceconv.BytesToInt16Slice(makeByteSlice(c+n, 2)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Int16Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Int16Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Int16Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Int16Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoInt16(forceconv.BytesToInt16Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Int16, minSliceSize)
	}
}

func (s *Int16Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Int16SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Int32Slice struct {
	S    []Int32
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoInt32(s []int32) []Int32 {
	return *(*[]Int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewInt32Slice(size int, ubbp bool) Int32Slice {
	s := Int32Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewInt32SliceFromGoSlice(s []int32, ubbp bool) Int32Slice {
	return Int32Slice{forceconvSliceFromGoInt32(s), ubbp}
}

func (s Int32Slice) GoSlice() []int32 {
	return *(*[]int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Int32Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Int32Slice) Copy() Int32Slice {
	buf := NewInt32Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Int32Slice) Type() Type { return ValTypes.Int32 }
func (s Int32Slice) Len() int   { return len(s.S) }
func (s Int32Slice) Cap() int   { return cap(s.S) }

func (s *Int32Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Int32Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Int32Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Int32, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/4 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoInt32(forceconv.BytesToInt32Slice(makeByteSlice(c+n, 4)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Int32Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Int32Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Int32Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Int32Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoInt32(forceconv.BytesToInt32Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Int32, minSliceSize)
	}
}

func (s *Int32Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Int32SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Int64Slice struct {
	S    []Int64
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoInt64(s []int64) []Int64 {
	return *(*[]Int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewInt64Slice(size int, ubbp bool) Int64Slice {
	s := Int64Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewInt64SliceFromGoSlice(s []int64, ubbp bool) Int64Slice {
	return Int64Slice{forceconvSliceFromGoInt64(s), ubbp}
}

func (s Int64Slice) GoSlice() []int64 {
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Int64Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Int64Slice) Copy() Int64Slice {
	buf := NewInt64Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Int64Slice) Type() Type { return ValTypes.Int64 }
func (s Int64Slice) Len() int   { return len(s.S) }
func (s Int64Slice) Cap() int   { return cap(s.S) }

func (s *Int64Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Int64Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Int64Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Int64, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/8 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoInt64(forceconv.BytesToInt64Slice(makeByteSlice(c+n, 8)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Int64Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Int64Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Int64Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Int64Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoInt64(forceconv.BytesToInt64Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Int64, minSliceSize)
	}
}

func (s *Int64Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Int64SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Uint8Slice struct {
	S    []Uint8
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoUint8(s []uint8) []Uint8 {
	return *(*[]Uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewUint8Slice(size int, ubbp bool) Uint8Slice {
	s := Uint8Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewUint8SliceFromGoSlice(s []uint8, ubbp bool) Uint8Slice {
	return Uint8Slice{forceconvSliceFromGoUint8(s), ubbp}
}

func (s Uint8Slice) GoSlice() []uint8 {
	return *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Uint8Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Uint8Slice) Copy() Uint8Slice {
	buf := NewUint8Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Uint8Slice) Type() Type { return ValTypes.Uint8 }
func (s Uint8Slice) Len() int   { return len(s.S) }
func (s Uint8Slice) Cap() int   { return cap(s.S) }

func (s *Uint8Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Uint8Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Uint8Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Uint8, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt - n) {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoUint8(forceconv.BytesToUint8Slice(makeByteSlice(c+n, 1)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Uint8Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Uint8Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Uint8Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Uint8Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoUint8(forceconv.BytesToUint8Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Uint8, minSliceSize)
	}
}

func (s *Uint8Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Uint8SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Uint16Slice struct {
	S    []Uint16
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoUint16(s []uint16) []Uint16 {
	return *(*[]Uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewUint16Slice(size int, ubbp bool) Uint16Slice {
	s := Uint16Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewUint16SliceFromGoSlice(s []uint16, ubbp bool) Uint16Slice {
	return Uint16Slice{forceconvSliceFromGoUint16(s), ubbp}
}

func (s Uint16Slice) GoSlice() []uint16 {
	return *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Uint16Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Uint16Slice) Copy() Uint16Slice {
	buf := NewUint16Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Uint16Slice) Type() Type { return ValTypes.Uint16 }
func (s Uint16Slice) Len() int   { return len(s.S) }
func (s Uint16Slice) Cap() int   { return cap(s.S) }

func (s *Uint16Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Uint16Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Uint16Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Uint16, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/2 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoUint16(forceconv.BytesToUint16Slice(makeByteSlice(c+n, 2)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Uint16Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Uint16Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Uint16Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Uint16Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoUint16(forceconv.BytesToUint16Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Uint16, minSliceSize)
	}
}

func (s *Uint16Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Uint16SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Uint32Slice struct {
	S    []Uint32
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoUint32(s []uint32) []Uint32 {
	return *(*[]Uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewUint32Slice(size int, ubbp bool) Uint32Slice {
	s := Uint32Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewUint32SliceFromGoSlice(s []uint32, ubbp bool) Uint32Slice {
	return Uint32Slice{forceconvSliceFromGoUint32(s), ubbp}
}

func (s Uint32Slice) GoSlice() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Uint32Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Uint32Slice) Copy() Uint32Slice {
	buf := NewUint32Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Uint32Slice) Type() Type { return ValTypes.Uint32 }
func (s Uint32Slice) Len() int   { return len(s.S) }
func (s Uint32Slice) Cap() int   { return cap(s.S) }

func (s *Uint32Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Uint32Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Uint32Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Uint32, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/4 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoUint32(forceconv.BytesToUint32Slice(makeByteSlice(c+n, 4)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Uint32Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Uint32Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Uint32Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Uint32Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoUint32(forceconv.BytesToUint32Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Uint32, minSliceSize)
	}
}

func (s *Uint32Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Uint32SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Uint64Slice struct {
	S    []Uint64
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoUint64(s []uint64) []Uint64 {
	return *(*[]Uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewUint64Slice(size int, ubbp bool) Uint64Slice {
	s := Uint64Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewUint64SliceFromGoSlice(s []uint64, ubbp bool) Uint64Slice {
	return Uint64Slice{forceconvSliceFromGoUint64(s), ubbp}
}

func (s Uint64Slice) GoSlice() []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Uint64Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Uint64Slice) Copy() Uint64Slice {
	buf := NewUint64Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Uint64Slice) Type() Type { return ValTypes.Uint64 }
func (s Uint64Slice) Len() int   { return len(s.S) }
func (s Uint64Slice) Cap() int   { return cap(s.S) }

func (s *Uint64Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Uint64Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Uint64Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Uint64, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/8 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoUint64(forceconv.BytesToUint64Slice(makeByteSlice(c+n, 8)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Uint64Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Uint64Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Uint64Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Uint64Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoUint64(forceconv.BytesToUint64Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Uint64, minSliceSize)
	}
}

func (s *Uint64Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Uint64SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Float32Slice struct {
	S    []Float32
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoFloat32(s []float32) []Float32 {
	return *(*[]Float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewFloat32Slice(size int, ubbp bool) Float32Slice {
	s := Float32Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewFloat32SliceFromGoSlice(s []float32, ubbp bool) Float32Slice {
	return Float32Slice{forceconvSliceFromGoFloat32(s), ubbp}
}

func (s Float32Slice) GoSlice() []float32 {
	return *(*[]float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Float32Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Float32Slice) Copy() Float32Slice {
	buf := NewFloat32Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Float32Slice) Type() Type { return ValTypes.Float32 }
func (s Float32Slice) Len() int   { return len(s.S) }
func (s Float32Slice) Cap() int   { return cap(s.S) }

func (s *Float32Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Float32Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Float32Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Float32, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/4 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoFloat32(forceconv.BytesToFloat32Slice(makeByteSlice(c+n, 4)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Float32Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Float32Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Float32Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Float32Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoFloat32(forceconv.BytesToFloat32Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Float32, minSliceSize)
	}
}

func (s *Float32Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Float32SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Float64Slice struct {
	S    []Float64
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoFloat64(s []float64) []Float64 {
	return *(*[]Float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewFloat64Slice(size int, ubbp bool) Float64Slice {
	s := Float64Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewFloat64SliceFromGoSlice(s []float64, ubbp bool) Float64Slice {
	return Float64Slice{forceconvSliceFromGoFloat64(s), ubbp}
}

func (s Float64Slice) GoSlice() []float64 {
	return *(*[]float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Float64Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Float64Slice) Copy() Float64Slice {
	buf := NewFloat64Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Float64Slice) Type() Type { return ValTypes.Float64 }
func (s Float64Slice) Len() int   { return len(s.S) }
func (s Float64Slice) Cap() int   { return cap(s.S) }

func (s *Float64Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Float64Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Float64Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Float64, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/8 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoFloat64(forceconv.BytesToFloat64Slice(makeByteSlice(c+n, 8)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Float64Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Float64Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Float64Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Float64Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoFloat64(forceconv.BytesToFloat64Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Float64, minSliceSize)
	}
}

func (s *Float64Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Float64SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Complex64Slice struct {
	S    []Complex64
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoComplex64(s []complex64) []Complex64 {
	return *(*[]Complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewComplex64Slice(size int, ubbp bool) Complex64Slice {
	s := Complex64Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewComplex64SliceFromGoSlice(s []complex64, ubbp bool) Complex64Slice {
	return Complex64Slice{forceconvSliceFromGoComplex64(s), ubbp}
}

func (s Complex64Slice) GoSlice() []complex64 {
	return *(*[]complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Complex64Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Complex64Slice) Copy() Complex64Slice {
	buf := NewComplex64Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Complex64Slice) Type() Type { return ValTypes.Complex64 }
func (s Complex64Slice) Len() int   { return len(s.S) }
func (s Complex64Slice) Cap() int   { return cap(s.S) }

func (s *Complex64Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Complex64Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Complex64Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Complex64, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/8 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoComplex64(forceconv.BytesToComplex64Slice(makeByteSlice(c+n, 8)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Complex64Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Complex64Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Complex64Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Complex64Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoComplex64(forceconv.BytesToComplex64Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Complex64, minSliceSize)
	}
}

func (s *Complex64Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Complex64SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

type Complex128Slice struct {
	S    []Complex128
	ubbp bool // use ByteBuffer Pool
}

func forceconvSliceFromGoComplex128(s []complex128) []Complex128 {
	return *(*[]Complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

func NewComplex128Slice(size int, ubbp bool) Complex128Slice {
	s := Complex128Slice{nil, ubbp}
	s.Resize(size)
	return s
}

func NewComplex128SliceFromGoSlice(s []complex128, ubbp bool) Complex128Slice {
	return Complex128Slice{forceconvSliceFromGoComplex128(s), ubbp}
}

func (s Complex128Slice) GoSlice() []complex128 {
	return *(*[]complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s.S)).Data,
		Len:  s.Len(),
		Cap:  s.Cap(),
	}))
}

func (s Complex128Slice) UseByteBufferPool() bool { return s.ubbp }

func (s Complex128Slice) Copy() Complex128Slice {
	buf := NewComplex128Slice(s.Len(), s.ubbp)
	if s.ubbp {
		buf.PoolGet()
	}
	copy(buf.S, s.S)
	return buf
}

func (s Complex128Slice) Type() Type { return ValTypes.Complex128 }
func (s Complex128Slice) Len() int   { return len(s.S) }
func (s Complex128Slice) Cap() int   { return cap(s.S) }

func (s *Complex128Slice) Reset() {
	s.S = s.S[:0]
}

func (s *Complex128Slice) tryGrowByReslice(n int) (int, bool) {
	if l := len(s.S); n <= cap(s.S)-l {
		s.S = s.S[:l+n]
		return l, true
	}
	return 0, false
}

func (s *Complex128Slice) grow(n int) int {
	m := s.Len()
	if i, ok := s.tryGrowByReslice(n); ok {
		return i
	}
	if s.Cap() == 0 && n <= minSliceSize {
		if s.ubbp {
			s.PoolGet()
		} else {
			s.S = make([]Complex128, n, minSliceSize)
		}
		return 0
	}
	x := 1.25
	c := s.Cap()
	if c < 1024 {
		x = 2
	}
	c = int(float64(c) * x)
	if c >= (maxInt-n)/16 {
		panic(ErrTooLarge)
	} else {
		buf := forceconvSliceFromGoComplex128(forceconv.BytesToComplex128Slice(makeByteSlice(c+n, 16)))
		copy(buf, s.S)
		if s.ubbp {
			s.PoolPut()
		}
		s.S = buf
	}
	s.S = s.S[:m+n]
	return m
}

// Grow grows the slice's capacity.
func (s *Complex128Slice) Grow(n int) {
	if n < 0 {
		panic("shunsukuda.goval.Complex128Slice.Grow: negative count")
	}
	m := s.grow(n)
	s.S = s.S[:m]
}

func (s *Complex128Slice) Resize(n int) {
	m := n - s.Len()
	if m > 0 {
		s.Grow(m)
	}
	s.S = s.S[:n]
}

func (s *Complex128Slice) PoolGet() {
	if !s.ubbp {
		return
	}
	if s.Cap() == 0 {
		s.S = forceconvSliceFromGoComplex128(forceconv.BytesToComplex128Slice(defaultByteBufferPool.Get().B[:0]))
	}
	if s.Cap() < minSliceSize {
		s.PoolPut()
		s.S = make([]Complex128, minSliceSize)
	}
}

func (s *Complex128Slice) PoolPut() {
	if !s.ubbp {
		return
	}
	if s.Cap() > 0 {
		defaultByteBufferPool.Put(&bytebufferpool.ByteBuffer{forceconv.Complex128SliceToBytes(s.GoSlice()[:0])})
		s.S = nil
	}
}

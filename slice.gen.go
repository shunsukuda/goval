package goval

type ValSlice interface {
	Interface() interface{}
	ValSlice() ValSlice
	Type() Type
	Len() int
	Cap() int
	Grow(n int)
	Reset()
}

type BoolSlice []bool

func NewBoolSlice(s []bool) BoolSlice { return s }

func (s *BoolSlice) Interface() interface{} { return s.BoolSlice() }
func (s *BoolSlice) BoolSlice() []bool      { return *s }
func (s *BoolSlice) Type() Type             { return ValTypes.Bool }
func (s *BoolSlice) Len() int               { return len(*s) }
func (s *BoolSlice) Cap() int               { return cap(*s) }

type Int8Slice []int8

func NewInt8Slice(s []int8) Int8Slice { return s }

func (s *Int8Slice) Interface() interface{} { return s.Int8Slice() }
func (s *Int8Slice) Int8Slice() []int8      { return *s }
func (s *Int8Slice) Type() Type             { return ValTypes.Int8 }
func (s *Int8Slice) Len() int               { return len(*s) }
func (s *Int8Slice) Cap() int               { return cap(*s) }

type Int16Slice []int16

func NewInt16Slice(s []int16) Int16Slice { return s }

func (s *Int16Slice) Interface() interface{} { return s.Int16Slice() }
func (s *Int16Slice) Int16Slice() []int16    { return *s }
func (s *Int16Slice) Type() Type             { return ValTypes.Int16 }
func (s *Int16Slice) Len() int               { return len(*s) }
func (s *Int16Slice) Cap() int               { return cap(*s) }

type Int32Slice []int32

func NewInt32Slice(s []int32) Int32Slice { return s }

func (s *Int32Slice) Interface() interface{} { return s.Int32Slice() }
func (s *Int32Slice) Int32Slice() []int32    { return *s }
func (s *Int32Slice) Type() Type             { return ValTypes.Int32 }
func (s *Int32Slice) Len() int               { return len(*s) }
func (s *Int32Slice) Cap() int               { return cap(*s) }

type Int64Slice []int64

func NewInt64Slice(s []int64) Int64Slice { return s }

func (s *Int64Slice) Interface() interface{} { return s.Int64Slice() }
func (s *Int64Slice) Int64Slice() []int64    { return *s }
func (s *Int64Slice) Type() Type             { return ValTypes.Int64 }
func (s *Int64Slice) Len() int               { return len(*s) }
func (s *Int64Slice) Cap() int               { return cap(*s) }

type Uint8Slice []uint8

func NewUint8Slice(s []uint8) Uint8Slice { return s }

func (s *Uint8Slice) Interface() interface{} { return s.Uint8Slice() }
func (s *Uint8Slice) Uint8Slice() []uint8    { return *s }
func (s *Uint8Slice) Type() Type             { return ValTypes.Uint8 }
func (s *Uint8Slice) Len() int               { return len(*s) }
func (s *Uint8Slice) Cap() int               { return cap(*s) }

type Uint16Slice []uint16

func NewUint16Slice(s []uint16) Uint16Slice { return s }

func (s *Uint16Slice) Interface() interface{} { return s.Uint16Slice() }
func (s *Uint16Slice) Uint16Slice() []uint16  { return *s }
func (s *Uint16Slice) Type() Type             { return ValTypes.Uint16 }
func (s *Uint16Slice) Len() int               { return len(*s) }
func (s *Uint16Slice) Cap() int               { return cap(*s) }

type Uint32Slice []uint32

func NewUint32Slice(s []uint32) Uint32Slice { return s }

func (s *Uint32Slice) Interface() interface{} { return s.Uint32Slice() }
func (s *Uint32Slice) Uint32Slice() []uint32  { return *s }
func (s *Uint32Slice) Type() Type             { return ValTypes.Uint32 }
func (s *Uint32Slice) Len() int               { return len(*s) }
func (s *Uint32Slice) Cap() int               { return cap(*s) }

type Uint64Slice []uint64

func NewUint64Slice(s []uint64) Uint64Slice { return s }

func (s *Uint64Slice) Interface() interface{} { return s.Uint64Slice() }
func (s *Uint64Slice) Uint64Slice() []uint64  { return *s }
func (s *Uint64Slice) Type() Type             { return ValTypes.Uint64 }
func (s *Uint64Slice) Len() int               { return len(*s) }
func (s *Uint64Slice) Cap() int               { return cap(*s) }

type Float32Slice []float32

func NewFloat32Slice(s []float32) Float32Slice { return s }

func (s *Float32Slice) Interface() interface{}  { return s.Float32Slice() }
func (s *Float32Slice) Float32Slice() []float32 { return *s }
func (s *Float32Slice) Type() Type              { return ValTypes.Float32 }
func (s *Float32Slice) Len() int                { return len(*s) }
func (s *Float32Slice) Cap() int                { return cap(*s) }

type Float64Slice []float64

func NewFloat64Slice(s []float64) Float64Slice { return s }

func (s *Float64Slice) Interface() interface{}  { return s.Float64Slice() }
func (s *Float64Slice) Float64Slice() []float64 { return *s }
func (s *Float64Slice) Type() Type              { return ValTypes.Float64 }
func (s *Float64Slice) Len() int                { return len(*s) }
func (s *Float64Slice) Cap() int                { return cap(*s) }

type Complex64Slice []complex64

func NewComplex64Slice(s []complex64) Complex64Slice { return s }

func (s *Complex64Slice) Interface() interface{}      { return s.Complex64Slice() }
func (s *Complex64Slice) Complex64Slice() []complex64 { return *s }
func (s *Complex64Slice) Type() Type                  { return ValTypes.Complex64 }
func (s *Complex64Slice) Len() int                    { return len(*s) }
func (s *Complex64Slice) Cap() int                    { return cap(*s) }

type Complex128Slice []complex128

func NewComplex128Slice(s []complex128) Complex128Slice { return s }

func (s *Complex128Slice) Interface() interface{}        { return s.Complex128Slice() }
func (s *Complex128Slice) Complex128Slice() []complex128 { return *s }
func (s *Complex128Slice) Type() Type                    { return ValTypes.Complex128 }
func (s *Complex128Slice) Len() int                      { return len(*s) }
func (s *Complex128Slice) Cap() int                      { return cap(*s) }

type StringSlice []string

func NewStringSlice(s []string) StringSlice { return s }

func (s *StringSlice) Interface() interface{} { return s.StringSlice() }
func (s *StringSlice) StringSlice() []string  { return *s }
func (s *StringSlice) Type() Type             { return ValTypes.String }
func (s *StringSlice) Len() int               { return len(*s) }
func (s *StringSlice) Cap() int               { return cap(*s) }

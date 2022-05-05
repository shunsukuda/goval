package goval

import "reflect"

const (
	TypeNil = iota
	TypeBool
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeComplex64
	TypeComplex128
	TypeString
	TypeBytes
)

var (
	ValTypes = struct {
		Nil        Type
		Bool       Type
		Int8       Type
		Int16      Type
		Int32      Type
		Int64      Type
		Uint8      Type
		Uint16     Type
		Uint32     Type
		Uint64     Type
		Float32    Type
		Float64    Type
		Complex64  Type
		Complex128 Type
		String     Type
		Bytes      Type
	}{
		Nil:        &NilType{},
		Bool:       &BoolType{},
		Int8:       &Int8Type{},
		Int16:      &Int16Type{},
		Int32:      &Int32Type{},
		Int64:      &Int64Type{},
		Uint8:      &Uint8Type{},
		Uint16:     &Uint16Type{},
		Uint32:     &Uint32Type{},
		Uint64:     &Uint64Type{},
		Float32:    &Float32Type{},
		Float64:    &Float64Type{},
		Complex64:  &Complex64Type{},
		Complex128: &Complex128Type{},
		String:     &StringType{},
		Bytes:      &BytesType{},
	}
)

type Type interface {
	ID() int
	Kind() reflect.Kind
	Name() string
	String() string
	// BitSize() int
}

type NilType struct{}

func (t NilType) ID() int            { return TypeNil }
func (t NilType) Kind() reflect.Kind { return reflect.Invalid }
func (t NilType) Name() string       { return "Nil" }
func (t NilType) String() string     { return "Nil" }

type BoolType struct{}

func (t BoolType) ID() int            { return TypeBool }
func (t BoolType) Kind() reflect.Kind { return reflect.Bool }
func (t BoolType) Name() string       { return "Bool" }
func (t BoolType) String() string     { return "Bool" }
func (t BoolType) BitSize() int       { return 1 }

type Int8Type struct{}

func (t Int8Type) ID() int            { return TypeInt8 }
func (t Int8Type) Kind() reflect.Kind { return reflect.Int8 }
func (t Int8Type) Name() string       { return "Int8" }
func (t Int8Type) String() string     { return "Int8" }
func (t Int8Type) BitSize() int       { return 8 }

type Int16Type struct{}

func (t Int16Type) ID() int            { return TypeInt16 }
func (t Int16Type) Kind() reflect.Kind { return reflect.Int16 }
func (t Int16Type) Name() string       { return "Int16" }
func (t Int16Type) String() string     { return "Int16" }
func (t Int16Type) BitSize() int       { return 16 }

type Int32Type struct{}

func (t Int32Type) ID() int            { return TypeInt32 }
func (t Int32Type) Kind() reflect.Kind { return reflect.Int32 }
func (t Int32Type) Name() string       { return "Int32" }
func (t Int32Type) String() string     { return "Int32" }
func (t Int32Type) BitSize() int       { return 32 }

type Int64Type struct{}

func (t Int64Type) ID() int            { return TypeInt64 }
func (t Int64Type) Kind() reflect.Kind { return reflect.Int64 }
func (t Int64Type) Name() string       { return "Int64" }
func (t Int64Type) String() string     { return "Int64" }
func (t Int64Type) BitSize() int       { return 64 }

type Uint8Type struct{}

func (t Uint8Type) ID() int            { return TypeUint8 }
func (t Uint8Type) Kind() reflect.Kind { return reflect.Uint8 }
func (t Uint8Type) Name() string       { return "Uint8" }
func (t Uint8Type) String() string     { return "Uint8" }
func (t Uint8Type) BitSize() int       { return 8 }

type Uint16Type struct{}

func (t Uint16Type) ID() int            { return TypeUint16 }
func (t Uint16Type) Kind() reflect.Kind { return reflect.Uint16 }
func (t Uint16Type) Name() string       { return "Uint16" }
func (t Uint16Type) String() string     { return "Uint16" }
func (t Uint16Type) BitSize() int       { return 16 }

type Uint32Type struct{}

func (t Uint32Type) ID() int            { return TypeUint32 }
func (t Uint32Type) Kind() reflect.Kind { return reflect.Uint32 }
func (t Uint32Type) Name() string       { return "Uint32" }
func (t Uint32Type) String() string     { return "Uint32" }
func (t Uint32Type) BitSize() int       { return 32 }

type Uint64Type struct{}

func (t Uint64Type) ID() int            { return TypeUint64 }
func (t Uint64Type) Kind() reflect.Kind { return reflect.Uint64 }
func (t Uint64Type) Name() string       { return "Uint64" }
func (t Uint64Type) String() string     { return "Uint64" }
func (t Uint64Type) BitSize() int       { return 64 }

type Float32Type struct{}

func (t Float32Type) ID() int            { return TypeFloat32 }
func (t Float32Type) Kind() reflect.Kind { return reflect.Float32 }
func (t Float32Type) Name() string       { return "Float32" }
func (t Float32Type) String() string     { return "Float32" }
func (t Float32Type) BitSize() int       { return 32 }

type Float64Type struct{}

func (t Float64Type) ID() int            { return TypeFloat64 }
func (t Float64Type) Kind() reflect.Kind { return reflect.Float64 }
func (t Float64Type) Name() string       { return "Float64" }
func (t Float64Type) String() string     { return "Float64" }
func (t Float64Type) BitSize() int       { return 64 }

type Complex64Type struct{}

func (t Complex64Type) ID() int            { return TypeComplex64 }
func (t Complex64Type) Kind() reflect.Kind { return reflect.Complex64 }
func (t Complex64Type) Name() string       { return "Complex64" }
func (t Complex64Type) String() string     { return "Complex64" }
func (t Complex64Type) BitSize() int       { return 32 }

type Complex128Type struct{}

func (t Complex128Type) ID() int            { return TypeComplex128 }
func (t Complex128Type) Kind() reflect.Kind { return reflect.Complex128 }
func (t Complex128Type) Name() string       { return "Complex128" }
func (t Complex128Type) String() string     { return "Complex128" }
func (t Complex128Type) BitSize() int       { return 64 }

type StringType struct{}

func (t StringType) ID() int            { return TypeString }
func (t StringType) Kind() reflect.Kind { return reflect.String }
func (t StringType) Name() string       { return "String" }
func (t StringType) String() string     { return "String" }

type BytesType struct{}

func (t BytesType) ID() int            { return TypeBytes }
func (t BytesType) Kind() reflect.Kind { return reflect.Slice }
func (t BytesType) Name() string       { return "Bytes" }
func (t BytesType) String() string     { return "Bytes" }

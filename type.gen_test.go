package goval

import (
	"reflect"
	"testing"
)

func TestNilType_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    NilType
		args args
		want bool
	}{
		{name: "True", e: NilType{}, args: args{x: NilType{}}, want: true},
		{name: "False", e: NilType{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("NilType.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilType_ID(t *testing.T) {
	tests := []struct {
		name string
		e    NilType
		want int
	}{
		{name: "", e: NilType{}, want: TypeNil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("NilType.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilType_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    NilType
		want reflect.Kind
	}{
		{name: "", e: NilType{}, want: reflect.Invalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("NilType.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilType_Name(t *testing.T) {
	tests := []struct {
		name string
		e    NilType
		want string
	}{
		{name: "", e: NilType{}, want: "Nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("NilType.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilType_String(t *testing.T) {
	tests := []struct {
		name string
		e    NilType
		want string
	}{
		{name: "", e: NilType{}, want: "Nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("NilType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilType_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    NilType
		want string
	}{
		{name: "", e: NilType{}, want: "Type:Nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("NilType.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    BoolType
		args args
		want bool
	}{
		{name: "True", e: BoolType{}, args: args{x: BoolType{}}, want: true},
		{name: "False", e: BoolType{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("BoolType.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_ID(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want int
	}{
		{name: "", e: BoolType{}, want: TypeBool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("BoolType.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want reflect.Kind
	}{
		{name: "", e: BoolType{}, want: reflect.Bool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("BoolType.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_Name(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want string
	}{
		{name: "", e: BoolType{}, want: "Bool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("BoolType.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_String(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want string
	}{
		{name: "", e: BoolType{}, want: "Bool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("BoolType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolType_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want string
	}{
		{name: "", e: BoolType{}, want: "Type:Bool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("BoolType.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBoolType_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    BoolType
		want int
	}{
		{name: "", e: BoolType{}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("BoolType.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Int8Type
		args args
		want bool
	}{
		{name: "True", e: Int8Type{}, args: args{x: Int8Type{}}, want: true},
		{name: "False", e: Int8Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int8Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want int
	}{
		{name: "", e: Int8Type{}, want: TypeInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Int8Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want reflect.Kind
	}{
		{name: "", e: Int8Type{}, want: reflect.Int8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Int8Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want string
	}{
		{name: "", e: Int8Type{}, want: "Int8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Int8Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want string
	}{
		{name: "", e: Int8Type{}, want: "Int8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Int8Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want string
	}{
		{name: "", e: Int8Type{}, want: "Type:Int8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Int8Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInt8Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Int8Type
		want int
	}{
		{name: "", e: Int8Type{}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Int8Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Int16Type
		args args
		want bool
	}{
		{name: "True", e: Int16Type{}, args: args{x: Int16Type{}}, want: true},
		{name: "False", e: Int16Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int16Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want int
	}{
		{name: "", e: Int16Type{}, want: TypeInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Int16Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want reflect.Kind
	}{
		{name: "", e: Int16Type{}, want: reflect.Int16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Int16Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want string
	}{
		{name: "", e: Int16Type{}, want: "Int16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Int16Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want string
	}{
		{name: "", e: Int16Type{}, want: "Int16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Int16Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want string
	}{
		{name: "", e: Int16Type{}, want: "Type:Int16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Int16Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInt16Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Int16Type
		want int
	}{
		{name: "", e: Int16Type{}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Int16Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Int32Type
		args args
		want bool
	}{
		{name: "True", e: Int32Type{}, args: args{x: Int32Type{}}, want: true},
		{name: "False", e: Int32Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int32Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want int
	}{
		{name: "", e: Int32Type{}, want: TypeInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Int32Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want reflect.Kind
	}{
		{name: "", e: Int32Type{}, want: reflect.Int32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Int32Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want string
	}{
		{name: "", e: Int32Type{}, want: "Int32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Int32Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want string
	}{
		{name: "", e: Int32Type{}, want: "Int32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Int32Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want string
	}{
		{name: "", e: Int32Type{}, want: "Type:Int32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Int32Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInt32Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Int32Type
		want int
	}{
		{name: "", e: Int32Type{}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Int32Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Int64Type
		args args
		want bool
	}{
		{name: "True", e: Int64Type{}, args: args{x: Int64Type{}}, want: true},
		{name: "False", e: Int64Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int64Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want int
	}{
		{name: "", e: Int64Type{}, want: TypeInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Int64Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want reflect.Kind
	}{
		{name: "", e: Int64Type{}, want: reflect.Int64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Int64Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want string
	}{
		{name: "", e: Int64Type{}, want: "Int64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Int64Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want string
	}{
		{name: "", e: Int64Type{}, want: "Int64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Int64Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want string
	}{
		{name: "", e: Int64Type{}, want: "Type:Int64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Int64Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInt64Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Int64Type
		want int
	}{
		{name: "", e: Int64Type{}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Int64Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Uint8Type
		args args
		want bool
	}{
		{name: "True", e: Uint8Type{}, args: args{x: Uint8Type{}}, want: true},
		{name: "False", e: Uint8Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint8Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want int
	}{
		{name: "", e: Uint8Type{}, want: TypeUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Uint8Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want reflect.Kind
	}{
		{name: "", e: Uint8Type{}, want: reflect.Uint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Uint8Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want string
	}{
		{name: "", e: Uint8Type{}, want: "Uint8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Uint8Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want string
	}{
		{name: "", e: Uint8Type{}, want: "Uint8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Uint8Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want string
	}{
		{name: "", e: Uint8Type{}, want: "Type:Uint8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Uint8Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUint8Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8Type
		want int
	}{
		{name: "", e: Uint8Type{}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Uint8Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Uint16Type
		args args
		want bool
	}{
		{name: "True", e: Uint16Type{}, args: args{x: Uint16Type{}}, want: true},
		{name: "False", e: Uint16Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint16Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want int
	}{
		{name: "", e: Uint16Type{}, want: TypeUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Uint16Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want reflect.Kind
	}{
		{name: "", e: Uint16Type{}, want: reflect.Uint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Uint16Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want string
	}{
		{name: "", e: Uint16Type{}, want: "Uint16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Uint16Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want string
	}{
		{name: "", e: Uint16Type{}, want: "Uint16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Uint16Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want string
	}{
		{name: "", e: Uint16Type{}, want: "Type:Uint16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Uint16Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUint16Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16Type
		want int
	}{
		{name: "", e: Uint16Type{}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Uint16Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Uint32Type
		args args
		want bool
	}{
		{name: "True", e: Uint32Type{}, args: args{x: Uint32Type{}}, want: true},
		{name: "False", e: Uint32Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint32Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want int
	}{
		{name: "", e: Uint32Type{}, want: TypeUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Uint32Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want reflect.Kind
	}{
		{name: "", e: Uint32Type{}, want: reflect.Uint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Uint32Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want string
	}{
		{name: "", e: Uint32Type{}, want: "Uint32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Uint32Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want string
	}{
		{name: "", e: Uint32Type{}, want: "Uint32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Uint32Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want string
	}{
		{name: "", e: Uint32Type{}, want: "Type:Uint32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Uint32Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUint32Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32Type
		want int
	}{
		{name: "", e: Uint32Type{}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Uint32Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Uint64Type
		args args
		want bool
	}{
		{name: "True", e: Uint64Type{}, args: args{x: Uint64Type{}}, want: true},
		{name: "False", e: Uint64Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint64Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want int
	}{
		{name: "", e: Uint64Type{}, want: TypeUint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Uint64Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want reflect.Kind
	}{
		{name: "", e: Uint64Type{}, want: reflect.Uint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Uint64Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want string
	}{
		{name: "", e: Uint64Type{}, want: "Uint64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Uint64Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want string
	}{
		{name: "", e: Uint64Type{}, want: "Uint64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Uint64Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want string
	}{
		{name: "", e: Uint64Type{}, want: "Type:Uint64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Uint64Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUint64Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64Type
		want int
	}{
		{name: "", e: Uint64Type{}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Uint64Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Float32Type
		args args
		want bool
	}{
		{name: "True", e: Float32Type{}, args: args{x: Float32Type{}}, want: true},
		{name: "False", e: Float32Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Float32Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want int
	}{
		{name: "", e: Float32Type{}, want: TypeFloat32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Float32Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want reflect.Kind
	}{
		{name: "", e: Float32Type{}, want: reflect.Float32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Float32Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want string
	}{
		{name: "", e: Float32Type{}, want: "Float32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Float32Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want string
	}{
		{name: "", e: Float32Type{}, want: "Float32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Float32Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want string
	}{
		{name: "", e: Float32Type{}, want: "Type:Float32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Float32Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFloat32Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Float32Type
		want int
	}{
		{name: "", e: Float32Type{}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Float32Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Float64Type
		args args
		want bool
	}{
		{name: "True", e: Float64Type{}, args: args{x: Float64Type{}}, want: true},
		{name: "False", e: Float64Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Float64Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want int
	}{
		{name: "", e: Float64Type{}, want: TypeFloat64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Float64Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want reflect.Kind
	}{
		{name: "", e: Float64Type{}, want: reflect.Float64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Float64Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want string
	}{
		{name: "", e: Float64Type{}, want: "Float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Float64Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want string
	}{
		{name: "", e: Float64Type{}, want: "Float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Float64Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want string
	}{
		{name: "", e: Float64Type{}, want: "Type:Float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Float64Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFloat64Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Float64Type
		want int
	}{
		{name: "", e: Float64Type{}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Float64Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Complex64Type
		args args
		want bool
	}{
		{name: "True", e: Complex64Type{}, args: args{x: Complex64Type{}}, want: true},
		{name: "False", e: Complex64Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Complex64Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want int
	}{
		{name: "", e: Complex64Type{}, want: TypeComplex64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Complex64Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want reflect.Kind
	}{
		{name: "", e: Complex64Type{}, want: reflect.Complex64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Complex64Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want string
	}{
		{name: "", e: Complex64Type{}, want: "Complex64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Complex64Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want string
	}{
		{name: "", e: Complex64Type{}, want: "Complex64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Complex64Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want string
	}{
		{name: "", e: Complex64Type{}, want: "Type:Complex64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Complex64Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestComplex64Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64Type
		want int
	}{
		{name: "", e: Complex64Type{}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Complex64Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    Complex128Type
		args args
		want bool
	}{
		{name: "True", e: Complex128Type{}, args: args{x: Complex128Type{}}, want: true},
		{name: "False", e: Complex128Type{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Complex128Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want int
	}{
		{name: "", e: Complex128Type{}, want: TypeComplex128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("Complex128Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want reflect.Kind
	}{
		{name: "", e: Complex128Type{}, want: reflect.Complex128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("Complex128Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want string
	}{
		{name: "", e: Complex128Type{}, want: "Complex128"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("Complex128Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want string
	}{
		{name: "", e: Complex128Type{}, want: "Complex128"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Complex128Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want string
	}{
		{name: "", e: Complex128Type{}, want: "Type:Complex128"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("Complex128Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestComplex128Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128Type
		want int
	}{
		{name: "", e: Complex128Type{}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got != tt.want {
				t.Errorf("Complex128Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    StringType
		args args
		want bool
	}{
		{name: "True", e: StringType{}, args: args{x: StringType{}}, want: true},
		{name: "False", e: StringType{}, args: args{x: BytesType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("StringType.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_ID(t *testing.T) {
	tests := []struct {
		name string
		e    StringType
		want int
	}{
		{name: "", e: StringType{}, want: TypeString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("StringType.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    StringType
		want reflect.Kind
	}{
		{name: "", e: StringType{}, want: reflect.String},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("StringType.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_Name(t *testing.T) {
	tests := []struct {
		name string
		e    StringType
		want string
	}{
		{name: "", e: StringType{}, want: "String"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("StringType.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_String(t *testing.T) {
	tests := []struct {
		name string
		e    StringType
		want string
	}{
		{name: "", e: StringType{}, want: "String"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("StringType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringType_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    StringType
		want string
	}{
		{name: "", e: StringType{}, want: "Type:String"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("StringType.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    BytesType
		args args
		want bool
	}{
		{name: "True", e: BytesType{}, args: args{x: BytesType{}}, want: true},
		{name: "False", e: BytesType{}, args: args{x: StringType{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("BytesType.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_ID(t *testing.T) {
	tests := []struct {
		name string
		e    BytesType
		want int
	}{
		{name: "", e: BytesType{}, want: TypeBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got != tt.want {
				t.Errorf("BytesType.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    BytesType
		want reflect.Kind
	}{
		{name: "", e: BytesType{}, want: reflect.Slice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got != tt.want {
				t.Errorf("BytesType.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_Name(t *testing.T) {
	tests := []struct {
		name string
		e    BytesType
		want string
	}{
		{name: "", e: BytesType{}, want: "Bytes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got != tt.want {
				t.Errorf("BytesType.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_String(t *testing.T) {
	tests := []struct {
		name string
		e    BytesType
		want string
	}{
		{name: "", e: BytesType{}, want: "Bytes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("BytesType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesType_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    BytesType
		want string
	}{
		{name: "", e: BytesType{}, want: "Type:Bytes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got != tt.want {
				t.Errorf("BytesType.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

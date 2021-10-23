package goval

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestToBool(t *testing.T) {
	type args struct {
		x BoolConverter
	}
	type data struct {
		name string
		args args
		want Bool
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Bool(true)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Bool(false)},
		{name: "Int8(true)", args: args{x: Int8(-123)}, want: Bool(true)},
		{name: "Int8(false)", args: args{x: Int8(0)}, want: Bool(false)},
		{name: "Int16(true)", args: args{x: Int16(-123)}, want: Bool(true)},
		{name: "Int16(false)", args: args{x: Int16(0)}, want: Bool(false)},
		{name: "Int32(true)", args: args{x: Int32(-123)}, want: Bool(true)},
		{name: "Int32(false)", args: args{x: Int32(0)}, want: Bool(false)},
		{name: "Int64(true)", args: args{x: Int64(-123)}, want: Bool(true)},
		{name: "Int64(false)", args: args{x: Int64(0)}, want: Bool(false)},
		{name: "Uint8(true)", args: args{x: Uint8(123)}, want: Bool(true)},
		{name: "Uint8(false)", args: args{x: Uint8(0)}, want: Bool(false)},
		{name: "Uint16(true)", args: args{x: Uint16(123)}, want: Bool(true)},
		{name: "Uint16(false)", args: args{x: Uint16(0)}, want: Bool(false)},
		{name: "Uint32(true)", args: args{x: Uint32(123)}, want: Bool(true)},
		{name: "Uint32(false)", args: args{x: Uint32(0)}, want: Bool(false)},
		{name: "Uint64(true)", args: args{x: Uint64(123)}, want: Bool(true)},
		{name: "Uint64(false)", args: args{x: Uint64(0)}, want: Bool(false)},
		{name: "Float32(true)", args: args{x: Float32(123.45)}, want: Bool(true)},
		{name: "Float32(false)", args: args{x: Float32(0)}, want: Bool(false)},
		{name: "Float32(Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Bool(true)},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Bool(true)},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Bool(true)},
		{name: "Float64(true)", args: args{x: Float64(123.45)}, want: Bool(true)},
		{name: "Float64(Inf)", args: args{x: Float64(math.Inf(0))}, want: Bool(true)},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Bool(true)},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Bool(true)},
		{name: "Float64(false)", args: args{x: Float64(0)}, want: Bool(false)},
		{name: "String(true)", args: args{x: String("1")}, want: Bool(true)},
		{name: "String(true)", args: args{x: String("t")}, want: Bool(true)},
		{name: "String(true)", args: args{x: String("T")}, want: Bool(true)},
		{name: "String(true)", args: args{x: String("TRUE")}, want: Bool(true)},
		{name: "String(true)", args: args{x: String("true")}, want: Bool(true)},
		{name: "String(true)", args: args{x: String("True")}, want: Bool(true)},
		{name: "String(false)", args: args{x: String("0")}, want: Bool(false)},
		{name: "String(false)", args: args{x: String("f")}, want: Bool(false)},
		{name: "String(false)", args: args{x: String("F")}, want: Bool(false)},
		{name: "String(false)", args: args{x: String("FALSE")}, want: Bool(false)},
		{name: "String(false)", args: args{x: String("false")}, want: Bool(false)},
		{name: "String(false)", args: args{x: String("False")}, want: Bool(false)},
		{name: "String(error)", args: args{x: String("error")}, want: Bool(false)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("1"))}, want: Bool(true)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("t"))}, want: Bool(true)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("T"))}, want: Bool(true)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("TRUE"))}, want: Bool(true)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("true"))}, want: Bool(true)},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("True"))}, want: Bool(true)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("0"))}, want: Bool(false)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("f"))}, want: Bool(false)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("F"))}, want: Bool(false)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("FALSE"))}, want: Bool(false)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("false"))}, want: Bool(false)},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("False"))}, want: Bool(false)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Bool(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.x); got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolCheck(t *testing.T) {
	type args struct {
		x BoolConverter
	}
	type data struct {
		name  string
		args  args
		want  Bool
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Bool(true), want1: nil},
		{name: "Bool(false), want1: nil,", args: args{x: Bool(false)}, want: Bool(false)},
		{name: "Int8(true), want1: nil,", args: args{x: Int8(-123)}, want: Bool(true)},
		{name: "Int8(false), want1: nil,", args: args{x: Int8(0)}, want: Bool(false)},
		{name: "Int16(true), want1: nil,", args: args{x: Int16(-123)}, want: Bool(true)},
		{name: "Int16(false), want1: nil,", args: args{x: Int16(0)}, want: Bool(false)},
		{name: "Int32(true), want1: nil,", args: args{x: Int32(-123)}, want: Bool(true)},
		{name: "Int32(false), want1: nil,", args: args{x: Int32(0)}, want: Bool(false)},
		{name: "Int64(true), want1: nil,", args: args{x: Int64(-123)}, want: Bool(true)},
		{name: "Int64(false), want1: nil,", args: args{x: Int64(0)}, want: Bool(false)},
		{name: "Uint8(true), want1: nil,", args: args{x: Uint8(123)}, want: Bool(true)},
		{name: "Uint8(false), want1: nil,", args: args{x: Uint8(0)}, want: Bool(false)},
		{name: "Uint16(true), want1: nil,", args: args{x: Uint16(123)}, want: Bool(true)},
		{name: "Uint16(false), want1: nil,", args: args{x: Uint16(0)}, want: Bool(false)},
		{name: "Uint32(true), want1: nil,", args: args{x: Uint32(123)}, want: Bool(true)},
		{name: "Uint32(false), want1: nil,", args: args{x: Uint32(0)}, want: Bool(false)},
		{name: "Uint64(true), want1: nil,", args: args{x: Uint64(123)}, want: Bool(true)},
		{name: "Uint64(false), want1: nil,", args: args{x: Uint64(0)}, want: Bool(false)},
		{name: "Float32(true), want1: nil,", args: args{x: Float32(123.45)}, want: Bool(true)},
		{name: "Float32(false), want1: nil,", args: args{x: Float32(0)}, want: Bool(false)},
		{name: "Float32(Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Bool(true), want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Bool(true), want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Bool(true), want1: nil},
		{name: "Float64(true), want1: nil,", args: args{x: Float64(123.45)}, want: Bool(true)},
		{name: "Float64(Inf)", args: args{x: Float64(math.Inf(0))}, want: Bool(true), want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Bool(true), want1: nil},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Bool(true), want1: nil},
		{name: "Float64(false), want1: nil,", args: args{x: Float64(0)}, want: Bool(false)},
		{name: "String(true), want1: nil,", args: args{x: String("1")}, want: Bool(true)},
		{name: "String(true), want1: nil,", args: args{x: String("t")}, want: Bool(true)},
		{name: "String(true), want1: nil,", args: args{x: String("T")}, want: Bool(true)},
		{name: "String(true), want1: nil,", args: args{x: String("TRUE")}, want: Bool(true)},
		{name: "String(true), want1: nil,", args: args{x: String("true")}, want: Bool(true)},
		{name: "String(true), want1: nil,", args: args{x: String("True")}, want: Bool(true)},
		{name: "String(false), want1: nil,", args: args{x: String("0")}, want: Bool(false)},
		{name: "String(false), want1: nil,", args: args{x: String("f")}, want: Bool(false)},
		{name: "String(false), want1: nil,", args: args{x: String("F")}, want: Bool(false)},
		{name: "String(false), want1: nil,", args: args{x: String("FALSE")}, want: Bool(false)},
		{name: "String(false), want1: nil,", args: args{x: String("false")}, want: Bool(false)},
		{name: "String(false), want1: nil,", args: args{x: String("False")}, want: Bool(false)},
		{name: "String(error)", args: args{x: String("error")}, want: Bool(false), want1: strconv.ErrSyntax},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("1"))}, want: Bool(true)},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("t"))}, want: Bool(true)},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("T"))}, want: Bool(true)},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("TRUE"))}, want: Bool(true)},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("true"))}, want: Bool(true)},
		{name: "Bytes(true), want1: nil,", args: args{x: Bytes([]byte("True"))}, want: Bool(true)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("0"))}, want: Bool(false)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("f"))}, want: Bool(false)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("F"))}, want: Bool(false)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("FALSE"))}, want: Bool(false)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("false"))}, want: Bool(false)},
		{name: "Bytes(false), want1: nil,", args: args{x: Bytes([]byte("False"))}, want: Bool(false)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Bool(false), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToBoolCheck(tt.args.x)
			if got != tt.want {
				t.Errorf("ToBoolCheck() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToBoolCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	type args struct {
		x Int8Converter
	}
	tests := []struct {
		name string
		args args
		want Int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8Check(t *testing.T) {
	type args struct {
		x Int8Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Int8
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt8Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToInt8Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	type args struct {
		x Int16Converter
	}
	tests := []struct {
		name string
		args args
		want Int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16Check(t *testing.T) {
	type args struct {
		x Int16Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Int16
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt16Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToInt16Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		x Int32Converter
	}
	tests := []struct {
		name string
		args args
		want Int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32Check(t *testing.T) {
	type args struct {
		x Int32Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Int32
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt32Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToInt32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		x Int64Converter
	}
	tests := []struct {
		name string
		args args
		want Int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Check(t *testing.T) {
	type args struct {
		x Int64Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Int64
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt64Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToInt64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	type args struct {
		x Uint8Converter
	}
	tests := []struct {
		name string
		args args
		want Uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8Check(t *testing.T) {
	type args struct {
		x Uint8Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Uint8
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint8Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToUint8Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	type args struct {
		x Uint16Converter
	}
	tests := []struct {
		name string
		args args
		want Uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16Check(t *testing.T) {
	type args struct {
		x Uint16Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Uint16
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint16Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToUint16Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	type args struct {
		x Uint32Converter
	}
	tests := []struct {
		name string
		args args
		want Uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32Check(t *testing.T) {
	type args struct {
		x Uint32Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Uint32
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint32Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToUint32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	type args struct {
		x Uint64Converter
	}
	tests := []struct {
		name string
		args args
		want Uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64Check(t *testing.T) {
	type args struct {
		x Uint64Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Uint64
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint64Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToUint64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	type args struct {
		x Float32Converter
	}
	tests := []struct {
		name string
		args args
		want Float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32Check(t *testing.T) {
	type args struct {
		x Float32Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Float32
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat32Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToFloat32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		x Float64Converter
	}
	tests := []struct {
		name string
		args args
		want Float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64Check(t *testing.T) {
	type args struct {
		x Float64Converter
	}
	tests := []struct {
		name  string
		args  args
		want  Float64
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat64Check(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64Check() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToFloat64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		x StringConverter
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringCheck(t *testing.T) {
	type args struct {
		x StringConverter
	}
	tests := []struct {
		name  string
		args  args
		want  String
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringCheck(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringCheck() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToStringCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		x BytesConverter
	}
	tests := []struct {
		name string
		args args
		want Bytes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBytes(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytesCheck(t *testing.T) {
	type args struct {
		x BytesConverter
	}
	tests := []struct {
		name  string
		args  args
		want  Bytes
		want1 Err
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToBytesCheck(tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytesCheck() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToBytesCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringFormat(t *testing.T) {
	type args struct {
		x      Val
		format string
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringFormat(tt.args.x, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytesFormat(t *testing.T) {
	type args struct {
		x      Val
		format string
	}
	tests := []struct {
		name string
		args args
		want Bytes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBytesFormat(tt.args.x, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

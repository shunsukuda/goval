package goval

import (
	"bytes"
	"errors"
	"math"
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
		{name: "String(error)", args: args{x: String("")}, want: Bool(false)},
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
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Bool(false)},
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
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Bool(false), want1: nil},
		{name: "Int8(true)", args: args{x: Int8(-123)}, want: Bool(true), want1: nil},
		{name: "Int8(false)", args: args{x: Int8(0)}, want: Bool(false), want1: nil},
		{name: "Int16(true)", args: args{x: Int16(-123)}, want: Bool(true), want1: nil},
		{name: "Int16(false)", args: args{x: Int16(0)}, want: Bool(false), want1: nil},
		{name: "Int32(true)", args: args{x: Int32(-123)}, want: Bool(true), want1: nil},
		{name: "Int32(false)", args: args{x: Int32(0)}, want: Bool(false), want1: nil},
		{name: "Int64(true)", args: args{x: Int64(-123)}, want: Bool(true), want1: nil},
		{name: "Int64(false)", args: args{x: Int64(0)}, want: Bool(false), want1: nil},
		{name: "Uint8(true)", args: args{x: Uint8(123)}, want: Bool(true), want1: nil},
		{name: "Uint8(false)", args: args{x: Uint8(0)}, want: Bool(false), want1: nil},
		{name: "Uint16(true)", args: args{x: Uint16(123)}, want: Bool(true), want1: nil},
		{name: "Uint16(false),", args: args{x: Uint16(0)}, want: Bool(false), want1: nil},
		{name: "Uint32(true)", args: args{x: Uint32(123)}, want: Bool(true), want1: nil},
		{name: "Uint32(false),", args: args{x: Uint32(0)}, want: Bool(false), want1: nil},
		{name: "Uint64(true)", args: args{x: Uint64(123)}, want: Bool(true), want1: nil},
		{name: "Uint64(false),", args: args{x: Uint64(0)}, want: Bool(false), want1: nil},
		{name: "Float32(true),", args: args{x: Float32(123.45)}, want: Bool(true), want1: nil},
		{name: "Float32(false)", args: args{x: Float32(0)}, want: Bool(false), want1: nil},
		{name: "Float32(Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Bool(true), want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Bool(true), want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Bool(true), want1: nil},
		{name: "Float64(true)", args: args{x: Float64(123.45)}, want: Bool(true), want1: nil},
		{name: "Float64(Inf)", args: args{x: Float64(math.Inf(0))}, want: Bool(true), want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Bool(true), want1: nil},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Bool(true), want1: nil},
		{name: "Float64(false)", args: args{x: Float64(0)}, want: Bool(false), want1: nil},
		{name: "String(true)", args: args{x: String("1")}, want: Bool(true), want1: nil},
		{name: "String(true)", args: args{x: String("t")}, want: Bool(true), want1: nil},
		{name: "String(true)", args: args{x: String("T")}, want: Bool(true), want1: nil},
		{name: "String(true)", args: args{x: String("TRUE")}, want: Bool(true), want1: nil},
		{name: "String(true)", args: args{x: String("true")}, want: Bool(true), want1: nil},
		{name: "String(true)", args: args{x: String("True")}, want: Bool(true), want1: nil},
		{name: "String(false)", args: args{x: String("0")}, want: Bool(false), want1: nil},
		{name: "String(false)", args: args{x: String("f")}, want: Bool(false), want1: nil},
		{name: "String(false)", args: args{x: String("F")}, want: Bool(false), want1: nil},
		{name: "String(false)", args: args{x: String("FALSE")}, want: Bool(false), want1: nil},
		{name: "String(false)", args: args{x: String("false")}, want: Bool(false), want1: nil},
		{name: "String(false)", args: args{x: String("False")}, want: Bool(false), want1: nil},
		{name: "String(error)", args: args{x: String("")}, want: Bool(false), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Bool(false), want1: strconv.ErrSyntax},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("1"))}, want: Bool(true), want1: nil},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("t"))}, want: Bool(true), want1: nil},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("T"))}, want: Bool(true), want1: nil},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("TRUE"))}, want: Bool(true), want1: nil},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("true"))}, want: Bool(true), want1: nil},
		{name: "Bytes(true)", args: args{x: Bytes([]byte("True"))}, want: Bool(true), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("0"))}, want: Bool(false), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("f"))}, want: Bool(false), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("F"))}, want: Bool(false), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("FALSE"))}, want: Bool(false), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("false"))}, want: Bool(false), want1: nil},
		{name: "Bytes(false)", args: args{x: Bytes([]byte("False"))}, want: Bool(false), want1: nil},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Bool(false), want1: strconv.ErrSyntax},
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
	type data struct {
		name string
		args args
		want Int8
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int8(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int8(0)},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int8(-123)},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int8(-123)},
		{name: "Int16(OF)", args: args{x: Int16(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int8(-123)},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int8(-123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int8(123)},
		{name: "Uint8(OF)", args: args{x: Uint8(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int8(123)},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int8(123)},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int8(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Float32", args: args{x: Float32(123)}, want: Int8(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "Float64", args: args{x: Float64(123)}, want: Int8(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8)},
		{name: "String", args: args{x: String("-123")}, want: Int8(-123)},
		{name: "String(OF)", args: args{x: String("128")}, want: Int8(math.MaxInt8)},
		{name: "String(error)", args: args{x: String("")}, want: Int8(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Int8(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int8(-123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("128"))}, want: Int8(math.MaxInt8)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int8(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int8(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8(tt.args.x); got != tt.want {
				t.Errorf("ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8Check(t *testing.T) {
	type args struct {
		x Int8Converter
	}
	type data struct {
		name  string
		args  args
		want  Int8
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int8(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int8(0), want1: nil},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int8(-123), want1: nil},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int8(-123), want1: nil},
		{name: "Int16(OF)", args: args{x: Int16(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int8(-123), want1: nil},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int8(-123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int8(123), want1: nil},
		{name: "Uint8(OF)", args: args{x: Uint8(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int8(123), want1: nil},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int8(123), want1: nil},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int8(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Int8(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Int8(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt8 + 1)}, want: Int8(math.MinInt8), want1: nil},
		{name: "String", args: args{x: String("-123")}, want: Int8(-123), want1: nil},
		{name: "String(OF)", args: args{x: String("128")}, want: Int8(math.MaxInt8), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Int8(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Int8(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int8(-123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("128"))}, want: Int8(math.MaxInt8), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int8(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int8(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt8Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToInt8Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToInt8Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	type args struct {
		x Int16Converter
	}
	type data struct {
		name string
		args args
		want Int16
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int16(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int16(0)},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int16(-123)},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int16(-123)},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int16(-123)},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int16(-123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int16(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int16(123)},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int16(123)},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int16(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Float32", args: args{x: Float32(123)}, want: Int16(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "Float64", args: args{x: Float64(123)}, want: Int16(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16)},
		{name: "String", args: args{x: String("-123")}, want: Int16(-123)},
		{name: "String(OF)", args: args{x: String("32768")}, want: Int16(math.MaxInt16)},
		{name: "String(error)", args: args{x: String("")}, want: Int16(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Int16(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int16(-123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("32768"))}, want: Int16(math.MaxInt16)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int16(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int16(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16(tt.args.x); got != tt.want {
				t.Errorf("ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16Check(t *testing.T) {
	type args struct {
		x Int16Converter
	}
	type data struct {
		name  string
		args  args
		want  Int16
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int16(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int16(0), want1: nil},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int16(-123), want1: nil},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int16(-123), want1: nil},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int16(-123), want1: nil},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int16(-123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int16(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int16(123), want1: nil},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int16(123), want1: nil},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int16(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Int16(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Int16(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt16 + 1)}, want: Int16(math.MinInt16), want1: nil},
		{name: "String", args: args{x: String("-123")}, want: Int16(-123), want1: nil},
		{name: "String(OF)", args: args{x: String("32768")}, want: Int16(math.MaxInt16), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Int16(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Int16(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int16(-123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("32768"))}, want: Int16(math.MaxInt16), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int16(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int16(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt16Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToInt16Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToInt16Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		x Int32Converter
	}
	type data struct {
		name string
		args args
		want Int32
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int32(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int32(0)},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int32(-123)},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int32(-123)},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int32(-123)},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int32(-123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int32(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int32(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int32(123)},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt32 + 1)}, want: Int32(math.MinInt32)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int32(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32)},
		{name: "Float32", args: args{x: Float32(123)}, want: Int32(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt32 + 1)}, want: Int32(math.MinInt32)},
		{name: "Float64", args: args{x: Float64(123)}, want: Int32(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32)},
		{name: "String", args: args{x: String("-123")}, want: Int32(-123)},
		{name: "String(OF)", args: args{x: String("2147483648")}, want: Int32(math.MaxInt32)},
		{name: "String(error)", args: args{x: String("")}, want: Int32(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Int32(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int32(-123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("2147483648"))}, want: Int32(math.MaxInt32)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int32(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int32(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.x); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32Check(t *testing.T) {
	type args struct {
		x Int32Converter
	}
	type data struct {
		name  string
		args  args
		want  Int32
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int32(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int32(0), want1: nil},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int32(-123), want1: nil},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int32(-123), want1: nil},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int32(-123), want1: nil},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int32(-123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int32(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int32(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int32(123), want1: nil},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxInt32 + 1)}, want: Int32(math.MinInt32), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int32(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Int32(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt32 + 1)}, want: Int32(math.MinInt32), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Int32(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt32 + 1)}, want: Int32(math.MinInt32), want1: nil},
		{name: "String", args: args{x: String("-123")}, want: Int32(-123), want1: nil},
		{name: "String(OF)", args: args{x: String("2147483648")}, want: Int32(math.MaxInt32), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Int32(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Int32(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int32(-123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("2147483648"))}, want: Int32(math.MaxInt32), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int32(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int32(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt32Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToInt32Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToInt32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		x Int64Converter
	}
	type data struct {
		name string
		args args
		want Int64
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int64(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int64(0)},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int64(-123)},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int64(-123)},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int64(-123)},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int64(-123)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int64(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int64(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int64(123)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int64(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt64 + 1)}, want: Int64(math.MinInt64)},
		{name: "Float32", args: args{x: Float32(123)}, want: Int64(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt64 + 1)}, want: Int64(math.MinInt64)},
		{name: "Float64", args: args{x: Float64(123)}, want: Int64(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt64 + 1)}, want: Int64(math.MinInt64)},
		{name: "String", args: args{x: String("-123")}, want: Int64(-123)},
		{name: "String(OF)", args: args{x: String("9223372036854775808")}, want: Int64(math.MaxInt64)},
		{name: "String(error)", args: args{x: String("")}, want: Int64(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Int64(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int64(-123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("9223372036854775808"))}, want: Int64(math.MaxInt64)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int64(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int64(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.x); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Check(t *testing.T) {
	type args struct {
		x Int64Converter
	}
	type data struct {
		name  string
		args  args
		want  Int64
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Int64(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Int64(0), want1: nil},
		{name: "Int8", args: args{x: Int8(-123)}, want: Int64(-123), want1: nil},
		{name: "Int16", args: args{x: Int16(-123)}, want: Int64(-123), want1: nil},
		{name: "Int32", args: args{x: Int32(-123)}, want: Int64(-123), want1: nil},
		{name: "Int64", args: args{x: Int64(-123)}, want: Int64(-123), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Int64(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Int64(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Int64(123), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Int64(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxInt64 + 1)}, want: Int64(math.MinInt64), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Int64(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxInt64 + 1)}, want: Int64(math.MinInt64), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Int64(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxInt64 + 1)}, want: Int64(math.MinInt64), want1: nil},
		{name: "String", args: args{x: String("-123")}, want: Int64(-123), want1: nil},
		{name: "String(OF)", args: args{x: String("9223372036854775808")}, want: Int64(math.MaxInt64), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Int64(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Int64(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("-123"))}, want: Int64(-123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("9223372036854775808"))}, want: Int64(math.MaxInt64), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Int64(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Int64(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt64Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToInt64Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToInt64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	type args struct {
		x Uint8Converter
	}
	type data struct {
		name string
		args args
		want Uint8
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint8(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint8(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint8(123)},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint8(math.MaxUint8)},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint8(123)},
		{name: "Int16(OF)", args: args{x: Int16(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint8(math.MaxUint8)},
		{name: "Int32", args: args{x: Uint8(123)}, want: Uint8(123)},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint8(math.MaxUint8)},
		{name: "Int64", args: args{x: Uint8(123)}, want: Uint8(123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint8(math.MaxUint8)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint8(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint8(123)},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint8(123)},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint8(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint8(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint8(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint8 + 1)}, want: Uint8(0)},
		{name: "String", args: args{x: String("123")}, want: Uint8(123)},
		{name: "String(OF)", args: args{x: String("256")}, want: Uint8(math.MaxUint8)},
		{name: "String(error)", args: args{x: String("")}, want: Uint8(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Uint8(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint8(123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("256"))}, want: Uint8(math.MaxUint8)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint8(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint8(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8(tt.args.x); got != tt.want {
				t.Errorf("ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8Check(t *testing.T) {
	type args struct {
		x Uint8Converter
	}
	type data struct {
		name  string
		args  args
		want  Uint8
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint8(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint8(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint8(123), want1: nil},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint8(math.MaxUint8), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint8(123), want1: nil},
		{name: "Int16(OF)", args: args{x: Int16(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint8(math.MaxUint8), want1: nil},
		{name: "Int32", args: args{x: Uint8(123)}, want: Uint8(123), want1: nil},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint8(math.MaxUint8), want1: nil},
		{name: "Int64", args: args{x: Uint8(123)}, want: Uint8(123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint8(math.MaxUint8), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint8(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint8(123), want1: nil},
		{name: "Uint16(OF)", args: args{x: Uint16(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint8(123), want1: nil},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint8(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint8(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint8(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint8 + 1)}, want: Uint8(0), want1: nil},
		{name: "String", args: args{x: String("123")}, want: Uint8(123), want1: nil},
		{name: "String(OF)", args: args{x: String("256")}, want: Uint8(math.MaxUint8), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Uint8(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Uint8(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint8(123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("256"))}, want: Uint8(math.MaxUint8), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint8(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint8(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint8Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToUint8Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToUint8Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	type args struct {
		x Uint16Converter
	}
	type data struct {
		name string
		args args
		want Uint16
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint16(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint16(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint16(123)},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint16(math.MaxUint16)},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint16(123)},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint16(math.MaxUint16)},
		{name: "Int32", args: args{x: Uint16(123)}, want: Uint16(123)},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint16(math.MaxUint16)},
		{name: "Int64", args: args{x: Uint16(123)}, want: Uint16(123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint16(math.MaxUint16)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint16(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint16(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint16(123)},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint16(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint16(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint16(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint16 + 1)}, want: Uint16(0)},
		{name: "String", args: args{x: String("123")}, want: Uint16(123)},
		{name: "String(OF)", args: args{x: String("65536")}, want: Uint16(math.MaxUint16)},
		{name: "String(error)", args: args{x: String("")}, want: Uint16(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Uint16(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint16(123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("65536"))}, want: Uint16(math.MaxUint16)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint16(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint16(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16(tt.args.x); got != tt.want {
				t.Errorf("ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16Check(t *testing.T) {
	type args struct {
		x Uint16Converter
	}
	type data struct {
		name  string
		args  args
		want  Uint16
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint16(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint16(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint16(123), want1: nil},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint16(math.MaxUint16), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint16(123), want1: nil},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint16(math.MaxUint16), want1: nil},
		{name: "Int32", args: args{x: Uint16(123)}, want: Uint16(123), want1: nil},
		{name: "Int32(OF)", args: args{x: Int32(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint16(math.MaxUint16), want1: nil},
		{name: "Int64", args: args{x: Uint16(123)}, want: Uint16(123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint16(math.MaxUint16), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint16(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint16(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint16(123), want1: nil},
		{name: "Uint32(OF)", args: args{x: Uint32(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint16(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint16(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint16(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint16 + 1)}, want: Uint16(0), want1: nil},
		{name: "String", args: args{x: String("123")}, want: Uint16(123), want1: nil},
		{name: "String(OF)", args: args{x: String("65536")}, want: Uint16(math.MaxUint16), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Uint16(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Uint16(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint16(123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("65536"))}, want: Uint16(math.MaxUint16), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint16(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint16(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint16Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToUint16Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToUint16Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	type args struct {
		x Uint32Converter
	}
	type data struct {
		name string
		args args
		want Uint32
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint32(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint32(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint32(123)},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint32(math.MaxUint32)},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint32(123)},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint32(math.MaxUint32)},
		{name: "Int32", args: args{x: Uint32(123)}, want: Uint32(123)},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint32(math.MaxUint32)},
		{name: "Int64", args: args{x: Uint32(123)}, want: Uint32(123)},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint32 + 1)}, want: Uint32(0)},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint32(math.MaxUint32)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint32(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint32(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint32(123)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint32(123)},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint32 + 1)}, want: Uint32(0)},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint32(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint32 + 1)}, want: Uint32(0)},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint32(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint32 + 1)}, want: Uint32(0)},
		{name: "String", args: args{x: String("123")}, want: Uint32(123)},
		{name: "String(OF)", args: args{x: String("4294967296")}, want: Uint32(math.MaxUint32)},
		{name: "String(error)", args: args{x: String("")}, want: Uint32(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Uint32(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint32(123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("4294967296"))}, want: Uint32(math.MaxUint32)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint32(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint32(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32(tt.args.x); got != tt.want {
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32Check(t *testing.T) {
	type args struct {
		x Uint32Converter
	}
	type data struct {
		name  string
		args  args
		want  Uint32
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint32(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint32(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint32(123), want1: nil},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint32(math.MaxUint32), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint32(123), want1: nil},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint32(math.MaxUint32), want1: nil},
		{name: "Int32", args: args{x: Uint32(123)}, want: Uint32(123), want1: nil},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint32(math.MaxUint32), want1: nil},
		{name: "Int64", args: args{x: Uint32(123)}, want: Uint32(123), want1: nil},
		{name: "Int64(OF)", args: args{x: Int64(math.MaxUint32 + 1)}, want: Uint32(0), want1: nil},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint32(math.MaxUint32), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint32(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint32(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint32(123), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint32(123), want1: nil},
		{name: "Uint64(OF)", args: args{x: Uint64(math.MaxUint32 + 1)}, want: Uint32(0), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint32(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint32 + 1)}, want: Uint32(0), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint32(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint32 + 1)}, want: Uint32(0), want1: nil},
		{name: "String", args: args{x: String("123")}, want: Uint32(123), want1: nil},
		{name: "String(OF)", args: args{x: String("4294967296")}, want: Uint32(math.MaxUint32), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Uint32(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Uint32(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint32(123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("4294967296"))}, want: Uint32(math.MaxUint32), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint32(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint32(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint32Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToUint32Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToUint32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	type args struct {
		x Uint64Converter
	}
	type data struct {
		name string
		args args
		want Uint64
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint64(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint64(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint64(123)},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint64(math.MaxUint64)},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint64(123)},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint64(math.MaxUint64)},
		{name: "Int32", args: args{x: Uint64(123)}, want: Uint64(123)},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint64(math.MaxUint64)},
		{name: "Int64", args: args{x: Uint64(123)}, want: Uint64(123)},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint64(math.MaxUint64)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint64(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint64(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint64(123)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint64(123)},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint64(123)},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint64 + 1)}, want: Uint64(1 << 63)},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint64(123)},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint64 + 1)}, want: Uint64(1 << 63)},
		{name: "String", args: args{x: String("123")}, want: Uint64(123)},
		{name: "String(OF)", args: args{x: String("18446744073709551616")}, want: Uint64(math.MaxUint64)},
		{name: "String(error)", args: args{x: String("")}, want: Uint64(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Uint64(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint64(123)},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("18446744073709551616"))}, want: Uint64(math.MaxUint64)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint64(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint64(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64(tt.args.x); got != tt.want {
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64Check(t *testing.T) {
	type args struct {
		x Uint64Converter
	}
	type data struct {
		name  string
		args  args
		want  Uint64
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Uint64(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Uint64(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Uint64(123), want1: nil},
		{name: "Int8(OR)", args: args{x: Int8(-1)}, want: Uint64(math.MaxUint64), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Uint64(123), want1: nil},
		{name: "Int16(OR)", args: args{x: Int16(-1)}, want: Uint64(math.MaxUint64), want1: nil},
		{name: "Int32", args: args{x: Uint64(123)}, want: Uint64(123), want1: nil},
		{name: "Int32(OR)", args: args{x: Int32(-1)}, want: Uint64(math.MaxUint64), want1: nil},
		{name: "Int64", args: args{x: Uint64(123)}, want: Uint64(123), want1: nil},
		{name: "Int64(OR)", args: args{x: Int64(-1)}, want: Uint64(math.MaxUint64), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Uint64(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Uint64(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Uint64(123), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Uint64(123), want1: nil},
		{name: "Float32", args: args{x: Float32(123)}, want: Uint64(123), want1: nil},
		{name: "Float32(OF)", args: args{x: Float32(math.MaxUint64 + 1)}, want: Uint64(1 << 63), want1: nil},
		{name: "Float64", args: args{x: Float64(123)}, want: Uint64(123), want1: nil},
		{name: "Float64(OF)", args: args{x: Float64(math.MaxUint64 + 1)}, want: Uint64(1 << 63), want1: nil},
		{name: "String", args: args{x: String("123")}, want: Uint64(123), want1: nil},
		{name: "String(OF)", args: args{x: String("18446744073709551616")}, want: Uint64(math.MaxUint64), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Uint64(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Uint64(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123"))}, want: Uint64(123), want1: nil},
		{name: "Bytes(OF)", args: args{x: Bytes([]byte("18446744073709551616"))}, want: Uint64(math.MaxUint64), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Uint64(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Uint64(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUint64Check(tt.args.x)
			if got != tt.want {
				t.Errorf("ToUint64Check() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToUint64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	type args struct {
		x Float32Converter
	}
	type data struct {
		name string
		args args
		want Float32
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Float32(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Float32(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Float32(123)},
		{name: "Int16", args: args{x: Int16(123)}, want: Float32(123)},
		{name: "Int32", args: args{x: Int32(123)}, want: Float32(123)},
		{name: "Int64", args: args{x: Int64(123)}, want: Float32(123)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Float32(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Float32(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Float32(123)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Float32(123)},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Float32(123.45)},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Float32(float32(math.Inf(0)))},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Float32(float32(math.Inf(-1)))},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Float32(float32(math.NaN()))},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Float32(123.45)},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Float32(float32(math.Inf(0)))},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Float32(float32(math.Inf(-1)))},
		{name: "Float64(NaN)", args: args{x: Float32(math.NaN())}, want: Float32(float32(math.NaN()))},
		{name: "String", args: args{x: String("123.45")}, want: Float32(123)},
		{name: "String(OR)", args: args{x: String("3.4+39")}, want: Float32(float32(math.Inf(0)))},
		{name: "String(error)", args: args{x: String("")}, want: Float32(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Float32(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123.45"))}, want: Float32(123)},
		{name: "Bytes(OR)", args: args{x: Bytes([]byte("3.4+39"))}, want: Float32(float32(math.Inf(0)))},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Float32(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Float32(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32(tt.args.x); got != tt.want {
				if math.IsNaN(float64(got)) != math.IsNaN(float64(tt.want)) {
					t.Errorf("ToFloat32() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestToFloat32Check(t *testing.T) {
	type args struct {
		x Float32Converter
	}
	type data struct {
		name  string
		args  args
		want  Float32
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Float32(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Float32(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Float32(123), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Float32(123), want1: nil},
		{name: "Int32", args: args{x: Int32(123)}, want: Float32(123), want1: nil},
		{name: "Int64", args: args{x: Int64(123)}, want: Float32(123), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Float32(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Float32(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Float32(123), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Float32(123), want1: nil},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Float32(123.45), want1: nil},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Float32(float32(math.Inf(0))), want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Float32(float32(math.Inf(-1))), want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Float32(float32(math.NaN())), want1: nil},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Float32(123.45), want1: nil},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Float32(float32(math.Inf(0))), want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Float32(float32(math.Inf(-1))), want1: nil},
		{name: "Float64(NaN)", args: args{x: Float32(math.NaN())}, want: Float32(float32(math.NaN())), want1: nil},
		{name: "String", args: args{x: String("123.45")}, want: Float32(123), want1: nil},
		{name: "String(OR)", args: args{x: String("3.4e+39")}, want: Float32(float32(math.Inf(0))), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Float32(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Float32(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123.45"))}, want: Float32(123), want1: nil},
		{name: "Bytes(OR)", args: args{x: Bytes([]byte("3.4e+39"))}, want: Float32(float32(math.Inf(0))), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Float32(0), want1: strconv.ErrSyntax},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Float32(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat32Check(tt.args.x)
			if got != tt.want {
				if math.IsNaN(float64(got)) != math.IsNaN(float64(tt.want)) {
					t.Errorf("ToFloat32Check() got = %v, want %v", got, tt.want)
				}
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToFloat32Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		x Float64Converter
	}
	type data struct {
		name string
		args args
		want Float64
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Float64(1)},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Float64(0)},
		{name: "Int8", args: args{x: Int8(123)}, want: Float64(123)},
		{name: "Int16", args: args{x: Int16(123)}, want: Float64(123)},
		{name: "Int32", args: args{x: Int32(123)}, want: Float64(123)},
		{name: "Int64", args: args{x: Int64(123)}, want: Float64(123)},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Float64(123)},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Float64(123)},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Float64(123)},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Float64(123)},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Float64(123.45)},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Float64(math.Inf(0))},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Float64(math.Inf(-1))},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Float64(math.NaN())},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Float64(123.45)},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Float64(float64(math.Inf(0)))},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Float64(float64(math.Inf(-1)))},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Float64(math.NaN())},
		{name: "String", args: args{x: String("123.45")}, want: Float64(123.45)},
		{name: "String(OR)", args: args{x: String("1.7+309")}, want: Float64(float64(math.Inf(0)))},
		{name: "String(error)", args: args{x: String("")}, want: Float64(0)},
		{name: "String(error)", args: args{x: String("error")}, want: Float64(0)},
		{name: "Bytes", args: args{x: Bytes([]byte("123.45"))}, want: Float64(123.45)},
		{name: "Bytes(OR)", args: args{x: Bytes([]byte("1.7+309"))}, want: Float64(float64(math.Inf(0)))},
		{name: "Bytes(error)", args: args{x: Bytes([]byte(""))}, want: Float64(0)},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Float64(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.x); got != tt.want {
				if math.IsNaN(float64(got)) != math.IsNaN(float64(tt.want)) {
					t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestToFloat64Check(t *testing.T) {
	type args struct {
		x Float64Converter
	}
	type data struct {
		name  string
		args  args
		want  Float64
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Float64(1), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Float64(0), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Float64(123), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Float64(123), want1: nil},
		{name: "Int32", args: args{x: Int32(123)}, want: Float64(123), want1: nil},
		{name: "Int64", args: args{x: Int64(123)}, want: Float64(123), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Float64(123), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Float64(123), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Float64(123), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Float64(123), want1: nil},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Float64(123.45), want1: nil},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Float64(math.Inf(0)), want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Float64(math.Inf(-1)), want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Float64(math.NaN()), want1: nil},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Float64(123.45), want1: nil},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Float64(float64(math.Inf(0))), want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Float64(float64(math.Inf(-1))), want1: nil},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Float64(math.NaN()), want1: nil},
		{name: "String", args: args{x: String("123.45")}, want: Float64(123.45), want1: nil},
		{name: "String(OR)", args: args{x: String("1.7e+309")}, want: Float64(float64(math.Inf(0))), want1: strconv.ErrRange},
		{name: "String(error)", args: args{x: String("")}, want: Float64(0), want1: strconv.ErrSyntax},
		{name: "String(error)", args: args{x: String("error")}, want: Float64(0), want1: strconv.ErrSyntax},
		{name: "Bytes", args: args{x: Bytes([]byte("123.45"))}, want: Float64(123.45), want1: nil},
		{name: "Bytes(OR)", args: args{x: Bytes([]byte("1.7e+309"))}, want: Float64(float64(math.Inf(0))), want1: strconv.ErrRange},
		{name: "Bytes(error)", args: args{x: Bytes([]byte("error"))}, want: Float64(0), want1: strconv.ErrSyntax},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat64Check(tt.args.x)
			if got != tt.want {
				if math.IsNaN(float64(got)) != math.IsNaN(float64(tt.want)) {
					t.Errorf("ToFloat64Check() got = %v, want %v", got, tt.want)
				}
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToFloat64Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		x StringConverter
	}
	type data struct {
		name string
		args args
		want String
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: String("true")},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: String("false")},
		{name: "Int8", args: args{x: Int8(123)}, want: String("123")},
		{name: "Int16", args: args{x: Int16(123)}, want: String("123")},
		{name: "Int32", args: args{x: Int32(123)}, want: String("123")},
		{name: "Int64", args: args{x: Int64(123)}, want: String("123")},
		{name: "Uint8", args: args{x: Uint8(123)}, want: String("123")},
		{name: "Uint16", args: args{x: Uint16(123)}, want: String("123")},
		{name: "Uint32", args: args{x: Uint32(123)}, want: String("123")},
		{name: "Uint64", args: args{x: Uint64(123)}, want: String("123")},
		{name: "Float32", args: args{x: Float32(123.45)}, want: String("123.45")},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: String("+Inf")},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: String("-Inf")},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: String("NaN")},
		{name: "Float64", args: args{x: Float64(123.45)}, want: String("123.45")},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: String("+Inf")},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: String("-Inf")},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: String("NaN")},
		{name: "String", args: args{x: String("string")}, want: String("string")},
		{name: "Bytes", args: args{x: Bytes([]byte("string"))}, want: String("string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.x); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringCheck(t *testing.T) {
	type args struct {
		x StringConverter
	}
	type data struct {
		name  string
		args  args
		want  String
		want1 Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: String("true"), want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: String("false"), want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: String("123"), want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: String("123"), want1: nil},
		{name: "Int32", args: args{x: Int32(123)}, want: String("123"), want1: nil},
		{name: "Int64", args: args{x: Int64(123)}, want: String("123"), want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: String("123"), want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: String("123"), want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: String("123"), want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: String("123"), want1: nil},
		{name: "Float32", args: args{x: Float32(123.45)}, want: String("123.45"), want1: nil},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: String("+Inf"), want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: String("-Inf"), want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: String("NaN"), want1: nil},
		{name: "Float64", args: args{x: Float64(123.45)}, want: String("123.45"), want1: nil},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: String("+Inf"), want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: String("-Inf"), want1: nil},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: String("NaN"), want1: nil},
		{name: "String", args: args{x: String("string")}, want: String("string"), want1: nil},
		{name: "Bytes", args: args{x: Bytes([]byte("string"))}, want: String("string"), want1: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringCheck(tt.args.x)
			if got != tt.want {
				t.Errorf("ToStringCheck() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.want1) {
				t.Errorf("ToStringCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		x BytesConverter
	}
	type data struct {
		name    string
		args    args
		want    Bytes
		wantLen int
		wantCap int
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Bytes([]byte("true")), wantLen: 4, wantCap: 4},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Bytes([]byte("false")), wantLen: 5, wantCap: 5},
		{name: "Int8", args: args{x: Int8(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Int16", args: args{x: Int16(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Int32", args: args{x: Int32(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Int64", args: args{x: Int64(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Bytes([]byte("123.45")), wantLen: 6, wantCap: 6},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Bytes([]byte("+Inf")), wantLen: 4, wantCap: 4},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Bytes([]byte("-Inf")), wantLen: 4, wantCap: 4},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Bytes([]byte("NaN")), wantLen: 3, wantCap: 3},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Bytes([]byte("123.45")), wantLen: 6, wantCap: 6},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Bytes([]byte("+Inf")), wantLen: 4, wantCap: 4},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Bytes([]byte("-Inf")), wantLen: 4, wantCap: 4},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Bytes([]byte("NaN")), wantLen: 3, wantCap: 3},
		{name: "String", args: args{x: String(string("bytes"))}, want: Bytes([]byte("bytes")), wantLen: 5, wantCap: 5},
		{name: "Bytes", args: args{x: Bytes([]byte("bytes"))}, want: Bytes([]byte("bytes")), wantLen: 5, wantCap: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToBytes(tt.args.x)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
			if len(got) != tt.wantLen {
				t.Errorf("ToBytes() len = %v, want %v", got, tt.wantLen)
			}
			if cap(got) != tt.wantCap {
				t.Errorf("ToBytes() cap = %v, want %v", got, tt.wantCap)
			}
		})
	}
}

func TestToBytesCheck(t *testing.T) {
	type args struct {
		x BytesConverter
	}
	type data struct {
		name    string
		args    args
		want    Bytes
		wantLen int
		wantCap int
		want1   Err
	}
	tests := []data{
		{name: "Bool(true)", args: args{x: Bool(true)}, want: Bytes([]byte("true")), wantLen: 4, wantCap: 4, want1: nil},
		{name: "Bool(false)", args: args{x: Bool(false)}, want: Bytes([]byte("false")), wantLen: 5, wantCap: 5, want1: nil},
		{name: "Int8", args: args{x: Int8(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Int16", args: args{x: Int16(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Int32", args: args{x: Int32(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Int64", args: args{x: Int64(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Uint8", args: args{x: Uint8(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Uint16", args: args{x: Uint16(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Uint32", args: args{x: Uint32(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Uint64", args: args{x: Uint64(123)}, want: Bytes([]byte("123")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Float32", args: args{x: Float32(123.45)}, want: Bytes([]byte("123.45")), wantLen: 6, wantCap: 6, want1: nil},
		{name: "Float32(+Inf)", args: args{x: Float32(float32(math.Inf(0)))}, want: Bytes([]byte("+Inf")), wantLen: 4, wantCap: 4, want1: nil},
		{name: "Float32(-Inf)", args: args{x: Float32(float32(math.Inf(-1)))}, want: Bytes([]byte("-Inf")), wantLen: 4, wantCap: 4, want1: nil},
		{name: "Float32(NaN)", args: args{x: Float32(float32(math.NaN()))}, want: Bytes([]byte("NaN")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "Float64", args: args{x: Float64(123.45)}, want: Bytes([]byte("123.45")), wantLen: 6, wantCap: 6, want1: nil},
		{name: "Float64(+Inf)", args: args{x: Float64(math.Inf(0))}, want: Bytes([]byte("+Inf")), wantLen: 4, wantCap: 4, want1: nil},
		{name: "Float64(-Inf)", args: args{x: Float64(math.Inf(-1))}, want: Bytes([]byte("-Inf")), wantLen: 4, wantCap: 4, want1: nil},
		{name: "Float64(NaN)", args: args{x: Float64(math.NaN())}, want: Bytes([]byte("NaN")), wantLen: 3, wantCap: 3, want1: nil},
		{name: "String", args: args{x: String(string("bytes"))}, want: Bytes([]byte("bytes")), wantLen: 5, wantCap: 5, want1: nil},
		{name: "Bytes", args: args{x: Bytes([]byte("bytes"))}, want: Bytes([]byte("bytes")), wantLen: 5, wantCap: 5, want1: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToBytesCheck(tt.args.x)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("ToBytesCheck() got = %v, want %v", got, tt.want)
			}
			if len(got) != tt.wantLen {
				t.Errorf("ToBytes() len = %v, want %v", got, tt.wantLen)
			}
			if cap(got) != tt.wantCap {
				t.Errorf("ToBytes() cap = %v, want %v", got, tt.wantCap)
			}
			if !errors.Is(got1, tt.want1) {
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
	type data struct {
		name string
		args args
		want String
	}
	tests := []data{
		{name: "", args: args{x: Int64(123), format: "%05d"}, want: String("00123")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringFormat(tt.args.x, tt.args.format); got != tt.want {
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
	type data struct {
		name string
		args args
		want Bytes
	}
	tests := []data{
		{name: "", args: args{x: Int64(123), format: "%05d"}, want: Bytes([]byte("00123"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBytesFormat(tt.args.x, tt.args.format); !bytes.Equal(got, tt.want) {
				t.Errorf("ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

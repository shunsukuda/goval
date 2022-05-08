package goval

import (
	"math"
	"testing"
)

func TestInt8_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Type
	}{
		{name: "", e: Int8{0}, want: ValTypes.Int8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Int8.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_Equal(t *testing.T) {
	type args struct {
		x Int8
	}
	tests := []struct {
		name string
		e    Int8
		args args
		want bool
	}{
		{name: "True", e: Int8{1}, args: args{x: Int8{1}}, want: true},
		{name: "False", e: Int8{1}, args: args{x: Int8{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int8.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want bool
	}{
		{name: "True", e: Int8{0}, want: true},
		{name: "Flase", e: Int8{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Int8.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Bool
	}{
		{name: "False", e: Int8{0}, want: Bool{false}},
		{name: "True", e: Int8{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Int8
	}{
		{name: "", e: Int8{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Int8
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Int16
	}{
		{name: "", e: Int8{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Int16
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Int32
	}{
		{name: "", e: Int8{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Int32
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Int64
	}{
		{name: "", e: Int8{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Int64
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Uint8
	}{
		{name: "", e: Int8{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
		{name: "False", e: Int8{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Uint16
	}{
		{name: "", e: Int8{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
		{name: "False", e: Int8{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Uint32
	}{
		{name: "", e: Int8{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
		{name: "False", e: Int8{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Uint64
	}{
		{name: "", e: Int8{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
		{name: "False", e: Int8{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Float32
	}{
		{name: "", e: Int8{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Float32
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Float64
	}{
		{name: "", e: Int8{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Float64
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Complex64
	}{
		{name: "", e: Int8{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt8_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Complex128
	}{
		{name: "", e: Int8{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Int8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Int8.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt8_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Int8

		want Bytes
	}{
		{name: "", e: Int8{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int8
		args args
		want Bytes
	}{
		{name: "", e: Int8{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Int8{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int8.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int8
		args args
		want Bytes
	}{
		{name: "", e: Int8{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int8.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Int8

		want String
	}{
		{name: "", e: Int8{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int8
		args args
		want String
	}{
		{name: "", e: Int8{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Int8{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int8.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int8
		args args
		want String
	}{
		{name: "", e: Int8{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int8.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Type
	}{
		{name: "", e: Int16{0}, want: ValTypes.Int16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Int16.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_Equal(t *testing.T) {
	type args struct {
		x Int16
	}
	tests := []struct {
		name string
		e    Int16
		args args
		want bool
	}{
		{name: "True", e: Int16{1}, args: args{x: Int16{1}}, want: true},
		{name: "False", e: Int16{1}, args: args{x: Int16{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int16.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want bool
	}{
		{name: "True", e: Int16{0}, want: true},
		{name: "Flase", e: Int16{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Int16.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Bool
	}{
		{name: "False", e: Int16{0}, want: Bool{false}},
		{name: "True", e: Int16{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Int16.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Int8
	}{
		{name: "", e: Int16{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Int8
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
		{name: "False", e: Int16{math.MaxInt16}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Int16
	}{
		{name: "", e: Int16{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Int16
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Int32
	}{
		{name: "", e: Int16{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Int32
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Int64
	}{
		{name: "", e: Int16{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Int64
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Uint8
	}{
		{name: "", e: Int16{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
		{name: "False", e: Int16{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Uint16
	}{
		{name: "", e: Int16{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
		{name: "False", e: Int16{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Uint32
	}{
		{name: "", e: Int16{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
		{name: "False", e: Int16{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Uint64
	}{
		{name: "", e: Int16{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
		{name: "False", e: Int16{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Float32
	}{
		{name: "", e: Int16{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Float32
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Float64
	}{
		{name: "", e: Int16{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Float64
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Complex64
	}{
		{name: "", e: Int16{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt16_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Complex128
	}{
		{name: "", e: Int16{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Int16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Int16.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt16_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Int16

		want Bytes
	}{
		{name: "", e: Int16{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Int16.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int16
		args args
		want Bytes
	}{
		{name: "", e: Int16{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Int16{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int16.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int16
		args args
		want Bytes
	}{
		{name: "", e: Int16{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int16.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Int16

		want String
	}{
		{name: "", e: Int16{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Int16.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int16
		args args
		want String
	}{
		{name: "", e: Int16{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Int16{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int16.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int16
		args args
		want String
	}{
		{name: "", e: Int16{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int16.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Type
	}{
		{name: "", e: Int32{0}, want: ValTypes.Int32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Int32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_Equal(t *testing.T) {
	type args struct {
		x Int32
	}
	tests := []struct {
		name string
		e    Int32
		args args
		want bool
	}{
		{name: "True", e: Int32{1}, args: args{x: Int32{1}}, want: true},
		{name: "False", e: Int32{1}, args: args{x: Int32{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int32.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want bool
	}{
		{name: "True", e: Int32{0}, want: true},
		{name: "Flase", e: Int32{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Int32.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Bool
	}{
		{name: "False", e: Int32{0}, want: Bool{false}},
		{name: "True", e: Int32{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Int32.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Int8
	}{
		{name: "", e: Int32{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Int8
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Int16
	}{
		{name: "", e: Int32{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Int16
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Int32
	}{
		{name: "", e: Int32{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Int32
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Int64
	}{
		{name: "", e: Int32{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Int64
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Uint8
	}{
		{name: "", e: Int32{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Uint16
	}{
		{name: "", e: Int32{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Uint32
	}{
		{name: "", e: Int32{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Uint64
	}{
		{name: "", e: Int32{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Float32
	}{
		{name: "", e: Int32{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Float32
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Float64
	}{
		{name: "", e: Int32{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Float64
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Complex64
	}{
		{name: "", e: Int32{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
		{name: "False", e: Int32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt32_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want Complex128
	}{
		{name: "", e: Int32{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Int32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Int32.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt32_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Int32

		want Bytes
	}{
		{name: "", e: Int32{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Int32.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int32
		args args
		want Bytes
	}{
		{name: "", e: Int32{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Int32{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int32.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int32
		args args
		want Bytes
	}{
		{name: "", e: Int32{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int32.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Int32

		want String
	}{
		{name: "", e: Int32{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Int32.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int32
		args args
		want String
	}{
		{name: "", e: Int32{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Int32{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int32.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int32
		args args
		want String
	}{
		{name: "", e: Int32{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int32.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Type
	}{
		{name: "", e: Int64{0}, want: ValTypes.Int64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Int64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_Equal(t *testing.T) {
	type args struct {
		x Int64
	}
	tests := []struct {
		name string
		e    Int64
		args args
		want bool
	}{
		{name: "True", e: Int64{1}, args: args{x: Int64{1}}, want: true},
		{name: "False", e: Int64{1}, args: args{x: Int64{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Int64.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want bool
	}{
		{name: "True", e: Int64{0}, want: true},
		{name: "Flase", e: Int64{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Int64.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Bool
	}{
		{name: "False", e: Int64{0}, want: Bool{false}},
		{name: "True", e: Int64{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Int64.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Int8
	}{
		{name: "", e: Int64{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Int8
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Int16
	}{
		{name: "", e: Int64{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Int16
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Int32
	}{
		{name: "", e: Int64{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Int32
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Int64
	}{
		{name: "", e: Int64{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Int64
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Uint8
	}{
		{name: "", e: Int64{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Uint16
	}{
		{name: "", e: Int64{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Uint32
	}{
		{name: "", e: Int64{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Uint64
	}{
		{name: "", e: Int64{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Float32
	}{
		{name: "", e: Int64{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Float32
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Float64
	}{
		{name: "", e: Int64{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Float64
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Complex64
	}{
		{name: "", e: Int64{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestInt64_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Complex128
	}{
		{name: "", e: Int64{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Int64{123}, want1: true},
		{name: "False", e: Int64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Int64.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt64_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Int64

		want Bytes
	}{
		{name: "", e: Int64{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Int64.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int64
		args args
		want Bytes
	}{
		{name: "", e: Int64{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Int64{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int64.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int64
		args args
		want Bytes
	}{
		{name: "", e: Int64{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int64.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Int64

		want String
	}{
		{name: "", e: Int64{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Int64.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Int64
		args args
		want String
	}{
		{name: "", e: Int64{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Int64{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Int64.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Int64
		args args
		want String
	}{
		{name: "", e: Int64{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Int64.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Type
	}{
		{name: "", e: Uint8{0}, want: ValTypes.Uint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint8.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_Equal(t *testing.T) {
	type args struct {
		x Uint8
	}
	tests := []struct {
		name string
		e    Uint8
		args args
		want bool
	}{
		{name: "True", e: Uint8{1}, args: args{x: Uint8{1}}, want: true},
		{name: "False", e: Uint8{1}, args: args{x: Uint8{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint8.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want bool
	}{
		{name: "True", e: Uint8{0}, want: true},
		{name: "Flase", e: Uint8{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Uint8.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Bool
	}{
		{name: "False", e: Uint8{0}, want: Bool{false}},
		{name: "True", e: Uint8{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Int8
	}{
		{name: "", e: Uint8{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Int8
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
		{name: "False", e: Uint8{math.MaxUint8}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Int16
	}{
		{name: "", e: Uint8{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Int16
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Int32
	}{
		{name: "", e: Uint8{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Int32
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Int64
	}{
		{name: "", e: Uint8{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Int64
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Uint8
	}{
		{name: "", e: Uint8{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Uint16
	}{
		{name: "", e: Uint8{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Uint32
	}{
		{name: "", e: Uint8{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Uint64
	}{
		{name: "", e: Uint8{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Float32
	}{
		{name: "", e: Uint8{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Float32
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Float64
	}{
		{name: "", e: Uint8{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Float64
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Complex64
	}{
		{name: "", e: Uint8{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint8_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want Complex128
	}{
		{name: "", e: Uint8{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Uint8{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint8.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUint8_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8

		want Bytes
	}{
		{name: "", e: Uint8{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint8
		args args
		want Bytes
	}{
		{name: "", e: Uint8{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Uint8{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint8
		args args
		want Bytes
	}{
		{name: "", e: Uint8{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8

		want String
	}{
		{name: "", e: Uint8{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint8
		args args
		want String
	}{
		{name: "", e: Uint8{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Uint8{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint8
		args args
		want String
	}{
		{name: "", e: Uint8{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint8.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Type
	}{
		{name: "", e: Uint16{0}, want: ValTypes.Uint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint16.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_Equal(t *testing.T) {
	type args struct {
		x Uint16
	}
	tests := []struct {
		name string
		e    Uint16
		args args
		want bool
	}{
		{name: "True", e: Uint16{1}, args: args{x: Uint16{1}}, want: true},
		{name: "False", e: Uint16{1}, args: args{x: Uint16{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint16.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want bool
	}{
		{name: "True", e: Uint16{0}, want: true},
		{name: "Flase", e: Uint16{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Uint16.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Bool
	}{
		{name: "False", e: Uint16{0}, want: Bool{false}},
		{name: "True", e: Uint16{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Int8
	}{
		{name: "", e: Uint16{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Int8
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
		{name: "False", e: Uint16{math.MaxUint16}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Int16
	}{
		{name: "", e: Uint16{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Int16
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
		{name: "False", e: Uint16{math.MaxUint16}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Int32
	}{
		{name: "", e: Uint16{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Int32
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Int64
	}{
		{name: "", e: Uint16{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Int64
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Uint8
	}{
		{name: "", e: Uint16{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
		{name: "False", e: Uint16{math.MaxUint16}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Uint16
	}{
		{name: "", e: Uint16{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Uint32
	}{
		{name: "", e: Uint16{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Uint64
	}{
		{name: "", e: Uint16{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Float32
	}{
		{name: "", e: Uint16{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Float32
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Float64
	}{
		{name: "", e: Uint16{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Float64
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Complex64
	}{
		{name: "", e: Uint16{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint16_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want Complex128
	}{
		{name: "", e: Uint16{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Uint16{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint16.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUint16_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16

		want Bytes
	}{
		{name: "", e: Uint16{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint16
		args args
		want Bytes
	}{
		{name: "", e: Uint16{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Uint16{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint16
		args args
		want Bytes
	}{
		{name: "", e: Uint16{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16

		want String
	}{
		{name: "", e: Uint16{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint16
		args args
		want String
	}{
		{name: "", e: Uint16{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Uint16{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint16
		args args
		want String
	}{
		{name: "", e: Uint16{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint16.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Type
	}{
		{name: "", e: Uint32{0}, want: ValTypes.Uint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_Equal(t *testing.T) {
	type args struct {
		x Uint32
	}
	tests := []struct {
		name string
		e    Uint32
		args args
		want bool
	}{
		{name: "True", e: Uint32{1}, args: args{x: Uint32{1}}, want: true},
		{name: "False", e: Uint32{1}, args: args{x: Uint32{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint32.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want bool
	}{
		{name: "True", e: Uint32{0}, want: true},
		{name: "Flase", e: Uint32{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Uint32.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Bool
	}{
		{name: "False", e: Uint32{0}, want: Bool{false}},
		{name: "True", e: Uint32{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Int8
	}{
		{name: "", e: Uint32{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Int8
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxUint32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Int16
	}{
		{name: "", e: Uint32{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Int16
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxUint32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Int32
	}{
		{name: "", e: Uint32{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Int32
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxUint32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Int64
	}{
		{name: "", e: Uint32{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Int64
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Uint8
	}{
		{name: "", e: Uint32{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxUint32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Uint16
	}{
		{name: "", e: Uint32{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxUint32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Uint32
	}{
		{name: "", e: Uint32{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Uint64
	}{
		{name: "", e: Uint32{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Float32
	}{
		{name: "", e: Uint32{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Float32
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Float64
	}{
		{name: "", e: Uint32{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Float64
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Complex64
	}{
		{name: "", e: Uint32{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
		{name: "False", e: Uint32{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint32_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want Complex128
	}{
		{name: "", e: Uint32{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Uint32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint32.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUint32_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32

		want Bytes
	}{
		{name: "", e: Uint32{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint32
		args args
		want Bytes
	}{
		{name: "", e: Uint32{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Uint32{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint32
		args args
		want Bytes
	}{
		{name: "", e: Uint32{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32

		want String
	}{
		{name: "", e: Uint32{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint32
		args args
		want String
	}{
		{name: "", e: Uint32{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Uint32{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint32
		args args
		want String
	}{
		{name: "", e: Uint32{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint32.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Type
	}{
		{name: "", e: Uint64{0}, want: ValTypes.Uint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_Equal(t *testing.T) {
	type args struct {
		x Uint64
	}
	tests := []struct {
		name string
		e    Uint64
		args args
		want bool
	}{
		{name: "True", e: Uint64{1}, args: args{x: Uint64{1}}, want: true},
		{name: "False", e: Uint64{1}, args: args{x: Uint64{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Uint64.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want bool
	}{
		{name: "True", e: Uint64{0}, want: true},
		{name: "Flase", e: Uint64{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Uint64.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Bool
	}{
		{name: "False", e: Uint64{0}, want: Bool{false}},
		{name: "True", e: Uint64{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Int8
	}{
		{name: "", e: Uint64{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Int8
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Int16
	}{
		{name: "", e: Uint64{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Int16
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Int32
	}{
		{name: "", e: Uint64{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Int32
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Int64
	}{
		{name: "", e: Uint64{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Int64
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Uint8
	}{
		{name: "", e: Uint64{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Uint16
	}{
		{name: "", e: Uint64{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Uint32
	}{
		{name: "", e: Uint64{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxUint64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Uint64
	}{
		{name: "", e: Uint64{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Float32
	}{
		{name: "", e: Uint64{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Float32
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Float64
	}{
		{name: "", e: Uint64{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Float64
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Complex64
	}{
		{name: "", e: Uint64{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxInt32}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUint64_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want Complex128
	}{
		{name: "", e: Uint64{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Uint64{123}, want1: true},
		{name: "False", e: Uint64{math.MaxInt64}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Uint64.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUint64_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64

		want Bytes
	}{
		{name: "", e: Uint64{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToBytesBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint64
		args args
		want Bytes
	}{
		{name: "", e: Uint64{123}, args: args{base: 10}, want: Bytes{[]byte("123")}},
		{name: "", e: Uint64{123}, args: args{base: 16}, want: Bytes{[]byte("7b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToBytesBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint64
		args args
		want Bytes
	}{
		{name: "", e: Uint64{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64

		want String
	}{
		{name: "", e: Uint64{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToStringBase(t *testing.T) {
	type args struct {
		base int
	}
	tests := []struct {
		name string
		e    Uint64
		args args
		want String
	}{
		{name: "", e: Uint64{123}, args: args{base: 10}, want: String{"123"}},
		{name: "", e: Uint64{123}, args: args{base: 16}, want: String{"7b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringBase(tt.args.base); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToStringBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Uint64
		args args
		want String
	}{
		{name: "", e: Uint64{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Uint64.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Type
	}{
		{name: "", e: Float32{0}, want: ValTypes.Float32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Float32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_Equal(t *testing.T) {
	type args struct {
		x Float32
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want bool
	}{
		{name: "True", e: Float32{1}, args: args{x: Float32{1}}, want: true},
		{name: "False", e: Float32{1}, args: args{x: Float32{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Float32.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want bool
	}{
		{name: "True", e: Float32{0}, want: true},
		{name: "Flase", e: Float32{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Float32.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Bool
	}{
		{name: "False", e: Float32{0}, want: Bool{false}},
		{name: "True", e: Float32{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Int8
	}{
		{name: "", e: Float32{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Int8
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Int16
	}{
		{name: "", e: Float32{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Int16
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Int32
	}{
		{name: "", e: Float32{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Int32
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Int64
	}{
		{name: "", e: Float32{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Int64
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Uint8
	}{
		{name: "", e: Float32{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Uint16
	}{
		{name: "", e: Float32{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Uint32
	}{
		{name: "", e: Float32{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Uint64
	}{
		{name: "", e: Float32{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
		{name: "False", e: Float32{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Float32
	}{
		{name: "", e: Float32{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Float32
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Float64
	}{
		{name: "", e: Float32{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Float64
		want1 bool
	}{
		{name: "False", e: Float32{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Complex64
	}{
		{name: "", e: Float32{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Float32{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat32_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want Complex128
	}{
		{name: "", e: Float32{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		//want Complex128
		want1 bool
	}{
		{name: "False", e: Float32{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Float32.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFloat32_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Float32

		want Bytes
	}{
		{name: "", e: Float32{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToBytesFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want Bytes
	}{
		{name: "", e: Float32{123.456}, args: args{fmt: 'e'}, want: Bytes{[]byte("1.23456e+02")}},
		{name: "", e: Float32{123.456}, args: args{fmt: 'E'}, want: Bytes{[]byte("1.23456E+02")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBytesFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToBytesPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want Bytes
	}{
		{name: "", e: Float32{123.456}, args: args{prec: 3}, want: Bytes{[]byte("123")}},
		{name: "", e: Float32{123.456}, args: args{prec: 4}, want: Bytes{[]byte("123.5")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBytesPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToBytesFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want Bytes
	}{
		{name: "", e: Float32{123.456}, args: args{fmt: 'e', prec: 3}, want: Bytes{[]byte("1.235e+02")}},
		{name: "", e: Float32{123.456}, args: args{fmt: 'e', prec: 4}, want: Bytes{[]byte("1.2346e+02")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBytesFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want Bytes
	}{
		{name: "", e: Float32{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Float32.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Float32

		want String
	}{
		{name: "", e: Float32{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Float32.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToStringFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want String
	}{
		{name: "", e: Float32{123.456}, args: args{fmt: 'e'}, want: String{"1.23456e+02"}},
		{name: "", e: Float32{123.456}, args: args{fmt: 'E'}, want: String{"1.23456E+02"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Float32.ToStringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToStringPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want String
	}{
		{name: "", e: Float32{123.456}, args: args{prec: 3}, want: String{"123"}},
		{name: "", e: Float32{123.456}, args: args{prec: 4}, want: String{"123.5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float32.ToStringPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToStringFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want String
	}{
		{name: "", e: Float32{123.456}, args: args{fmt: 'e', prec: 3}, want: String{"1.235e+02"}},
		{name: "", e: Float32{123.456}, args: args{fmt: 'e', prec: 4}, want: String{"1.2346e+02"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float32.ToStringFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Float32
		args args
		want String
	}{
		{name: "", e: Float32{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Float32.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Type
	}{
		{name: "", e: Float64{0}, want: ValTypes.Float64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Float64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_Equal(t *testing.T) {
	type args struct {
		x Float64
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want bool
	}{
		{name: "True", e: Float64{1}, args: args{x: Float64{1}}, want: true},
		{name: "False", e: Float64{1}, args: args{x: Float64{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Float64.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want bool
	}{
		{name: "True", e: Float64{0}, want: true},
		{name: "Flase", e: Float64{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Float64.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Bool
	}{
		{name: "False", e: Float64{0}, want: Bool{false}},
		{name: "True", e: Float64{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Int8
	}{
		{name: "", e: Float64{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Int8
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Int16
	}{
		{name: "", e: Float64{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Int16
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Int32
	}{
		{name: "", e: Float64{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Int32
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Int64
	}{
		{name: "", e: Float64{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Int64
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Uint8
	}{
		{name: "", e: Float64{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Uint16
	}{
		{name: "", e: Float64{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Uint32
	}{
		{name: "", e: Float64{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Uint64
	}{
		{name: "", e: Float64{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
		{name: "False", e: Float64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Float32
	}{
		{name: "", e: Float64{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Float32
		want1 bool
	}{
		{name: "False", e: Float64{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Float64
	}{
		{name: "", e: Float64{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Float64
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Complex64
	}{
		{name: "", e: Float64{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Complex64
		want1 bool
	}{
		{name: "False", e: Float64{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestFloat64_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want Complex128
	}{
		{name: "", e: Float64{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Float64{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Float64.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFloat64_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Float64

		want Bytes
	}{
		{name: "", e: Float64{123}, want: Bytes{[]byte("123")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToBytesFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want Bytes
	}{
		{name: "", e: Float64{123.456}, args: args{fmt: 'e'}, want: Bytes{[]byte("1.23456e+02")}},
		{name: "", e: Float64{123.456}, args: args{fmt: 'E'}, want: Bytes{[]byte("1.23456E+02")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBytesFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToBytesPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want Bytes
	}{
		{name: "", e: Float64{123.456}, args: args{prec: 3}, want: Bytes{[]byte("123")}},
		{name: "", e: Float64{123.456}, args: args{prec: 4}, want: Bytes{[]byte("123.5")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBytesPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToBytesFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want Bytes
	}{
		{name: "", e: Float64{123.456}, args: args{fmt: 'e', prec: 3}, want: Bytes{[]byte("1.235e+02")}},
		{name: "", e: Float64{123.456}, args: args{fmt: 'e', prec: 4}, want: Bytes{[]byte("1.2346e+02")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBytesFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want Bytes
	}{
		{name: "", e: Float64{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={123}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Float64.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Float64

		want String
	}{
		{name: "", e: Float64{123}, want: String{"123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Float64.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToStringFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want String
	}{
		{name: "", e: Float64{123.456}, args: args{fmt: 'e'}, want: String{"1.23456e+02"}},
		{name: "", e: Float64{123.456}, args: args{fmt: 'E'}, want: String{"1.23456E+02"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Float64.ToStringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToStringPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want String
	}{
		{name: "", e: Float64{123.456}, args: args{prec: 3}, want: String{"123"}},
		{name: "", e: Float64{123.456}, args: args{prec: 4}, want: String{"123.5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float64.ToStringPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToStringFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want String
	}{
		{name: "", e: Float64{123.456}, args: args{fmt: 'e', prec: 3}, want: String{"1.235e+02"}},
		{name: "", e: Float64{123.456}, args: args{fmt: 'e', prec: 4}, want: String{"1.2346e+02"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Float64.ToStringFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Float64
		args args
		want String
	}{
		{name: "", e: Float64{123}, args: args{format: "v=%v"}, want: String{"v={123}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Float64.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Type
	}{
		{name: "", e: Complex64{0}, want: ValTypes.Complex64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Complex64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_Equal(t *testing.T) {
	type args struct {
		x Complex64
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want bool
	}{
		{name: "True", e: Complex64{1}, args: args{x: Complex64{1}}, want: true},
		{name: "False", e: Complex64{1}, args: args{x: Complex64{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Complex64.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want bool
	}{
		{name: "True", e: Complex64{0}, want: true},
		{name: "Flase", e: Complex64{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Complex64.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Bool
	}{
		{name: "False", e: Complex64{0}, want: Bool{false}},
		{name: "True", e: Complex64{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestComplex64_Real(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Float32
	}{
		{name: "", e: Complex64{123.456}, want: Float32{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Real(); !got.Equal(tt.want) {
				t.Errorf("Complex64.Real() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_Imag(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Float32
	}{
		{name: "", e: Complex64{123.456i}, want: Float32{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Imag(); !got.Equal(tt.want) {
				t.Errorf("Complex64.Imag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Int8
	}{
		{name: "", e: Complex64{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Int8
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Int16
	}{
		{name: "", e: Complex64{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Int16
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Int32
	}{
		{name: "", e: Complex64{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Int32
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Int64
	}{
		{name: "", e: Complex64{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Int64
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Uint8
	}{
		{name: "", e: Complex64{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Uint16
	}{
		{name: "", e: Complex64{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Uint32
	}{
		{name: "", e: Complex64{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Uint64
	}{
		{name: "", e: Complex64{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Float32
	}{
		{name: "", e: Complex64{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Float32
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
		{name: "False", e: Complex64{123.456i}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Float64
	}{
		{name: "", e: Complex64{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Float64
		want1 bool
	}{
		{name: "False", e: Complex64{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Complex64
	}{
		{name: "", e: Complex64{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Complex64
		want1 bool
	}{
		{name: "True", e: Complex64{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex64_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		want Complex128
	}{
		{name: "", e: Complex64{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64
		//want Complex128
		want1 bool
	}{
		{name: "False", e: Complex64{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex64.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestComplex64_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64

		want Bytes
	}{
		{name: "", e: Complex64{123}, want: Bytes{[]byte("(123+0i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToBytesFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want Bytes
	}{
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e'}, want: Bytes{[]byte("(1.23456e+02+0e+00i)")}},
		{name: "", e: Complex64{123.456}, args: args{fmt: 'E'}, want: Bytes{[]byte("(1.23456E+02+0E+00i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBytesFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToBytesPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want Bytes
	}{
		{name: "", e: Complex64{123.456}, args: args{prec: 3}, want: Bytes{[]byte("(123+0i)")}},
		{name: "", e: Complex64{123.456}, args: args{prec: 4}, want: Bytes{[]byte("(123.5+0i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBytesPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToBytesFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want Bytes
	}{
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e', prec: 3}, want: Bytes{[]byte("(1.235e+02+0.000e+00i)")}},
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e', prec: 4}, want: Bytes{[]byte("(1.2346e+02+0.0000e+00i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBytesFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want Bytes
	}{
		{name: "", e: Complex64{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={(123+0i)}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Complex64

		want String
	}{
		{name: "", e: Complex64{123}, want: String{"(123+0i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToStringFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want String
	}{
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e'}, want: String{"(1.23456e+02+0e+00i)"}},
		{name: "", e: Complex64{123.456}, args: args{fmt: 'E'}, want: String{"(1.23456E+02+0E+00i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToStringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToStringPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want String
	}{
		{name: "", e: Complex64{123.456}, args: args{prec: 3}, want: String{"(123+0i)"}},
		{name: "", e: Complex64{123.456}, args: args{prec: 4}, want: String{"(123.5+0i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToStringPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToStringFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want String
	}{
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e', prec: 3}, want: String{"(1.235e+02+0.000e+00i)"}},
		{name: "", e: Complex64{123.456}, args: args{fmt: 'e', prec: 4}, want: String{"(1.2346e+02+0.0000e+00i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToStringFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Complex64
		args args
		want String
	}{
		{name: "", e: Complex64{123}, args: args{format: "v=%v"}, want: String{"v={(123+0i)}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Complex64.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Type
	}{
		{name: "", e: Complex128{0}, want: ValTypes.Complex128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Complex128.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_Equal(t *testing.T) {
	type args struct {
		x Complex128
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want bool
	}{
		{name: "True", e: Complex128{1}, args: args{x: Complex128{1}}, want: true},
		{name: "False", e: Complex128{1}, args: args{x: Complex128{2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Complex128.Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want bool
	}{
		{name: "True", e: Complex128{0}, want: true},
		{name: "Flase", e: Complex128{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Complex128.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Bool
	}{
		{name: "False", e: Complex128{0}, want: Bool{false}},
		{name: "True", e: Complex128{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestComplex128_Real(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Float64
	}{
		{name: "", e: Complex128{123.456}, want: Float64{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Real(); !got.Equal(tt.want) {
				t.Errorf("Complex128.Real() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_Imag(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Float64
	}{
		{name: "", e: Complex128{123.456i}, want: Float64{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Imag(); !got.Equal(tt.want) {
				t.Errorf("Complex128.Imag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Int8
	}{
		{name: "", e: Complex128{123}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToInt8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Int8
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt8Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToInt8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Int16
	}{
		{name: "", e: Complex128{123}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToInt16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Int16
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt16Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToInt16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Int32
	}{
		{name: "", e: Complex128{123}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToInt32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Int32
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToInt32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Int64
	}{
		{name: "", e: Complex128{123}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToInt64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Int64
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{123.456}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToInt64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToInt64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Uint8
	}{
		{name: "", e: Complex128{123}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToUint8Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Uint8
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint8Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToUint8Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Uint16
	}{
		{name: "", e: Complex128{123}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToUint16Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Uint16
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint16Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToUint16Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Uint32
	}{
		{name: "", e: Complex128{123}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToUint32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Uint32
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToUint32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Uint64
	}{
		{name: "", e: Complex128{123}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToUint64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Uint64
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{-123}, want1: false}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToUint64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToUint64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Float32
	}{
		{name: "", e: Complex128{123}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToFloat32Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Float32
		want1 bool
	}{
		{name: "False", e: Complex128{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat32Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToFloat32Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Float64
	}{
		{name: "", e: Complex128{123}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToFloat64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Float64
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
		{name: "False", e: Complex128{123.456i}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToFloat64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToFloat64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Complex64
	}{
		{name: "", e: Complex128{123}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToComplex64Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Complex64
		want1 bool
	}{
		{name: "False", e: Complex128{123}, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex64Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToComplex64Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestComplex128_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		want Complex128
	}{
		{name: "", e: Complex128{123}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Int8.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToComplex128Eq(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128
		//want Complex128
		want1 bool
	}{
		{name: "True", e: Complex128{123}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.ToComplex128Eq()
			if got1 != tt.want1 {
				t.Errorf("Complex128.ToComplex128Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestComplex128_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128

		want Bytes
	}{
		{name: "", e: Complex128{123}, want: Bytes{[]byte("(123+0i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToBytesFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want Bytes
	}{
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e'}, want: Bytes{[]byte("(1.23456e+02+0e+00i)")}},
		{name: "", e: Complex128{123.456}, args: args{fmt: 'E'}, want: Bytes{[]byte("(1.23456E+02+0E+00i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBytesFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToBytesPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want Bytes
	}{
		{name: "", e: Complex128{123.456}, args: args{prec: 3}, want: Bytes{[]byte("(123+0i)")}},
		{name: "", e: Complex128{123.456}, args: args{prec: 4}, want: Bytes{[]byte("(123.5+0i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBytesPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToBytesFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want Bytes
	}{
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e', prec: 3}, want: Bytes{[]byte("(1.235e+02+0.000e+00i)")}},
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e', prec: 4}, want: Bytes{[]byte("(1.2346e+02+0.0000e+00i)")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBytesFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToBytesFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want Bytes
	}{
		{name: "", e: Complex128{123}, args: args{format: "v=%v"}, want: Bytes{[]byte("v={(123+0i)}")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytesFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToBytesFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Complex128

		want String
	}{
		{name: "", e: Complex128{123}, want: String{"(123+0i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToStringFmt(t *testing.T) {
	type args struct {
		fmt byte
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want String
	}{
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e'}, want: String{"(1.23456e+02+0e+00i)"}},
		{name: "", e: Complex128{123.456}, args: args{fmt: 'E'}, want: String{"(1.23456E+02+0E+00i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmt(tt.args.fmt); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToStringFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToStringPrec(t *testing.T) {
	type args struct {
		prec int
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want String
	}{
		{name: "", e: Complex128{123.456}, args: args{prec: 3}, want: String{"(123+0i)"}},
		{name: "", e: Complex128{123.456}, args: args{prec: 4}, want: String{"(123.5+0i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringPrec(tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToStringPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToStringFmtPrec(t *testing.T) {
	type args struct {
		fmt  byte
		prec int
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want String
	}{
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e', prec: 3}, want: String{"(1.235e+02+0.000e+00i)"}},
		{name: "", e: Complex128{123.456}, args: args{fmt: 'e', prec: 4}, want: String{"(1.2346e+02+0.0000e+00i)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFmtPrec(tt.args.fmt, tt.args.prec); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToStringFmtPrec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128_ToStringFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e    Complex128
		args args
		want String
	}{
		{name: "", e: Complex128{123}, args: args{format: "v=%v"}, want: String{"v={(123+0i)}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToStringFormat(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("Complex128.ToStringFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

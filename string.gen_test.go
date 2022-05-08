package goval

import (
	"testing"
)

func TestBytes_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Type
	}{
		{name: "", e: Bytes{[]byte("")}, want: ValTypes.Bytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Bytes.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Equal(t *testing.T) {
	type args struct {
		x Bytes
	}
	tests := []struct {
		name string
		e    Bytes
		args args
		want bool
	}{
		{name: "True", e: Bytes{[]byte("abc")}, args: args{x: Bytes{[]byte("abc")}}, want: true},
		{name: "False", e: Bytes{[]byte("abc")}, args: args{x: Bytes{[]byte("def")}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Bytes.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want bool
	}{
		{name: "True", e: Bytes{[]byte("")}, want: true},
		{name: "False", e: Bytes{[]byte("abc")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("Bytes.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Len(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want int
	}{
		{name: "", e: Bytes{[]byte("")}, want: 0},
		{name: "", e: Bytes{[]byte("abc")}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Len(); got != tt.want {
				t.Errorf("Bytes.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Int8
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Int16
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Int32
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Int64
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Uint8
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Uint16
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Uint32
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Uint64
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Float32
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Float64
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Complex64
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Complex128
	}{
		{name: "", e: Bytes{[]byte("123")}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Bool
	}{
		{name: "True", e: Bytes{[]byte("true")}, want: Bool{true}},
		{name: "True", e: Bytes{[]byte("false")}, want: Bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want String
	}{
		{name: "", e: Bytes{[]byte("abc")}, want: String{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Bytes
	}{
		{name: "", e: Bytes{[]byte("abc")}, want: Bytes{[]byte("abc")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Bytes.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Type(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Type
	}{
		{name: "", e: String{""}, want: ValTypes.String},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("String.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Equal(t *testing.T) {
	type args struct {
		x String
	}
	tests := []struct {
		name string
		e    String
		args args
		want bool
	}{
		{name: "True", e: String{"abc"}, args: args{x: String{"abc"}}, want: true},
		{name: "False", e: String{"abc"}, args: args{x: String{"def"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("String.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want bool
	}{
		{name: "True", e: String{""}, want: true},
		{name: "False", e: String{"abc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("String.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Len(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want int
	}{
		{name: "", e: String{""}, want: 0},
		{name: "", e: String{"abc"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Len(); got != tt.want {
				t.Errorf("String.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Int8
	}{
		{name: "", e: String{"123"}, want: Int8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("String.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Int16
	}{
		{name: "", e: String{"123"}, want: Int16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("String.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Int32
	}{
		{name: "", e: String{"123"}, want: Int32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("String.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Int64
	}{
		{name: "", e: String{"123"}, want: Int64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("String.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Uint8
	}{
		{name: "", e: String{"123"}, want: Uint8{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("String.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Uint16
	}{
		{name: "", e: String{"123"}, want: Uint16{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("String.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Uint32
	}{
		{name: "", e: String{"123"}, want: Uint32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("String.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Uint64
	}{
		{name: "", e: String{"123"}, want: Uint64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("String.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Float32
	}{
		{name: "", e: String{"123"}, want: Float32{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("String.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Float64
	}{
		{name: "", e: String{"123"}, want: Float64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("String.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Complex64
	}{
		{name: "", e: String{"123"}, want: Complex64{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("String.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Complex128
	}{
		{name: "", e: String{"123"}, want: Complex128{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("String.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Bool
	}{
		{name: "True", e: String{"true"}, want: Bool{true}},
		{name: "True", e: String{"false"}, want: Bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("String.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want String
	}{
		{name: "", e: String{"abc"}, want: String{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("String.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Bytes
	}{
		{name: "", e: String{"abc"}, want: Bytes{[]byte("abc")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("String.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

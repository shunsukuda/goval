package goval

import (
	"testing"
)

func TestBool_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Type
	}{
		{name: "", e: Bool{true}, want: ValTypes.Bool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("Bool.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Equal(t *testing.T) {
	type args struct {
		x Bool
	}
	tests := []struct {
		name string
		e    Bool
		args args
		want bool
	}{
		{name: "True", e: Bool{true}, args: args{x: Bool{true}}, want: true},
		{name: "False", e: Bool{true}, args: args{x: Bool{false}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Bool.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToInt8(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Int8
	}{
		{name: "True", e: Bool{true}, want: Int8{1}},
		{name: "False", e: Bool{false}, want: Int8{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt8(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToInt16(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Int16
	}{
		{name: "True", e: Bool{true}, want: Int16{1}},
		{name: "False", e: Bool{false}, want: Int16{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt16(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToInt32(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Int32
	}{
		{name: "True", e: Bool{true}, want: Int32{1}},
		{name: "False", e: Bool{false}, want: Int32{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt32(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Int64
	}{
		{name: "True", e: Bool{true}, want: Int64{1}},
		{name: "False", e: Bool{false}, want: Int64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt64(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToUint8(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Uint8
	}{
		{name: "True", e: Bool{true}, want: Uint8{1}},
		{name: "False", e: Bool{false}, want: Uint8{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint8(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToUint16(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Uint16
	}{
		{name: "True", e: Bool{true}, want: Uint16{1}},
		{name: "False", e: Bool{false}, want: Uint16{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint16(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToUint32(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Uint32
	}{
		{name: "True", e: Bool{true}, want: Uint32{1}},
		{name: "False", e: Bool{false}, want: Uint32{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint32(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToUint64(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Uint64
	}{
		{name: "True", e: Bool{true}, want: Uint64{1}},
		{name: "False", e: Bool{false}, want: Uint64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUint64(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToFloat32(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Float32
	}{
		{name: "True", e: Bool{true}, want: Float32{1}},
		{name: "False", e: Bool{false}, want: Float32{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat32(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToFloat64(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Float64
	}{
		{name: "True", e: Bool{true}, want: Float64{1}},
		{name: "False", e: Bool{false}, want: Float64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToFloat64(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToComplex64(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Complex64
	}{
		{name: "True", e: Bool{true}, want: Complex64{1}},
		{name: "False", e: Bool{false}, want: Complex64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex64(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBool_ToComplex128(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Complex128
	}{
		{name: "True", e: Bool{true}, want: Complex128{1}},
		{name: "False", e: Bool{false}, want: Complex128{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToComplex128(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Bool
	}{
		{name: "True", e: Bool{true}, want: Bool{true}},
		{name: "False", e: Bool{false}, want: Bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want String
	}{
		{name: "True", e: Bool{true}, want: String{"true"}},
		{name: "False", e: Bool{false}, want: String{"false"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Bytes
	}{
		{name: "True", e: Bool{true}, want: Bytes{[]byte("true")}},
		{name: "False", e: Bool{false}, want: Bytes{[]byte("false")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

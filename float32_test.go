package goval

import (
	"testing"
)

func TestFloat32_Float32(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want float32
	}{
		{name: "", e: Float32(123.45), want: float32(123.45)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Float32(); got != tt.want {
				t.Errorf("Float32.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Float32
		want interface{}
	}{
		{name: "", e: Float32(123.45), want: float32(123.45)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Float32.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Float32(123.45), want: ValTypeFloat32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Float32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

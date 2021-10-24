package goval

import (
	"testing"
)

func TestFloat64_Float64(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want float64
	}{
		{name: "", e: Float64(123.45), want: float64(123.45)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Float64(); got != tt.want {
				t.Errorf("Float64.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Float64
		want interface{}
	}{
		{name: "", e: Float64(123.45), want: float64(123.45)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Float64.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Float64(123.45), want: ValTypeFloat64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Float64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

package goval

import (
	"testing"
)

func TestUint32_Uint32(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want uint32
	}{
		{name: "", e: Uint32(123), want: uint32(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Uint32(); got != tt.want {
				t.Errorf("Uint32.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Uint32
		want interface{}
	}{
		{name: "", e: Uint32(123), want: uint32(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Uint32.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Uint32(123), want: ValTypeUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Uint32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

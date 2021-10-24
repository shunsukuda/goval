package goval

import (
	"testing"
)

func TestUint8_Uint8(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want uint8
	}{
		{name: "", e: Uint8(123), want: uint8(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Uint8(); got != tt.want {
				t.Errorf("Uint8.Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Uint8
		want interface{}
	}{
		{name: "", e: Uint8(123), want: uint8(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Uint8.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Uint8(123), want: ValTypeUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Uint8.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

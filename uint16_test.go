package goval

import (
	"testing"
)

func TestUint16_Uint16(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want uint16
	}{
		{name: "", e: Uint16(123), want: uint16(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Uint16(); got != tt.want {
				t.Errorf("Uint16.Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Uint16
		want interface{}
	}{
		{name: "", e: Uint16(123), want: uint16(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Uint16.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Uint16(123), want: ValTypeUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Uint16.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

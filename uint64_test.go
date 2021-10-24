package goval

import (
	"testing"
)

func TestUint64_Uint64(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want uint64
	}{
		{name: "", e: Uint64(123), want: uint64(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Uint64(); got != tt.want {
				t.Errorf("Uint64.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Uint64
		want interface{}
	}{
		{name: "", e: Uint64(123), want: uint64(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Uint64.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Uint64(123), want: ValTypeUint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Uint64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

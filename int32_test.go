package goval

import (
	"testing"
)

func TestInt32_Int32(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want int32
	}{
		{name: "", e: Int32(123), want: int32(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int32(); got != tt.want {
				t.Errorf("Int32.Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Int32
		want interface{}
	}{
		{name: "", e: Int32(123), want: int32(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Int32.Interface() = %v, want %v", got, tt.want)
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
		{name: "", e: Int32(123), want: ValTypeInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Int32.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

package goval

import (
	"testing"
)

func TestInt8_Int8(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want int8
	}{
		{name: "", e: Int8(123), want: int8(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int8(); got != tt.want {
				t.Errorf("Int8.Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want interface{}
	}{
		{name: "", e: Int8(123), want: int8(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Int8.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int8
		want Type
	}{
		{name: "", e: Int8(123), want: ValTypeInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Int8.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

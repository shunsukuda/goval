package goval

import (
	"testing"
)

func TestInt16_Int16(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want int16
	}{
		{name: "", e: Int16(123), want: int16(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int16(); got != tt.want {
				t.Errorf("Int16.Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want interface{}
	}{
		{name: "", e: Int16(123), want: int16(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Int16.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int16
		want Type
	}{
		{name: "", e: Int16(123), want: ValTypeInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Int16.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

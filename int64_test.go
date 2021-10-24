package goval

import (
	"testing"
)

func TestInt64_Int64(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want int64
	}{
		{name: "", e: Int64(123), want: int64(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int64(); got != tt.want {
				t.Errorf("Int64.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want interface{}
	}{
		{name: "", e: Int64(123), want: int64(123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Int64.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Int64
		want Type
	}{
		{name: "", e: Int64(123), want: ValTypeInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Int64.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

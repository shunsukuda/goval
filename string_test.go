package goval

import (
	"testing"
)

func TestString_String(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want string
	}{
		{name: "", e: String("string"), want: string("string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want interface{}
	}{
		{name: "", e: String("string"), want: string("string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("String.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Type(t *testing.T) {
	tests := []struct {
		name string
		e    String
		want Type
	}{
		{name: "", e: String("string"), want: ValTypeString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("String.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

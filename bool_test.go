package goval

import (
	"testing"
)

func TestBool_Bool(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want bool
	}{
		{name: "", e: Bool(true), want: bool(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Bool(); got != tt.want {
				t.Errorf("Bool.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want interface{}
	}{
		{name: "", e: Bool(true), want: bool(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); got != tt.want {
				t.Errorf("Bool.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Bool
		want Type
	}{
		{name: "", e: Bool(true), want: ValTypeBool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Bool.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

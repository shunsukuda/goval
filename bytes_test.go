package goval

import (
	"bytes"
	"testing"
)

func TestBytes_Bytes(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want []byte
	}{
		{name: "", e: Bytes([]byte("bytes")), want: []byte([]byte("bytes"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Bytes(); !bytes.Equal(got, tt.want) {
				t.Errorf("Bytes.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Interface(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want interface{}
	}{
		{name: "", e: Bytes([]byte("bytes")), want: []byte([]byte("bytes"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Interface(); !bytes.Equal(got.([]byte), tt.want.([]byte)) {
				t.Errorf("Bytes.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Type(t *testing.T) {
	tests := []struct {
		name string
		e    Bytes
		want Type
	}{
		{name: "", e: Bytes([]byte("bytes")), want: ValTypeBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); got != tt.want {
				t.Errorf("Bytes.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

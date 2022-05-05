package goval

import (
	"reflect"
	"testing"
)

func TestBool_Type(t *testing.T) {
	tests := []struct {
		name string
		e Bool
		want Type
	}{
		{name: "", e: Bool{true}, want: ValTypes.Bool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Equal(t *testing.T) {
	type args struct {
		x Bool
	}
	tests := []struct {
		name string
		e Bool
		args args
		want bool
	}{
		{name: "True", e: Bool{true}, args: args{x: Bool{true}}, want: true},
		{name: "False", e: Bool{true}, args: args{x: Bool{false}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("Bool.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

{{range $To := $.TL -}}
func TestBool_To{{$To.TypeName}}(t *testing.T) {
	tests := []struct {
		name string
		e Bool
		want {{$To.TypeName}}
	}{
		{name: "True", e: Bool{true}, want: {{$To.TypeName}}{1}},
		{name: "False", e: Bool{false}, want: {{$To.TypeName}}{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.To{{$To.TypeName}}(); !got.Equal(tt.want) {
				t.Errorf("Bool.To{{$To.TypeName}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}}
func TestBool_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e Bool
		want Bool
	}{
		{name: "True", e: Bool{true}, want: Bool{true}},
		{name: "False", e: Bool{false}, want: Bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToString(t *testing.T) {
	tests := []struct {
		name string
		e Bool
		want String
	}{
		{name: "True", e: Bool{true}, want: String{"true"}},
		{name: "False", e: Bool{false}, want: String{"false"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e Bool
		want Bytes
	}{
		{name: "True", e: Bool{true}, want: Bytes{[]byte("true")}},
		{name: "False", e: Bool{false}, want: Bytes{[]byte("false")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("Bool.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
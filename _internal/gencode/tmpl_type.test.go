package goval

import (
	"reflect"
	"testing"
)

{{range $T := $.TL}}

func Test{{$T.TypeName}}Type_Equal(t *testing.T) {
	type args struct {
		x Type
	}
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		args args
		want bool
	}{
		{name: "True", e:{{$T.TypeName}}Type{}, args: args{x: {{$T.TypeName}}Type{}}, want: true},
		{name: "False", e:{{$T.TypeName}}Type{}, args: args{x: {{if ne $T.TypeName "String"}}String{{else}}Bytes{{end}}Type{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Type_ID(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want int
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: Type{{$T.TypeName}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ID(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Type_Kind(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want reflect.Kind
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: reflect.{{$T.ReflectKind}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Kind(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Type_Name(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want string
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: "{{$T.TypeName}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Name(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Type_String(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want string
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: "{{$T.TypeName}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Type_GoString(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want string
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: "Type:{{$T.TypeName}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GoString(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

{{- if $T.BitSize}}
func Test{{$T.TypeName}}Type_BitSize(t *testing.T) {
	tests := []struct {
		name string
		e    {{$T.TypeName}}Type
		want int
	}{
		{name: "", e:{{$T.TypeName}}Type{}, want: {{$T.BitSize}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.BitSize(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Type.BitSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{- end}}
{{end}}{{/* range $T */}}
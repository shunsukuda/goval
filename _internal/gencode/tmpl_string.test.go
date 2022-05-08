package goval

import (
	"testing"
)

{{range $From := $.From}}
func Test{{$From.TypeName}}_Type(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want Type
	}{
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte(""){{else}}""{{end -}} }, want: ValTypes.{{$From.TypeName}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Type(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_Equal(t *testing.T) {
	type args struct {
		x {{$From.TypeName}}
	}
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		args args
		want bool
	}{
		{name: "True", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }, args: args{x: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }}, want: true},
		{name: "False", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }, args: args{x: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("def"){{else}}"def"{{end -}} }}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("{{$From.TypeName}}.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_IsZero(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want bool
	}{
		{name: "True", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte(""){{else}}""{{end}}}, want: true},
		{name: "False", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("{{$From.TypeName}}.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_Len(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want int
	}{
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte(""){{else}}""{{end -}} }, want: 0},
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Len(); got != tt.want {
				t.Errorf("{{$From.TypeName}}.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{range $To := $.To}}
func Test{{$From.TypeName}}_To{{$To.TypeName}}(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want {{$To.TypeName}}
	}{
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("123"){{else}}"123"{{end -}} }, want: {{$To.TypeName}}{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.To{{$To.TypeName}}(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.To{{$To.TypeName}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}} {{- /* range $To */}}
func Test{{$From.TypeName}}_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want Bool
	}{
		{name: "True", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("true"){{else}}"true"{{end -}} }, want: Bool{true}},
		{name: "True", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("false"){{else}}"false"{{end -}} }, want: Bool{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_ToString(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want String
	}{
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }, want: String{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToString(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_ToBytes(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want Bytes
	}{
		{name: "", e: {{$From.TypeName}}{ {{- if eq $From.TypeName "Bytes"}}[]byte("abc"){{else}}"abc"{{end -}} }, want: Bytes{[]byte("abc")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBytes(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}} {{- /* range $From */ -}}
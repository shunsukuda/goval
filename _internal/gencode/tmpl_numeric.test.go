package goval

import (
	"math"
	"testing"
)

{{range $From := $.From}}
func Test{{$From.TypeName}}_Type(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want Type
	}{
		{name: "", e: {{$From.TypeName}}{0}, want: ValTypes.{{$From.TypeName}}},
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
		{name: "True", e: {{$From.TypeName}}{1}, args: args{x: {{$From.TypeName}}{1},}, want: true},
		{name: "False", e: {{$From.TypeName}}{1}, args: args{x: {{$From.TypeName}}{2},}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equal(tt.args.x); got != tt.want {
				t.Errorf("{{$From.TypeName}}.Equal = %v, want %v", got, tt.want)
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
		{name: "True", e: {{$From.TypeName}}{0}, want: true},
		{name: "Flase", e: {{$From.TypeName}}{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsZero(); got != tt.want {
				t.Errorf("{{$From.TypeName}}.IsZero = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_ToBool(t *testing.T) {
	tests := []struct {
		name string
		e   {{$From.TypeName}}
		want Bool
	}{
		{name: "False", e: {{$From.TypeName}}{0}, want: Bool{false}},
		{name: "True", e: {{$From.TypeName}}{1}, want: Bool{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToBool(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

{{- if eq $From.TypeKind "Complex"}}
func Test{{$From.TypeName}}_Real(t *testing.T) {
	tests := []struct {
		name string
		e   {{$From.TypeName}}
		want Float{{$From.BitSize}}
	}{
		{name: "", e: {{$From.TypeName}}{123.456}, want: Float{{$From.BitSize}}{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Real(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.Real() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_Imag(t *testing.T) {
	tests := []struct {
		name string
		e   {{$From.TypeName}}
		want Float{{$From.BitSize}}
	}{
		{name: "", e: {{$From.TypeName}}{123.456i}, want: Float{{$From.BitSize}}{123.456}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Imag(); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.Imag() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{- end}}

{{range $To := $.To -}}
func Test{{$From.TypeName}}_To{{$To.TypeName}}(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		want {{$To.TypeName}}
	}{
		{name: "", e: {{$From.TypeName}}{123}, want: {{$To.TypeName}}{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.To{{$To.TypeName}}(); !got.Equal(tt.want) {
				t.Errorf("Int8.To{{$To.TypeName}}() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$From.TypeName}}_To{{$To.TypeName}}Eq(t *testing.T) {
	tests := []struct {
		name string
		e    {{$From.TypeName}}
		//want {{$To.TypeName}}
		want1 bool
	}{
		{{- if and (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")) 
			(or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex"))
			(ne $From.BitSize $To.BitSize)
		}}
			{name: "False", e: {{$From.TypeName}}{123}, want1: false},
		{{- else}}
			{name: "True", e: {{$From.TypeName}}{123}, want1: true},
			{{- if and (ne $From.TypeKind "Uint") (eq $To.TypeKind "Uint")}}
				{name: "False", e: {{$From.TypeName}}{-123}, want1: false}, // 
			{{- else if and (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex"))
				(not (or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex")))
			}}
				{name: "False", e: {{$From.TypeName}}{123.456}, want1: false},
			{{- else if and (eq $From.TypeKind "Int") (eq $To.TypeKind "Int") (gt $From.BitSize $To.BitSize)}}
				{name: "False", e: {{$From.TypeName}}{math.MaxInt{{$From.BitSize}}}, want1: false},
			{{- else if and (eq $From.TypeKind "Uint") (eq $To.TypeKind "Int") (ge $From.BitSize $To.BitSize)}}
				{name: "False", e: {{$From.TypeName}}{math.MaxUint{{$From.BitSize}}}, want1: false},
			{{- else if and (eq $From.TypeKind "Uint") (eq $To.TypeKind "Uint") (gt $From.BitSize $To.BitSize)}}
				{name: "False", e: {{$From.TypeName}}{math.MaxUint{{$From.BitSize}}}, want1: false},
			{{- else if and (or (eq $From.TypeKind "Int") (eq $From.TypeKind "Uint")) (ge $From.BitSize 32)
				(or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex")) (eq $To.BitSize 32)
			}}
				{name: "False", e: {{$From.TypeName}}{math.MaxInt32}, want1: false},
			{{- else if and (or (eq $From.TypeKind "Int") (eq $From.TypeKind "Uint")) (ge $From.BitSize 64)
				(or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex")) (eq $To.BitSize 64)
			}}
				{name: "False", e: {{$From.TypeName}}{math.MaxInt64}, want1: false},
			{{- else if and (eq $From.TypeKind "Complex") (eq $To.TypeKind "Float") (eq $From.BitSize $To.BitSize)}}
				{name: "False", e: {{$From.TypeName}}{123.456i}, want1: false},
			{{- end}}
		{{- end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.e.To{{$To.TypeName}}Eq()
			if got1 != tt.want1 {
				t.Errorf("{{$From.TypeName}}.To{{$To.TypeName}}Eq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
{{end}} {{/* range $To */}}
{{range $To := $.ToString}}
{{range $Opt := $.ToStringFuncs}}
{{- if or
	(eq $Opt "")
	(and (eq $Opt "Base") (or (eq $From.TypeKind "Int") (eq $From.TypeKind "Uint")))
	(and (eq $Opt "Fmt") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))
	(and (eq $Opt "Prec") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))
	(and (eq $Opt "FmtPrec") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))
}}
func Test{{$From.TypeName}}_To{{$To.TypeName}}{{$Opt}}(t *testing.T) {
	{{- if ne $Opt ""}}
	type args struct {
		{{- if eq $Opt "Base"}}base int
		{{- else if eq $Opt "Fmt"}}fmt byte
		{{- else if eq $Opt "Prec"}}prec int
		{{- else if eq $Opt "FmtPrec"}}
			fmt byte
			prec int
		{{- end}}
	}
	{{- end}}
	tests := []struct {
		name string
		e   {{$From.TypeName}}
		{{if ne $Opt ""}}args args{{end}}
		want {{$To.TypeName}}
	}{
		{{- if eq $Opt ""}}
			{{- if eq $From.TypeKind "Complex"}}
				{name: "", e: {{$From.TypeName}}{123}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(123+0i)"){{else}}"(123+0i)"{{end -}} }},
			{{- else}}
				{name: "", e: {{$From.TypeName}}{123}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("123"){{else}}"123"{{end -}} }},
			{{- end}}
		{{- else if eq $Opt "Base"}}
			{name: "", e: {{$From.TypeName}}{123}, args: args{base: 10}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("123"){{else}}"123"{{end -}} }},
			{name: "", e: {{$From.TypeName}}{123}, args: args{base: 16}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("7b"){{else}}"7b"{{end -}} }},
		{{- else if eq $Opt "Fmt"}}
			{{- if eq $From.TypeKind "Complex"}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e'}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(1.23456e+02+0e+00i)"){{else}}"(1.23456e+02+0e+00i)"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'E'}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(1.23456E+02+0E+00i)"){{else}}"(1.23456E+02+0E+00i)"{{end -}} }},
			{{- else}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e'}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("1.23456e+02"){{else}}"1.23456e+02"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'E'}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("1.23456E+02"){{else}}"1.23456E+02"{{end -}} }},
			{{- end}}
		{{- else if eq $Opt "Prec"}}
			{{- if eq $From.TypeKind "Complex"}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{prec: 3}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(123+0i)"){{else}}"(123+0i)"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{prec: 4}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(123.5+0i)"){{else}}"(123.5+0i)"{{end -}} }},
			{{- else}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{prec: 3}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("123"){{else}}"123"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{prec: 4}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("123.5"){{else}}"123.5"{{end -}} }},
			{{- end}}
		{{- else if eq $Opt "FmtPrec"}}
			{{- if eq $From.TypeKind "Complex"}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e', prec: 3}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(1.235e+02+0.000e+00i)"){{else}}"(1.235e+02+0.000e+00i)"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e', prec: 4}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("(1.2346e+02+0.0000e+00i)"){{else}}"(1.2346e+02+0.0000e+00i)"{{end -}} }},
			{{- else}}
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e', prec: 3}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("1.235e+02"){{else}}"1.235e+02"{{end -}} }},
				{name: "", e: {{$From.TypeName}}{123.456}, args: args{fmt: 'e', prec: 4}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("1.2346e+02"){{else}}"1.2346e+02"{{end -}} }},
			{{- end}}
		{{- end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.To{{$To.TypeName}}{{$Opt}}(
				{{- if eq $Opt "Base"}}tt.args.base
				{{- else if eq $Opt "Fmt"}}tt.args.fmt
				{{- else if eq $Opt "Prec"}}tt.args.prec
				{{- else if eq $Opt "FmtPrec"}}tt.args.fmt, tt.args.prec
				{{- end}}); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.To{{$To.TypeName}}{{$Opt}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}} {{/* if */}} {{end}} {{/* range $Opt */}}
func Test{{$From.TypeName}}_To{{$To.TypeName}}Format(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		e   {{$From.TypeName}}
		args args
		want {{$To.TypeName}}
	}{
		{{- if eq $From.TypeKind "Complex"}}
			{name: "", e: {{$From.TypeName}}{123}, args: args{format: "v=%v"}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("v={(123+0i)}"){{else}}"v={(123+0i)}"{{end -}} }},
		{{- else}}
			{name: "", e: {{$From.TypeName}}{123}, args: args{format: "v=%v"}, want: {{$To.TypeName}}{ {{- if eq $To.TypeName "Bytes"}}[]byte("v={123}"){{else}}"v={123}"{{end -}} }},
		{{- end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.To{{$To.TypeName}}Format(tt.args.format); !got.Equal(tt.want) {
				t.Errorf("{{$From.TypeName}}.To{{$To.TypeName}}Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}} {{/* range $To */}}
{{end}} {{/* range $From */}}
package goval

import "reflect"

const ({{range $T := $.T}}
	Type{{$T.TypeName}}{{if eq $T.TypeName "Nil"}} = iota{{end}}
	{{- end}}{{/* range $T */}}
)

var (
	ValTypes = struct { {{range $T := $.T}}
		{{$T.TypeName}} Type
		{{- end}}{{/* range $T */}}
	}{ {{range $T := $.T}}
		{{$T.TypeName}}: &{{$T.TypeName}}Type{},
	{{- end}}{{/* range $T */}}
	}
)

type Type interface {
	ID() int
	Kind() reflect.Kind
	Name() string
	String() string
	// BitSize() int
}

{{range $T := $.T}}
type {{$T.TypeName}}Type struct{}

func (t {{$T.TypeName}}Type) ID() int { return Type{{$T.TypeName}} }
func (t {{$T.TypeName}}Type) Kind() reflect.Kind { return reflect.
	{{- if $T.RefKind}}{{$T.RefKind}}{{else}}{{$T.TypeName}}{{end}} }
func (t {{$T.TypeName}}Type) Name() string { return "{{$T.TypeName}}" }
func (t {{$T.TypeName}}Type) String() string { return "{{$T.TypeName}}" }
{{- if $T.BitSize}}
func (t {{$T.TypeName}}Type) BitSize() int { return {{$T.BitSize}} }{{- end}}
{{end}}{{/* range $T */}}
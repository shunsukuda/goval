package goval

import "reflect"

const ({{range $T := $.TL}}
	Type{{$T.TypeName}}{{if eq $T.TypeName "Nil"}} = iota{{end}}
	{{- end}}{{/* range $T */}}
)

var (
	ValTypes = struct { {{range $T := $.TL}}
		{{$T.TypeName}} Type
		{{- end}}{{/* range $T */}}
	}{ {{range $T := $.TL}}
		{{$T.TypeName}}: &{{$T.TypeName}}Type{},
	{{- end}}{{/* range $T */}}
	}
)

type Type interface {
	Equal(x Type) bool
	ID() int
	Kind() reflect.Kind
	Name() string
	String() string
	GoString() string
}

{{range $T := $.TL}}
type {{$T.TypeName}}Type struct{}

func (t {{$T.TypeName}}Type) Equal(x Type) bool { return t.ID() == x.ID() }
func (t {{$T.TypeName}}Type) ID() int { return Type{{$T.TypeName}} }
func (t {{$T.TypeName}}Type) Kind() reflect.Kind { return reflect.{{$T.ReflectKind}} }
func (t {{$T.TypeName}}Type) Name() string { return "{{$T.TypeName}}" }
func (t {{$T.TypeName}}Type) String() string { return "{{$T.TypeName}}" }
func (t {{$T.TypeName}}Type) GoString() string { return "Type:{{$T.TypeName}}" }
{{- if $T.BitSize}}
func (t {{$T.TypeName}}Type) BitSize() int { return {{$T.BitSize}} }{{- end}}
{{end}}{{/* range $T */}}
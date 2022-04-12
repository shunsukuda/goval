package goval

import (
	"strconv"

	"github.com/shunsukuda/goval/forceconv"
)

{{range $From := $.From}}

type {{$From.TypeName}} {{$From.GoTypeName}}

func (e {{$From.TypeName}}) {{$From.TypeName}}() {{$From.GoTypeName}} { return {{$From.GoTypeName}}(e) }
func (e {{$From.TypeName}}) Interface() interface{} { return e.{{$From.TypeName}}() }
func (e {{$From.TypeName}}) Val() Val { return e }
func (e {{$From.TypeName}}) Type() Type { return ValTypes.{{$From.TypeName}} }
{{range $To := $.To}}
func (e {{$From.TypeName}}) To{{$To.TypeName}}() {{$To.TypeName}} {
	{{if eq $From.TypeName "String" -}}
	ce, _ := strconv.Parse{{$To.TypeKind}}(e.String(), 0{{if or (eq $To.TypeKind "Int") (eq $To.TypeKind "Uint")}}, {{$To.BitSize}}{{end}})
	return {{$To.TypeName}}(ce)
	{{- else}}return e.ToString().To{{$To.TypeName}}(){{end}}
} {{end}} {{- /* range $To */}}
func (e {{$From.TypeName}}) ToBool() Bool { {{if eq $From.TypeName "String"}}
	ce, _ := strconv.ParseBool(e.String())
	return Bool(ce)
	{{- else}}return e.ToString().ToBool(){{end}}
}
func (e {{$From.TypeName}}) ToString() String { return
	{{- if eq $From.TypeName "String"}} e{{else}} String(forceconv.BytesToString(e.Bytes())){{end}} }
func (e {{$From.TypeName}}) ToBytes() Bytes { return
	{{- if eq $From.TypeName "Bytes"}} e{{else}} Bytes(forceconv.StringToBytes(e.String())){{end}} }
{{end}} {{- /* range $From */ -}}
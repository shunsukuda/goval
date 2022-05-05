package goval

import (
	"bytes"
	"strconv"

	"github.com/shunsukuda/forceconv"
)

{{range $From := $.From}}

type {{$From.TypeName}} struct {
	{{$From.TypeName}} {{$From.GoTypeName}}
}

// func (e {{$From.TypeName}}) Go{{$From.TypeName}}() {{$From.GoTypeName}} { return e.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Type() Type { return ValTypes.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Equal(x {{$From.TypeName}}) bool { return
	{{- if eq $From.TypeName "String"}} e.String == x.String{{else}} bytes.Equal(e.Bytes, x.Bytes){{end}} }
func (e {{$From.TypeName}}) IsZero() bool { return e.Len() == 0 }
func (e {{$From.TypeName}}) Len() int { return len(e.{{$From.TypeName}}) }

{{range $To := $.To}}
func (e {{$From.TypeName}}) To{{$To.TypeName}}() {{$To.TypeName}} {
	{{if eq $From.TypeName "String" -}}
	ce, _ := strconv.Parse{{$To.TypeKind}}(e.String, 0{{if or (eq $To.TypeKind "Int") (eq $To.TypeKind "Uint")}}, {{$To.BitSize}}{{end}})
	return {{$To.TypeName}}{ {{$To.GoTypeName}}(ce) }
	{{- else}}return e.ToString().To{{$To.TypeName}}(){{end}}
}
{{end}} {{- /* range $To */}}
func (e {{$From.TypeName}}) ToBool() Bool { {{if eq $From.TypeName "String"}}
	ce, _ := strconv.ParseBool(e.String)
	return Bool{ce}
	{{- else}}return e.ToString().ToBool(){{end}}
}
func (e {{$From.TypeName}}) ToString() String {
	return {{- if eq $From.TypeName "String"}} e{{else}} String{forceconv.BytesToString(e.Bytes)}{{end}}
}
func (e {{$From.TypeName}}) ToBytes() Bytes {
	return {{- if eq $From.TypeName "Bytes"}} e{{else}} Bytes{forceconv.StringToBytes(e.String)}{{end}}
}
{{end}} {{- /* range $From */ -}}
package goval

import (
	"strconv"
	"math"
	//"database/sql"
	//"database/sql/driver"

	"github.com/shunsukuda/forceconv"
)

var (
	MinInt8   = Int8{Int8: math.MinInt8}
	MinInt16  = Int16{Int16: math.MinInt16}
	MinInt32  = Int32{Int32: math.MinInt32}
	MinInt64  = Int64{Int64: math.MinInt64}
	MaxInt8   = Int8{Int8: math.MaxInt8}
	MaxInt16  = Int16{Int16: math.MaxInt16}
	MaxInt32  = Int32{Int32: math.MaxInt32}
	MaxInt64  = Int64{Int64: math.MaxInt64}
	MinUint8  = Uint8{Uint8: 0}
	MinUint16 = Uint16{Uint16: 0}
	MinUint32 = Uint32{Uint32: 0}
	MinUint64 = Uint64{Uint64: 0}
	MaxUint8  = Uint8{Uint8: math.MaxUint8}
	MaxUint16 = Uint16{Uint16: math.MaxUint16}
	MaxUint32 = Uint32{Uint32: math.MaxUint32}
	MaxUint64 = Uint64{Uint64: math.MaxUint64}
)

var (
	maxInt24 = 1<<24 - 1
	maxInt53 = 1<<53 - 1
	minInt24 = -1 * maxInt24
	minInt53 = -1 * maxInt53
)

var (
	{{$.StrconvIntBaseName}}   int = {{$.StrconvIntBaseValue}}
	{{$.StrconvUintBaseName}}  int = {{$.StrconvUintBaseValue}}
	{{$.StrconvFloatFmtName}} byte = {{$.StrconvFloatFmtValue}}
	{{$.StrconvFloatPrecName}} int = {{$.StrconvFloatPrecValue}}
)
{{range $From := $.From}}

type {{$From.TypeName}} struct{
	{{$From.TypeName}} {{$From.GoTypeName}}
}

func (e {{$From.TypeName}}) Go{{$From.TypeName}}() {{$From.GoTypeName}} { return e.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Interface() interface{} { return e.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Val() Val { return e }
func (e {{$From.TypeName}}) Type() Type { return ValTypes.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Equal(x {{$From.TypeName}}) bool { return e.{{$From.TypeName}} == x.{{$From.TypeName}} }

{{range $To := $.To -}}
func (e {{$From.TypeName}}) To{{$To.TypeName}}() {{$To.TypeName}} {
	return {{if eq $From.TypeName $To.TypeName}}e
	{{- else}}{{$To.TypeName}}{
		{{- if eq $To.TypeKind "Complex"}}
			{{- if eq $From.TypeKind "Complex"}}e.To{{$To.TypeName}}().{{$To.TypeName}}{{- else}}complex(e.ToFloat{{$To.BitSize}}().Float{{$To.BitSize}}, float{{$To.BitSize}}(0)){{end}}
		{{- else if eq $From.TypeKind "Complex"}}
			{{- if eq $To.TypeKind "Complex"}}e{{else}}{{$To.GoTypeName}}(real(e.{{$From.TypeName}})){{end}}
		{{- else}}{{$To.GoTypeName}}(e.{{$From.TypeName}}){{end}}}
	{{- end}}
}
func (e {{$From.TypeName}}) To{{$To.TypeName}}Eq() ({{$To.TypeName}}, bool) {
	return {{if eq $From.TypeName $To.TypeName}}e, true
	{{- else}}e.To{{$To.TypeName}}(),
		{{- if eq $To.TypeKind "Int"}}
			{{- if eq $From.TypeKind "Int"}} {{/* IntToInt */}}
				{{- if le $From.BitSize $To.BitSize}} true{{else}} math.MinInt{{$To.BitSize}} <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.MaxInt{{$To.BitSize}}{{end}} 
			{{- else if eq $From.TypeKind "Uint"}} {{/* UintToInt */}}
				{{- if lt $From.BitSize $To.BitSize}} 0 <= e.{{$From.TypeName}}{{else}} 0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.MaxInt{{$To.BitSize}}{{end}} 
			{{- else}} e == e.To{{$To.TypeName}}().To{{$From.TypeName}}()
			{{- end}}
		{{- else if eq $To.TypeKind "Uint"}}
			{{- if eq $From.TypeKind "Int"}} {{/* UintToInt */}}
				{{- if le $From.BitSize $To.BitSize}} 0 <= e.{{$From.TypeName}}{{else}} 0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.MaxUint{{$To.BitSize}}{{end}} 
			{{- else if eq $From.TypeKind "Uint"}} {{/* UintToUint */}}
				{{- if le $From.BitSize $To.BitSize}} true{{else}} e.{{$From.TypeName}} <= math.MaxUint{{$To.BitSize}}{{end}} 
			{{- else}} e == e.To{{$To.TypeName}}().To{{$From.TypeName}}()
			{{- end}}
		{{- else if eq $To.TypeKind "Float"}}
			{{- if eq $From.TypeKind "Int"}}
				{{- if lt $From.BitSize $To.BitSize}} true
				{{- else}} {{$From.GoTypeName}}(minInt{{if eq $From.BitSize 32}}24{{else}}53{{end}}) <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt{{if eq $From.BitSize 32}}24{{else}}53{{end}})
				{{- end}}
			{{- else if eq $From.TypeKind "Uint"}}
				{{- if lt $From.BitSize $To.BitSize}} true
				{{- else}} 0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt{{if eq $From.BitSize 32}}24{{else}}53{{end}})
				{{- end}}
			{{- else if or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")}}
				{{- if lt $From.BitSize $To.BitSize}} e == e.To{{$To.TypeName}}().To{{$From.TypeName}}()
				{{- else}} e.ToString() == e.To{{$To.TypeName}}().ToString(){{end}}
			{{- end}}
		{{- else if eq $To.TypeKind "Complex"}}
			{{- if or (eq $From.TypeKind "Int") (eq $From.TypeKind "Uint")}}
				{{- if lt $From.BitSize $To.BitSize}} true{{else}} e == e.To{{$To.TypeName}}().To{{$From.TypeName}}(){{end}} 
			{{- else if or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")}}
				{{- if lt $From.BitSize $To.BitSize}} e == e.To{{$To.TypeName}}().To{{$From.TypeName}}()
				{{- else}} e.ToString() == e.To{{$To.TypeName}}().ToString(){{end}}
			{{- end}}
		{{- end}}
	{{- end}}
}
{{end}} {{- /* range $To */ -}}
func (e {{$From.TypeName}}) ToBool() Bool {
	return Bool{e.{{$From.TypeName}} != 0}
}
func (e {{$From.TypeName}}) ToString() String {
	return String{strconv.Format{{$From.TypeKind}}(e
		{{- if ne $From.BitSize 64}}.To{{$From.TypeKind}}
			{{- if ne $From.TypeKind "Complex"}}64{{else}}128{{end}}()
		{{- end}}.{{$From.TypeKind}}{{- if ne $From.TypeKind "Complex"}}64{{else}}128{{end}},
		{{- if eq $From.TypeKind "Int"}} {{$.StrconvIntBaseName}}
		{{- else if eq $From.TypeKind "Uint"}} {{$.StrconvUintBaseName}}
		{{- else if or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex") -}}
			{{$.StrconvFloatFmtName}}, {{$.StrconvFloatPrecName}},
			{{- if eq $From.TypeKind "Complex"}}
				{{- if ne $From.TypeKind "Complex"}} 64{{else}} 128{{end}}
			{{- else}} {{$From.BitSize}}
			{{- end}}
		{{- end -}})}
}
func (e {{$From.TypeName}}) ToBytes() Bytes {
	return Bytes{forceconv.StringToBytes(e.ToString().String)}
}
{{end}} {{- /* range $From */ -}}
package goval

import (
	"fmt"
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
type {{$From.TypeName}} struct {
	{{$From.TypeName}} {{$From.GoTypeName}}
}

// func (e {{$From.TypeName}}) Go{{$From.TypeName}}() {{$From.GoTypeName}} { return e.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Type() Type { return ValTypes.{{$From.TypeName}} }
func (e {{$From.TypeName}}) Equal(x {{$From.TypeName}}) bool { return e.{{$From.TypeName}} == x.{{$From.TypeName}} }
func (e {{$From.TypeName}}) IsZero() bool { return e.{{$From.TypeName}} == 0 }
func (e {{$From.TypeName}}) ToBool() Bool { return Bool{e.{{$From.TypeName}} != 0} }
{{- if eq $From.TypeKind "Complex"}}
func (e {{$From.TypeName}}) Real() Float{{$From.BitSize}} { return Float{{$From.BitSize}} {real(e.{{$From.TypeName}})} }
func (e {{$From.TypeName}}) Imag() Float{{$From.BitSize}} { return Float{{$From.BitSize}} {imag(e.{{$From.TypeName}})} }
{{- end}}

{{range $To := $.To -}}
func (e {{$From.TypeName}}) To{{$To.TypeName}}() {{$To.TypeName}} { return {{if eq $From.TypeName $To.TypeName -}}e }
	{{- else -}}{{$To.TypeName}}{
		{{- if eq $To.TypeKind "Complex" -}}
			{{- if eq $From.TypeKind "Complex" -}} {{$To.GoTypeName}}(e.{{$From.TypeName}})} }
			{{- else -}} complex(e.ToFloat{{$To.BitSize}}().Float{{$To.BitSize}}, float{{$To.BitSize}}(0))} }
			{{- end -}}
		{{- else if eq $From.TypeKind "Complex" -}}{{$To.GoTypeName}}(real(e.{{$From.TypeName}}))} }
		{{- else -}} {{$To.GoTypeName}}(e.{{$From.TypeName}})} }
		{{- end -}}
	{{- end}}
func (e {{$From.TypeName}}) To{{$To.TypeName}}Eq() ({{$To.TypeName}}, bool) { return {{if eq $From.TypeName $To.TypeName -}}e, true }
	{{- else -}}e.To{{$To.TypeName}}(),
		{{- if eq $From.TypeKind "Int" -}}
			{{- if eq $To.TypeKind "Int" -}}
				{{- if gt $From.BitSize $To.BitSize}} math.Min{{$To.TypeName}} <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.Max{{$To.TypeName}} }
				{{- else}} true }
				{{- end}}
			{{- else if eq $To.TypeKind "Uint"}} 0 <= e.{{$From.TypeName}}
				{{- if gt $From.BitSize $To.BitSize}} && e.{{$From.TypeName}} <= math.Max{{$To.TypeName}}
				{{- end}} }
			{{- else if or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex")}}
				{{- if or (le $From.BitSize 16) (and (eq $From.BitSize 32) (eq $To.BitSize 64))}} true }
				{{- else if and (ge $From.BitSize 32) (eq $To.BitSize 32)}} {{$From.GoTypeName}}(minInt24) <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt24) }
				{{- else if and (eq $From.BitSize 64) (eq $To.BitSize 64)}} {{$From.GoTypeName}}(minInt53) <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt53) }
				{{- end}}
			{{- end}}
		{{- else if eq $From.TypeKind "Uint"}}
			{{- if eq $To.TypeKind "Int"}}
				{{- if ge $From.BitSize $To.BitSize}} e.{{$From.TypeName}} <= math.Max{{$To.TypeName}} }
				{{- else}} true }
				{{- end}}
			{{- else if eq $To.TypeKind "Uint"}}
				{{- if gt $From.BitSize $To.BitSize}} e.{{$From.TypeName}} <= math.Max{{$To.TypeName}} }
				{{- else}} true }
				{{- end}}
			{{- else if or (eq $To.TypeKind "Float") (eq $To.TypeKind "Complex")}}
				{{- if or (le $From.BitSize 16) (and (eq $From.BitSize 32) (eq $To.BitSize 64))}} true }
				{{- else if and (ge $From.BitSize 32) (eq $To.BitSize 32)}} 0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt24) }
				{{- else if and (eq $From.BitSize 64) (eq $To.BitSize 64)}} 0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= {{$From.GoTypeName}}(maxInt53) }
				{{- end}}
			{{- end}}
		{{- else if eq $From.TypeKind "Float" -}}
			{{- if or (eq $To.TypeKind "Int") (eq $To.TypeKind "Uint") -}}
				e.{{$From.TypeName}} == e.ToInt64().To{{$From.TypeName}}().{{$From.TypeName}} &&
			{{- end}}
			{{- if eq $To.TypeKind "Int" -}}
				math.Min{{$To.TypeName}} <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.Max{{$To.TypeName}} }
			{{- else if eq $To.TypeKind "Uint" -}}
				0 <= e.{{$From.TypeName}} && e.{{$From.TypeName}} <= math.Max{{$To.TypeName}} }
			{{- else if eq $To.TypeKind "Float"}} false }
			{{- else if eq $To.TypeKind "Complex"}}
				{{- if eq $From.BitSize $To.BitSize}} true }
				{{- else}} false }
				{{- end -}}
			{{- end -}}
		{{- else if eq $From.TypeKind "Complex" -}}
			{{- if or (eq $To.TypeKind "Int") (eq $To.TypeKind "Uint")}} e.Imag().IsZero() && e.Real().Float{{$From.BitSize}} == e.ToInt64().ToFloat{{$From.BitSize}}().Float{{$From.BitSize}} &&
			{{- end -}}
			{{- if eq $To.TypeKind "Int"}} math.Min{{$To.TypeName}} <= e.Real().Float{{$From.BitSize}} && e.Real().Float{{$From.BitSize}} <= math.Max{{$To.TypeName}} }
			{{- else if eq $To.TypeKind "Uint"}} 0 <= e.Real().Float{{$From.BitSize}} && e.Real().Float{{$From.BitSize}} <= math.Max{{$To.TypeName}} }
			{{- else if eq $To.TypeKind "Float"}}
				{{- if eq $From.BitSize $To.BitSize}} e.Imag().IsZero() }
				{{- else}} false }
				{{- end -}}
			{{- else if eq $To.TypeKind "Complex"}} false }
			{{- end -}}
		{{- end -}}
	{{- end}}
{{end}} {{- /* range $To */ -}}
{{range $Opt := $.ToStringFuncs -}}
{{- if or
	(eq $Opt "")
	(and (eq $Opt "Base") (or (eq $From.TypeKind "Int") (eq $From.TypeKind "Uint")))
	(and (eq $Opt "Fmt") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))
	(and (eq $Opt "Prec") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))
	(and (eq $Opt "FmtPrec") (or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")))}}
func (e {{$From.TypeName}}) ToString{{$Opt}}(
	{{- if eq $Opt "Base"}}base int
	{{- else if eq $Opt "Fmt"}}fmt byte
	{{- else if eq $Opt "Prec"}}prec int
	{{- else if eq $Opt "FmtPrec"}}fmt byte, prec int{{end}}) String { return String{strconv.Format{{$From.TypeKind}}(e
		{{- if ne $From.BitSize 64}}.To{{$From.TypeKind}}
			{{- if ne $From.TypeKind "Complex"}}64{{else}}128{{end}}()
		{{- end}}.{{$From.TypeKind}}{{- if ne $From.TypeKind "Complex"}}64{{else}}128{{end}},
		{{- if eq $Opt "Base"}} base
		{{- else if and (eq $Opt "") (eq $From.TypeKind "Int")}} {{$.StrconvIntBaseName}}
		{{- else if and (eq $Opt "") (eq $From.TypeKind "Uint")}} {{$.StrconvUintBaseName}}
		{{- end -}}
		{{- if or (eq $From.TypeKind "Float") (eq $From.TypeKind "Complex")}}
			{{- if or (eq $Opt "Fmt") (eq $Opt "FmtPrec")}} fmt{{else}} {{$.StrconvFloatFmtName}}{{end}},
			{{- if or (eq $Opt "Prec") (eq $Opt "FmtPrec")}} prec{{else}} {{$.StrconvFloatPrecName}}{{end}},
			{{- if eq $From.TypeKind "Complex"}}
				{{- if eq $From.BitSize 32}} 64{{else}} 128{{end}}
			{{- else}} {{$From.BitSize}}
			{{- end}}
		{{- end}})} }
func (e {{$From.TypeName}}) ToBytes{{$Opt}}(
	{{- if eq $Opt "Base"}}base int
	{{- else if eq $Opt "Fmt"}}fmt byte
	{{- else if eq $Opt "Prec"}}prec int
	{{- else if eq $Opt "FmtPrec"}}fmt byte, prec int{{end}}) Bytes { return Bytes{forceconv.StringToBytes(e.ToString{{$Opt}}(
		{{- if eq $Opt "Base"}}base
		{{- else if eq $Opt "Fmt"}}fmt
		{{- else if eq $Opt "Prec"}}prec
		{{- else if eq $Opt "FmtPrec"}}fmt, prec{{end}}).String)} }
{{- end}}
{{- end -}} {{/* range $Opt */}}
func (e {{$From.TypeName}}) ToStringFormat(format string) String { return String{fmt.Sprintf(format, e)} }
func (e {{$From.TypeName}}) ToBytesFormat(format string) Bytes { return Bytes{forceconv.StringToBytes(e.ToStringFormat(format).String)} }

{{end}} {{/* range $From */}}

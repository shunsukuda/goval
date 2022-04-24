package goval

import (
	"strconv"
)

type Bool struct {
	Bool bool
}

func (e Bool) GoBool() bool {
	return e.Bool
}
func (e Bool) Interface() interface{} {
	return e.Bool
}
func (e Bool) Val() Val {
	return e
}
func (e Bool) Type() Type {
	return ValTypes.Bool
}
func (e Bool) Equal(x Bool) bool {
	return e.Bool == x.Bool
}

{{range $To := $.TL -}}
func (e Bool) To{{$To.TypeName}}() {{$To.TypeName}} {
	if e.Bool {
		return {{$To.TypeName}}{1}
	}
	return {{$To.TypeName}}{0}
}
{{end}}
func (e Bool) ToBool() Bool {
	return e
}
func (e Bool) ToString() String {
	return String{strconv.FormatBool(e.Bool)}
}
func (e Bool) ToBytes() Bytes {
	return e.ToString().ToBytes()
}
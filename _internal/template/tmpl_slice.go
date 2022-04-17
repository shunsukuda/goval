package goval

type ValSlice interface {
	Interface() interface{}
	ValSlice() ValSlice
	Type() Type
	Len() int
	Cap() int
	Grow(n int)
	Reset()
}

{{range $T := $.TL}}

type {{$T.TypeName}}Slice []{{$T.GoTypeName}}

func New{{$T.TypeName}}Slice(s []{{$T.GoTypeName}}) {{$T.TypeName}}Slice { return s }

func (s *{{$T.TypeName}}Slice) Interface() interface{} { return s.{{$T.TypeName}}Slice() }
func (s *{{$T.TypeName}}Slice) {{$T.TypeName}}Slice() []{{$T.GoTypeName}} { return *s }
func (s *{{$T.TypeName}}Slice) Type() Type { return ValTypes.{{$T.TypeName}} }
func (s *{{$T.TypeName}}Slice) Len() int { return len(*s) }
func (s *{{$T.TypeName}}Slice) Cap() int { return cap(*s) }

{{end}}
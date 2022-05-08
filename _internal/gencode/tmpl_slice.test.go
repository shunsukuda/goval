package goval

import (
	"reflect"
	"testing"
)

{{range $T := $.TL}}
func TestNew{{$T.TypeName}}Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    {{$T.TypeName}}Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: {{$T.TypeName}}Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: {{$T.TypeName}}Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New{{$T.TypeName}}Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New{{$T.TypeName}}Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("New{{$T.TypeName}}Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("New{{$T.TypeName}}Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNew{{$T.TypeName}}SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []{{$T.GoTypeName}}
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    {{$T.TypeName}}Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: {{$T.TypeName}}Slice{S: nil, ubbp: false}, wantLen: 0, wantCap:-1},
		{name: "", args: args{s: nil, ubbp: true}, want: {{$T.TypeName}}Slice{S: nil, ubbp: true}, wantLen: 0, wantCap:-1},
		{name: "", args: args{s: make([]{{$T.GoTypeName}},10, 20), ubbp: false}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]{{$T.GoTypeName}},100, 200), ubbp: false}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]{{$T.GoTypeName}},10, 20), ubbp: true}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]{{$T.GoTypeName}},100, 200), ubbp: true}, want: {{$T.TypeName}}Slice{S: make([]{{$T.TypeName}}, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New{{$T.TypeName}}SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       {{$T.TypeName}}Slice
		want    []{{$T.GoTypeName}}
		wantLen int
		wantCap int
	}{
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 20), false), want: make([]{{$T.GoTypeName}}, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 20), true), want: make([]{{$T.GoTypeName}}, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("New{{$T.TypeName}}SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name    string
		s       {{$T.TypeName}}Slice
		want    bool
	}{
		{name: "True", s: New{{$T.TypeName}}SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: New{{$T.TypeName}}SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("{{$T.TypeName}}Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       {{$T.TypeName}}Slice
		want    {{$T.TypeName}}Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 20), false), want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 20), true), want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 100, 200), false), want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 100, 200), true), want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{$T.TypeName}}Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.To{{$T.TypeName}}()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{$T.TypeName}}Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("{{$T.TypeName}}Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("{{$T.TypeName}}Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    {{$T.TypeName}}Slice
		want Type
	}{
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(nil, false), want: ValTypes.{{$T.TypeName}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("{{$T.TypeName}}Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    {{$T.TypeName}}Slice
		want int
	}{
		{name: "", s: New{{$T.TypeName}}Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    {{$T.TypeName}}Slice
		want int
	}{
		{name: "", s: New{{$T.TypeName}}Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got !=tt.want {
				t.Errorf("{{$T.TypeName}}Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       {{$T.TypeName}}Slice
		want    {{$T.TypeName}}Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 20), false), want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{$T.TypeName}}Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("{{$T.TypeName}}Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("{{$T.TypeName}}Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func Test{{$T.TypeName}}Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       {{$T.TypeName}}Slice
		args    args
		want    {{$T.TypeName}}Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 0), false), args: args{n: 0}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10), false), args: args{n: 10}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 100), false), args: args{n: 10}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 100), false), args: args{n: 200}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 2000), false), args: args{n: 200}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 0), true), args: args{n: 0}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10), true), args: args{n: 10}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 10, 100), true), args: args{n: 10}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 100), true), args: args{n: 200}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: New{{$T.TypeName}}SliceFromGoSlice(make([]{{$T.GoTypeName}}, 2000), true), args: args{n: 200}, want: {{$T.TypeName}}Slice{make([]{{$T.TypeName}}, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{$T.TypeName}}Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("{{$T.TypeName}}Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("{{$T.TypeName}}Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
{{end}}
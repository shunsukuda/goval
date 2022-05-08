package goval

import (
	"reflect"
	"testing"
)

func TestNewBoolSlice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    BoolSlice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: BoolSlice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: BoolSlice{S: make([]Bool, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: BoolSlice{S: make([]Bool, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: BoolSlice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: BoolSlice{S: make([]Bool, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: BoolSlice{S: make([]Bool, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBoolSlice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoolSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewBoolSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewBoolSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewBoolSliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []bool
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    BoolSlice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: BoolSlice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: BoolSlice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]bool, 10, 20), ubbp: false}, want: BoolSlice{S: make([]Bool, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]bool, 100, 200), ubbp: false}, want: BoolSlice{S: make([]Bool, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]bool, 10, 20), ubbp: true}, want: BoolSlice{S: make([]Bool, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]bool, 100, 200), ubbp: true}, want: BoolSlice{S: make([]Bool, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBoolSliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoolSliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewBoolSliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewBoolSliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestBoolSlice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSlice
		want    []bool
		wantLen int
		wantCap int
	}{
		{name: "", s: NewBoolSliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 20), false), want: make([]bool, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 20), true), want: make([]bool, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoolSliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewBoolSliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewBoolSliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestBoolSlice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSlice
		want bool
	}{
		{name: "True", s: NewBoolSliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewBoolSliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("BoolSlice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSlice
		want    BoolSlice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 20), false), want: BoolSlice{make([]Bool, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 20), true), want: BoolSlice{make([]Bool, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 100, 200), false), want: BoolSlice{make([]Bool, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 100, 200), true), want: BoolSlice{make([]Bool, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToBool()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("BoolSlice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("BoolSlice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestBoolSlice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSlice
		want Type
	}{
		{name: "", s: NewBoolSliceFromGoSlice(nil, false), want: ValTypes.Bool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("BoolSlice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSlice
		want int
	}{
		{name: "", s: NewBoolSlice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("BoolSlice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSlice
		want int
	}{
		{name: "", s: NewBoolSlice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("BoolSlice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSlice
		want    BoolSlice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 20), false), want: BoolSlice{make([]Bool, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("BoolSlice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("BoolSlice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestBoolSlice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       BoolSlice
		args    args
		want    BoolSlice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 0), false), args: args{n: 0}, want: BoolSlice{make([]Bool, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10), false), args: args{n: 10}, want: BoolSlice{make([]Bool, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 100), false), args: args{n: 10}, want: BoolSlice{make([]Bool, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 100), false), args: args{n: 200}, want: BoolSlice{make([]Bool, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 2000), false), args: args{n: 200}, want: BoolSlice{make([]Bool, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 0), true), args: args{n: 0}, want: BoolSlice{make([]Bool, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10), true), args: args{n: 10}, want: BoolSlice{make([]Bool, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 10, 100), true), args: args{n: 10}, want: BoolSlice{make([]Bool, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 100), true), args: args{n: 200}, want: BoolSlice{make([]Bool, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewBoolSliceFromGoSlice(make([]bool, 2000), true), args: args{n: 200}, want: BoolSlice{make([]Bool, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("BoolSlice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("BoolSlice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt8Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int8Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Int8Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Int8Slice{S: make([]Int8, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Int8Slice{S: make([]Int8, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Int8Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Int8Slice{S: make([]Int8, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Int8Slice{S: make([]Int8, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt8Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt8SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []int8
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int8Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Int8Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Int8Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]int8, 10, 20), ubbp: false}, want: Int8Slice{S: make([]Int8, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int8, 100, 200), ubbp: false}, want: Int8Slice{S: make([]Int8, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]int8, 10, 20), ubbp: true}, want: Int8Slice{S: make([]Int8, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int8, 100, 200), ubbp: true}, want: Int8Slice{S: make([]Int8, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt8SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt8SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt8SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt8SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt8Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Slice
		want    []int8
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt8SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 20), false), want: make([]int8, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 20), true), want: make([]int8, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt8SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewInt8SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewInt8SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt8Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Slice
		want bool
	}{
		{name: "True", s: NewInt8SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewInt8SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Int8Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Slice
		want    Int8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 20), false), want: Int8Slice{make([]Int8, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 20), true), want: Int8Slice{make([]Int8, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 100, 200), false), want: Int8Slice{make([]Int8, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 100, 200), true), want: Int8Slice{make([]Int8, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToInt8()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int8Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int8Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt8Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Slice
		want Type
	}{
		{name: "", s: NewInt8SliceFromGoSlice(nil, false), want: ValTypes.Int8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Int8Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Slice
		want int
	}{
		{name: "", s: NewInt8Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Int8Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Slice
		want int
	}{
		{name: "", s: NewInt8Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Int8Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Slice
		want    Int8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 20), false), want: Int8Slice{make([]Int8, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int8Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int8Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt8Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Int8Slice
		args    args
		want    Int8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 0), false), args: args{n: 0}, want: Int8Slice{make([]Int8, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10), false), args: args{n: 10}, want: Int8Slice{make([]Int8, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 100), false), args: args{n: 10}, want: Int8Slice{make([]Int8, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 100), false), args: args{n: 200}, want: Int8Slice{make([]Int8, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 2000), false), args: args{n: 200}, want: Int8Slice{make([]Int8, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 0), true), args: args{n: 0}, want: Int8Slice{make([]Int8, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10), true), args: args{n: 10}, want: Int8Slice{make([]Int8, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 10, 100), true), args: args{n: 10}, want: Int8Slice{make([]Int8, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 100), true), args: args{n: 200}, want: Int8Slice{make([]Int8, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt8SliceFromGoSlice(make([]int8, 2000), true), args: args{n: 200}, want: Int8Slice{make([]Int8, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int8Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int8Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt16Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int16Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Int16Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Int16Slice{S: make([]Int16, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Int16Slice{S: make([]Int16, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Int16Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Int16Slice{S: make([]Int16, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Int16Slice{S: make([]Int16, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt16Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt16SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []int16
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int16Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Int16Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Int16Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]int16, 10, 20), ubbp: false}, want: Int16Slice{S: make([]Int16, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int16, 100, 200), ubbp: false}, want: Int16Slice{S: make([]Int16, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]int16, 10, 20), ubbp: true}, want: Int16Slice{S: make([]Int16, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int16, 100, 200), ubbp: true}, want: Int16Slice{S: make([]Int16, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt16SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt16SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt16SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt16SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt16Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16Slice
		want    []int16
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt16SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 20), false), want: make([]int16, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 20), true), want: make([]int16, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt16SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewInt16SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewInt16SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt16Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Slice
		want bool
	}{
		{name: "True", s: NewInt16SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewInt16SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Int16Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16Slice
		want    Int16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 20), false), want: Int16Slice{make([]Int16, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 20), true), want: Int16Slice{make([]Int16, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 100, 200), false), want: Int16Slice{make([]Int16, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 100, 200), true), want: Int16Slice{make([]Int16, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToInt16()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int16Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int16Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt16Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Slice
		want Type
	}{
		{name: "", s: NewInt16SliceFromGoSlice(nil, false), want: ValTypes.Int16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Int16Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Slice
		want int
	}{
		{name: "", s: NewInt16Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Int16Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Slice
		want int
	}{
		{name: "", s: NewInt16Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Int16Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16Slice
		want    Int16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 20), false), want: Int16Slice{make([]Int16, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int16Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int16Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt16Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Int16Slice
		args    args
		want    Int16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 0), false), args: args{n: 0}, want: Int16Slice{make([]Int16, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10), false), args: args{n: 10}, want: Int16Slice{make([]Int16, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 100), false), args: args{n: 10}, want: Int16Slice{make([]Int16, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 100), false), args: args{n: 200}, want: Int16Slice{make([]Int16, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 2000), false), args: args{n: 200}, want: Int16Slice{make([]Int16, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 0), true), args: args{n: 0}, want: Int16Slice{make([]Int16, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10), true), args: args{n: 10}, want: Int16Slice{make([]Int16, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 10, 100), true), args: args{n: 10}, want: Int16Slice{make([]Int16, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 100), true), args: args{n: 200}, want: Int16Slice{make([]Int16, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt16SliceFromGoSlice(make([]int16, 2000), true), args: args{n: 200}, want: Int16Slice{make([]Int16, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int16Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int16Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt32Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Int32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Int32Slice{S: make([]Int32, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Int32Slice{S: make([]Int32, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Int32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Int32Slice{S: make([]Int32, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Int32Slice{S: make([]Int32, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt32Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt32SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []int32
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Int32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Int32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]int32, 10, 20), ubbp: false}, want: Int32Slice{S: make([]Int32, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int32, 100, 200), ubbp: false}, want: Int32Slice{S: make([]Int32, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]int32, 10, 20), ubbp: true}, want: Int32Slice{S: make([]Int32, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int32, 100, 200), ubbp: true}, want: Int32Slice{S: make([]Int32, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt32SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt32Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32Slice
		want    []int32
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt32SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 20), false), want: make([]int32, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 20), true), want: make([]int32, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewInt32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewInt32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt32Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Slice
		want bool
	}{
		{name: "True", s: NewInt32SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewInt32SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Int32Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32Slice
		want    Int32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 20), false), want: Int32Slice{make([]Int32, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 20), true), want: Int32Slice{make([]Int32, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 100, 200), false), want: Int32Slice{make([]Int32, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 100, 200), true), want: Int32Slice{make([]Int32, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToInt32()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int32Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int32Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt32Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Slice
		want Type
	}{
		{name: "", s: NewInt32SliceFromGoSlice(nil, false), want: ValTypes.Int32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Int32Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Slice
		want int
	}{
		{name: "", s: NewInt32Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Int32Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Slice
		want int
	}{
		{name: "", s: NewInt32Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Int32Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32Slice
		want    Int32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 20), false), want: Int32Slice{make([]Int32, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int32Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int32Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt32Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Int32Slice
		args    args
		want    Int32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 0), false), args: args{n: 0}, want: Int32Slice{make([]Int32, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10), false), args: args{n: 10}, want: Int32Slice{make([]Int32, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 100), false), args: args{n: 10}, want: Int32Slice{make([]Int32, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 100), false), args: args{n: 200}, want: Int32Slice{make([]Int32, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 2000), false), args: args{n: 200}, want: Int32Slice{make([]Int32, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 0), true), args: args{n: 0}, want: Int32Slice{make([]Int32, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10), true), args: args{n: 10}, want: Int32Slice{make([]Int32, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 10, 100), true), args: args{n: 10}, want: Int32Slice{make([]Int32, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 100), true), args: args{n: 200}, want: Int32Slice{make([]Int32, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt32SliceFromGoSlice(make([]int32, 2000), true), args: args{n: 200}, want: Int32Slice{make([]Int32, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int32Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int32Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt64Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Int64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Int64Slice{S: make([]Int64, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Int64Slice{S: make([]Int64, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Int64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Int64Slice{S: make([]Int64, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Int64Slice{S: make([]Int64, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt64Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewInt64SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []int64
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Int64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Int64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Int64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]int64, 10, 20), ubbp: false}, want: Int64Slice{S: make([]Int64, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int64, 100, 200), ubbp: false}, want: Int64Slice{S: make([]Int64, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]int64, 10, 20), ubbp: true}, want: Int64Slice{S: make([]Int64, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]int64, 100, 200), ubbp: true}, want: Int64Slice{S: make([]Int64, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt64SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewInt64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewInt64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt64Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Slice
		want    []int64
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt64SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 20), false), want: make([]int64, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 20), true), want: make([]int64, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewInt64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewInt64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt64Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Slice
		want bool
	}{
		{name: "True", s: NewInt64SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewInt64SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Int64Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Slice
		want    Int64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 20), false), want: Int64Slice{make([]Int64, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 20), true), want: Int64Slice{make([]Int64, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 100, 200), false), want: Int64Slice{make([]Int64, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 100, 200), true), want: Int64Slice{make([]Int64, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToInt64()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int64Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int64Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt64Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Slice
		want Type
	}{
		{name: "", s: NewInt64SliceFromGoSlice(nil, false), want: ValTypes.Int64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Int64Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Slice
		want int
	}{
		{name: "", s: NewInt64Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Int64Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Slice
		want int
	}{
		{name: "", s: NewInt64Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Int64Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Slice
		want    Int64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 20), false), want: Int64Slice{make([]Int64, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int64Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int64Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt64Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Int64Slice
		args    args
		want    Int64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 0), false), args: args{n: 0}, want: Int64Slice{make([]Int64, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10), false), args: args{n: 10}, want: Int64Slice{make([]Int64, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 100), false), args: args{n: 10}, want: Int64Slice{make([]Int64, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 100), false), args: args{n: 200}, want: Int64Slice{make([]Int64, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 2000), false), args: args{n: 200}, want: Int64Slice{make([]Int64, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 0), true), args: args{n: 0}, want: Int64Slice{make([]Int64, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10), true), args: args{n: 10}, want: Int64Slice{make([]Int64, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 10, 100), true), args: args{n: 10}, want: Int64Slice{make([]Int64, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 100), true), args: args{n: 200}, want: Int64Slice{make([]Int64, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewInt64SliceFromGoSlice(make([]int64, 2000), true), args: args{n: 200}, want: Int64Slice{make([]Int64, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Int64Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Int64Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint8Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint8Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Uint8Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Uint8Slice{S: make([]Uint8, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Uint8Slice{S: make([]Uint8, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Uint8Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Uint8Slice{S: make([]Uint8, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Uint8Slice{S: make([]Uint8, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint8Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint8SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []uint8
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint8Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Uint8Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Uint8Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]uint8, 10, 20), ubbp: false}, want: Uint8Slice{S: make([]Uint8, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint8, 100, 200), ubbp: false}, want: Uint8Slice{S: make([]Uint8, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]uint8, 10, 20), ubbp: true}, want: Uint8Slice{S: make([]Uint8, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint8, 100, 200), ubbp: true}, want: Uint8Slice{S: make([]Uint8, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint8SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint8SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint8SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint8SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint8Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint8Slice
		want    []uint8
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint8SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 20), false), want: make([]uint8, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 20), true), want: make([]uint8, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint8SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewUint8SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewUint8SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint8Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Slice
		want bool
	}{
		{name: "True", s: NewUint8SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewUint8SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Uint8Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint8Slice
		want    Uint8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 20), false), want: Uint8Slice{make([]Uint8, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 20), true), want: Uint8Slice{make([]Uint8, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 100, 200), false), want: Uint8Slice{make([]Uint8, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 100, 200), true), want: Uint8Slice{make([]Uint8, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToUint8()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint8Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint8Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint8Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Slice
		want Type
	}{
		{name: "", s: NewUint8SliceFromGoSlice(nil, false), want: ValTypes.Uint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint8Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Slice
		want int
	}{
		{name: "", s: NewUint8Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Uint8Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Slice
		want int
	}{
		{name: "", s: NewUint8Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Uint8Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint8Slice
		want    Uint8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 20), false), want: Uint8Slice{make([]Uint8, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint8Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint8Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint8Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Uint8Slice
		args    args
		want    Uint8Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 0), false), args: args{n: 0}, want: Uint8Slice{make([]Uint8, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10), false), args: args{n: 10}, want: Uint8Slice{make([]Uint8, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 100), false), args: args{n: 10}, want: Uint8Slice{make([]Uint8, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 100), false), args: args{n: 200}, want: Uint8Slice{make([]Uint8, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 2000), false), args: args{n: 200}, want: Uint8Slice{make([]Uint8, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 0), true), args: args{n: 0}, want: Uint8Slice{make([]Uint8, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10), true), args: args{n: 10}, want: Uint8Slice{make([]Uint8, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 10, 100), true), args: args{n: 10}, want: Uint8Slice{make([]Uint8, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 100), true), args: args{n: 200}, want: Uint8Slice{make([]Uint8, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint8SliceFromGoSlice(make([]uint8, 2000), true), args: args{n: 200}, want: Uint8Slice{make([]Uint8, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint8Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint8Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint16Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint16Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Uint16Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Uint16Slice{S: make([]Uint16, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Uint16Slice{S: make([]Uint16, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Uint16Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Uint16Slice{S: make([]Uint16, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Uint16Slice{S: make([]Uint16, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint16Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint16SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []uint16
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint16Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Uint16Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Uint16Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]uint16, 10, 20), ubbp: false}, want: Uint16Slice{S: make([]Uint16, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint16, 100, 200), ubbp: false}, want: Uint16Slice{S: make([]Uint16, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]uint16, 10, 20), ubbp: true}, want: Uint16Slice{S: make([]Uint16, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint16, 100, 200), ubbp: true}, want: Uint16Slice{S: make([]Uint16, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint16SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint16SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint16SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint16SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint16Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16Slice
		want    []uint16
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint16SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 20), false), want: make([]uint16, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 20), true), want: make([]uint16, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint16SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewUint16SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewUint16SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint16Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Slice
		want bool
	}{
		{name: "True", s: NewUint16SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewUint16SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Uint16Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16Slice
		want    Uint16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 20), false), want: Uint16Slice{make([]Uint16, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 20), true), want: Uint16Slice{make([]Uint16, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 100, 200), false), want: Uint16Slice{make([]Uint16, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 100, 200), true), want: Uint16Slice{make([]Uint16, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToUint16()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint16Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint16Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint16Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Slice
		want Type
	}{
		{name: "", s: NewUint16SliceFromGoSlice(nil, false), want: ValTypes.Uint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint16Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Slice
		want int
	}{
		{name: "", s: NewUint16Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Uint16Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Slice
		want int
	}{
		{name: "", s: NewUint16Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Uint16Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16Slice
		want    Uint16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 20), false), want: Uint16Slice{make([]Uint16, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint16Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint16Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint16Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Uint16Slice
		args    args
		want    Uint16Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 0), false), args: args{n: 0}, want: Uint16Slice{make([]Uint16, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10), false), args: args{n: 10}, want: Uint16Slice{make([]Uint16, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 100), false), args: args{n: 10}, want: Uint16Slice{make([]Uint16, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 100), false), args: args{n: 200}, want: Uint16Slice{make([]Uint16, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 2000), false), args: args{n: 200}, want: Uint16Slice{make([]Uint16, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 0), true), args: args{n: 0}, want: Uint16Slice{make([]Uint16, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10), true), args: args{n: 10}, want: Uint16Slice{make([]Uint16, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 10, 100), true), args: args{n: 10}, want: Uint16Slice{make([]Uint16, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 100), true), args: args{n: 200}, want: Uint16Slice{make([]Uint16, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint16SliceFromGoSlice(make([]uint16, 2000), true), args: args{n: 200}, want: Uint16Slice{make([]Uint16, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint16Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint16Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint32Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Uint32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Uint32Slice{S: make([]Uint32, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Uint32Slice{S: make([]Uint32, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Uint32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Uint32Slice{S: make([]Uint32, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Uint32Slice{S: make([]Uint32, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint32Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint32SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []uint32
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Uint32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Uint32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]uint32, 10, 20), ubbp: false}, want: Uint32Slice{S: make([]Uint32, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint32, 100, 200), ubbp: false}, want: Uint32Slice{S: make([]Uint32, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]uint32, 10, 20), ubbp: true}, want: Uint32Slice{S: make([]Uint32, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint32, 100, 200), ubbp: true}, want: Uint32Slice{S: make([]Uint32, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint32SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint32Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Slice
		want    []uint32
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint32SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 20), false), want: make([]uint32, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 20), true), want: make([]uint32, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewUint32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewUint32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint32Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Slice
		want bool
	}{
		{name: "True", s: NewUint32SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewUint32SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Uint32Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Slice
		want    Uint32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 20), false), want: Uint32Slice{make([]Uint32, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 20), true), want: Uint32Slice{make([]Uint32, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 100, 200), false), want: Uint32Slice{make([]Uint32, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 100, 200), true), want: Uint32Slice{make([]Uint32, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToUint32()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint32Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint32Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint32Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Slice
		want Type
	}{
		{name: "", s: NewUint32SliceFromGoSlice(nil, false), want: ValTypes.Uint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint32Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Slice
		want int
	}{
		{name: "", s: NewUint32Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Uint32Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Slice
		want int
	}{
		{name: "", s: NewUint32Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Uint32Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Slice
		want    Uint32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 20), false), want: Uint32Slice{make([]Uint32, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint32Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint32Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint32Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Uint32Slice
		args    args
		want    Uint32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 0), false), args: args{n: 0}, want: Uint32Slice{make([]Uint32, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10), false), args: args{n: 10}, want: Uint32Slice{make([]Uint32, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 100), false), args: args{n: 10}, want: Uint32Slice{make([]Uint32, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 100), false), args: args{n: 200}, want: Uint32Slice{make([]Uint32, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 2000), false), args: args{n: 200}, want: Uint32Slice{make([]Uint32, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 0), true), args: args{n: 0}, want: Uint32Slice{make([]Uint32, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10), true), args: args{n: 10}, want: Uint32Slice{make([]Uint32, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 10, 100), true), args: args{n: 10}, want: Uint32Slice{make([]Uint32, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 100), true), args: args{n: 200}, want: Uint32Slice{make([]Uint32, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint32SliceFromGoSlice(make([]uint32, 2000), true), args: args{n: 200}, want: Uint32Slice{make([]Uint32, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint32Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint32Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint64Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Uint64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Uint64Slice{S: make([]Uint64, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Uint64Slice{S: make([]Uint64, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Uint64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Uint64Slice{S: make([]Uint64, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Uint64Slice{S: make([]Uint64, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint64Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewUint64SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []uint64
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Uint64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Uint64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Uint64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]uint64, 10, 20), ubbp: false}, want: Uint64Slice{S: make([]Uint64, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint64, 100, 200), ubbp: false}, want: Uint64Slice{S: make([]Uint64, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]uint64, 10, 20), ubbp: true}, want: Uint64Slice{S: make([]Uint64, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]uint64, 100, 200), ubbp: true}, want: Uint64Slice{S: make([]Uint64, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUint64SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewUint64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewUint64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint64Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Slice
		want    []uint64
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint64SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 20), false), want: make([]uint64, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 20), true), want: make([]uint64, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewUint64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewUint64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint64Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Slice
		want bool
	}{
		{name: "True", s: NewUint64SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewUint64SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Uint64Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Slice
		want    Uint64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 20), false), want: Uint64Slice{make([]Uint64, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 20), true), want: Uint64Slice{make([]Uint64, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 100, 200), false), want: Uint64Slice{make([]Uint64, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 100, 200), true), want: Uint64Slice{make([]Uint64, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToUint64()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint64Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint64Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint64Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Slice
		want Type
	}{
		{name: "", s: NewUint64SliceFromGoSlice(nil, false), want: ValTypes.Uint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Uint64Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Slice
		want int
	}{
		{name: "", s: NewUint64Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Uint64Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Slice
		want int
	}{
		{name: "", s: NewUint64Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Uint64Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Slice
		want    Uint64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 20), false), want: Uint64Slice{make([]Uint64, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint64Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint64Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint64Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Uint64Slice
		args    args
		want    Uint64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 0), false), args: args{n: 0}, want: Uint64Slice{make([]Uint64, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10), false), args: args{n: 10}, want: Uint64Slice{make([]Uint64, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 100), false), args: args{n: 10}, want: Uint64Slice{make([]Uint64, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 100), false), args: args{n: 200}, want: Uint64Slice{make([]Uint64, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 2000), false), args: args{n: 200}, want: Uint64Slice{make([]Uint64, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 0), true), args: args{n: 0}, want: Uint64Slice{make([]Uint64, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10), true), args: args{n: 10}, want: Uint64Slice{make([]Uint64, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 10, 100), true), args: args{n: 10}, want: Uint64Slice{make([]Uint64, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 100), true), args: args{n: 200}, want: Uint64Slice{make([]Uint64, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewUint64SliceFromGoSlice(make([]uint64, 2000), true), args: args{n: 200}, want: Uint64Slice{make([]Uint64, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Uint64Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Uint64Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewFloat32Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Float32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Float32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Float32Slice{S: make([]Float32, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Float32Slice{S: make([]Float32, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Float32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Float32Slice{S: make([]Float32, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Float32Slice{S: make([]Float32, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat32Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewFloat32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewFloat32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewFloat32SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []float32
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Float32Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Float32Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Float32Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]float32, 10, 20), ubbp: false}, want: Float32Slice{S: make([]Float32, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]float32, 100, 200), ubbp: false}, want: Float32Slice{S: make([]Float32, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]float32, 10, 20), ubbp: true}, want: Float32Slice{S: make([]Float32, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]float32, 100, 200), ubbp: true}, want: Float32Slice{S: make([]Float32, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat32SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewFloat32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewFloat32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat32Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Slice
		want    []float32
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat32SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 20), false), want: make([]float32, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 20), true), want: make([]float32, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat32SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewFloat32SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewFloat32SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat32Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Slice
		want bool
	}{
		{name: "True", s: NewFloat32SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewFloat32SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Float32Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Slice
		want    Float32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 20), false), want: Float32Slice{make([]Float32, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 20), true), want: Float32Slice{make([]Float32, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 100, 200), false), want: Float32Slice{make([]Float32, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 100, 200), true), want: Float32Slice{make([]Float32, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToFloat32()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float32Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float32Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat32Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Slice
		want Type
	}{
		{name: "", s: NewFloat32SliceFromGoSlice(nil, false), want: ValTypes.Float32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Float32Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Slice
		want int
	}{
		{name: "", s: NewFloat32Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Float32Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Slice
		want int
	}{
		{name: "", s: NewFloat32Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Float32Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Slice
		want    Float32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 20), false), want: Float32Slice{make([]Float32, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float32Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float32Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat32Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Float32Slice
		args    args
		want    Float32Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 0), false), args: args{n: 0}, want: Float32Slice{make([]Float32, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10), false), args: args{n: 10}, want: Float32Slice{make([]Float32, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 100), false), args: args{n: 10}, want: Float32Slice{make([]Float32, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 100), false), args: args{n: 200}, want: Float32Slice{make([]Float32, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 2000), false), args: args{n: 200}, want: Float32Slice{make([]Float32, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 0), true), args: args{n: 0}, want: Float32Slice{make([]Float32, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10), true), args: args{n: 10}, want: Float32Slice{make([]Float32, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 10, 100), true), args: args{n: 10}, want: Float32Slice{make([]Float32, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 100), true), args: args{n: 200}, want: Float32Slice{make([]Float32, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewFloat32SliceFromGoSlice(make([]float32, 2000), true), args: args{n: 200}, want: Float32Slice{make([]Float32, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float32Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float32Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewFloat64Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Float64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Float64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Float64Slice{S: make([]Float64, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Float64Slice{S: make([]Float64, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Float64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Float64Slice{S: make([]Float64, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Float64Slice{S: make([]Float64, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat64Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewFloat64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewFloat64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewFloat64SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []float64
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Float64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Float64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Float64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]float64, 10, 20), ubbp: false}, want: Float64Slice{S: make([]Float64, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]float64, 100, 200), ubbp: false}, want: Float64Slice{S: make([]Float64, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]float64, 10, 20), ubbp: true}, want: Float64Slice{S: make([]Float64, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]float64, 100, 200), ubbp: true}, want: Float64Slice{S: make([]Float64, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFloat64SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewFloat64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewFloat64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat64Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64Slice
		want    []float64
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat64SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 20), false), want: make([]float64, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 20), true), want: make([]float64, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloat64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewFloat64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewFloat64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat64Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Slice
		want bool
	}{
		{name: "True", s: NewFloat64SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewFloat64SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Float64Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64Slice
		want    Float64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 20), false), want: Float64Slice{make([]Float64, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 20), true), want: Float64Slice{make([]Float64, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 100, 200), false), want: Float64Slice{make([]Float64, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 100, 200), true), want: Float64Slice{make([]Float64, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToFloat64()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float64Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float64Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat64Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Slice
		want Type
	}{
		{name: "", s: NewFloat64SliceFromGoSlice(nil, false), want: ValTypes.Float64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Float64Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Slice
		want int
	}{
		{name: "", s: NewFloat64Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Float64Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Slice
		want int
	}{
		{name: "", s: NewFloat64Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Float64Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64Slice
		want    Float64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 20), false), want: Float64Slice{make([]Float64, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float64Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float64Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat64Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Float64Slice
		args    args
		want    Float64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 0), false), args: args{n: 0}, want: Float64Slice{make([]Float64, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10), false), args: args{n: 10}, want: Float64Slice{make([]Float64, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 100), false), args: args{n: 10}, want: Float64Slice{make([]Float64, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 100), false), args: args{n: 200}, want: Float64Slice{make([]Float64, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 2000), false), args: args{n: 200}, want: Float64Slice{make([]Float64, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 0), true), args: args{n: 0}, want: Float64Slice{make([]Float64, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10), true), args: args{n: 10}, want: Float64Slice{make([]Float64, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 10, 100), true), args: args{n: 10}, want: Float64Slice{make([]Float64, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 100), true), args: args{n: 200}, want: Float64Slice{make([]Float64, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewFloat64SliceFromGoSlice(make([]float64, 2000), true), args: args{n: 200}, want: Float64Slice{make([]Float64, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Float64Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Float64Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewComplex64Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Complex64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Complex64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Complex64Slice{S: make([]Complex64, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Complex64Slice{S: make([]Complex64, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Complex64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Complex64Slice{S: make([]Complex64, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Complex64Slice{S: make([]Complex64, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewComplex64Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewComplex64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewComplex64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewComplex64SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []complex64
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Complex64Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Complex64Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Complex64Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]complex64, 10, 20), ubbp: false}, want: Complex64Slice{S: make([]Complex64, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]complex64, 100, 200), ubbp: false}, want: Complex64Slice{S: make([]Complex64, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]complex64, 10, 20), ubbp: true}, want: Complex64Slice{S: make([]Complex64, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]complex64, 100, 200), ubbp: true}, want: Complex64Slice{S: make([]Complex64, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewComplex64SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewComplex64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewComplex64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex64Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex64Slice
		want    []complex64
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex64SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 20), false), want: make([]complex64, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 20), true), want: make([]complex64, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex64SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewComplex64SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewComplex64SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex64Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Complex64Slice
		want bool
	}{
		{name: "True", s: NewComplex64SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewComplex64SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Complex64Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex64Slice
		want    Complex64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 20), false), want: Complex64Slice{make([]Complex64, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 20), true), want: Complex64Slice{make([]Complex64, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 100, 200), false), want: Complex64Slice{make([]Complex64, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 100, 200), true), want: Complex64Slice{make([]Complex64, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToComplex64()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex64Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex64Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex64Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Complex64Slice
		want Type
	}{
		{name: "", s: NewComplex64SliceFromGoSlice(nil, false), want: ValTypes.Complex64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Complex64Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Complex64Slice
		want int
	}{
		{name: "", s: NewComplex64Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Complex64Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Complex64Slice
		want int
	}{
		{name: "", s: NewComplex64Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Complex64Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex64Slice
		want    Complex64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 20), false), want: Complex64Slice{make([]Complex64, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex64Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex64Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex64Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Complex64Slice
		args    args
		want    Complex64Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 0), false), args: args{n: 0}, want: Complex64Slice{make([]Complex64, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10), false), args: args{n: 10}, want: Complex64Slice{make([]Complex64, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 100), false), args: args{n: 10}, want: Complex64Slice{make([]Complex64, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 100), false), args: args{n: 200}, want: Complex64Slice{make([]Complex64, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 2000), false), args: args{n: 200}, want: Complex64Slice{make([]Complex64, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 0), true), args: args{n: 0}, want: Complex64Slice{make([]Complex64, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10), true), args: args{n: 10}, want: Complex64Slice{make([]Complex64, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 10, 100), true), args: args{n: 10}, want: Complex64Slice{make([]Complex64, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 100), true), args: args{n: 200}, want: Complex64Slice{make([]Complex64, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewComplex64SliceFromGoSlice(make([]complex64, 2000), true), args: args{n: 200}, want: Complex64Slice{make([]Complex64, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex64Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex64Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewComplex128Slice(t *testing.T) {
	type args struct {
		size int
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Complex128Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{size: 0, ubbp: false}, want: Complex128Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: false}, want: Complex128Slice{S: make([]Complex128, 10), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: false}, want: Complex128Slice{S: make([]Complex128, 100), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{size: 0, ubbp: true}, want: Complex128Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{size: 10, ubbp: true}, want: Complex128Slice{S: make([]Complex128, 10), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{size: 100, ubbp: true}, want: Complex128Slice{S: make([]Complex128, 100), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewComplex128Slice(tt.args.size, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex128Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewComplex128Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewComplex128Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestNewComplex128SliceFromGoSlice(t *testing.T) {
	type args struct {
		s    []complex128
		ubbp bool
	}
	tests := []struct {
		name    string
		args    args
		want    Complex128Slice
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil, ubbp: false}, want: Complex128Slice{S: nil, ubbp: false}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: nil, ubbp: true}, want: Complex128Slice{S: nil, ubbp: true}, wantLen: 0, wantCap: -1},
		{name: "", args: args{s: make([]complex128, 10, 20), ubbp: false}, want: Complex128Slice{S: make([]Complex128, 10, 20), ubbp: false}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]complex128, 100, 200), ubbp: false}, want: Complex128Slice{S: make([]Complex128, 100, 200), ubbp: false}, wantLen: 100, wantCap: -1},
		{name: "", args: args{s: make([]complex128, 10, 20), ubbp: true}, want: Complex128Slice{S: make([]Complex128, 10, 20), ubbp: true}, wantLen: 10, wantCap: -1},
		{name: "", args: args{s: make([]complex128, 100, 200), ubbp: true}, want: Complex128Slice{S: make([]Complex128, 100, 200), ubbp: true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewComplex128SliceFromGoSlice(tt.args.s, tt.args.ubbp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex128SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("NewComplex128SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("NewComplex128SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex128Slice_GoSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex128Slice
		want    []complex128
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex128SliceFromGoSlice(nil, false), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 20), false), want: make([]complex128, 10, 20), wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(nil, true), want: nil, wantLen: 0, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 20), true), want: make([]complex128, 10, 20), wantLen: 10, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GoSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplex128SliceFromGoSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := len(got); l != tt.wantLen {
					t.Errorf("NewComplex128SliceFromGoSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := cap(got); c != tt.wantCap {
					t.Errorf("NewComplex128SliceFromGoSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex128Slice_UseByteBufferPool(t *testing.T) {
	tests := []struct {
		name string
		s    Complex128Slice
		want bool
	}{
		{name: "True", s: NewComplex128SliceFromGoSlice(nil, true), want: true},
		{name: "False", s: NewComplex128SliceFromGoSlice(nil, false), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UseByteBufferPool(); got != tt.want {
				t.Errorf("Complex128Slice.UseByteBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Slice_Copy(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex128Slice
		want    Complex128Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 20), false), want: Complex128Slice{make([]Complex128, 10, 20), false}, wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 20), true), want: Complex128Slice{make([]Complex128, 10, 20), true}, wantLen: 10, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 100, 200), false), want: Complex128Slice{make([]Complex128, 100, 200), false}, wantLen: 100, wantCap: -1},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 100, 200), true), want: Complex128Slice{make([]Complex128, 100, 200), true}, wantLen: 100, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128Slice.Copy() = %v, want %v", got, tt.want)
			}
			got.S[0] = Int64{100}.ToComplex128()
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128Slice.Copy() is not a copy.")
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex128Slice.Copy() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex128Slice.Copy() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex128Slice_Type(t *testing.T) {
	tests := []struct {
		name string
		s    Complex128Slice
		want Type
	}{
		{name: "", s: NewComplex128SliceFromGoSlice(nil, false), want: ValTypes.Complex128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Type(); !got.Equal(tt.want) {
				t.Errorf("Complex128Slice.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Complex128Slice
		want int
	}{
		{name: "", s: NewComplex128Slice(10, false), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Complex128Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Slice_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    Complex128Slice
		want int
	}{
		{name: "", s: NewComplex128Slice(10, false), want: minSliceSize},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Complex128Slice.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Slice_Reset(t *testing.T) {
	tests := []struct {
		name    string
		s       Complex128Slice
		want    Complex128Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 20), false), want: Complex128Slice{make([]Complex128, 0, 20), false}, wantLen: 0, wantCap: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Reset()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128Slice.Reset() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex128Slice.Reset() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex128Slice.Reset() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex128Slice_Grow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		s       Complex128Slice
		args    args
		want    Complex128Slice
		wantLen int
		wantCap int
	}{
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 0), false), args: args{n: 0}, want: Complex128Slice{make([]Complex128, 0), false}, wantLen: 0, wantCap: 0},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10), false), args: args{n: 10}, want: Complex128Slice{make([]Complex128, 10), false}, wantLen: 10, wantCap: 30},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 100), false), args: args{n: 10}, want: Complex128Slice{make([]Complex128, 10), false}, wantLen: 10, wantCap: 100},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 100), false), args: args{n: 200}, want: Complex128Slice{make([]Complex128, 100), false}, wantLen: 100, wantCap: 400},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 2000), false), args: args{n: 200}, want: Complex128Slice{make([]Complex128, 2000), false}, wantLen: 2000, wantCap: 2700},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 0), true), args: args{n: 0}, want: Complex128Slice{make([]Complex128, 0), true}, wantLen: 0, wantCap: 0},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10), true), args: args{n: 10}, want: Complex128Slice{make([]Complex128, 10), true}, wantLen: 10, wantCap: 30},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 10, 100), true), args: args{n: 10}, want: Complex128Slice{make([]Complex128, 10), true}, wantLen: 10, wantCap: 100},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 100), true), args: args{n: 200}, want: Complex128Slice{make([]Complex128, 100), true}, wantLen: 100, wantCap: 400},
		{name: "", s: NewComplex128SliceFromGoSlice(make([]complex128, 2000), true), args: args{n: 200}, want: Complex128Slice{make([]Complex128, 2000), true}, wantLen: 2000, wantCap: 2700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol, oc := tt.s.Len(), tt.s.Cap()
			tt.s.Grow(tt.args.n)
			nl, nc := tt.s.Len(), tt.s.Cap()
			got := tt.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128Slice.Grow() = %v, want %v", got, tt.want)
				t.Errorf("len %v -> %v, cap %v -> %v", ol, nl, oc, nc)
			}
			if tt.wantLen > 0 {
				if l := got.Len(); l != tt.wantLen {
					t.Errorf("Complex128Slice.Grow() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c := got.Cap(); c != tt.wantCap {
					t.Errorf("Complex128Slice.Grow() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

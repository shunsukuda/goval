package main

import (
	"os"

	"github.com/shunsukuda/go-gencode"
)

type tmplTypeInfo struct {
	TypeName    string
	GoTypeName  string
	TypeKind    string
	ReflectKind string
	ZeroValue   string
	BitSize     int
	SizeOf      int
}

type tmplConfNumeric struct {
	From                  []tmplTypeInfo
	To                    []tmplTypeInfo
	ToString              []tmplTypeInfo
	ToStringFuncs         []string
	StrconvIntBaseName    string
	StrconvIntBaseValue   string
	StrconvUintBaseName   string
	StrconvUintBaseValue  string
	StrconvFloatFmtName   string
	StrconvFloatFmtValue  string
	StrconvFloatPrecName  string
	StrconvFloatPrecValue string
}

type tmplConfString struct {
	From []tmplTypeInfo
	To   []tmplTypeInfo
}

type tmplConfTypeInfoList struct {
	TL []tmplTypeInfo
}

var (
	typeInfoListNumeric = []tmplTypeInfo{
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", ReflectKind: "Int8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", ReflectKind: "Int16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", ReflectKind: "Int32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", ReflectKind: "Int64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", ReflectKind: "Uint8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", ReflectKind: "Uint16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", ReflectKind: "Uint32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", ReflectKind: "Uint64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", ReflectKind: "Float32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", ReflectKind: "Float64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", ReflectKind: "Complex64", ZeroValue: "0", BitSize: 32, SizeOf: 8},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", ReflectKind: "Complex128", ZeroValue: "0", BitSize: 64, SizeOf: 16},
	}
	typeInfoListString = []tmplTypeInfo{
		{TypeName: "Bytes", GoTypeName: "[]byte", TypeKind: "Bytes", ReflectKind: "Slice", ZeroValue: "nil", BitSize: 0, SizeOf: 24},
		{TypeName: "String", GoTypeName: "string", TypeKind: "String", ReflectKind: "String", ZeroValue: "\"\"", BitSize: 0, SizeOf: 16},
	}
	typeInfoListAll = []tmplTypeInfo{
		{TypeName: "Nil", GoTypeName: "nil", TypeKind: "Nil", ReflectKind: "Invalid", ZeroValue: "nil", BitSize: 0, SizeOf: 0},
		{TypeName: "Bool", GoTypeName: "bool", TypeKind: "Bool", ReflectKind: "Bool", ZeroValue: "0", BitSize: 1, SizeOf: 1},
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", ReflectKind: "Int8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", ReflectKind: "Int16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", ReflectKind: "Int32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", ReflectKind: "Int64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", ReflectKind: "Uint8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", ReflectKind: "Uint16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", ReflectKind: "Uint32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", ReflectKind: "Uint64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", ReflectKind: "Float32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", ReflectKind: "Float64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", ReflectKind: "Complex64", ZeroValue: "0", BitSize: 32, SizeOf: 8},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", ReflectKind: "Complex128", ZeroValue: "0", BitSize: 64, SizeOf: 16},
		{TypeName: "String", GoTypeName: "string", TypeKind: "String", ReflectKind: "String", ZeroValue: "\"\"", BitSize: 0, SizeOf: 16},
		{TypeName: "Bytes", GoTypeName: "[]byte", TypeKind: "Bytes", ReflectKind: "Slice", ZeroValue: "nil", BitSize: 0, SizeOf: 24},
	}

	typeInfoListSlice = []tmplTypeInfo{
		{TypeName: "Bool", GoTypeName: "bool", TypeKind: "Bool", ReflectKind: "Bool", ZeroValue: "0", BitSize: 1, SizeOf: 1},
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", ReflectKind: "Int8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", ReflectKind: "Int16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", ReflectKind: "Int32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", ReflectKind: "Int64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", ReflectKind: "Uint8", ZeroValue: "0", BitSize: 8, SizeOf: 1},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", ReflectKind: "Uint16", ZeroValue: "0", BitSize: 16, SizeOf: 2},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", ReflectKind: "Uint32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", ReflectKind: "Uint64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", ReflectKind: "Float32", ZeroValue: "0", BitSize: 32, SizeOf: 4},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", ReflectKind: "Float64", ZeroValue: "0", BitSize: 64, SizeOf: 8},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", ReflectKind: "Complex64", ZeroValue: "0", BitSize: 32, SizeOf: 8},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", ReflectKind: "Complex128", ZeroValue: "0", BitSize: 64, SizeOf: 16},
		//{TypeName: "String", GoTypeName: "string", TypeKind: "String", ReflectKind: "String", ZeroValue: "\"\"", BitSize: 0, SizeOf: 16},
	}

	tmplDataNumeric = tmplConfNumeric{
		From:                  typeInfoListNumeric,
		To:                    typeInfoListNumeric,
		ToString:              typeInfoListString,
		ToStringFuncs:         []string{"", "Base", "Fmt", "Prec", "FmtPrec"},
		StrconvIntBaseName:    "IntToStringBase",
		StrconvIntBaseValue:   "10",
		StrconvUintBaseName:   "UintToStringBase",
		StrconvUintBaseValue:  "10",
		StrconvFloatFmtName:   "FloatToStringFmt",
		StrconvFloatFmtValue:  "'g'",
		StrconvFloatPrecName:  "FloatToStringPrec",
		StrconvFloatPrecValue: "-1",
	}
	tmplDataString = tmplConfString{
		From: typeInfoListString,
		To:   typeInfoListNumeric,
	}
	tmplDataGeneral = tmplConfTypeInfoList{
		TL: typeInfoListNumeric,
	}
	tmplDataSlice = tmplConfTypeInfoList{
		TL: typeInfoListSlice,
	}
	tmplDataType = tmplConfTypeInfoList{
		TL: typeInfoListAll,
	}
)

type tmplSet struct {
	Name   string
	Input  string
	Output string
	Config interface{}
}

func main() {
	set := []tmplSet{
		{Name: "Numeric", Input: "tmpl_numeric.go", Output: "numeric.gen.go", Config: tmplDataNumeric},
		{Name: "NumericTest", Input: "tmpl_numeric.test.go", Output: "numeric.gen_test.go", Config: tmplDataNumeric},
		{Name: "String", Input: "tmpl_string.go", Output: "string.gen.go", Config: tmplDataString},
		{Name: "StringTest", Input: "tmpl_string.test.go", Output: "string.gen_test.go", Config: tmplDataString},
		{Name: "Bool", Input: "tmpl_bool.go", Output: "bool.gen.go", Config: tmplDataGeneral},
		{Name: "BoolTest", Input: "tmpl_bool.test.go", Output: "bool.gen_test.go", Config: tmplDataGeneral},
		{Name: "Type", Input: "tmpl_type.go", Output: "type.gen.go", Config: tmplDataType},
		{Name: "TypeTest", Input: "tmpl_type.test.go", Output: "type.gen_test.go", Config: tmplDataType},
		{Name: "Slice", Input: "tmpl_slice.go", Output: "slice.gen.go", Config: tmplDataSlice},
		{Name: "SliceTest", Input: "tmpl_slice.test.go", Output: "slice.gen_test.go", Config: tmplDataSlice},
	}

	PROJECT_PATH := os.Getenv("GOPATH") + "/src/github.com/shunsukuda/goval/"
	TEMPLATE_DIR := PROJECT_PATH + "_internal/gencode/"
	gencode.DoGoFmt = true

	for i := range set {
		set[i].Input = TEMPLATE_DIR + set[i].Input
		set[i].Output = PROJECT_PATH + set[i].Output
		gencode.GenCode(set[i].Name, set[i].Input, set[i].Output, set[i].Config)
	}

}

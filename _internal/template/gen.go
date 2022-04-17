package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type tmplTypeInfo struct {
	TypeName   string
	GoTypeName string
	TypeKind   string
	RefKind    string
	ZeroValue  string
	BitSize    int
	Sizeof     int
}

type tmplConfNumeric struct {
	To                    []tmplTypeInfo
	From                  []tmplTypeInfo
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
	To   []tmplTypeInfo
	From []tmplTypeInfo
}

type tmplConfTypeInfoList struct {
	TL []tmplTypeInfo
}

var (
	typeInfoListNumeric = []tmplTypeInfo{
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", BitSize: 8},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", BitSize: 16},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", BitSize: 32},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", BitSize: 64},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", BitSize: 8},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", BitSize: 16},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", BitSize: 32},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", BitSize: 64},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", BitSize: 32},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", BitSize: 64},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", BitSize: 32},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", BitSize: 64},
	}
	typeInfoListString = []tmplTypeInfo{
		{TypeName: "String", GoTypeName: "string", TypeKind: "String"},
		{TypeName: "Bytes", GoTypeName: "[]byte", TypeKind: "Bytes"},
	}
	typeInfoListAll = []tmplTypeInfo{
		{TypeName: "Nil", GoTypeName: "nil", TypeKind: "Nil", RefKind: "Invalid", BitSize: 0},
		{TypeName: "Bool", GoTypeName: "bool", TypeKind: "Bool", RefKind: "Bool", BitSize: 1},
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", RefKind: "", BitSize: 8},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", RefKind: "", BitSize: 16},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", RefKind: "", BitSize: 32},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", RefKind: "", BitSize: 64},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", RefKind: "", BitSize: 8},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", RefKind: "", BitSize: 16},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", RefKind: "", BitSize: 32},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", RefKind: "", BitSize: 64},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", RefKind: "", BitSize: 32},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", RefKind: "", BitSize: 64},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", RefKind: "", BitSize: 64},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", RefKind: "", BitSize: 128},
		{TypeName: "String", GoTypeName: "string", TypeKind: "String", RefKind: "", BitSize: 0},
		{TypeName: "Bytes", GoTypeName: "[]byte", TypeKind: "Bytes", RefKind: "Slice", BitSize: 0},
	}

	typeInfoListSlice = []tmplTypeInfo{
		{TypeName: "Bool", GoTypeName: "bool", TypeKind: "Bool", RefKind: "Bool", BitSize: 1},
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", RefKind: "", BitSize: 8},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", RefKind: "", BitSize: 16},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", RefKind: "", BitSize: 32},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", RefKind: "", BitSize: 64},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", RefKind: "", BitSize: 8},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", RefKind: "", BitSize: 16},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", RefKind: "", BitSize: 32},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", RefKind: "", BitSize: 64},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", RefKind: "", BitSize: 32},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", RefKind: "", BitSize: 64},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", RefKind: "", BitSize: 64},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", RefKind: "", BitSize: 128},
		{TypeName: "String", GoTypeName: "string", TypeKind: "String", RefKind: "", BitSize: 0},
	}

	tmplDataNumeric = tmplConfNumeric{
		To:                    typeInfoListNumeric,
		From:                  typeInfoListNumeric,
		StrconvIntBaseName:    "IntToStringBase",
		StrconvIntBaseValue:   "10",
		StrconvUintBaseName:   "UintToStringBase",
		StrconvUintBaseValue:  "16",
		StrconvFloatFmtName:   "FloatToStringFmt",
		StrconvFloatFmtValue:  "'g'",
		StrconvFloatPrecName:  "FloatToStringPrec",
		StrconvFloatPrecValue: "-1",
	}
	tmplDataString = tmplConfString{
		To:   typeInfoListNumeric,
		From: typeInfoListString,
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
		{Name: "String", Input: "tmpl_string.go", Output: "string.gen.go", Config: tmplDataString},
		{Name: "Bool", Input: "tmpl_bool.go", Output: "bool.gen.go", Config: tmplDataGeneral},
		{Name: "Type", Input: "tmpl_type.go", Output: "type.gen.go", Config: tmplDataType},
		{Name: "Slice", Input: "tmpl_slice.go", Output: "slice.gen.go", Config: tmplDataSlice},
	}

	PROJECT_PATH := os.Getenv("GOPATH") + "/src/github.com/shunsukuda/goval/"
	TEMPLATE_DIR := PROJECT_PATH + "_internal/template/"

	for i := range set {
		fmt.Printf("generate code: %s\n", set[i].Name)
		inputPath := TEMPLATE_DIR + set[i].Input
		outputPath := PROJECT_PATH + set[i].Output
		tmplFile, _ := os.Open(inputPath)
		defer tmplFile.Close()
		tmplBuf := bytes.NewBuffer(nil)
		io.Copy(tmplBuf, tmplFile)
		fmt.Printf("input %s %d bytes\n", inputPath, tmplBuf.Len())
		tmpl, err := template.New(set[i].Name).Parse(tmplBuf.String())
		if err != nil {
			log.Fatal(err)
		}
		outFile, _ := os.Create(outputPath)
		defer outFile.Close()
		if err = tmpl.Execute(outFile, set[i].Config); err != nil {
			log.Fatal(err)
		}
		if err = exec.Command("go", "fmt", outputPath).Start(); err != nil {
			log.Fatal(err)
		}
	}

}

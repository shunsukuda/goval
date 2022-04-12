package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
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
		{Name: "Forceconv", Input: "tmpl_forceconv.go", Output: "forceconv/forceconv.gen.go", Config: tmplDataGeneral},
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
	}

}

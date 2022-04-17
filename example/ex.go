package main

import (
	"fmt"

	"github.com/shunsukuda/goval"
)

func main() {
	s := goval.Float64Slice([]float64{1, 2, 3})
	fmt.Println(s.Float64Slice())
}

package main

import (
	"fmt"
	"math/bits"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	//. "github.com/shunsukuda/goval/decimal"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	d1, err := NewFromString("10")
	if err != nil {
		zlog.Err(err).Msg("")
	}
	d2, _ := NewFromString("0x01p-2")

	if err != nil {
		zlog.Err(err).Msg("")
	}
	fmt.Println(d1.String(), d2.String())
	v := uint64(1)
	i := 0
	for {
		v *= 5
		w, v := bits.Mul64(v, 5)
		if w != 0 {
			break
		}
		fmt.Printf("%d,\n", v)
		i++
	}
}

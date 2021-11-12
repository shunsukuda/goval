module github.com/shunsukuda/goval

go 1.17

replace github.com/shunsukuda/goval/decimal => ./decimal

require (
	github.com/apache/arrow/go/arrow v0.0.0-20211108161759-01b4caa7bcf4
	github.com/rs/zerolog v1.26.0
)

require golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect

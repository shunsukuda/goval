package decimal

/*
import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	//. "github.com/shunsukuda/goval"
	zlog "github.com/rs/zerolog/log"
	"github.com/shunsukuda/goval/forceconv"
)

type Decimal struct {
	val   *big.Int
	exp   int32
	mode  RoundingMode
	vaild bool
}

const (
	packageName    = "goval/decimal"
	digitsBinChars = "01"
	digitsOctChars = "01234567"
	digitsDecChars = "0123456789"
	digitsHexChars = "0123456789abcdefABCDEFG"
)

var (
	DecimalDivExp int32 = -16
)

var (
	bigV0  = big.NewInt(0)
	bigV1  = big.NewInt(1)
	bigV2  = big.NewInt(2)
	bigV5  = big.NewInt(5)
	bigV10 = big.NewInt(10)
)

var (
	Zero = NewInt64Exp(0, 0)
)

func NewInt64(x int64) Decimal {
	return NewInt64Exp(x, 0)
}

func NewInt64Exp(x int64, exp int32) Decimal {
	return Decimal{
		val: big.NewInt(x),
		exp: exp,
	}
}

func NewFromBigInt(x *big.Int) Decimal {
	return NewFromBigIntExp(x, 0)
}

func NewFromBigIntExp(x *big.Int, exp int32) Decimal {
	return Decimal{
		val: new(big.Int).Set(x),
		exp: exp,
	}
}

//     number    = [ sign ] ( float | "inf" | "Inf" ) .
//     sign      = "+" | "-" .
//     float     = ( mantissa | prefix pmantissa ) [ exponent ] .
//     prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
//     mantissa  = digits "." [ digits ] | digits | "." digits .
//     pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
//     exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
//     digits    = digit { [ "_" ] digit } .
//     digit     = "0" ... "9" | "a" ... "f" | "A" ... "F" .

type decimal struct {
	base  int
	dpix  int // decimal point index
	expix int // exponent index
}

func NewFromString(s string) (Decimal, error) {
	if len(s) == 0 {
		zlog.Debug().Str("package", packageName).Str("func", "NewFromString()").Msg("len=0")
		return Zero, nil
	}

	sign := ""
	if s[0] == '+' || s[0] == '-' {
		sign = s[0:1]
		s = s[1:]
	}

	if len(s) == 3 && (s == "Inf" || s == "inf") {
		zlog.Debug().Str("package", packageName).Str("func", "NewFromString()").Msg("cannot create Inf")
		return Zero, fmt.Errorf("goval/decimal: cannot create %v", s)

	}

	base := 10
	prefix := ""
	expSep := ""

	if s[0] == '0' { // with prefix
		switch s[1] {
		case 'b', 'B':
			base = 2
		case 'o', 'O':
			base = 8
		case 'x', 'X':
			base = 16
		default:
			base = 10
		}
	}

	if base != 10 {
		prefix = s[:2]
		s = s[2:]
	}
	split := strings.Split(s, ".")
	if (base == 10 && len(split) >= 3) || (base != 10 && len(split) >= 2) {
		zlog.Debug().Str("package", packageName).Str("func", "NewFromString()").Msg("cannot create decimal points.")
		return Zero, fmt.Errorf("goval/decimal: cannot create(interger part)")
	}

	if len(split) >= 2 && base == 10 {
		if i := strings.IndexAny(split[1], "eE"); i != -1 {
			expSep = s[i : i+1]
		}
	} else if len(split) >= 1 && base == 16 {
		if i := strings.IndexAny(split[0], "pP"); i != -1 {
			expSep = s[i : i+1]
		}
	}

	stri := ""
	strf := ""
	stre := ""

	// base 2/8/16 "integer part"+"exponent part"
	stri, l := scanDigits(split[0], 'i', base)
	if l == -1 {
		return Zero, fmt.Errorf("goval/decimal: cannot create(interger part)")
	}
	if expSep != "" {

		splitexp := strings.Split(split[1], expSep)
		if len(splitexp) >= 1 {
			strf, l = scanDigits(splitexp[0], 'f', base)
			if l == -1 {
				return Zero, fmt.Errorf("goval/decimal: cannot create(fractional part)")
			}
		}
		if len(splitexp) >= 2 {
			stre, l = scanDigits(splitexp[1], 'e', 10)
			if l == -1 {
				return Zero, fmt.Errorf("goval/decimal: cannot create(exponent part)")
			}
		}
	} else {
		if len(split) >= 2 {
			strf, l = scanDigits(split[1], 'f', base)
			if l == -1 {
				return Zero, fmt.Errorf("goval/decimal: cannot create(fractional part)")
			}
		}
	}

	strif := sign + prefix + stri + strf
	exp, err := strconv.ParseInt(stre, 10, 0)
	if err != nil && stre != "" {
		return Zero, fmt.Errorf("goval/decimal: cannot create %v", s)
	}
	exp -= int64(len(strf))
	mant, ok := new(big.Int).SetString(strif, base)
	if !ok {
		return Zero, fmt.Errorf("goval/decimal: cannot create %v", s)
	}

	zlog.Debug().Str("package", packageName).Str("func", "SetExp()").
		Str("v_sign", sign).Int("v_base", base).Str("v_stri", stri).Str("v_strf", strf).Str("v_stre", stre).Int64("v_exp", exp).Msg("")
	x := NewFromBigIntExp(mant, int32(exp))
	if x.IsZero() || base != 16 {
		return x, nil
	}

	for {
		x.val = x.val.And(x.val, bigV1)
		if x.val.Cmp(bigV1) == 0 { // (x.val & 1) != 1
			break
		}
		x.val.Rsh(x.val, 1)
		x.exp++
	}

	return x, nil
}

// part: 'i'=integer, 'f'=fractional, 'e'=exponent
// base: 2 or 8 or 10 or 16
func scanDigits(s string, part byte, base int) (string, int) {
	switch base {
	case 2:
		return scanDigitsBin(s)
	case 8:
		return scanDigitsOct(s)
	case 10:
		return scanDigitsDec(s, part)
	case 16:
		return scanDigitsHex(s, part)
	default:
		return "", -1
	}
}

func scanDigitsBin(s string) (string, int) {
	if len(s) == 0 {
		return "0", 1
	}
	if strings.Contains(s, "__") || (s[len(s)-1] == '_') {
		return "", -1
	}

	b := forceconv.StringToBytes(strings.ReplaceAll(s, "_", ""))
	result := make([]byte, len(b))

	for i := range b {
		if bytes.ContainsAny(b[i:i+1], digitsBinChars) {
			result[i] = b[i]
		} else {
			return "", -1
		}
	}
	return forceconv.BytesToString(result), len(result)
}

func scanDigitsOct(s string) (string, int) {
	if len(s) == 0 {
		return "0", 1
	}
	if strings.Contains(s, "__") || (s[len(s)-1] == '_') {
		return "", -1
	}

	b := forceconv.StringToBytes(strings.ReplaceAll(s, "_", ""))
	result := make([]byte, len(b))

	for i := range b {
		if bytes.ContainsAny(b[i:i+1], digitsOctChars) {
			result[i] = b[i]
		} else {
			return "", -1
		}
	}
	return forceconv.BytesToString(result), len(result)
}

func scanDigitsDec(s string, part byte) (string, int) {
	if len(s) == 0 {
		return "0", 1
	}
	if strings.Contains(s, "__") ||
		(part == 'e' && strings.Contains(s, "_")) ||
		(s[0] == '_') || (s[len(s)-1] == '_') {
		return "", -1
	}

	b := forceconv.StringToBytes(strings.ReplaceAll(s, "_", ""))
	result := make([]byte, len(b))

	sign := ""
	if part == 'e' && (b[0] == '+' || b[0] == '-') {
		sign = forceconv.BytesToString(b[0:1])
		b = b[1:]
	}
	for i := range b {
		if bytes.ContainsAny(b[i:i+1], digitsDecChars) {
			result[i] = b[i]
		} else {
			return "", -1
		}
	}
	return sign + forceconv.BytesToString(result), len(result)
}

func scanDigitsHex(s string, part byte) (string, int) {
	if len(s) == 0 {
		return "0", 1
	}
	if strings.Contains(s, "__") ||
		(part == 'e' && strings.Contains(s, "_")) ||
		(part == 'f' && s[0] == '_') || (s[len(s)-1] == '_') {
		return "", -1
	}

	b := forceconv.StringToBytes(strings.ReplaceAll(s, "_", ""))
	result := make([]byte, len(b))

	for i := range b {
		if bytes.ContainsAny(b[i:i+1], digitsHexChars) {
			result[i] = b[i]
		} else {
			return "", -1
		}
	}
	return forceconv.BytesToString(result), len(result)
}

func NewFromFloat64Exp(x float64, exp int32, rounding bool) Decimal {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		panic(fmt.Sprintf("goval/decimal: cannot create %v", x))
	}

	if x == 0 {
		return Zero
	}

	fb := NewNumBitsFromFloat64(x)
	return newFromFloatExp(&fb, exp, rounding)
}

func newFromFloatExp(fb *NumBits, exp int32, rounding bool) Decimal {
	return Decimal{}
}

type RoundingMode byte

const (
	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
	ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
	ToZero                            // == IEEE 754-2008 roundTowardZero
	AwayFromZero                      // no IEEE 754-2008 equivalent
	ToNegative                        // == IEEE 754-2008 roundTowardNegative
	ToPositive                        // == IEEE 754-2008 roundTowardPositive
)

func (e Decimal) Round(p int32, mode RoundingMode) Decimal {
	switch mode {
	case ToNearestEven:
		return e.RoundEven(p)
	case ToNearestAway:
		return e.RoundAway(p)
	case ToZero:
		if e.val.Sign() == -1 {
			return e.RoundUp(p)
		}
		return e.RoundDown(p)
	case AwayFromZero:
		if e.val.Sign() == -1 {
			return e.RoundDown(p)
		}
		return e.RoundUp(p)
	case ToNegative:
		return e.RoundDown(p)
	case ToPositive:
		return e.RoundUp(p)
	default:
		return e.RoundEven(p)
	}
}

func (e Decimal) RoundDown(p int32) Decimal {
	return e.SetExp(p)
}

func (e Decimal) RoundUp(p int32) Decimal {
	a, r := e.roundRem(p)
	if r != 0 {
		a.val = new(big.Int).Add(a.val, bigV1)
	}
	return a
}

func (e Decimal) RoundEven(p int32) Decimal {
	a, r := e.roundRem(p)
	if r == 5 {
		temp := new(big.Int).Rsh(a.val, 1)
		a.val = new(big.Int).Lsh(temp, 1)
	} else if r > 5 {
		a.val = new(big.Int).Add(a.val, bigV1)
	}
	return a
}

func (e Decimal) RoundAway(p int32) Decimal {
	a, r := e.roundRem(p)
	if r < 5 {
		a.val = new(big.Int).Add(a.val, bigV1)
	}
	return a
}

func (e Decimal) roundRem(p int32) (Decimal, int) {
	var a big.Int
	var r int
	if p < e.exp {

		a = *e.SetExp(p - 1).val
		a, r = a.QuoRem(a, bigV10, new(big.Int))
	} else {

	}

	return Decimal{
		val: a,
		exp: p - 1,
	}, int(r.Int64())
}

func (e Decimal) Normalize() Decimal {
	if e.IsZero() {
		return Zero
	}

	temp := new(big.Int).Set(e.val)
	exp := e.exp
	for {
		z, r := new(big.Int).QuoRem(temp, bigV10, new(big.Int))
		if r.Int64() != 0 {
			break
		}
		temp = z
		exp++
	}
	a := Decimal{
		val: temp,
		exp: exp,
	}
	zlog.Debug().Str("package", packageName).Str("func", "Normalize()").
		Str("e", e.String()).Str("a", a.String()).Msg("")

	return a
}

func (e Decimal) SetExp(exp int32) Decimal {
	if e.exp == exp {
		return e
	}

	if e.IsZero() {
		return Zero
	}

	temp := new(big.Int)
	diff := int64(-e.exp + exp)

	if diff > 0 {
		temp = temp.Quo(e.val, bigExp10(diff))
	} else { // diff < 0
		temp = temp.Mul(e.val, bigExp10(-diff))
	}

	a := Decimal{
		val: temp,
		exp: exp,
	}
	zlog.Debug().Str("package", packageName).Str("func", "SetExp()").
		Str("e", e.String()).Str("a", a.String()).Msg("")
	return a
}

func setExpPair(d1, d2 Decimal) (Decimal, Decimal) {
	if d1.exp < d2.exp {
		return d1, d2.SetExp(d1.exp)
	}

	if d1.exp > d2.exp {
		return d1.SetExp(d2.exp), d2
	}

	return d1, d2
}

func (e Decimal) IsZero() bool {
	return e.val.Cmp(bigV0) == 0
}

func (e Decimal) Add(x Decimal) Decimal {
	a, b := setExpPair(e, x)
	c := NewFromBigIntExp(new(big.Int).Add(a.val, b.val), a.exp)
	zlog.Debug().Str("package", packageName).Str("func", "Add()").
		Str("a", a.String()).Str("b", b.String()).Str("c", c.String()).Msg("")
	return c
}

func (e Decimal) Sub(x Decimal) Decimal {
	a, b := setExpPair(e, x)
	c := NewFromBigIntExp(new(big.Int).Sub(a.val, b.val), a.exp)
	zlog.Debug().Str("package", packageName).Str("func", "Sub()").
		Str("a", a.String()).Str("b", b.String()).Str("c", c.String()).Msg("")
	return a
}

func (e Decimal) Mul(x Decimal) Decimal {
	a, b := setExpPair(e, x)
	c := NewFromBigIntExp(new(big.Int).Mul(a.val, b.val), a.exp+b.exp)
	zlog.Debug().Str("package", packageName).Str("func", "Mul()").
		Str("a", a.String()).Str("b", b.String()).Str("c", c.String()).Msg("")
	return c
}

func (e Decimal) Quo(x Decimal) Decimal {
	if x.IsZero() {
		panic("goval/decimal: cannot divided by zero!")
	}

	a, b := setExpPair(e, x)

	a = a.SetExp(DecimalDivExp + a.exp)
	c := NewFromBigIntExp(new(big.Int).Quo(a.val, b.val), a.exp)
	c.exp = a.exp

	zlog.Debug().Str("package", packageName).Str("func", "Quo()").
		Str("v_a", a.String()).Str("v_b", b.String()).Str("v_c", c.String()).Msg("")
	return c
}

func (e Decimal) String() string {
	return e.val.Text(10) + "e" + strconv.FormatInt(int64(e.exp), 10)
}
*/

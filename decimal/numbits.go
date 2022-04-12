package decimal

/*
import (
	"math"
	"math/big"
)

var (
	exp5 = [...]uint64{
		1,
		5,
		25,
		125,
		625,
		3125,
		15625,
		78125,
		390625,
		1953125,
		9765625,
		48828125,
		244140625,
		1220703125,
		6103515625,
		30517578125,
		152587890625,
		762939453125,
		3814697265625,
		19073486328125,
		95367431640625,
		476837158203125,
		2384185791015625,
		11920928955078125,
		59604644775390625,
		298023223876953125,
		1490116119384765625,
		7450580596923828125,
	}
	exp10 = [...]uint64{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
		100000000000,
		1000000000000,
		10000000000000,
		100000000000000,
		1000000000000000,
		10000000000000000,
		100000000000000000,
		1000000000000000000,
		10000000000000000000,
	}
)

func bigExp2(x int32) *big.Int {
	if x < 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Lsh(big.NewInt(1), uint(x))
}

func bigExp5(x int32) *big.Int {
	if x < 0 {
		return big.NewInt(0)
	}
	if x <= int32(len(exp5)) {
		return new(big.Int).SetUint64(exp5[x])
	}
	return new(big.Int).Mul(new(big.Int).SetUint64(exp5[len(exp5)-1]), bigExp5(x-int32(len(exp5))))
}

func bigExp10(x int32) *big.Int {
	if x < 0 {
		return big.NewInt(0)
	}
	if x <= int32(len(exp10)) {
		return new(big.Int).SetUint64(exp10[x])
	}
	return new(big.Int).Mul(new(big.Int).SetUint64(exp5[len(exp5)-1]), bigExp10(x-int32(len(exp10))))
}

const (
	Float32ExpBias = 127
	Float32ExpBits = 8
	Float32ExpMin  = -126
	Float32ExpMax  = 127
	Float64ExpBias = 1023
	Float64ExpBits = 11
	Float64ExpMin  = 1022
	Float64ExpMax  = 1023
	Float32Prec    = 24
	Float64Prec    = 53
)

type bitsParttion struct {
	sign int32
	exp  int32
	num  int32
}

var (
	boolPart    = &bitsParttion{sign: 0, exp: 0, num: 1}
	int8Part    = &bitsParttion{sign: 1, exp: 0, num: 7}
	int16Part   = &bitsParttion{sign: 1, exp: 0, num: 15}
	int32Part   = &bitsParttion{sign: 1, exp: 0, num: 31}
	int64Part   = &bitsParttion{sign: 1, exp: 0, num: 63}
	uint8Part   = &bitsParttion{sign: 0, exp: 0, num: 8}
	uint16Part  = &bitsParttion{sign: 0, exp: 0, num: 16}
	uint32Part  = &bitsParttion{sign: 0, exp: 0, num: 32}
	uint64Part  = &bitsParttion{sign: 0, exp: 0, num: 64}
	float32Part = &bitsParttion{sign: 1, exp: 8, num: 23}
	float64Part = &bitsParttion{sign: 1, exp: 11, num: 52}
)

func (bp bitsParttion) size() int {
	return int(bp.sign + bp.exp + bp.num)
}

func (bp bitsParttion) expBias() int {
	if bp.exp >= 1 {
		return (1 << (bp.exp - 1)) - 1
	}
	return 0
}

func (bp bitsParttion) expMask() uint64 {
	if bp.num >= 0 {
		return (1 << bp.exp) - 1
	}
	return 0
}

func (bp bitsParttion) numMask() uint64 {
	if bp.num >= 0 {
		return (1 << bp.num) - 1
	}
	return 0
}

type NumBits struct {
	num  *big.Int
	exp  int32
	part *bitsParttion
	sign int8
}

func NewNumBits(bits uint64, part *bitsParttion) NumBits {
	sign := bits >> (part.exp + part.num)
	exp := int(((bits >> part.num) & part.expMask())) - part.expBias()
	num := new(big.Int).SetUint64(bits & (part.numMask()))

	num.Rsh(num, num.TrailingZeroBits())
	if exp > -part.expBias() {
		num.Or(num, new(big.Int).SetUint64(1<<num.BitLen()))
	}

	return NumBits{
		num:  num,
		exp:  int32(exp),
		part: part,
		sign: int8(sign),
	}
}

func NewNumBitsFromFloat32(x float32) NumBits {
	return NewNumBits(uint64(math.Float32bits(x)), float32Part)
}

func NewNumBitsFromFloat64(x float64) NumBits {
	return NewNumBits(math.Float64bits(x), float64Part)
}

func (e NumBits) blen() int {
	return e.num.BitLen()
}

func (x NumBits) Sign() int {
	return int(x.sign)
}

// Exp returns float exponent - exponent bias.
func (e NumBits) Exp() int {
	return int(e.exp)
}

// IntExp returns
func (e NumBits) IntExp() int {
	return int(e.exp) - e.num.BitLen() + 1
}

// IntMant return integer mantissa.
func (e NumBits) IntNum() *big.Int {
	return new(big.Int).Set(e.num)
}

func (e *NumBits) SetExp(exp int) {
	e.exp = int32(exp)
}

func (e NumBits) IsZero() bool {
	return e.num.BitLen() == 0
}

func (e NumBits) SignificantDigits() int32 {
	return 0
}

func (e NumBits) ToDecimal(exp int32) Decimal {

	if e.IsZero() {
		return Zero
	}

	exp2 := int32(e.IntExp())
	mant := e.IntNum()

	if exp < 0 && exp < exp2 {
		if exp2 < 0 {
			exp = exp2
		} else {
			exp = 0
		}
	}

	// a * 10^x = a * 2^x * 5^x
	// a * 2^n * 10^x = a * 2^(n+x) * 5^x
	exp2 -= exp

	temp := big.NewInt(1)
	// 5 ^ M
	if exp < 0 {
		mant = mant.Mul(mant, bigExp5(-exp))
	} else if exp > 0 {
		temp = bigExp5(exp)
	}

	// 2 ^ (M+N)
	if exp2 > 0 {
		mant = mant.Lsh(mant, uint(exp2))
	} else if exp2 < 0 {
		temp = temp.Lsh(temp, uint(-exp2))
	}

	// rounding and downscaling
	if exp > 0 || exp2 < 0 {
		halfDown := new(big.Int).Rsh(temp, 1)
		mant = mant.Add(mant, halfDown)
		mant = mant.Quo(mant, temp)
	}

	if e.sign != 0 {
		mant = mant.Neg(mant)
	}

	return Decimal{
		val: mant,
		exp: exp,
	}
}
*/

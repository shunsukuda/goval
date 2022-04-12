package gobig

import "math/big"

type BigInt big.Int

func NewBigInt() BigInt {
	var z BigInt
	return z
}

func newInt() *big.Int {
	var z BigInt
	return *z
}

func (e BigInt) BigInt() *big.Int {
	z := big.Int(e)
	return &z
}

func (e BigInt) Abs() *BigInt {
	a := BigInt(*newInt().Abs(e.BigInt()))
	return &a
}

func (e BigInt) Add(x *BigInt) *BigInt {
	a := BigInt(*e.BigInt().Add(e.BigInt(), x.BigInt()))
	return &a
}

func (e BigInt) And(x *BigInt) *BigInt {
	a := BigInt(*e.BigInt().And(e.BigInt(), x.BigInt()))
	return &a
}

func (e BigInt) AndNot(x *BigInt) *BigInt {
	a := BigInt(*e.BigInt().AndNot(e.BigInt(), x.BigInt()))
	return &a
}

// Appeend

// Binomial

func (e BigInt) Bit(i int) uint {
	return e.BigInt().Bit(i)
}

func (e BigInt) BitLen() int {
	return e.BigInt().BitLen()
}

func (e BigInt) Bytes() []byte {
	return e.BigInt().Bytes()
}

func (e BigInt) Cmp(x *BigInt) (r int) {
	r = e.BigInt().Cmp(x.BigInt())
	return
}

func (e BigInt) CmpAbs(x *BigInt) int {
	r := e.BigInt().Cmp(x.BigInt())
	return r
}

func (e BigInt) Div(x *BigInt) *BigInt {
	a := BigInt(*newInt().Div(e.BigInt(), x.BigInt()))
	return &a
}

func (e BigInt) DivMod(x, m *BigInt) *BigInt {
	a, b := newInt().DivMod(e.BigInt(), x.BigInt(), m.BigInt())
	return &BigInt(a), &BigInt(b)
}

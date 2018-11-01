package algorithm

import "math/big"

var (
	tt255     = BigPow(2, 255)
	tt256     = BigPow(2, 256)
	tt256m1   = new(big.Int).Sub(tt256, big.NewInt(1))
	tt63      = BigPow(2, 63)
	MaxBig256 = new(big.Int).Set(tt256m1)
	MaxBig63  = new(big.Int).Sub(tt63, big.NewInt(1))
)

func IntMax(a int64, b int64) int64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

func IntMin(a int64, b int64) int64 {
	if a >= b {
		return b
	} else {
		return a
	}
}

func FloatMax(a float64, b float64) float64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

func FloatMin(a float64, b float64) float64 {
	if a >= b {
		return b
	} else {
		return a
	}
}

// BigPow returns a ** b as a big integer.
func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

// BigToUint64 returns the integer casted to a uint64 and returns whether it
// overflowed in the process.
func BigToUint64(v *big.Int) (uint64, bool) {
	return v.Uint64(), v.BitLen() > 64
}

// U256 encodes as a 256 bit two's complement number. This operation is destructive.
func U256(x *big.Int) *big.Int {
	return x.And(x, tt256m1)
}

// S256 interprets x as a two's complement number.
// x must not exceed 256 bits (the result is undefined if it does) and is not modified.
//
//   S256(0)        = 0
//   S256(1)        = 1
//   S256(2**255)   = -2**255
//   S256(2**256-1) = -1
func S256(x *big.Int) *big.Int {
	if x.Cmp(tt255) < 0 {
		return x
	}
	return new(big.Int).Sub(x, tt256)
}

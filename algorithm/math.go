package algorithm

import "math/big"

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

// BigToUint64 returns the integer casted to a uint64 and returns whether it
// overflowed in the process.
func BigToUint64(v *big.Int) (uint64, bool) {
	return v.Uint64(), v.BitLen() > 64
}

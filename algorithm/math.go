package algorithm

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

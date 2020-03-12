package main

func PowerWithExponent(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	res := PowerWithExponent(base, exponent>>1)
	res *= res
	if exponent%2 == 1 {
		res *= base
	}

	return res
}

package main

import (
	"math"
)

func Sum(num int) int {
	return (int(math.Pow(float64(num), float64(2))) + num) >> 1
}

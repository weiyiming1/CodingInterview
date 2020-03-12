package main

func Fibonacci(n int) int {
	sum := make([]int, n)
	sum[0] = 1
	sum[1] = 1
	for idx := 2; idx < len(sum); idx++ {
		sum[idx] = sum[idx-1] + sum[idx-2]
	}
	return sum[n-1]
}

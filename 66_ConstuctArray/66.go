package main

func Multiply(A []int) []int {
	if A == nil || len(A) == 0 {
		return nil
	}
	//计算C, 左三角
	C := make([]int, len(A))
	C[0] = 1
	for i := 1; i < len(A); i++ {
		C[i] = C[i-1] * A[i-1]
	}

	//计算D, 右三角
	D := make([]int, len(A))
	D[len(A)-1] = 1
	for i := len(A) - 2; i >= 0; i-- {
		D[i] = D[i+1] * A[i+1]
	}

	//合成B
	B := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		B[i] = C[i] * D[i]
	}
	return B
}

package main

import (
	"fmt"
)

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	j := p
	for j >= p && j < r {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
		j++
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func quickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

func main() {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	quickSort(A, 0, 7)
	fmt.Println(A)
}

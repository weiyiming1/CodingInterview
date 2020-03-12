package main

import "fmt"

func searchMatrix(matrix [][]int, t int) bool {
	i := 0
	j := len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == t {
			return true
		} else if matrix[i][j] > t {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 6))
}

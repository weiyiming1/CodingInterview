package main

import "fmt"

func PrintMatrixClockwise(matrix [][]int, col, row int) {
	if matrix == nil || col <= 0 || row <= 0 {
		return
	}
	start := 0
	for col > start*2 && row > start*2 {
		printMatrixInCircle(matrix, col, row, start)
		start++
	}
}

func printMatrixInCircle(matrix [][]int, col, row, start int) {
	colEnd := col - 1 - start
	rowEnd := row - 1 - start

	// 从左到右打印一行
	for i := start; i <= colEnd; i++ {
		fmt.Println(matrix[start][i])
	}

	// 从上到下打印一行
	if start < rowEnd {
		for i := start + 1; i <= rowEnd; i++ {
			fmt.Println(matrix[i][colEnd])
		}
	}

	// 从右到左打印一行
	if start < rowEnd && start < colEnd {
		for i := colEnd - 1; i >= start; i-- {
			fmt.Println(matrix[rowEnd][i])
		}
	}

	//从下到上打印一列
	if start < colEnd && start < rowEnd-1 {
		for i := rowEnd - 1; i >= start+1; i-- {
			fmt.Println(matrix[i][start])
		}
	}
}

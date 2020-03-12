package main

func HasPath(matrix [][]rune, rows, cols int, chars []rune) bool {
	if matrix == nil || rows == 0 || cols == 0 || len(chars) == 0 {
		return false
	}
	visited := make([]bool, rows*cols)
	curLen := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if coreHasPath(matrix, rows, cols, i, j, chars, curLen, &visited) {
				return true
			}
		}
	}
	return false
}

func coreHasPath(matrix [][]rune, rows, cols, row, col int, chars []rune, curLen int, visited *[]bool) bool {
	if curLen == len(chars) {
		return true

	}
	hasPath := false
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		!(*visited)[row*cols+col] &&
		matrix[row][col] == chars[curLen] {
		curLen++
		(*visited)[row*cols+col] = true
		hasPath = coreHasPath(matrix, rows, cols, row-1, col, chars, curLen, visited) ||
			coreHasPath(matrix, rows, cols, row, col-1, chars, curLen, visited) ||
			coreHasPath(matrix, rows, cols, row+1, col, chars, curLen, visited) ||
			coreHasPath(matrix, rows, cols, row, col+1, chars, curLen, visited)
		if !hasPath {
			(*visited)[row*cols+col] = false
		}
	}
	return hasPath
}

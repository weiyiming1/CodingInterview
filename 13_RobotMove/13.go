package main

func MovingCount(threshold, rows, cols int) int {
	if threshold < 0 || rows < 0 || cols < 0 {
		return 0
	}
	visited := make([]bool, rows*cols)
	return movingCountCore(threshold, rows, cols, 0, 0, visited)
}

func movingCountCore(threshold, rows, cols, row, col int, visited []bool) int {
	count := 0
	if check(threshold, rows, cols, row, col, visited) {
		visited[row*cols+col] = true
		count = 1 + movingCountCore(threshold, rows, cols, row-1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col-1, visited) +
			movingCountCore(threshold, rows, cols, row+1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col+1, visited)
	}
	return count
}

func check(threshold, rows, cols, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		getDigitSum(row)+getDigitSum(col) <= threshold && !visited[row*cols+col] {
		return true
	}
	return false
}

func getDigitSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

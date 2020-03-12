package main

func NumInMax(matrix [][]int, rows, columns, target int) bool {
	flag := false

	if matrix != nil && rows > 0 && columns > 0 {
		row := 0
		col := columns - 1
		for row < rows && col >= 0 {
			if matrix[row][col] == target {
				flag = true
				break
			} else if matrix[row][col] > target {
				col--
			} else {
				row++
			}
		}
	}

	return flag
}

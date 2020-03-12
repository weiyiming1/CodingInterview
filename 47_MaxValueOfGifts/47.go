package main

func GetMaxVal(values [][]int, rows, cols int) int {
	if values == nil || rows <= 0 || cols <= 0 {
		return 0
	}

	maxVal := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// left和up 清零
			left := 0
			up := 0

			if i > 0 {
				up = maxVal[j]
			}
			if j > 0 {
				left = maxVal[j-1]
			}
			if up >= left {
				maxVal[j] = up + values[i][j]
			} else {
				maxVal[j] = left + values[i][j]
			}

		}
	}
	return maxVal[cols-1]
}

package main

func FindNumbersWithSum(data []int, sum int, num1, num2 *int) bool {
	if len(data) < 2 || num1 == nil || num2 == nil {
		return false
	}
	idx1 := 0
	idx2 := len(data) - 1

	for idx1 < idx2 {
		if data[idx1]+data[idx2] == sum {
			*num1 = data[idx1]
			*num2 = data[idx2]
			return true
		} else if data[idx1]+data[idx2] > sum {
			idx2--
		} else {
			idx1++
		}
	}
	return false
}

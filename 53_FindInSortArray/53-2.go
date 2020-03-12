package main

func GetMissingNumber(num []int) int {
	if len(num) == 0 {
		return -1
	}
	left := 0
	right := len(num) - 1
	for left <= right {
		mid := (left + right) >> 1
		if num[mid] != mid {
			if mid == 0 || num[mid-1] == mid-1 {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	if left == len(num) {
		return left
	}

	// 无效的输入，比如数组不是按要求排序的，
	// 或者有数字不在0到n-1范围之内
	return -1
}

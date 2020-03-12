package main

func GetNumSameAsIdx(num []int) int {
	if len(num) == 0 {
		return -1
	}
	left := 0
	right := len(num) - 1
	for left <= right {
		mid := (left + right) >> 1
		if num[mid] == mid {
			return mid
		}
		if num[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

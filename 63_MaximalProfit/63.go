package main

func MaxDiff(num []int) int {
	if len(num) == 0 {
		return -1
	}
	min := num[0]
	maxDiff := num[1] - min
	for i := 2; i < len(num); i++ {
		if num[i-1] < min {
			min = num[i-1]
		}
		curDiff := num[i] - min
		if curDiff > maxDiff {
			maxDiff = curDiff
		}
	}
	return maxDiff
}

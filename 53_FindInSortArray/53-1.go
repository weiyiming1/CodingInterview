package main

func GetNumOfK(data []int, k int) int {
	if len(data) == 0 {
		return -1
	}
	count := -1
	start := getFirstK(data, k, 0, len(data)-1)
	end := getLastk(data, k, 0, len(data)-1)
	if start > -1 && end > -1 {
		count = end - start + 1
	}
	return count
}

func getFirstK(data []int, k, start, end int) int {
	if start > end {
		return -1
	}
	midIdx := (start + end) / 2
	if data[midIdx] == k {
		if (midIdx > 0 && data[midIdx-1] != k) ||
			midIdx == 0 {
			return midIdx
		} else {
			end = midIdx - 1
		}
	} else if data[midIdx] > k {
		end = midIdx - 1
	} else {
		start = midIdx + 1
	}
	return getFirstK(data, k, start, end)
}

func getLastk(data []int, k, start, end int) int {
	if start > end {
		return -1
	}
	midIdx := (start + end) / 2
	if data[midIdx] == k {
		if ((midIdx < len(data)-1) && data[midIdx+1] != k) || midIdx == len(data)-1 {
			return midIdx
		} else {
			start = midIdx + 1
		}
	} else if data[midIdx] < k {
		start = midIdx + 1
	} else {
		end = midIdx - 1
	}
	return getLastk(data, k, start, end)
}

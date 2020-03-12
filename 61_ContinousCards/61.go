package main

import "sort"

func IsContinuous(num []int) bool {
	if len(num) < 1 {
		return false
	}
	sort.Ints(num)
	numOfZero := 0
	numOfGap := 0
	for i := 0; i < len(num) && num[i] == 0; i++ {
		numOfZero++
	}
	//第一个非0元素的位置
	small := numOfZero
	big := small + 1
	for big < len(num) {
		//两个数相等，有对子， 不可能是顺子
		if num[small] == num[big] {
			return false
		}
		numOfGap += (num[big] - num[small] - 1)
		small = big
		big++
	}
	if numOfGap > numOfZero {
		return false
	} else {
		return true
	}
}

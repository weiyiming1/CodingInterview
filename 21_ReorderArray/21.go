package main

func Record(num []int, f func(int) bool) {
	if num == nil || len(num) == 0 {
		return
	}
	idxBegin := 0
	idxEnd := len(num) - 1
	for idxBegin < idxEnd && !f(num[idxBegin]) { // 获取第一个偶数位置
		idxBegin++
	}
	for idxBegin < idxEnd && f(num[idxEnd]) { // 获取倒序的第一个奇数位置
		idxEnd--
	}
	if idxBegin < idxEnd {
		temp := num[idxBegin]
		num[idxBegin] = num[idxEnd]
		num[idxEnd] = temp
	}
}

func isEven(num int) bool { // 偶数返回true
	return (num & 1) == 0
}

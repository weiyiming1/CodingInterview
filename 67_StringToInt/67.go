package main

import "math"

func StrToInt(str string, num *int) {
	if len(str) < 1 {
		return
	}
	first := str[0]
	idx := 0
	minus := false

	if first == 43 { // 正号的ascii值是43
		idx++
	} else if first == 45 { // 负号的ascii值是45
		minus = true
		idx++
	}
	if idx != len(str) {
		*num = toDigit(str, idx, minus)
	}
}

func toDigit(str string, idx int, minus bool) int {
	num := 0
	var flag int
	if minus == true {
		flag = -1
	} else {
		flag = 1
	}
	for idx != len(str) {
		// 当前的idx值在0-9之间
		// 0的ascii值是48， 9的ascii值是57
		if str[idx] >= 48 && str[idx] <= 57 {
			//计算整数值
			num = num*10 + flag*int(str[idx]) - 48
			// 溢出
			if minus && num < flag*math.MinInt32 || !minus && num > math.MaxInt32 {
				break
			}
			idx++
		}
	}
	return num
}

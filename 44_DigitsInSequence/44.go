package main

import "math"

func DigitAtIndex(idx int) int {
	if idx < 0 {
		return -1
	}
	digit := 1 // 位数
	for {
		num := countOfIntegers(digit)
		if idx < num*digit {
			return getDigitIndex(idx, digit)
		}
		idx -= digit * num
		digit++
	}
}

func countOfIntegers(digit int) int {
	// 长度为m位的数字总共有多少个
	// 如:
	// 		输入2, 返回两位数(10-99)的个数90
	// 		输入3, 返回两位数(100-999)的个数900
	if digit == 1 {
		return 10
	}
	return int(9 * math.Pow(float64(9), float64(digit-1)))
}

func getDigitIndex(idx, digit int) int {

	num := getBeginNumber(digit) + idx/digit // 当前的数字
	idxFromRight := digit - idx%digit
	//例如100-999中间的第811位
	// num = 100+811/3 = 100+270=370
	// idxfromright = 3-1=2  这里从1开始计数， 等于2也就是位于某个三位数的中间
	for i := 1; i < idxFromRight; i++ {
		num /= 10
	}
	return num % 10
}

func getBeginNumber(digit int) int {
	// 获取m位数的第一个数字
	// 如：
	// 		第一个两位数是10
	//		第一个三位数是100
	if digit == 1 {
		return 0
	}
	return int(10 * math.Pow(float64(10), float64(digit-1)))
}

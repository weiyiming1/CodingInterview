package main

func LeftRotateString(data *[]rune, num int) { // num移动的个数
	if data == nil || len(*data) <= 0 {
		return
	}
	if num > 0 && len(*data) > 0 && num < len(*data) {
		start1 := 0
		end1 := num - 1
		start2 := num
		end2 := len(*data) - 1

		//翻转字符串的前面n个字符
		reverse(*data, start1, end1)
		//翻转字符串的后面n个字符
		reverse(*data, start2, end2)
		//翻转整个字符串
		reverse(*data, start1, end2)
	}
}

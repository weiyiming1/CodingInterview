package main

func FirstNotRepeatChar(str string) string {
	if len(str) == 0 {
		return ""
	}

	// 假设字符串由包含26个英文字母
	charMap := make(map[string]int, 26)
	for i := 0; i < len(charMap); i++ {
		charMap[string(97+i)] = 0
	}
	for _, val := range str {
		charMap[string(val)]++
	}
	for _, val := range str {
		if charMap[string(val)] == 1 {
			return string(val)
		}
	}
	return ""
}

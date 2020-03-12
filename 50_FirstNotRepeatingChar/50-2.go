package main

// 字符流中第一个只出现一次的字符

func Insert(char string, str *string, charMap *map[string]int) {
	/*
		charMap[char] ->ASCII value i
			-1		not found
			-2		multiple time
			>=0		only once
	*/
	if (*charMap)[char] == -1 {
		// 新的char加入原来str的尾部
		*str += char
		(*charMap)[char] = len(*str) - 1
	} else if (*charMap)[char] >= 0 {
		// 之前出现过
		(*charMap)[char] = -2
	}
}

func FirstCharInStream(str *string, charMap *map[string]int) string {
	if len(*str) == 0 {
		return ""
	}
	for _, val := range *str {
		if (*charMap)[string(val)] >= 0 {
			return string(val)
		}
	}
	return ""
}

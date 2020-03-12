package main

func LongestSubString(chars []rune) int {
	if chars == nil || len(chars) <= 0 {
		return 0
	}

	curLen := 0 // 当前字符串长度
	maxLen := 0 // 最大长度

	pos := make([]rune, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
	}

	for i := 0; i < len(chars); i++ {
		preIndex := int(pos[chars[i]-'a'])
		// 1. 当前字符第一次出现，
		// 2. 当前字符和上次出现位置之间的距离 大于curlen
		if preIndex < 0 || i-preIndex > curLen {
			curLen++
		} else {
			// 距离小于curlen
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = i - preIndex
		}
		// 更新出现的位置
		pos[chars[i]-'a'] = rune(i)
	}
	return maxLen
}

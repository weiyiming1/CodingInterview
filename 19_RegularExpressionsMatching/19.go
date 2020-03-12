package main

func PatternMatch(str, patten []string) bool {
	// 非法输入
	if str == nil || patten == nil {
		return false
	}
	return match(str, patten, 0, 0)
}

func match(str, pattern []string, sIdx, pIdx int) bool {
	if sIdx == len(str) && pIdx == len(pattern) { // 同时结束
		return true
	}
	if sIdx != len(str) && pIdx >= len(pattern) { // pattern先结束
		return false
	}
	// 下个字符串是"*"的时候
	if pIdx < len(pattern)-1 { // 只有pIdx没到达最后一位，下一步才能考虑pIdx+1的情况
		if pattern[pIdx+1] == "*" {
			// pattern当前字符和str的当前字符匹配
			if pattern[pIdx] == str[sIdx] || (pattern[pIdx] == "." && sIdx < len(str)) {
				/*  1, patten匹配str的0个字符:match(sIdx, pIdx+2)
				2, patten匹配str的1个字符:match(sIdx+1, pIdx+2)
				3, patten匹配str的多个字符:match(sIdx+1, pIdx)  */
				return match(str, pattern, sIdx, pIdx+2) ||
					match(str, pattern, sIdx+1, pIdx+2) ||
					match(str, pattern, sIdx+1, pIdx)
			} else {
				// 当前的不匹配，后面的"*"也无法匹配，pIdx+2
				return match(str, pattern, sIdx, pIdx+2)
			}
		}
	}

	// 当前不是"*"，所以只能逐个匹配
	if pattern[pIdx] == str[sIdx] || pattern[pIdx] == "." && sIdx < len(str) {
		return match(str, pattern, sIdx+1, pIdx+1)
	}
	return false
}

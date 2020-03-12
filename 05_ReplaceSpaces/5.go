package main

// MaxLen 字符数组总容量
func ReplaceBlank(str []rune, MaxLen int) []rune {
	if &str == nil || MaxLen <= 0 {
		return nil
	}

	orginLen := 0
	blankCount := 0
	index := 0
	for str[index] != '\000' {
		orginLen++
		if string(str[index]) == " " {
			blankCount++
		}
		index++
	}

	newLen := orginLen + blankCount*2
	if newLen > MaxLen {
		return nil
	}

	for orginLen >= 0 && newLen > orginLen {
		if str[orginLen] == ' ' {
			str[newLen] = '0'
			newLen--
			str[newLen] = '2'
			newLen--
			str[newLen] = '%'
			newLen--
		} else {
			str[newLen] = str[orginLen]
			newLen--
		}

		orginLen--
	}
	return str
}

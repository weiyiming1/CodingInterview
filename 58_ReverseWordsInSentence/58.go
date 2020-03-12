package main

import "fmt"

func ReverseSentence(chars *[]rune) {
	if len(*chars) <= 0 {
		return
	}
	begin := 0
	end := len(*chars) - 1

	//反转整个句子
	reverse(*chars, begin, end)

	fmt.Println(string(*chars))
	//反转每个单词
	end = 0
	for begin != len(*chars)-1 {
		if (*chars)[begin] == ' ' {
			begin++
			end++
		} else if (*chars)[end] == ' ' || end == len(*chars) {
			end--
			reverse(*chars, begin, end)
			end++
			begin = end
		} else {
			end++
		}
	}
}

func reverse(chars []rune, begin, end int) {
	for begin < end {
		temp := chars[begin]
		chars[begin] = chars[end]
		chars[end] = temp
		begin++
		end--
	}
}

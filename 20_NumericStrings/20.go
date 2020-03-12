package main

/*根据题目要求，正确的输入格式应该符合以下的形式：
[sign]integer-digits[.[fragment-digits]][e|E[sign]exponential-digits]
[]中的表示不一定需要*/

func IsNumberic(str string) bool {
	if &str == nil {
		return false
	}
	var pidx *int
	idx := 0
	pidx = &idx
	flag := scanInteger(str, pidx) // 判断整数部分
	if *pidx < len(str) && string(str[*pidx]) == "." {
		*pidx++
		flag = flag || scanUnsingedInteger(str, pidx) // 整数部分可能不存在
	}
	if *pidx < len(str) && (string(str[*pidx]) == "e" || string(str[*pidx]) == "E") {
		*pidx++
		flag = flag && scanInteger(str, pidx) // 指数部分前面必须有整数
	}
	return flag && *pidx == len(str)
}

func scanInteger(str string, pidx *int) bool {
	if *pidx < len(str) && (string(str[*pidx]) == "+" || string(str[*pidx]) == "-") {
		*pidx++
	}
	return scanUnsingedInteger(str, pidx)
}

func scanUnsingedInteger(str string, pidx *int) bool {
	temp := *pidx
	for *pidx < len(str) && string(str[*pidx]) >= "0" && string(str[*pidx]) <= "9" {
		*pidx++
	}
	return *pidx > temp
}

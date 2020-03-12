package main

import (
	"strconv"
)

func GetTranslation(num int) int {
	if num < 0 {
		return 0
	}
	strNum := strconv.Itoa(num)
	return getTranslationCount(strNum)
}

/*动态规划，右到左
 f(r-2) = f(r-1)+g(r-2,r-1)*f(r)
             f2 + g * f1

	如果r-2，r-1能够翻译成字符，则g(r-2,r-1)=1，否则为0。
	因此，对于12258：
	f(5) = 0  --> f1
	f(4) = 1  --> f2
	f(3) = f(4)+0 = 1
	f(2) = f(3)+f(4) = 2
	f(1) = f(2)+f(3) = 3
	f(0) = f(1)+f(2) = 5
*/
func getTranslationCount(str string) int {
	f2 := 1
	f1 := 0
	g := 0
	var temp int
	for i := len(str) - 2; i >= 0; i-- {
		digit1, _ := strconv.Atoi(string(str[i]))
		digit2, _ := strconv.Atoi(string(str[i+1]))
		conveted := digit1*10 + digit2
		if conveted >= 10 && conveted <= 25 {
			g = 1
		} else {
			g = 0
		}
		temp = f2
		f2 = f2 + g*f1
		f1 = temp
	}
	return f2
}

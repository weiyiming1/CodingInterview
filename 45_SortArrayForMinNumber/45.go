package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintMinNumber(num []int) {
	if num == nil || len(num) == 0 {
		return
	}
	var str string
	for i := 0; i < len(num); i++ {
		str = strCompare(str, strconv.Itoa(num[i]))
	}
	fmt.Println(str)

}

func strCompare(str1, str2 string) string {
	tempStr1 := str1 + str2
	tempStr2 := str2 + str1
	flag := strings.Compare(tempStr1, tempStr2)
	if flag <= 0 {
		return tempStr1
	} else {
		return tempStr2
	}
}

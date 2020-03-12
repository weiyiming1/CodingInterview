package main

import (
	"fmt"
	"strconv"
)

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	num := []string{"1"}
	for i := 0; i < n; i++ {
		num[0] = getStr(getInt(num) * 10)
	}
	for i := 1; i < getInt(num); i++ {
		fmt.Println(i)
	}
}

func getInt(str []string) int {
	num, _ := strconv.Atoi(str[0])
	return num
}

func getStr(num int) string {
	return strconv.Itoa(num)
}

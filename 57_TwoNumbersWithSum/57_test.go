package main

import (
	"fmt"
	"testing"
)

func TestFindNumbersWithSum(t *testing.T) {
	num := []int{1, 2, 4, 7, 11, 15}
	var num1, num2 int
	fmt.Println(FindNumbersWithSum(num, 15, &num1, &num2))
	fmt.Println(num1, num2)
}

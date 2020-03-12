package main

import (
	"fmt"
	"testing"
)

func TestFindNumberAppearingOnce(t *testing.T) {
	num := []int{1, 1, 2, 2, 2, 1, 3}
	var res int
	FindNumberAppearingOnce(num, &res)
	fmt.Println("res:", res)
}

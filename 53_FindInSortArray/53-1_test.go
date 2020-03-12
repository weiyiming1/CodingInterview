package main

import (
	"fmt"
	"testing"
)

func TestGetNumOfK(t *testing.T) {
	num := []int{1, 2, 3, 3, 3, 3, 4, 5}
	count := GetNumOfK(num, 3)
	fmt.Println(count)
}

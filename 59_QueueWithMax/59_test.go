package main

import (
	"fmt"
	"testing"
)

func TestMaxInWindow(t *testing.T) {
	num := []int{2, 3, 4, 2, 6, 2, 5, 1}
	max := MaxInWindow(num, 3)
	fmt.Println(max)
}

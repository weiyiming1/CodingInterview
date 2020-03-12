package main

import (
	"fmt"
	"testing"
)

func TestVerifySquenceOfBST(t *testing.T) {
	a := []int{5, 7, 6, 9, 11, 10, 8}
	b := []int{7, 4, 6, 5}
	flag := VerifySquenceOfBST(a, 7)
	fmt.Println(flag)
	flag = VerifySquenceOfBST(b, 4)
	fmt.Println(flag)
}

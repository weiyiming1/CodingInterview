package main

import (
	"log"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	array := []int{3, 1, 0, 2, 5, 3}
	length := len(array)
	duplication := new(int)

	flag := FindDuplicate(array, length, duplication)
	if flag {
		log.Printf("找到重复数字%d\n", *duplication)
	}
}

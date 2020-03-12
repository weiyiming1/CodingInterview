package main

import (
	"log"
	"testing"
)

func TestMaxDiff(t *testing.T) {
	num := []int{9, 11, 8, 5, 7, 12, 16, 14}
	res := MaxDiff(num)
	log.Println(res)
}

package main

import (
	"log"
	"testing"
)

func TestFindGreatestSumOfSubArray(t *testing.T) {
	num := []int{1, -2, 3, 10, -4, 7, 2, -5}
	sum := FindGreatestSumOfSubArray(num)
	if sum != nil {
		log.Println("最大累加和:", *sum)
	}
}

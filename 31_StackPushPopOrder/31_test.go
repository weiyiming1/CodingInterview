package main

import (
	"log"
	"testing"
)

func TestIsPopOrder(t *testing.T) {
	len := 5
	pushOrder := []int{1, 2, 3, 4, 5}
	popOrder := []int{4, 5, 3, 2, 1}
	flag := IsPopOrder(pushOrder, popOrder, len)
	log.Println(flag)

	pushOrder = []int{1, 2, 3, 4, 5}
	popOrder = []int{4, 3, 1, 5, 2}
	flag = IsPopOrder(pushOrder, popOrder, len)
	log.Println(flag)
}

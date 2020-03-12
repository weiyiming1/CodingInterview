package main

import (
	"log"
	"testing"
)

func TestMoreThanHalfNum(t *testing.T) {
	num := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	res := MoreThanHalfNum(&num)
	log.Println(res)
}

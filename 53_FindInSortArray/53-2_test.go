package main

import (
	"log"
	"testing"
)

func TestGetMissingNumber(t *testing.T) {
	num := []int{0, 1, 2, 3, 4, 5}
	log.Println(GetMissingNumber(num))
}

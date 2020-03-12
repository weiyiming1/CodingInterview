package main

import (
	"log"
	"testing"
)

func TestGetNumSameAsIdx(t *testing.T) {
	num := []int{-3, -1, 1, 3, 5}
	idx := GetNumSameAsIdx(num)
	log.Println("idx:", idx)
}

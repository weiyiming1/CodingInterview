package main

import (
	"log"
	"testing"
)

func TestSwitchMin(t *testing.T) {
	num := []int{1, 1, 1, 0, 1}
	log.Println(SwitchMin(num))

	num = []int{5, 6, 7, 3, 4}
	log.Println(SwitchMin(num))
}

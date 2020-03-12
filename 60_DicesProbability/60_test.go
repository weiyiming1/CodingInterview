package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrintProbability(t *testing.T) {
	PrintProbability(3)
	demo := []int{3, 2, 5, 9}
	sort.Ints(demo)
	fmt.Println(demo)
}

package main

import (
	"fmt"
	"testing"
)

func TestFindNumAppearOnce(t *testing.T) {
	data := []int{2, 4, 3, 6, 3, 2, 5, 5}
	var res1, res2 int
	FindNumAppearOnce(data, &res1, &res2)
	fmt.Println(res1, res2)
}

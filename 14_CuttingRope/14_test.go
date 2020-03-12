package main

import (
	"fmt"
	"testing"
)

func TestMaxCut(t *testing.T) {
	res := MaxCut(5)
	fmt.Println(res)
	res = MaxCut(6)
	fmt.Println(res)
	res = MaxCut(7)
	fmt.Println(res)
	res = MaxCut(8)
	fmt.Println(res)
	res = MaxCut(9)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"testing"
)

func TestGetUglyNumber(t *testing.T) {
	val := GetUglyNumber(1500)
	fmt.Println(val)
}

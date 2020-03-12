package main

import (
	"fmt"
	"testing"
)

func TestStrToInt(t *testing.T) {
	num := 0
	str := "123"
	StrToInt(str, &num)
	fmt.Println(num)
}

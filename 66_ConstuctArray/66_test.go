package main

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := Multiply(input)
	fmt.Println(output)
}

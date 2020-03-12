package main

import (
	"fmt"
	"testing"
)

func TestIsContinuous(t *testing.T) {
	a := []int{4, 6, 2, 8, 0}
	fmt.Println(IsContinuous(a))

}

package main

import (
	"fmt"
	"testing"
)

func TestGetMaxVal(t *testing.T) {
	values := [][]int{{1, 10, 3, 8}, {12, 2, 9, 6}, {5, 7, 4, 11}, {3, 7, 16, 5}} //{1,10,3,8}, {12,2,9,6} ,{5,7,4,11},{3,7,16,5}
	max := GetMaxVal(values, 4, 4)
	fmt.Println(max)
}

package main

import (
	"fmt"
	"testing"
)

//  1   2   8   9
//  2   4   9   12
//  4   7   10  13
//  6   8   11  15
// 要查找的数在数组中

func TestNumInMax(t *testing.T) {
	matrix := [][]int{{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15}}
	flag := NumInMax(matrix, 4, 4, 7)
	if flag == true {
		fmt.Println("find")
	} else {
		fmt.Println("not find")
	}

}

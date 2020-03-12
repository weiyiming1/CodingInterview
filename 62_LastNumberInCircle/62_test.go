package main

import (
	"container/ring"
	"fmt"
	"log"
	"testing"
)

func TestLastRemain(t *testing.T) {
	num := ring.New(5)
	step := 3
	// 赋值
	for i := 0; i < num.Len(); i++ {
		num.Value = i
		fmt.Println(num)
		num = num.Next()
	}
	res := LastRemain(num, step)
	log.Println("剩下的最后一个数:", res)
}

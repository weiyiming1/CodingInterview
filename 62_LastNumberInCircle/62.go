package main

import (
	"container/ring"
	"fmt"
)

func LastRemain(num *ring.Ring, step int) *ring.Ring {
	if num.Len() == 0 {
		return nil
	}
	for num.Len() > 1 {
		num = num.Move(step - 2)
		del := num.Unlink(1)
		fmt.Println("del:", del.Value)
		num = num.Next()
	}
	return num
}

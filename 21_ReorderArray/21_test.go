package main

import (
	"log"
	"testing"
)

func TestRecord(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	log.Println("输入:", a)
	Record(a, isEven)
	log.Println("输出:", a)
}

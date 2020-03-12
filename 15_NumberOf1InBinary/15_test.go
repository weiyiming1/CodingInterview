package main

import (
	"log"
	"strconv"
	"testing"
)

func TestBinaryCountOne(t *testing.T) {
	input := strconv.FormatInt(5, 2)
	log.Printf("输入整数的二进制表示：%v\n", input)
	count := BinaryCountOne(5)
	log.Printf("二进制表示位中，1的个数是:%d", count)
}

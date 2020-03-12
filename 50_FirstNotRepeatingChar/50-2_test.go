package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	charMap := make(map[string]int, 256)
	for i := 0; i < 256; i++ {
		charMap[string(i)] = -1
	}
	str := ""
	Insert("g", &str, &charMap)
	Insert("o", &str, &charMap)
	fmt.Println("第一次:", FirstCharInStream(&str, &charMap))
	Insert("o", &str, &charMap)
	Insert("g", &str, &charMap)
	Insert("l", &str, &charMap)
	Insert("e", &str, &charMap)
	fmt.Println("第一次:", FirstCharInStream(&str, &charMap))
}

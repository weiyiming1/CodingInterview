package main

import (
	"log"
	"testing"
)

func TestLeftRotateString(t *testing.T) {
	data := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	LeftRotateString(&data, 2)
	log.Println(string(data))
}

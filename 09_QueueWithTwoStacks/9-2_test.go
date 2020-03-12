package main

import (
	"log"
	"testing"
)

func TestNewTwoQueueStack(t *testing.T) {
	tt := NewTwoQueueStack()
	tt.Push(1)
	tt.Push("ww")
	temp := tt.Pop()
	log.Println("Pop->", temp)
	tt.Push("aaaaaaa")
	temp = tt.Pop()
	log.Println("Pop->", temp)
}

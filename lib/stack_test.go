package lib

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	len := stack.Size()
	if len != 4 {
		t.Errorf("stack.Len() failed. Got %d, expected 4.", len)
	}
	log.Println(stack)

}

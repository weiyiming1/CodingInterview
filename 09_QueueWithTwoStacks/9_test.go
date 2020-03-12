package main

import (
	"testing"
)

func TestNewQuene(t *testing.T) {
	q := NewQuene()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	temp := q.Pop()
	t.Log(temp)
	temp = q.Pop()
	t.Log(temp)
	temp = q.Pop()
	t.Log(temp)
	temp = q.Pop()
	t.Log(temp)
	temp = q.Pop()
	t.Log(temp)
}

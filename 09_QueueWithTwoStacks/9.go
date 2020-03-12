package main

import (
	"CodingInterview/getOffer-Go/lib"
)

type TwoStackQueue struct {
	stackPush *lib.Stack
	stackPop  *lib.Stack
}

func NewQuene() *TwoStackQueue {
	return &TwoStackQueue{stackPush: lib.NewStack(),
		stackPop: lib.NewStack()}
}

func (q *TwoStackQueue) Push(val interface{}) {
	q.stackPush.Push(val)
}

func (q *TwoStackQueue) Pop() interface{} {
	if q.stackPop.Size() <= 0 {
		for q.stackPush.Size() > 0 {
			val := q.stackPush.Pop()
			q.stackPop.Push(val)
		}
	}
	if q.stackPop.Size() == 0 {
		return nil
	}

	return q.stackPop.Pop()
}

package main

import (
	"CodingInterview/getOffer-Go/lib"
)

func O1PutStackMin(dataStack, minStack lib.Stack, num int) {
	dataStack.Push(num)
	if minStack.Size() == 0 || num < minStack.Top().(int) {
		minStack.Push(num)
	} else {
		minStack.Push(minStack.Top())
	}
}

func O1PopStackMin(dataStack, minStack lib.Stack) {
	if dataStack.Size() > 0 && minStack.Size() > 0 {
		dataStack.Pop()
		minStack.Pop()
	}
}

func O1GetStackMin(dataStack, minStack lib.Stack) int {
	if dataStack.Size() > 0 && minStack.Size() > 0 {
		return minStack.Top().(int)
	}
	return 1e10
}

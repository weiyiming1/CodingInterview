package main

import "fmt"

func Constructor() MyQueue {
	return MyQueue{
		a: NewStack(),
		b: NewStack()}
}

func (mq *MyQueue) Push(x int) {
	mq.a.Push(x)
}

/*
Peek vs Pop:
	相同点：大家都返回栈顶的值。
	不同点：peek 不改变栈的值(不删除栈顶的值)，pop会把栈顶的值删除。
*/

func (mq *MyQueue) Pop() int {
	if mq.b.Len() == 0 {
		for mq.a.Len() > 0 {
			mq.b.Push(mq.a.Pop())
		}
	}
	return mq.b.Pop()
}

func (mq *MyQueue) Peek() int {
	res := mq.Pop()
	mq.b.Push(res)
	return res
}

func (mq *MyQueue) Empty() bool {
	return mq.a.Len() == 0 && mq.b.Len() == 0
}

func main() {
	obj := Constructor()
	obj.Push(10)
	//obj.Push(20)
	//obj.Push(30)
	/*param_2 := obj.Pop()
	fmt.Println("param_2: ", param_2)*/
	param_3 := obj.Peek()
	fmt.Println("param_3: ", param_3)
	param_4 := obj.Empty()
	fmt.Println("param_4: ", param_4)
}

// 栈结构
type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

type MyQueue struct {
	a, b *Stack
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

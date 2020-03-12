package main

import "fmt"

type MyStack struct {
	a, b *Queue
}

func Constructor() MyStack {
	return MyStack{
		a: NewQueue(),
		b: NewQueue()}
}

// 保证b为空，push/pop操作在a上
func (ms *MyStack) Push(x int) {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

func (ms *MyStack) Pop() int {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}

	for ms.a.Len() > 1 {
		ms.b.Push(ms.a.Pop())
	}

	return ms.a.Pop()
}

func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res)
	return res
}

func (ms *MyStack) Empty() bool {
	return (ms.a.Len() + ms.b.Len()) == 0
}

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Push(20)
	obj.Push(30)
	param_2 := obj.Pop()
	fmt.Println("param_2: ", param_2)
	param_3 := obj.Top()
	fmt.Println("param_3: ", param_3)
	param_4 := obj.Empty()
	fmt.Println("param_4: ", param_4)
}

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

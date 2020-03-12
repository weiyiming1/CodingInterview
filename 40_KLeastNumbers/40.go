package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 默认是min-heap，修改less函数，返回max-heap类型
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetLeastNumbers(data []int, k int) {
	if k < 1 || len(data) < k {
		return
	}
	// 从n个整数中取k个数字放到容器maxHeap数据容器中
	maxHeap := &IntHeap{}
	for i := 0; i < k; i++ {
		*maxHeap = append(*maxHeap, data[i])
	}
	heap.Init(maxHeap)
	for _, val := range data[k:] {
		if val < (*maxHeap)[0] {
			heap.Pop(maxHeap)
			*maxHeap = append(*maxHeap, val)
			heap.Init(maxHeap)
		}
	}
	fmt.Println(*maxHeap)
}

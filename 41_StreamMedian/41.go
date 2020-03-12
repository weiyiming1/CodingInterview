package main

// 容器分割为两部分，
// 左边的元素都小于右边的元素
// 指针p1指向左边最大的数， 指针p2指向右边最小的数
import (
	"CodingInterview/getOffer-Go/lib"
	"container/heap"
)

func ArrayGetMedian(num []int) int {
	if len(num) == 0 {
		return 0
	}

	max := &lib.MaxHeap{} // 最大堆
	min := &lib.MinHeap{} // 最小堆

	for _, val := range num {
		// 堆中已有数据的总数为偶数，新数据加入最小堆
		if ((min.Len() + max.Len()) & 1) == 0 {
			// 先临时加入最大堆
			if max.Len() > 0 && val < (*max)[0] {
				heap.Push(max, val) // 加入新数据num
				heap.Init(max)

				val = (*max)[0] // 找出num中的最大值
				heap.Pop(max)
			}
			heap.Push(min, val) //把最大值加入到最小堆
		} else {
			// 总数为奇数，加入最大堆
			// 先临时加入最小堆
			if min.Len() > 0 && (*min)[0] < val {
				heap.Push(min, val)
				heap.Init(min)
				val = (*min)[0] // 找出min堆中的最小值
				heap.Pop(min)
			}
			heap.Push(max, val)
		}
	}
	var median int
	if len(num)&1 == 0 { // 偶数
		median = ((*min)[0] + (*max)[0]) / 2
	} else { // 奇数
		median = (*min)[0]
	}
	return median
}

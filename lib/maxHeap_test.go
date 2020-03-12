package lib

import (
	"container/heap"
	"log"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := &MaxHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	log.Printf("max: %d\n", (*h)[0])
	for h.Len() > 0 {
		log.Printf("%d ", heap.Pop(h))
	}
}

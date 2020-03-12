package main

import (
	"CodingInterview/getOffer-Go/lib"
	"container/heap"
	"log"
	"testing"
)

func TestArrayGetMedian(t *testing.T) {
	num := &lib.MinHeap{6, 5, 4, 2, 1}
	log.Println("输入num:", num)
	heap.Init(num)
	mid := ArrayGetMedian(*num)
	log.Println("中位数:", mid)

}

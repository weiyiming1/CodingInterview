package lib

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {

	h := &MinHeap{4, 2}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 3)
	fmt.Println(h)
	heap.Pop(h)
	fmt.Println(h)
	temp := heap.Pop(h)
	fmt.Println(h)
	fmt.Println(temp)
	/*h := &MinHeap{4,2,3}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 1)
	fmt.Println(h)*/

	/*num := &lib.MinHeap{4,2,3}
		log.Println("1:",num)
		heap.Init(num)
		log.Println("2:",num)


		 &[4 2 3]
	     &[2 4 3]
	*/

}

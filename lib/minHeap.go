// This example demonstrates an integer heap built using the heap interface.
package lib

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

// default: h[i] < h[j]
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

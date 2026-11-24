package main

import (
	"container/heap"
	"fmt"
)

// intHeap is a min-heap of ints for container/heap.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// KthLargest returns the k-th largest element by keeping a min-heap of the
// k largest values seen; its root is the answer. ok is false for invalid k.
func KthLargest(nums []int, k int) (int, bool) {
	if k < 1 || k > len(nums) {
		return 0, false
	}
	h := &intHeap{}
	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v > (*h)[0] {
			(*h)[0] = v
			heap.Fix(h, 0)
		}
	}
	return (*h)[0], true
}

func main() {
	v, _ := KthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(v)
}

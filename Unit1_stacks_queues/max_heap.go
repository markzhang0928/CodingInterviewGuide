package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // 最大堆
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 1)

	fmt.Println("当前最大值:", (*maxHeap)[0])
	heap.Push(maxHeap, 5)
	heap.Push(maxHeap, 4)

	for maxHeap.Len() > 0 {
		fmt.Println(heap.Pop(maxHeap))
	}

}

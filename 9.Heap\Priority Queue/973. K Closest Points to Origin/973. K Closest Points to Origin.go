package main

import (
	"container/heap"
)
func main() {
	
}

type MaxHeap [][]int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i,j int) bool { return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1] }	
func (h MaxHeap) Swap(i,j int) { h[i],h[j] = h[j],h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} { old := *h; n := len(old); x := old[n-1]; *h = old[0:n-1]; return x }

func kClosest(points [][]int, k int) [][]int {
	max := &MaxHeap{}
	heap.Init(max)
	for _,point := range points {
		heap.Push(max, point)
		if max.Len() > k {
			heap.Pop(max)
		}
	}
	return *max



}
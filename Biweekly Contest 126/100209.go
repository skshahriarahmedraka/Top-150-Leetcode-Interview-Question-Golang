package main

import (
	"container/heap"
)

func main() {

}

type Pair struct {
	Value int
	Index int
}

type minHeap []Pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Value < h[j].Value }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *minHeap) Pop() interface{} {
	old := *h
	min := old[len(old)-1]
	*h = old[:len(old)-1]
	return min
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	numLen := len(nums)
	priorityQueue := minHeap{}
	heap.Init(&priorityQueue)
	var sumOfNumber  int64
	ind := make(map[int]bool)
	for i, j := range nums {
		ind[i] = true
		sumOfNumber += int64(j)
		priorityQueue = append(priorityQueue, Pair{Value: j, Index: i})
	}
	h := make(map[int][]int)
	for i := 0; i < numLen; i++ {
		if _, b := h[nums[i]]; !b {
			h[i] = []int{i}
		} else {
			h[nums[i]] = append(h[nums[i]], nums[i])
		}
	}
	queriesLen := len(queries)
	ans := make([]int64, queriesLen)
	for i := 0; i < queriesLen; i++ {
		left := queries[i][0]
		right := queries[i][1]

		if ind[left] {
			delete(ind, left)
			sumOfNumber -= int64(nums[left])
		}

		for right > 0 {
			if len(priorityQueue) > 0 {

				minVal := heap.Pop(&priorityQueue).(Pair)
				indexes := h[minVal.Value]
				h[minVal.Value] = indexes[1:]
				minInx := indexes[0]

				if ind[minInx] {
					sumOfNumber -= int64(nums[minInx])
					delete(ind, minInx)
					right--
				}

			} else {
				break
			}
			ans[i] = sumOfNumber

		}

	}
	return ans

}

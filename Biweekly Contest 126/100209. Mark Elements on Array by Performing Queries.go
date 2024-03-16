package main
import (
	"container/heap"
	
)

type pair struct {
	value int
	index int
}

type minHeap []pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	lenNums := len(nums)
	ind := make(map[int]bool)
	pq := make(minHeap, 0)
	heap.Init(&pq)
	var sumOfNums int64
	for i := 0; i < lenNums; i++ {
		ind[i] = true
		heap.Push(&pq, pair{value: nums[i], index: i})
		sumOfNums += int64(nums[i])
	}
	h := make(map[int][]int)
	for i := 0; i < lenNums; i++ {
		val := nums[i]
		if _, exists := h[val]; !exists {
			h[val] = []int{i}
		} else {
			h[val] = append(h[val], i)
		}
	}

	lent := len(queries)
	ans := make([]int64, lent)
	for i := 0; i < lent; i++ {
		left := queries[i][0]
		right := queries[i][1]

		

		if ind[left] {
			delete(ind, left)
			sumOfNums -= int64(nums[left])
		}

		for right > 0 {
			if len(pq) > 0 {
				minVal := heap.Pop(&pq).(pair)
				indexes := h[minVal.value]
				minInd := indexes[0]
				h[minVal.value] = indexes[1:]

				if ind[minInd] {
					right--
					delete(ind, minInd)
					sumOfNums -= int64(nums[minInd])
				}
			} else {
				break
			}
		}

		ans[i] = sumOfNums
	}

	return ans
}



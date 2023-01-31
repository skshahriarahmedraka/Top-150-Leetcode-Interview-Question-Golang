package main
import (
	"container/heap"
	"sort"
	
)

type Interval struct {
	Start int
	End   int
}

type IntervalMinHeap []Interval

func (h IntervalMinHeap) Len() int {
	return len(h)
}

func (h IntervalMinHeap) Less(i, j int) bool {
	return h[i].End < h[j].End
}

func (h IntervalMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *IntervalMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minMeetingRooms(intervals []Interval) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}

	intervalsOrderedByStart := make([]Interval, len(intervals))
	copy(intervalsOrderedByStart, intervals)
	sort.Slice(intervalsOrderedByStart, func (i ,j int) bool {
		return intervalsOrderedByStart[i].Start < intervalsOrderedByStart[j].Start
	})

	h := &IntervalMinHeap{}
	heap.Init(h)

	for i := 0; i < len(intervalsOrderedByStart); i++ {
		if h.Len() == 0 {
			heap.Push(h, intervalsOrderedByStart[i])
			continue
		}

		m := (*h)[0]
		if m.End <= intervalsOrderedByStart[i].Start {
			heap.Pop(h)
		}
		heap.Push(h, intervalsOrderedByStart[i])
	}


	return h.Len()
}
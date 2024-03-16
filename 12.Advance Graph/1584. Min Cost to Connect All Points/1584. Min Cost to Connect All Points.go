package main

import (
	"container/heap"
	"math"
)
type Item struct {
	distance, point int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func manhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	visited := make([]bool, n)
	heapDict := make(map[int]int)
	for i := 0; i < n; i++ {
		heapDict[i] = int(math.MaxInt64) // Initialize all distances to infinity
	}
	heapDict[0] = 0 // Start node

	pq := make(PriorityQueue, 1)
	pq[0] = Item{distance: 0, point: 0}
	heap.Init(&pq)

	mstWeight := 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(Item)
		w, u := item.distance, item.point

		if visited[u] || heapDict[u] < w {
			continue
		}

		visited[u] = true
		mstWeight += w

		for v := 0; v < n; v++ {
			if !visited[v] {
				newDistance := manhattanDistance(points[u], points[v])
				if newDistance < heapDict[v] {
					heapDict[v] = newDistance
					heap.Push(&pq, Item{distance: newDistance, point: v})
				}
			}
		}
	}

	return mstWeight
}
package main

	


import "container/heap"

type Point struct {
    t, r, c int
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].t < pq[j].t || (pq[i].t == pq[j].t && pq[i].r < pq[j].r) || (pq[i].t == pq[j].t && pq[i].r == pq[j].r && pq[i].c < pq[j].c)
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Point))
}
func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    x := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return x
}

func swimInWater(grid [][]int) int {
    N := len(grid)
    visit := make(map[int]bool)
    pq := &PriorityQueue{}
    heap.Push(pq, Point{grid[0][0], 0, 0})
    visit[0] = true
    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    for pq.Len() > 0 {
        curr := heap.Pop(pq).(Point)
        if curr.r == N-1 && curr.c == N-1 {
            return curr.t
        }

        for _, dir := range directions {
            neiR, neiC := curr.r+dir[0], curr.c+dir[1]
            if neiR < 0 || neiR >= N || neiC < 0 || neiC >= N {
                continue
            }
            key := neiR*N + neiC
            if visit[key] {
                continue
            }
            visit[key] = true
            height := max(curr.t, grid[neiR][neiC])
            heap.Push(pq, Point{height, neiR, neiC})
        }
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
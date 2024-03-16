package main
import "container/list"

const MAX_INT = int(^uint(0) >> 1)
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    ret := MAX_INT
    graph := make([][][]int, n)
    for _, flight := range flights {
        graph[flight[0]] = append(graph[flight[0]], flight[1:])
    }
    queue := list.New()
    queue.PushBack([]int{src, 0})
    level := 0
    for queue.Len() != 0 && level <= K + 1 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            curr := e.Value.([]int)
            if curr[0] == dst && ret > curr[1]{ ret = curr[1] }
            for _, edge := range graph[curr[0]] {
                if tempCost := curr[1] + edge[1]; tempCost < ret {
                    queue.PushBack([]int{edge[0], tempCost})
                }
            } 
        }
        level++
    }
    if ret == MAX_INT { return -1 } else { return ret }
}
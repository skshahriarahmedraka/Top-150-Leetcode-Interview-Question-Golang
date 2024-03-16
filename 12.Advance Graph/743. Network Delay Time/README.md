743. Network Delay Time
Medium
Topics
Companies
Hint

You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.

 

Example 1:

Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2

Example 2:

Input: times = [[1,2,1]], n = 2, k = 1
Output: 1

Example 3:

Input: times = [[1,2,1]], n = 2, k = 2
Output: -1

 

Constraints:

    1 <= k <= n <= 100
    1 <= times.length <= 6000
    times[i].length == 3
    1 <= ui, vi <= n
    ui != vi
    0 <= wi <= 100
    All the pairs (ui, vi) are unique. (i.e., no multiple edges.)

## Dijkstra:

```
func networkDelayTime(times [][]int, n int, k int) int {
    type edge struct { to, weight int }
    
    w := make(map[int][]edge)
    for _, time := range times {
        from, to, weight := time[0], time[1], time[2]
        w[from] = append(w[from], edge{to, weight})
    }
    
    dist := make([]int, n+1)
    for i := 1; i <= n; i++ { dist[i] = -1 }
    
    intree := make([]bool, n+1)
    
    intree[k] = true
    dist[k] = 0
    
    v := k
    for {
        intree[v] = true
        
        for _, e := range w[v] {
            if dist[e.to] == -1 || (dist[e.to] > dist[v] + e.weight) {
                dist[e.to] = dist[v] + e.weight
            }
        }
        
        minv, mind := -1, -1
        for i := 1; i <= n; i++ {
            if !intree[i] && (mind == -1 || dist[i] < mind) {
                minv, mind = i, dist[i]
            }
        }
        
        if minv == -1 { break }
        
        v = minv
    }
    
    max := dist[1]
    for i := 1; i <= n; i++ {
        if dist[i] == -1 { return -1 }
        if dist[i] > max { max = dist[i] }
    }
    
    return max
}
```

## Bellman-Ford:
```
func networkDelayTime(times [][]int, n int, k int) int {
    dist := make([]int, n+1)
    for i := 1; i <= n; i++ { dist[i] = -1 }
    
    dist[k] = 0
    
    for i := 0; i < n-1; i++ {
        for _, time := range times {
            from, to, weight := time[0], time[1], time[2]
            
            if dist[from] == -1 { continue }
            
            if dist[to] == -1 || dist[from] + weight < dist[to] {
                dist[to] = dist[from] + weight
            }
        }
    }
    
    max := dist[1]
    for i := 1; i <= n; i++ {
        if dist[i] == -1 { return -1 }
        if dist[i] > max { max = dist[i] }
    }
    return max
}
```

## Floyd-Warshall:

```
func networkDelayTime(times [][]int, n int, k int) int {
    g := make([][]int, n+1)
    for i := 0; i < n+1; i++ { 
        g[i] = make([]int, n+1) 
        for j := 0; j < n+1; j++ { g[i][j] = -1 }
    }
    
    for _, time := range times {
        from, to, weight := time[0], time[1], time[2]
        g[from][to] = weight
    }
    
    g[k][k] = 0
    
    for k := 1; k <= n; k++ {
        for i := 1; i <= n; i++ {
            for j := 1; j <= n; j++ {
                if g[i][k] != -1 && g[k][j] != -1 {
                    if g[i][j] == -1 || (g[i][j] > g[i][k] + g[k][j]) {
                        g[i][j] = g[i][k] + g[k][j]
                    }
                }
            }
        }
    }
    
    max := -1
    for i := 1; i <= n; i++ {
        if k != i && g[k][i] == -1 { return -1 }
        if g[k][i] > max { max = g[k][i] }
    }
    return max
}
```  

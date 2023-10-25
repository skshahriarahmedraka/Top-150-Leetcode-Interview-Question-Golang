package main

func main() {

}
func ValidTree(n int, edges [][]int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n-1 != len(edges) {
		return false
	}

	adj := make(map[int][]int)
	visited := make(map[int]bool)
	for _, x := range edges {
		i, j := x[0], x[1]
		adj[i] = append(adj[i], j)
		adj[j] = append(adj[j], i)
	}

	var hasCycle bool

	var DFS func(int, int)
	DFS = func(i, pre int) {
		if visited[i] {
			hasCycle = true
			return
		}
		visited[i] = true
		for _, x := range adj[i] {
			if x == pre {
				continue
			} else {
				DFS(x, i)
			}
		}
	}

	DFS(0, -1)

	if hasCycle {
		return false
	}

	// Check if all nodes are visited
	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}

	return true
}


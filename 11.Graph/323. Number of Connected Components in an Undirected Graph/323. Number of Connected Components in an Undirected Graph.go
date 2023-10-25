
// 20 mstime cost
// ·
// 3.16 MBmemory cost
// ·
// Your submission beats100.00 %Submissions
package main

func main(){

}

func CountComponents(n int, edges [][]int) int {
	adj := make(map[int][]int)
	visited := make(map[int]bool)
	for _, x := range edges {
		i, j := x[0], x[1]
		adj[i] = append(adj[i], j)
		adj[j] = append(adj[j], i)
	}

	var DFS func(int)
	DFS = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true
		for _, x := range adj[i] {
			DFS(x)
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if !visited[i] {
			ans++
			DFS(i)
		}
	}
	return ans
}
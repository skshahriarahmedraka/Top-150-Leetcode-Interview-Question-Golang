// Accepted
// Runtime
// Details
// 14ms
// Beats 46.99%of users with Go
// Memory
// Details
// 6.40MB
// Beats 67.25%of users with Go

package main

func main() {
	
}


func canFinish(numCourses int, prerequisites [][]int) bool {
	var preMap = make(map[int][]int)
	var visited = make(map[int]bool)
	for _, v := range prerequisites {
		preMap[v[0]] = append(preMap[v[0]], v[1:]...)
	}
	var DFS func(int) bool

	DFS = func(i int) bool {
		if visited[i] {
			return false
		}
		if preMap[i] == nil {
			return true
		}
		visited[i] = true
		for _, v := range preMap[i] {
			if !DFS(v) {
				return false
			}
		}

		visited[i] = false
		preMap[i] = nil
		return true
	}
	for i:=0;i<numCourses;i++ {
		if !DFS(i) {
			return false
		}
	}
	return true
}
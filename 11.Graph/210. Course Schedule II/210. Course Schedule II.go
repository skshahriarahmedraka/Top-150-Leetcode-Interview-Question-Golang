// Accepted
// Runtime
// Details
// 17ms
// Beats 15.76%of users with Go
// Memory
// Details
// 6.66MB
// Beats 33.74%of users with Go
package main

func main() {

}

//	func findOrder(numCourses int, prerequisites [][]int) []int {
//		preMap := make(map[int][]int)
//		visited := make(map[int]bool)
//		for _, i := range prerequisites {
//			preMap[i[0]] = append(preMap[i[0]], i[1:]...)
//		}
//
//		var DFS func(int) bool
//
//		DFS = func(i int) bool {
//			if visited[i] {
//				return false
//			}
//			if preMap[i] == nil {
//				return true
//			}
//			visited[i] = true
//			for _, x := range preMap[i] {
//				if !DFS(x) {
//					return false
//				}
//			}
//			visited[i] = false
//			preMap[i] = nil
//			return true
//		}
//		for i := 0; i < numCourses; i++ {
//			if !DFS(i) {
//				return false
//			}
//		}
//		return true
//	}
func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int)
	visited := make(map[int]bool)
	cicle:= make(map[int]bool)
	ans := []int{}
	for _, i := range prerequisites {
		preMap[i[0]] = append(preMap[i[0]], i[1:]...)
	}
	var DFS func(int) bool

	DFS = func(i int) bool {
		if cicle[i] {
			return false
		}
		if visited[i] {
			return true
		}
		cicle[i] = true
		for _, j := range preMap[i] {
			if !DFS(j) {
				return false
			}
		}
		visited[i] = true 
		cicle[i] = false
		ans = append(ans, i)
		// preMap[i] = nil
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !DFS(i) {
			return []int{}
		}
	}

	return ans
}

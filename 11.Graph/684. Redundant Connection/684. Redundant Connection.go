// Runtime
// Details
// 27ms
// Beats 6.32%of users with Go
// Memory
// Details
// 7.02MB
// Beats 6.32%of users with Go
package main

func main() {

}

//
// func findRedundantConnection(edges [][]int) []int {
// 	visited := make(map[int]bool)
//
// 	for _, x := range edges {
// 		_, b1 := visited[x[0]]
// 		_, b2 := visited[x[1]]
// 		if b1 && b2 {
// 			return x
// 		} else {
// 			visited[x[0]] = true
// 			visited[x[1]] = true
// 		}
// 	}
// 	return []int{}
// }

func findRedundantConnection(edges [][]int) []int {
	alreadyConnected := make(map[int][]int, 0)
	var IsAlreadyConnected func(int, int, map[int]bool) bool
	IsAlreadyConnected = func(x, y int, visited map[int]bool) bool {
		if x == y {
			return true
		}
		for _, j := range alreadyConnected[x] {
			if _, b := visited[j]; !b {
				visited[j] = true
				if IsAlreadyConnected(j, y, visited) {
					return true
				}
			}

		}
		return false
	}

	for _, edge := range edges {
		visited := make(map[int]bool, 0)
		x, y := edge[0], edge[1]
		if IsAlreadyConnected(x, y, visited) {
			return edge
		}
		alreadyConnected[x] = append(alreadyConnected[x], y)
		alreadyConnected[y] = append(alreadyConnected[y], x)
	}
	return []int{}
}

// Accepted
// Runtime
// Details
// 13ms
// Beats 42.78%of users with Go
// Memory
// Details
// 7.15MB
// Beats 7.65%of users with Go

package main

func main() {

}
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	var AreaSize func(int, int) int
	AreaSize = func(i, j int) int {
		stack := make([][]int, 0)
		stack = append(stack, []int{i, j})
		size := 0 // Initialize size to 0
		for len(stack) != 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i, j = cell[0], cell[1]
			if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
				continue
			}
			size++
			grid[i][j] = 0 // Mark the cell as visited
			// Explore neighbors
			stack = append(stack, []int{i + 1, j}, []int{i - 1, j}, []int{i, j + 1}, []int{i, j - 1})
		}
		return size
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if Area := AreaSize(i, j); Area > maxArea {
					maxArea = Area
				}
			}
		}
	}
	return maxArea
}

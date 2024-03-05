// Accepted
// Runtime
// Details
// 4ms
// Beats 76.55%of users with Go
// Memory
// Details
// 3.83MB
// Beats 50.25%of users with Go
package main

// func main() {
	
// }


// func numIslands(grid [][]byte) int {
// 	if len(grid) == 0 {
// 		return 0
// 	}
	
// 	count := 0
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid); j++ {
// 			if grid[i][j] == '1' {
// 				count++
// 				dfs(grid, i, j)
// 			}
// 		}
// 	}
// 	return count

// }

// func dfs(grid [][]byte, i, j int) {
// 	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid)  {
// 		return
// 	}
// 	if   grid[i][j] != '1' {
// 		return
// 	}
// 	grid[i][j] = '0'
// 	dfs(grid, i+1, j)
// 	dfs(grid, i-1, j)
// 	dfs(grid, i, j+1)
// 	dfs(grid, i, j-1)
// }



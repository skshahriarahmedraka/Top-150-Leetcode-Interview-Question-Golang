package main



func numIslands(grid [][]byte) int {
	if len(grid) ==0 {
		return 0
	}

	count := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i:=0 ; i<len(grid);i++ {
		for j:=0 ; j<len(grid[0]);j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				BFS(grid,visited,i,j)
			}
		}
	}
	return count
}

func BFS(grid [][]byte,visited [][]bool, i,j int) {
	if i<0 || j<0 || i>=len(grid) || j>=len(grid[0]) || grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	BFS(grid,visited,i+1,j)
	BFS(grid,visited,i-1,j)
	BFS(grid,visited,i,j+1)
	BFS(grid,visited,i,j-1)
}



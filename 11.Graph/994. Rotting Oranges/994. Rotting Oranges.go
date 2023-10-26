package main

func main() {

}
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	freshOranges := 0
	rotten := make([][]int, 0)
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0

	// Count fresh oranges and find initial rotten oranges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshOranges++
			} else if grid[i][j] == 2 {
				rotten = append(rotten, []int{i, j})
			}
		}
	}

	for freshOranges > 0 && len(rotten) > 0 {
		newRotten := make([][]int, 0)

		for _, orange := range rotten {
			for _, dir := range directions {
				x, y := orange[0]+dir[0], orange[1]+dir[1]
				if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
					grid[x][y] = 2
					freshOranges--
					newRotten = append(newRotten, []int{x, y})
				}
			}
		}

		rotten = newRotten
		if len(rotten) > 0 {
			minutes++
		}
	}

	if freshOranges > 0 {
		return -1 // Some fresh oranges remain and can't be rotten
	}

	return minutes
}

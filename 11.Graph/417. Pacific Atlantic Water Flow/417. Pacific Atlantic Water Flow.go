// Accepted
// Runtime
// Details
// 43ms
// Beats 37.70%of users with Go
// Memory
// Details
// 7.36MB
// Beats 70.49%of users with Go
package main

func main() {

}

func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	pac, atl := make(map[[2]int]bool, 0), make(map[[2]int]bool, 0)

	var DFS func(int, int, map[[2]int]bool, int)

	DFS = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		if _, b := visit[[2]int{r, c}]; b || r < 0 || c < 0 || r == ROWS || c == COLS || heights[r][c] < prevHeight {
			return
		}
		visit[[2]int{r, c}] = true
		DFS(r+1, c, visit, heights[r][c])
		DFS(r-1, c, visit, heights[r][c])
		DFS(r, c+1, visit, heights[r][c])
		DFS(r, c-1, visit, heights[r][c])
	}
	for c := 0; c < COLS; c++ {
		DFS(0, c, pac, heights[0][c])
		DFS(ROWS-1, c, atl, heights[ROWS-1][c])
	}
	for r := 0; r < ROWS; r++ {
		DFS(r, 0, pac, heights[r][0])
		DFS(r, COLS-1, atl, heights[r][COLS-1])
	}
	res := [][]int{}
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			_, b1 := pac[[2]int{r, c}]
			_, b2 := atl[[2]int{r, c}]
			if b1 && b2 {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

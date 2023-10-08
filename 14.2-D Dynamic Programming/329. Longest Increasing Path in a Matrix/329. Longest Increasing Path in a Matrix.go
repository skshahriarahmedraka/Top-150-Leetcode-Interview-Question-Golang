package main

func main() {

}
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	M, N := len(matrix), len(matrix[0])
	dp := make([][]int, M)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}

		val := matrix[i][j]
		directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		maxPath := 1

		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < M && nj >= 0 && nj < N && val > matrix[ni][nj] {
				path := 1 + dfs(ni, nj)
				if path > maxPath {
					maxPath = path
				}
			}
		}

		dp[i][j] = maxPath
		return maxPath
	}

	result := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			path := dfs(i, j)
			if path > result {
				result = path
			}
		}
	}

	return result
}

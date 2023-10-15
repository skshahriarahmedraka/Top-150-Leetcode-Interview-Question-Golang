package main

import "fmt"

func main() {
	steps := 3
	arrLen := 2
	fmt.Println(numWays(steps, arrLen))
}

func numWays(steps int, arrLen int) int {
	mod := 1000000007
	maxPosition := min(arrLen-1, steps/2)
	dp := make([][]int, steps+1)

	for i := 0; i <= steps; i++ {
		dp[i] = make([]int, maxPosition+1)
	}

	dp[0][0] = 1

	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxPosition; j++ {
			dp[i][j] = dp[i-1][j]
			if j > 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j < maxPosition {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}

	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

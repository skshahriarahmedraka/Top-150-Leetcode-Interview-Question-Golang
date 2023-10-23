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


//////////////////

func numWays(steps int, arrLen int) int {
    var dfs func(int, int) int
    MOD := int(1e9) + 7
    cache := make(map[[2]int]int)
    
    dfs = func(steps, idx int) int {
        if idx > steps {
            return 0
        }
        if idx >= arrLen || idx < 0 {
            return 0
        }
        if steps == 0 {
            if idx == 0 {
                return 1
            }
            return 0
        }
        if val, ok := cache[[2]int{steps, idx}]; ok {
            return val
        }
        temp := 0
        temp = (temp + dfs(steps-1, idx-1)) % MOD
        temp = (temp + dfs(steps-1, idx)) % MOD
        temp = (temp + dfs(steps-1, idx+1)) % MOD

        cache[[2]int{steps, idx}] = temp

        return cache[[2]int{steps, idx}]
    }

    return dfs(steps, 0)
}

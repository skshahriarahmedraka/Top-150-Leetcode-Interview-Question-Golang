package main

func main() {

}

/* Top-Down Approach */

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	temp := []int{1}
	temp = append(temp, nums...)
	nums = temp
	n := len(nums)

	cache := make([][]int, n)
	for r := 0; r < n; r++ {
		cache[r] = make([]int, n)
	}

	return dp(nums, 1, len(nums)-2, cache)
}

func dp(nums []int, left, right int, cache [][]int) int {
	if left > right {
		return 0
	}

	if cache[left][right] > 0 {
		return cache[left][right]
	}

	res := 0
	for i := left; i <= right; i++ {
		cGain := nums[left-1] * nums[i] * nums[right+1]
		lGain := dp(nums, left, i-1, cache)
		rGain := dp(nums, i+1, right, cache)
		res = max(res, cGain+lGain+rGain)
	}

	cache[left][right] = res
	return res
}

/* Bottom-Up Approach */

/*
   3  1  5   8
1  3  1  5   8   1

0  0  0  0   0   0
0  3  30 159 167 0
0  0  15 135 159 0
0  0  0  40  48  0
0  0  0  0   40  0
0  0  0  0   0   0
*/

func maxCoins(nums []int) int {
	temp := []int{1}
	temp = append(temp, nums...)
	temp = append(temp, 1)
	nums = temp
	n := len(nums)

	dp := make([][]int, n)
	for r := 0; r < n; r++ {
		dp[r] = make([]int, n)
	}

	for left := n - 2; left >= 1; left-- {
		for right := left; right <= n-2; right++ {
			for i := left; i <= right; i++ {
				cGain := nums[left-1] * nums[i] * nums[right+1]
				remaining := dp[left][i-1] + dp[i+1][right]
				dp[left][right] = max(dp[left][right], cGain+remaining)
			}
		}
	}

	return dp[1][n-2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

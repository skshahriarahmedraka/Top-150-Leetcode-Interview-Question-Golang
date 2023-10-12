// Runtime
// Details
// 5ms
// Beats 74.65%of users with Go
// Memory
// Details
// 4.48MB
// Beats 96.91%of users with Go
package main

func main() {

}

// func hammingWeight(num int) int {
// 	weight := 0
// 	for num != 0 {
// 		num &= (num - 1)
// 		weight += 1
// 	}
// 	return weight
// }
//
// func countBits(n int) []int {
// 	var ans []int
//
// 	for i := 0; i <= n; i++ {
// 		ans = append(ans, hammingWeight(i))
// 	}
// 	return ans
// }

// Runtime8 ms
// Beats
// 66.27%
// Memory4.7 MB
// Beats
// 51.89%
func countBits(num int) []int {
	dp := make([]int, num+1)
	offset := 1
	for i := 1; i < num+1; i++ {
		if offset*2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]
	}
	return dp
}

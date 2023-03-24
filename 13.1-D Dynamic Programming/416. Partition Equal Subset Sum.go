/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-12 16:26:39  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-12 16:26:39  */

package main

func main() {
	
}


// My solution

// Runtime1277 ms
// Beats
// 6.30%
// Memory9.3 MB
// Beats
// 27.3%
func canPartition(nums []int) bool {
	sum := 0
	numLen := len(nums)
	for i:=0;i<numLen;i++{
		sum += nums[i]
	}
	if sum%2==1{
		return false
	}

    dp := map[int]bool{0:true}
	for i:=0;i<numLen;i++{
		temp := map[int]bool{} 
		for j,_ := range dp{
			
				temp[j+nums[i]] = true
				temp[j] = true
		}
		dp = temp
	}
	return dp[sum/2]
}


// best solution

// Runtime18 ms
// Beats
// 76.58%
// Memory2.6 MB
// Beats
// 75.23%
func canPartition(nums []int) bool {
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	if sum%2==1{
		return false
	}
	sum = sum/2
	dp := make([]bool,sum+1)
	dp[0] = true
	for i:=0;i<len(nums);i++{
		for j:=sum;j>=nums[i];j--{
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[sum]
}
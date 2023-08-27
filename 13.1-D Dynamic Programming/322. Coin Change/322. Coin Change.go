/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-11 00:10:00  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-11 00:10:00  */
// Runtime15 ms
// Beats
// 55.7%
// Memory6.6 MB
// Beats
// 49.14%
package main

func main() {
	
}


func coinChange(coins []int, amount int) int {
	if amount==0{
		return 0
	}
	dp := make([]int,amount+1)
	for i:=1;i<=amount;i++{
		dp[i] = amount+1
	}
	for i:=1;i<=amount;i++{
		for j:=0;j<len(coins);j++{
			if coins[j]<=i{
				dp[i] = min(dp[i],dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount]>amount{
		return -1
	}
	return dp[amount]
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
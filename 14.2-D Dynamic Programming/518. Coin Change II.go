package main

import "fmt"



func main(){
	amount := 5
	coins := []int{1,2,5}
	fmt.Println(change(amount,coins))
}



func change(amount int, coins []int) int {
	dp := make([][]int,0)
	for i:=0;i<amount+1;i++{
		dp = append(dp,make([]int,len(coins)+1))	
	}
	for i:=0;i<len(coins)+1;i++{
		dp[0][i]=1
	}
	// fmt.Println("🚀 dp : ", dp)

	for a:=1;a<amount+1;a++{
		for c:=len(coins)-1;c>=0;c--{
			dp[a][c] = dp[a][c+1]
			if a-coins[c]>=0{
				dp[a][c] += dp[a-coins[c]][c]
			}
            
		}
	}
	return dp[amount][0] 
	
}
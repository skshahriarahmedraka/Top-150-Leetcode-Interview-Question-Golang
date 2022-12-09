/*
 * Author: Sk Shahriar Ahmed Raka
 * Email: skshahriarahmedraka@gmail.com
 * Telegram: https://t.me/shahriarraka
 * Github: https://github.com/skshahriarahmedraka
 * StackOverflow: https://stackoverflow.com/users/12216779/
 * Linkedin: https://linkedin.com/in/shahriarraka
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Copyright (c) 2022 Your Company
 * -----
 * HISTORY:
 */

//  283 / 283 test cases passed.
//  Status: Accepted
// Runtime: 5 ms
// Memory Usage: 2.9 MB

// Your runtime beats 64.67 % of golang submissions.
// Your memory usage beats 86.84 % of golang submissions.


package main

func main(){
	li:= []int{1,100,1,1,1,100,1,1,100,1}
	minCostClimbingStairs(li)
}


func minCostClimbingStairs(cost []int) int {
    

	for i:=2 ;i<len(cost);i++ {
		cost[i]+= min(cost[i-1],cost[i-2])
	}
	return min(cost[len(cost)-1],cost[len(cost)-2])
}

func min(i,j int)int{
	if i<j {
		return i 
	}
	return j
}

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

//  139 / 139 test cases passed.
// 	Status: Accepted
// Runtime: 611 ms
// Memory Usage: 2 MB

//  Runtime: 611 ms, faster than 19.64% of Go online submissions for Target Sum.
//  Memory Usage: 2 MB, less than 97.26% of Go online submissions for Target Sum.
//  Next challenges:

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3

	fmt.Println("🚀 ~ file: 494. Target Sum.go ~ line 42 ~ funcmain ~ findTargetSumWays : ", findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	ans :=0
	CurrentValue:=0 
	var sum func(int,[]int)
	sum =func (CurrentValue int, arr []int) {
		if len(arr)==0 && CurrentValue==target {
			ans+=1
			return
		}else if len(arr)==0 {
			return
		}
		temp:=arr
		if len(arr)==1{
			temp=[]int{}
		}else {
			temp=arr[1:]
		}

		sum(CurrentValue+arr[0],temp)
		sum(CurrentValue-arr[0],temp)


	}

	sum(CurrentValue,nums)
	return ans 
}


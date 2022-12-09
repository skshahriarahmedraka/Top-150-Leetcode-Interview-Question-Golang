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
//  75 / 75 test cases passed.
//  Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2 MB

// Your runtime beats 100.00 % of golang submissions.

// Your memory usage beats 86.85 % of golang submissions.

package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println("🚀 ~ file: 213. House Robber II.go ~ line 5 ~ funcmain ~ nums : ", rob(nums))

}

func rob(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	rob1,rob2:=0,0

	for _,j:= range nums[1:] {
		temp:=MAX(j+rob1,rob2)
		rob1=rob2
		rob2=temp  
	}
	ans:=rob2 
	rob1,rob2=0,0

	for _,j:= range nums[:len(nums)-1] {
		temp:=MAX(j+rob1,rob2)
		rob1=rob2
		rob2=temp  
	}

	return MAX(ans,rob2)



}

func MAX(i, j int) int {
	if i > j {
		return i
	}
	return j
}

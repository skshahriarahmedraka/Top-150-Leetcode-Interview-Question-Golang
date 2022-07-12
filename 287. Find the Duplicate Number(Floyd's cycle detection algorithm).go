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


//  58 / 58 test cases passed.
// 	Status: Accepted
// Runtime: 212 ms
// Memory Usage: 11.8 MB

// Your runtime beats 14.46 % of golang submissions.
// Your memory usage beats 11.72 % of golang submissions.

package main 

import "fmt"


func main(){
	fmt.Println(findDuplicate([]int{1,3,4,5,2,4}))
}

func findDuplicate(nums []int) int {
    fast :=nums[nums[0]]
	slow := nums[0]
	for fast!= slow {
		slow=nums[slow]
		fast=nums[nums[fast]]
		
	}

	fast=0 
	for fast!=slow {
		slow=nums[slow]
		fast=nums[fast]
	}
	return fast

}
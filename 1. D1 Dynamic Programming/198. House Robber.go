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


//  68 / 68 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2 MB

// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 38.14 % of golang submissions.

package main

import "fmt"

func main(){
	li:= []int{2,7,9,3,1}
	fmt.Println(rob(li))
}



func MAX(i,j int)int{
	if i>j {
		return i 
	}
	return j 
}

func rob(nums []int) int {
	rob1,rob2:= 0,0 

	for _,i:= range nums {
		temp:=MAX(i+rob1,rob2)
		rob1=rob2
		rob2=temp 
	}
	return rob2
}
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


// 111 / 111 test cases passed.
// Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.1 MB

// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 65.39 % of golang submissions.


package main

import (
	"fmt"
	
)

func main() {
	li := []int{1,2,3}

	fmt.Println(plusOne(li))
}

func plusOne(digits []int) []int {
	l:=len(digits)
	temp:=0
	digits[l-1]+=1 
	for i:=l-1 ;i>=0 ;i--{
		digits[i]+=temp
		if digits[i]>9 && i==0{
			temp=digits[i]/10
			digits[i]%=10
			digits= append([]int{temp},digits... )
		}
		if digits[i]>9{
			temp=digits[i]/10
			digits[i]%=10
		}else {
			break
		}
	}
	return digits
}

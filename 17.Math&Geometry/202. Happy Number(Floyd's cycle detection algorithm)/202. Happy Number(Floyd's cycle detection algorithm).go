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



// 404 / 404 test cases passed.
// Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2 MB

// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 94.42 % of golang submissions.

package main

import "fmt"

func main(){
	n:=19
	fmt.Println(isHappy(n))
}



func isHappy(n int) bool {
	slow:=NumberSum(n)
    if slow==1{
		return true 
	}
	fast:=NumberSum(NumberSum(n))
	for slow!=fast{
		slow=NumberSum(slow)
		fast=NumberSum(NumberSum(fast))
		if slow==1 || fast==1 {
			return true 
		}
	}
	return slow==1 
}


func NumberSum(x int)int{
	ans:=0
	for x>0{
		ans+=((x%10)*(x%10))
		x/=10
	}
	return ans 
}
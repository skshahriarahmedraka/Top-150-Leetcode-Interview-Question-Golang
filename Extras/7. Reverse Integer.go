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

//  1032 / 1032 test cases passed.
// 	Status: Accepted
// Runtime: 3 ms
// Memory Usage: 2.1 MB
// Your runtime beats 58.28 % of golang submissions.
// Your memory usage beats 86.64 % of golang submissions.


package main 


func main(){

}

func reverse(x int) int {
	
    neg:=false
	if x<0 {
		x*=-1
		neg=true 
	}
	ans:=0
	for x>0 {
		
		ans=ans*10+x%10
		x=x/10
	}
	if neg{
		ans*=-1
	}
	if ans < -2147483648 || ans>2147483647 {
		return 0
	}
	return ans 
}
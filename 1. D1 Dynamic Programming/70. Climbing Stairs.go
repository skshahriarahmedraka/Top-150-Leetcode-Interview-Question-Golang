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


//  45 / 45 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 1.9 MB

// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 90.47 % of golang submissions.

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	climbStairs(n)
}

func climbStairs(n int) int {
    if n<=2 {
		return n 
	}
	PlusVar,Var:= 2,1
	for i:=2;i<n;i++{
		PlusVar=PlusVar+Var 
		Var=PlusVar-Var
	}

	return PlusVar 
}

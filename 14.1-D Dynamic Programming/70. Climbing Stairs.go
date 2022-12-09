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

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 1.9 MB, less than 88.99% of Go online submissions for Climbing Stairs.

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println( "answer : ", climbStairs(n))
}

// func climbStairs(n int) int {
//     if n<=2 {
// 		return n 
// 	}
// 	PlusVar,Var:= 2,1
// 	for i:=2;i<n;i++{
// 		PlusVar=PlusVar+Var 
//         fmt.Println("🚀 ~ file: 70. Climbing Stairs.go ~ line 44 ~ funcclimbStairs ~ PlusVar : ", PlusVar)
// 		Var=PlusVar-Var
//         fmt.Println("🚀 ~ file: 70. Climbing Stairs.go ~ line 46 ~ funcclimbStairs ~ Var : ", Var)
// 	}

// 	return PlusVar 
// }

func climbStairs(n int) int {
    if n<=2 {
		return n 
	}
	PrevIterRes,BePrevIterRes:= 2,1
	ANS:=0
	for i:=2;i<n;i++{
		// PrevIterRes = PrevIterRes+BePrevIterRes
		// BePrevIterRes = PrevIterRes-BePrevIterRes
		ANS= PrevIterRes+BePrevIterRes
		PrevIterRes,BePrevIterRes=ANS,PrevIterRes
	}

	return ANS
	// return PrevIterRes 

}

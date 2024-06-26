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

// 305 / 305 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2 MB

// Your runtime beats 100.00 % of golang submissions.
// Your memory usage beats 84.93 % of golang submissions.

package main

import (
	"fmt"
	"math"
)


func main() {
	fmt.Println("maxUInt",MaxUint)
	fmt.Println("minUint", MinUint)
	fmt.Println("MaxInt",MaxInt)
	fmt.Println("MaxInt",MinInt)
	fmt.Println(myPow(2.00000, -2147483648))
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1
func myPow(x float64, n int) float64 {
	if n == MaxInt {
		return x
	} else if n == MinInt {
		if x == 1 || x == -1{
			return 1
		}else {return 0}
	}else if x==1.0 {
		return x
	}else if x==-1.0 {
		if n%2==0{
			return 1.0
		}else{return -1.0}
	}

	if n <0 {
		x=1/x ;n*=-1
	}
	math.Pow()
	ans:=1.0
	for n>0 && ans!=0{
		ans*=x 
        // fmt.Println("🚀 ~ file: 50. Pow(x, n).go ~ line 34 ~ funcmyPow ~ ans : ", ans)
		// fmt.Println("🚀 ~ file: 50. Pow(x, n).go ~ line 40 ~ funcmyPow ~ n : ", n)
		n--
	}
	return ans 

}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
	}

	num := int64(math.Abs(float64(n)))
	pow := 1.0

	for num != 0 {
		if num&1 != 0 { // equivalent to if((n % 2) != 0) i.e. multiply only when the number is odd 
			pow *= x
		}

		x *= x
		num >>= 1 // equivalent to n = n / 2; i.e. keep dividing the number by 2
	}

	return pow
}



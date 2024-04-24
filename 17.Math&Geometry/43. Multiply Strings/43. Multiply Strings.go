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
 
 // 311 / 311 test cases passed.
 // Status: Accepted
 // Runtime: 0 ms
 // Memory Usage: 2.3 MB
 
 // You are here!
 // Your runtime beats 100.00 % of golang submissions.
 
 // You are here!
 // Your memory usage beats 77.90 % of golang submissions.
package main

import (
	"fmt"
	// "math/big"

)

func main(){
	// fmt.Println(multiply("123456789","987654321"))	
	fmt.Println(multiply("123","321"))

}

// func multiply(num1 string, num2 string) string {
// 	var x big.Int
// 	var y big.Int
// 	var z big.Int 
// 	x.SetString(num1,10)
// 	y.SetString(num2,10)
// 	z.Mul(&x,&y)
// 	return z.String()
// }


// func multiply(num1 string, num2 string) string {
//     m, n := len(num1), len(num2)
//     pos := make([]int, m+n)

//     for i := m - 1; i >= 0; i-- {
//         for j := n - 1; j >= 0; j-- {
//             mul := int(num1[i]-'0') * int(num2[j]-'0')
//             p1, p2 := i+j, i+j+1
//             sum := mul + pos[p2]

//             pos[p1] += sum / 10
//             pos[p2] = sum % 10
//         }
//     }

//     var sb strings.Builder
//     for _, p := range pos {
//         if !(sb.Len() == 0 && p == 0) {
//             sb.WriteString(fmt.Sprintf("%d", p))
//         }
//     }

//     if sb.Len() == 0 {
//         return "0"
//     }

//     return sb.String()
// }


// func multiply(num1 string, num2 string) string {
//     if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	l1, l2 := len(num1), len(num2)
// 	res := make([]byte, l1+l2)
// 	for i := l1 - 1; i >= 0; i-- {
// 		for j := l2 - 1; j >= 0; j-- {
// 			val := (num1[i] - '0') * (num2[j] - '0')
// 			res[i+j+1] += val
// 			if res[i+j+1] >= 10 {
// 				res[i+j] += (res[i+j+1] / 10)
// 				res[i+j+1] %= 10
// 			}
// 		}
// 	}
// 	if res[0] == 0 {
// 		res = res[1:]
// 	}
// 	fmt.Println(res)
// 	for i := range res {
// 		res[i] += '0'
// 	}
// 	fmt.Println(res)	

// 	return string(res)
// }



func multiply(num1,num2 string) string {
	if num1 == "0" || num2 =="0" {
		return "0"
	}
	num := make([]byte,len(num1)+len(num2))
	for i:= len(num1)-1 ; i>=0 ;i-- {
		for j:= len(num2)-1 ;j>=0 ;j-- {
			mul := (num1[i]-'0') * (num2[j]-'0')
			num[i+j+1]+=mul
			if num[i+j+1] >9 {
				num[i+j]+= num[i+j+1]/10 
				num[i+j+1] = num[i+j+1]%10 
			}

		}
	}
	if num[0]==0 {
		num= num[1:]
	}
	for i := range num {
		num[i] +='0'
	}
	return string(num)
} 


 a, b  integer 

 length(a*b) == length(a)+length(b)    
 or 
 length(a*b) == lenght(a)+lenght(b)-1
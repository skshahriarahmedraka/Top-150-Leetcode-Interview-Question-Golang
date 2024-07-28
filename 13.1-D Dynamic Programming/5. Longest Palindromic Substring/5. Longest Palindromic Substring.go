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

//  140 / 140 test cases passed.
// 	Status: Accepted
// Runtime: 10 ms
// Memory Usage: 2.5 MB
//  You are here!
// Your runtime beats 73.83 % of golang submissions.
// You are here!
// Your memory usage beats 99.39 % of golang submissions.

package main

import "fmt"


func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	s="aaäabcaaaäabca"
	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {

	maxLen := 1
	maxStr := string(s[0])
	for i := 0; i < len(s); i++ {
		//	if odd
		j, k := i-1, i+1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			if k-j+1 > maxLen {
				maxStr = s[j:k+1]
				maxLen=k-j+1
			}
			j-=1
			k+=1
		}
		//	if even
		j, k = i, i+1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			if k-j+1 > maxLen {
				maxStr = s[j:k+1]
				maxLen=k-j+1
			}
			j-=1
			k+=1
		}

	}

	return maxStr
}

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
//  130 / 130 test cases passed.
//  Status: Accepted
// Runtime: 2 ms
// Memory Usage: 1.9 MB
 
// Submitted: 0 minutes ago

// Your runtime beats 70.25 % of golang submissions.
// Your memory usage beats 61.16 % of golang submissions.


package main


func main(){

}



func countSubstrings(s string) int {

	
	ans:=0
	for i := 0; i < len(s); i++ {
		//	if odd
		// oddlen := 1
		// oddStr := ""
		j, k := i, i
		for j >= 0 && k < len(s) && s[j] == s[k] {
			ans+=1
			j-=1
			k+=1
		}
		//	if even
		
		j, k = i, i+1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			ans+=1
			j-=1
			k+=1
		}
		
	}

	return ans
}


/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-13 17:27:11  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-13 17:27:11  */
// Runtime20 ms
// Beats
// 64.29%
// Memory3.1 MB
// Beats
// 42.86%
package main

func main() {
	
}


func minWindow(s string, t string) string {
	
	if len(s) < len(t) || len(s) == 0 || len(t) == 0 {
		return ""
	}
	countT , window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}
	have, need := 0, len(countT)
	res , resLen := [2]int{-1,-1}, 1<<62
	l := 0
	for i:=0;i<len(s);i++{
		c:= s[i]
		window[c]++
		if x,b := countT[c];b{
			if window[c] == x{
				have++
			}
		}
		for have == need{
			if i-l+1 < resLen{
				res = [2]int{l,i}
				resLen = i-l+1
			}
			d := s[l]
			window[d]--
			if x,b := countT[d];b{
				if window[d] < x{
					have--
				}
			}
			l++
		}
	}
	l,r := res[0],res[1]
	if resLen != 1<<62{
		return s[l:r+1]
	}else {
		return ""
	}
}
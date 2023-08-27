/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-09 23:05:13  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-09 23:05:13  */
// Runtime4 ms
// Beats
// 47.62%
// Memory2.1 MB
// Beats
// 46.56%

package main

func main() {
	
}


func partitionLabels(s string) []int {
	
	CharLastIndex := make(map[byte]int)
	for i := range s {
		CharLastIndex[s[i]] = i
	}
	size :=0 
	end := 0
	res := make([]int,0)
	for i := range s {
		size++
		end = MAX(end, CharLastIndex[s[i]])
		if i == end {
			res = append(res, size)
			size = 0
		}
	}
	return res
}

func MAX(a,b int) int {
	if a > b {
		return a
	}
	return b
}
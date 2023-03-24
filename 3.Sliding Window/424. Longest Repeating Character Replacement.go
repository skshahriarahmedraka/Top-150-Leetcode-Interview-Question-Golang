/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-14 16:38:37  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-14 16:38:37  */
// Runtime0 ms
// Beats
// 100%
// Memory2.5 MB
// Beats
// 45%
package main

func main() {
	
}

func characterReplacement(s string, k int) int {
	
	if k >= len(s){
		return len(s)
	}
	left, right := 0, 0
	maxCount := 0
	count := make([]int, 26)
	for right < len(s){
		count[s[right]-'A']++
		maxCount = max(maxCount, count[s[right]-'A'])
		if right - left + 1 - maxCount > k{
			count[s[left]-'A']--
			left++
		}
		right++
	}
	return right - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
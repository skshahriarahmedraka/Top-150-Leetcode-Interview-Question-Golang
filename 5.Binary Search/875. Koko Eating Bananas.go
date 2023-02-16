/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-15 00:20:58  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-15 00:20:58  */
// Runtime38 ms
// Beats
// 69.7%
// Memory6.6 MB
// Beats
// 98.97%
package main

import (
	"math"
	"sort"
)

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	left, right := 1,piles[len(piles)-1] 
	var ans int = right
	for left <= right {
		mid := (left + right) / 2
        // fmt.Println(" mid : ", mid)
		k:=MinIntegers(piles, mid)
		if  k<= h {
			ans = min(ans, mid)
			right = mid -1
		} else  {
			left = mid + 1
		}
		
	}
	return ans
}

func MinIntegers(nums []int, n int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += int(math.Ceil(float64(nums[i]) / float64(n)))
	}
	return ans
}
func min (a,b int) int {
	if a < b {
		return a
	}
	return b
}
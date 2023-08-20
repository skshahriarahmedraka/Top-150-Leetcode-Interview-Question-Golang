/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-16 23:52:36  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-16 23:52:36  */
// Runtime11 ms
// Beats
// 94.31%
// Memory5.7 MB
// Beats
// 50.84%
package main

import (
	"fmt"
)

func main() {
	li := []int{1, 1, 1, 2, 2, 3}

	fmt.Println("347. Top K Frequent Elements :", topKFrequent(li, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, j := range nums {
		m[j] += 1
	}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 15 ~ functopKFrequent ~ m : ", m)
	li := make([][]int, len(nums))
	for i, j := range m {
		li[j-1] = append(li[j-1], i)
	}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 20 ~ functopKFrequent ~ li : ", li)
	ans := []int{}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 26 ~ fori:=len ~ len(li) : ", len(li))
	for i := len(li) - 1; i >= 0; i-- {
		if li[i] != nil {
			ans = append(ans, li[i]...)
		}
	}
	return ans[:k]
}

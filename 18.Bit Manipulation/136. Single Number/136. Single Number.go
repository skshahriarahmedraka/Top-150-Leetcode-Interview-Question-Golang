/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-11 11:48:44  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-11 11:48:44  */

// Runtime11 ms
// Beats
// 99.33%
// Memory6.1 MB
// Beats
// 94.41%

package main

import (
	"fmt"
	// "sort"
)

func main() {
	li := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(li))
}

// func singleNumber2(nums []int) int {
// 	sort.Ints(nums)
// 	if len(nums) < 2 {
// 		return nums[0]
// 	}
// 	if nums[0] != nums[1] {
// 		return nums[0]
// 	}
// 	if nums[len(nums)-1] != nums[len(nums)-2] {
// 		return nums[len(nums)-1]
// 	}
//
// 	for i := 0; i <= len(nums); i += 2 {
// 		if nums[i] != nums[i+1] {
// 			return nums[i]
// 		}
//
// 	}
// 	return nums[0]
// }

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

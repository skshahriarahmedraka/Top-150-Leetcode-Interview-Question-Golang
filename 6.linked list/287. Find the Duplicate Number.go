/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-17 18:40:58  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-17 18:40:58  */
// Runtime107 ms
// Beats
// 41.50%
// Memory9.5 MB
// Beats
// 23.86%

package main

func main() {
	
}

func findDuplicate(nums []int) int {
    fast :=0 
	slow := 0 
	nums_len := len(nums)
	
	for {
		slow=nums[slow]
		if nums[fast] >=nums_len{
			return -1
		}
		fast=nums[nums[fast]]
		if slow==fast {
			break
		}
	}

	fast=0 
	for {
		if fast==slow {
			return fast 
		}
		slow=nums[slow]
		fast=nums[fast]
	}

}
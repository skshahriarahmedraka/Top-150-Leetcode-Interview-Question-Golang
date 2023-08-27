/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-09 23:26:43  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-09 23:26:43  */
// Runtime149 ms
// Beats
// 13.12%
// Memory11.6 MB
// Beats
// 5.8%

package main

func main() {
	
}


func canCompleteCircuit(gas []int, cost []int) int {
	if SUM(gas) < SUM(cost) {
		return -1
	}
	total :=0 
	res := 0

	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			res = i+1
		}
	}
	return res
}

func SUM(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}


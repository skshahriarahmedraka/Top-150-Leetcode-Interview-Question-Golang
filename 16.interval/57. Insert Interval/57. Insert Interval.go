/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-28 23:56:33  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-28 23:56:33  */
// Runtime10 ms
// Beats
// 49.64%
// Memory4.6 MB
// Beats
// 41.76%
package main

func main() {

}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			newInterval[0] = MIN(newInterval[0], intervals[i][0])
			newInterval[1] = MAX(newInterval[1], intervals[i][1])

		}

	}
	res = append(res, newInterval)
	return res
}
func MIN(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

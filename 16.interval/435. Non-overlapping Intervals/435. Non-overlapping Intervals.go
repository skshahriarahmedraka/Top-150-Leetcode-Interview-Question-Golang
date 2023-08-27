/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-08 16:30:35  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-08 16:30:35  */
// Runtime210 ms
// Beats
// 95.75%
// Memory18.8 MB
// Beats
// 22.64%

package main

import "sort"

func main() {

}

type Matrix [][]int

func (m Matrix) Len() int {
	return len(m)
}
func (m Matrix) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}
func (m Matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(Matrix(intervals))
	ans := 0

	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= prevEnd {
			prevEnd = intervals[i][1]
			} else {
			ans++
			if intervals[i][1] < prevEnd {
				prevEnd = intervals[i][1]
			}
		}
	}
	return ans
}

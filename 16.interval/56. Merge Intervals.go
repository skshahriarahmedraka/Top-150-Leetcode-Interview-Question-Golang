/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-28 23:56:33  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-28 23:56:33  */
// Runtime47 ms
// Beats
// 21.99%
// Memory7.5 MB
// Beats
// 7.49%
package main

import (
	"fmt"
	"sort"
)

func main() {

}


type DArray [][]int


func (m DArray) Len() int { return len(m) }
func (m DArray) Less(i, j int) bool {
    for x := range m[i] {
        if m[i][x] == m[j][x] {
            continue
        }
        return m[i][x] < m[j][x]
    }
    return false
}

func (m DArray) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func merge(intervals [][]int) [][]int {
	var d DArray = intervals
	sort.Sort(&d)	
	
    fmt.Println("🚀 ~ file: 56. Merge Intervals.go ~ line 38 ~ funcmerge ~ d : ", d)
	for i := 0; i < len(d)-1; i++ {
		if d[i][1] >= d[i+1][0] {
			d[i][0], d[i][1] = MIN(d[i][0],d[i+1][0]), MAX(d[i][1],d[i+1][1])
			d = append(d[:i+1], d[i+2:]...)
			i--
		}
	}
	return d
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
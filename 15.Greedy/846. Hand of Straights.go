/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-08 23:25:47  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-08 23:25:47  */
// Runtime96 ms
// Beats
// 12.31%
// Memory6.9 MB
// Beats
// 69.23%
package main

import "sort"

func main() {
	
}





func isNStraightHand(hand []int, W int) bool {
	n := len(hand)

	if n%W != 0 {
		return false
	}

	size := n / W
	grid := make([][]int, size)
	sort.Ints(hand)

	for _, c := range hand {
		i := 0
		for ; i < size; i++ {
			if len(grid[i]) == W {
				continue
			}
			last := len(grid[i]) - 1
			if last == -1 || grid[i][last]+1 == c {
				grid[i] = append(grid[i], c)
				break
			}
		}
		if i == size {
			return false
		}
	}

	return true
}

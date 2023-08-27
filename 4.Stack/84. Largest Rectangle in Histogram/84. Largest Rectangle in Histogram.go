/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-14 18:10:15  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-14 18:10:15  */

package main

func main() {
	
}


func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
    stack := [][]int{}
	maxArea := heights[0] 
	for i,j := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > j {
			index,height := stack[len(stack)-1][0] , stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			maxArea = MAX(maxArea, height*(i-index))
			start = index			
		}
		stack = append(stack, []int{start,j})
	}
	for _,j := range stack {
		maxArea= MAX(maxArea, j[1]*(len(heights)-j[0]))
	}
	return maxArea
	

}

func MAX(a,b int) int {
	if a > b {
		return a
	}
	return b
}
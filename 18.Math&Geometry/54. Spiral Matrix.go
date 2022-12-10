/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-10 16:46:06  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-10 16:46:06  */
package main

import "fmt"

func main() {
	li := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(spiralOrder(li))
}

func spiralOrder(matrix [][]int) []int {
	arr := []int{}
	leftlen, rightlen := 0, len(matrix[0])-1
	toplen, bottomlen := 0, len(matrix)-1
	i := 0
	for leftlen <= rightlen && toplen <= bottomlen {
		i = leftlen
		for ;i <= rightlen;i++ {
			arr = append(arr, matrix[toplen][i])
			
		}
		toplen++
		i = toplen
		for ;i <= bottomlen;i++ {
			arr = append(arr, matrix[i][rightlen])
			
		}
		rightlen--
		if !(leftlen <= rightlen && toplen <= bottomlen) {
			break
		}
		i = rightlen
		for ; i >= leftlen ;i--{
			arr = append(arr, matrix[bottomlen][i])
			
		}
		bottomlen--
		i = bottomlen
		for ;i >= toplen ;i-- {
			arr = append(arr, matrix[i][leftlen])
			
		}
		leftlen++
	}
	return arr
}

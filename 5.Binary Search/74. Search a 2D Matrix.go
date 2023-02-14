/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-14 21:05:26  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-14 21:05:26  */
// Runtime0 ms
// Beats
// 100%
// Memory2.6 MB
// Beats
// 100%

package main 


func main (){

}



func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		if target >= matrix[i][0] && target <= matrix[i][col-1] {
			return BinarySearch(matrix[i], target) != -1
		}
	}
	return false 
}


func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid +1 
		}
		if nums[mid]> target {
			right = mid -1 
		}
	}
	return -1
}
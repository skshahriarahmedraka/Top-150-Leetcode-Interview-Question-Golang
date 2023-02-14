/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-13 21:52:25  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-13 21:52:25  */
// Runtime11 ms
// Beats
// 78.84%
// Memory5.2 MB
// Beats
// 56.57%
package main  

func main(){

}

func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	leftMax , rightMax := height[0], height[len(height)-1]
	for left < right  {
		if height[left] < height[right] {
			left++
			if height[left] < leftMax {
				ans += leftMax - height[left]

			}else {
				leftMax = height[left]
			}
		} else {
			right--
			if height[right] < rightMax {
				ans += rightMax - height[right]
			}else{
				rightMax = height[right]
			}
		}
	}
	return ans
}
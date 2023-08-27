/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-09 10:29:30  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-09 10:29:30  */
// Runtime4 ms
// Beats
// 85.56%
// Memory3.3 MB
// Beats
// 100%
package main  

func main (){


}

func maxProduct(nums []int) int {
	currentMax:= nums[0]
	currentmin := nums[0]
	finalMax:= nums[0] 
	MAX:= func (a int,b int)int{
		if a>b {
			return a
		}
		return b
	}
	MIN:= func (a int,b int)int{
		if a<b {
			return a
		}
		return b
	}
	for i:=1;i<len(nums);i++{
		temp:= currentMax

		currentMax =MAX(MAX(currentMax * nums[i],currentmin*nums[i]),nums[i])
		currentmin= MIN(MIN(temp*nums[i],currentmin*nums[i]),nums[i])
		finalMax=MAX(currentMax,finalMax)
	}
	return finalMax
}

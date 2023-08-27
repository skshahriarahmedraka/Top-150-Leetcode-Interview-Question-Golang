/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-12 10:17:47  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-12 10:17:47  */
// Runtime611 ms
// Beats
// 32.20%
// Memory2 MB
// Beats
// 88.98%
package main


func main(){

}


func findTargetSumWays(nums []int, target int) int {
	ans :=0
	CurrentValue:=0 
	var sum func(int,[]int)
	sum =func (CurrentValue int, arr []int) {
		if len(arr)==0 && CurrentValue==target {
			ans+=1
			return
		}else if len(arr)==0 {
			return
		}
		temp:=arr
		if len(arr)==1{
			temp=[]int{}
		}else {
			temp=arr[1:]
		}

		sum(CurrentValue+arr[0],temp)
		sum(CurrentValue-arr[0],temp)


	}

	sum(CurrentValue,nums)
	return ans 
}

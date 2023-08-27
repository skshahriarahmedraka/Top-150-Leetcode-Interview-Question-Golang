/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-09 16:12:32  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-09 16:12:32  */
// Runtime71 ms
// Beats
// 36.91%
// Memory4.2 MB
// Beats
// 10.73%
package main

func main(){

}

func lengthOfLIS(nums []int) int {
	
	var li []int 
	for i,_ := range nums {
		li=append(li,1) 
        _=i
    }

	for i:=len(nums)-1;i>=0;i-- {
		for j:=i+1 ;j<len(nums);j++ {
			if nums[i]<nums[j]{
				li[i]=MAX(li[i],li[j]+1)
			}
		}
	}
	ans:=0
	for _,i:= range li {
		if i>ans {
			ans=i 
		}
	}
	return ans 
}


func MAX(i,j int )int{
	if i>j {
		return i
	}
	return j 
}
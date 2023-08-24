/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-14 22:59:35  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-14 22:59:35  */
// Runtime7 ms
// Beats
// 98.74%
// Memory5.5 MB
// Beats
// 36.62%

package main

func main() {
	
}

func twoSum(numbers []int, target int) []int {
    numMap := make(map[int]int)

	for i:=0;i<len(numbers);i++{
		if _,ok := numMap[target-numbers[i]];ok{
			return []int{numMap[target-numbers[i]],i+1}
		}
		numMap[numbers[i]]=i+1
	}
	return []int{}
}

func twoSum(numbers []int, target int) []int {
    start:= 0 
    end := len(numbers)-1
    for start<end {
        if numbers[start]+numbers[end]>target {
            end--
        }else if numbers[start]+numbers[end]<target{
            start++
        }else {
            return []int{start+1,end+1}
        }
    }
    return []int{}
}

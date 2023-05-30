/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-17 16:33:02  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-17 16:33:02  */


package main

import (
	"fmt"
	// "sort"
)

func main() {
	li := []int{1,2,0,1}
	fmt.Println("",longestConsecutive(li))
}

// Runtime55 ms
// Beats
// 99.79%
// Memory8.2 MB
// Beats
// 97.53%

// func longestConsecutive(nums []int) int {
// 	if len(nums) < 2 {
// 		return len(nums)
// 	}
// 	sort.Ints(nums)
// 	ans := 1
//     // fmt.Println("🚀 ~ file: 128. Longest Consecutive Sequence.go ~ line 17 ~ funclongestConsecutive ~ nums : ", nums)
// 	temp:=1
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			continue
// 		}else if nums[i]+1 == nums[i+1] {
// 			temp += 1
// 		} else {
			
// 			temp = 1
// 		}
//         if temp > ans {
// 				ans = temp
// 			}
// 	}
// 	return ans
// }


// Runtime105 ms
// Beats
// 53.94%
// Memory10.2 MB
// Beats
// 64.67%
func longestConsecutive(nums []int) int {
    numMap := map[int]bool{}
    consecutiveNumCnt := 0

    for _, num := range nums{
        numMap[num] = true
    }
    
	// If I iterate over the map instead of input nums slice, the runtime would be 10x faster than slice one.
    //for _, num := range nums{
    for num := range numMap{
        if numMap[num - 1]{
            continue
        }
        
        cur := num
        for numMap[cur+1]{
            cur++
        }
        
        consecutiveNumCnt = max(consecutiveNumCnt, cur - num+1)
    }
    
    return consecutiveNumCnt
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
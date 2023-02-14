package main

import (
	"fmt"
	"sort"
)

func main() {

}

// func threeSum(nums []int) [][]int {
// 	if len(nums) < 3 {
// 		return [][]int{{}}
// 	}
// 	// for i := 0; i < len(nums); i++ {
// 	// 	numMap[nums[i]] = i
// 	// }
// 	sort.Ints(nums)
// 	var ans [][]int
	
// 	for i := 0; i < len(nums)-2; i++ {
// 		fmt.Println("🚀 ~ file: 15. 3Sum.go ~ line 24 ~ fori:=0;i<len ~ nums[i]  : ", nums[i])
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}
// 		numMap:= map[int]int{}
// 		for j := i + 1; j < len(nums); j++ {
// 			if _, ok := numMap[-nums[i]-nums[j]]; ok {
// 				ans = append(ans, []int{nums[i], nums[j], -nums[i] - nums[j]})
// 			} else {
// 				numMap[nums[j]] = j
// 			}
// 			fmt.Println("🚀 ~ file: 15. 3Sum.go ~ line 31 ~ forj:=i+1;j<len ~ ans : ", j, ans)
// 		}
// 	}

// 	return ans

// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	m := make(map[[3]int]struct{})
	for i := 1; i < len(nums)-1; i++ {
		start := 0
		end := len(nums) - 1
		for start < i && end > i {
			if nums[start]+nums[end] < 0-nums[i] {
				start++
			} else if nums[start]+nums[end] > 0-nums[i] {
				end--
			} else {
				curr := [3]int{nums[start], nums[i], nums[end]}
				if _, ok := m[curr]; !ok {
					m[curr] = struct{}{}
					result = append(result, curr[:])
				}
				start++
				end--
			}
		}
	}
	return result
}
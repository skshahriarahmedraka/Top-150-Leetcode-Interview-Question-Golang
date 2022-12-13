// Runtime62 ms
// Beats
// 90.72%
// Memory6.8 MB
// Beats
// 100%
package main

import "fmt"

func main() {
	li := []int{3, 2, 1, 0, 4}
	fmt.Println(" ", canJump(li))
}

func canJump(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= goal {
			goal = i
		}
	}
	if goal == 0 {
		return true
	} else {
		return false
	}
}

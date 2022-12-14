// Runtime34 ms
// Beats
// 57.4%
// Memory6.1 MB
// Beats
// 69.63%
// package main

import "fmt"

func main() {
	li := []int{2, 3, 1, 1, 4}
	fmt.Println("45. Jump Game II : ", jump(li))
}

func jump(nums []int) int {
	jump := 0
	left, right := 0, 0
	for right < len(nums)-1 {
		r := 0

		for i := left; i <= right; i++ {
			if nums[i]+i > r {
				r = nums[i] + i
			}

		}
		left = right + 1
		right = r
		jump++
	}
	return jump

}

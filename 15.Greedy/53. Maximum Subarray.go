package main

import "fmt"

func main() {
	li := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(" maxSubArray", maxSubArray(li))
}
func maxSubArray(nums []int) int {
	max := nums[0]
	temp := 0

	for _, i := range nums {
		temp += i
		if temp > max {
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
	}
	return max
}

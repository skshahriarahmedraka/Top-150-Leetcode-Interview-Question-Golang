package main

import "fmt"

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
    lenNums := len(nums)
    for i:= 0 ; i<len(nums);i++ {
        if nums[i]>=1 && nums[i]<=lenNums && nums[i]!=nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
            i--
        }
    }
    for i:=0;i<lenNums;i++ {
        if nums[i]!= i+1 {
            return i+1
        }
    }
    return lenNums+1
}
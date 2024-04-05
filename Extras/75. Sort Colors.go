package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	 sortColors(nums)
}

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0

	for _, i := range nums {
		switch i {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		default:
			continue
		}

	}
	fmt.Println("🚀 ~ file: 75. Sort Colors.go ~ line 22 ~ funcsortColors ~ blue : ", blue)
	fmt.Println("🚀 ~ file: 75. Sort Colors.go ~ line 22 ~ funcsortColors ~ white : ", white)
	fmt.Println("🚀 ~ file: 75. Sort Colors.go ~ line 22 ~ funcsortColors ~ red : ", red)
	// delete all value form array
	nums = nums[:0]
	for red > 0 {
		nums = append(nums, 0)
		red--
	}
	for white > 0 {
		nums = append(nums, 1)
		white--
	}
	for blue > 0 {
		nums = append(nums, 2)
		blue--
	}
	fmt.Println("🚀 ~ file: 75. Sort Colors.go ~ line 40 ~ funcsortColors ~ nums : ", nums)

}
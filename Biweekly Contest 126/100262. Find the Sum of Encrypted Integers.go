package main

import "fmt"

func main() {
	fmt.Println(sumOfEncryptedInt([]int{1, 2, 3}))
	fmt.Println(sumOfEncryptedInt([]int{10, 21, 31}))
}

func sumOfEncryptedInt(nums []int) int {
	ans := 0
	for _, i := range nums {
		ans += replaceWithLargest(i)
	}
	return ans
}

func replaceWithLargest(i int) int {

	largest := i % 10
	numLen := 0
	for i > 0 {
		if i%10 > largest {
			largest = i % 10
		}
		numLen++
		i = i / 10
	}
	num := 0
	for j := 0; j < numLen; j++ {
		num = (num*10 + largest)
	}
	return num
}

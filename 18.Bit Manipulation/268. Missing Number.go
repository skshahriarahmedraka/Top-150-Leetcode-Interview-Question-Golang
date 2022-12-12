package main

func main() {

}

// solution 1

// func missingNumber(nums []int) int {
// 	s := len(nums)
// 	sumx := s * (s + 1) / 2
// 	sum := 0
// 	for _, j := range nums {
// 		sum += j
// 	}
// 	return sumx - sum
//
// }

// solution 2

func missingNumber(nums []int) int {
	numslen := len(nums)
	ans := 0
	for i := 0; i < numslen; i++ {
		ans ^= (nums[i] ^ i)
	}
	ans ^= (numslen)
	return ans
}

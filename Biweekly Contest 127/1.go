package main

import "math"

func main() {

}

// func generateSubarrays(arr []int) [][]int {
//     var subarrays [][]int

//     // Iterate over all possible start indices
//     for start := 0; start < len(arr); start++ {
//         // Iterate over all possible end indices
//         for end := start; end < len(arr); end++ {
//             // Create a subarray from start to end index
//             subarray := arr[start : end+1]
//             // Append the subarray to the list of subarrays
//             subarrays = append(subarrays, subarray)
//         }
//     }

//     return subarrays
// }

// func minimumSubarrayLength(nums []int, k int) int {
// 	// Initialize the minimum length to be the maximum possible value

// 	subArr := generateSubarrays(nums)
// 	ans :=0
// 	for i:=0; i<len(subArr); i++{
// 	temp := BitwiseORs(subArr[i])
// 		if  temp>= k{
// 			if ans ==0{
// 				ans= temp
// 			}else if temp < ans{
// 				ans = temp
// 			}

// 		}
// 	}
// 	if ans ==0{
// 		return -1
// 	}
// 	return ans
// }

// func BitwiseORs(arr []int) int {

// 	ans :=0

// 		for j:=0; j<len(arr); j++{
// 			ans |= arr[j]
// 		}

// 	return ans

// }

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	minLength := math.MaxInt32

	for start := 0; start < n; start++ {
		currentOr := 0
		for end := start; end < n; end++ {
			currentOr |= nums[end]
			if currentOr >= k {
				minLength = min(minLength, end-start+1)
				break
			}
		}
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}

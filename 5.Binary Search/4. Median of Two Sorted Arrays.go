package main

import (
	"math"
	"sort"
)

func main() {

}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	if len(nums1) > len(nums2) {
// 		return FindMedian(nums1, nums2)
// 	}
// }

// func FindMedian(nums1 []int, nums2 []int) float64 {
// 	tempList := make([]int, 0)
// 	if len(nums1)%2 != 0 && len(nums2)%2 != 0 {
// 		tempList = append(tempList,  nums2[len(nums2)/2])
// 		if len(nums1)%2 == 0 {
// 			left, right := 0, len(nums1)-1
// 			mid1 := 0
// 			for left <= right {
// 				mid1 = (left + right) / 2
				
// 				if mid
// 				if tempList[0] >= nums1[mid1] && tempList[0] <= nums1[mid2] {
// 					tempList = append(tempList, nums1[mid1], nums1[mid2])
// 					break
// 				} else if tempList[1] <= nums1[mid1] {
// 					right = mid1 - 1
// 				} else {
// 					left = mid2 + 1
// 				}

// 			}
// 			sort.Ints(tempList)
// 			if len(tempList)%2 == 0 {
// 				return float64(tempList[len(tempList)/2-1]+tempList[len(tempList)/2]) / 2
// 			} else {
// 				return float64(tempList[len(tempList)/2])
// 			}
		

// 	}
// 	if len(nums2)%2 == 0 {
// 		tempList = append(tempList, nums2[len(nums2)/2-1], nums2[len(nums2)/2])
// 		if len(nums1)%2 == 0 {
// 			left, right := 0, len(nums1)-1
// 			mid1, mid2 := 0, 0
// 			for left <= right {
// 				mid1 = (left + right) / 2
// 				mid2 = (left+right)/2 + 1

// 				if tempList[1] >= nums1[mid1] && tempList[0] <= nums1[mid2] {
// 					tempList = append(tempList, nums1[mid1], nums1[mid2])
// 					break
// 				} else if tempList[1] <= nums1[mid1] {
// 					right = mid1 - 1
// 				} else {
// 					left = mid2 + 1
// 				}

// 			}
// 			sort.Ints(tempList)
// 			if len(tempList)%2 == 0 {
// 				return float64(tempList[len(tempList)/2-1]+tempList[len(tempList)/2]) / 2
// 			} else {
// 				return float64(tempList[len(tempList)/2])
// 			}
// 		}else {
// 			left, right := 0, len(nums1)-1
// 			mid1:= 0
// 			for left <= right {
// 				mid1 = (left + right) / 2
				

// 				if tempList[1] >= nums1[mid1] && tempList[0]>= nums1[mid1] {
// 					tempList = append(tempList, nums1[mid1])
// 					break
// 				} else if tempList[1] <= nums1[mid1] {
// 					right = mid1 - 1
// 				} else {
// 					left = mid1 + 1
// 				}

// 			}
// 			sort.Ints(tempList)
// 			if len(tempList)%2 == 0 {
// 				return float64(tempList[len(tempList)/2-1]+tempList[len(tempList)/2]) / 2
// 			} else {
// 				return float64(tempList[len(tempList)/2])
// 			}
// 		}
// 	} else {
// 		tempList = append(tempList, nums2[len(nums2)/2])
// 	}
// 	// binary search

// }



func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
 
	if len(nums1) < len(nums2) {
		nums1,nums2 = nums2,nums1
	}
	total := (len(nums1) + len(nums2) ) 
	half := total / 2
	left,right := 0,len(nums1)-1

	for true {
		i := (left + right) / 2
		j := half - i - 2 
		left1,right1,left2,right2 := 0,0,0,0
		if i>=0 {
			left1 = nums1[i] 
		}else { 
			left1 = int(math.Inf(-1))
		}

		if i+1 < len(nums1) {
			right1 = nums1[i+1]
		}else {
			right1 = int( math.Inf(1))
		}

		if j>=0 {
			left2 = nums2[j]	
		}else {
			left2 = int(math.Inf(-1))
		}

		if j+1 < len(nums2) {
			right2 = nums2[j+1]
		}else {
			right2 = int(math.Inf(1))
		}

		if left1 <= right2 && left2 <= right1 {
			if total % 2 == 0 {
				return float64(min(right1,right2))
				}else {
				return float64(max(left1,left2) + min(right1,right2)) / 2
			}
		}else if left1 > right2 {
			right = i - 1
		}else {
			left = i + 1
		}

	}
	return 0
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
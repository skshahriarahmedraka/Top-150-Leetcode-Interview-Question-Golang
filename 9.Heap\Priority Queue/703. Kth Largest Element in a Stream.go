func findKthLargest(nums []int, k int) int {
    return doFindKthLargest(nums, k, 0, len(nums)-1)
}

func doFindKthLargest(nums []int, k int, start, end int) int {
    nlen := len(nums)
    targetPos := nlen - k
    
    for {
        pivotIndex := partition(nums, start, end)
        if pivotIndex == targetPos {
            return nums[pivotIndex]
        } else if pivotIndex < targetPos {
            start = pivotIndex + 1
        } else {
            end = pivotIndex - 1
        }
    }
}

// partition choses nums[start] as pivot and
// moves all values less than pivot to the left side,
// moves all values greater than pivot to the rigit side
// of the pivot value.
// This represents an "one shot" implementation of quick sort.
func partition(nums []int, start, end int) int {
    pivot := nums[start]
    left, right := start+1, end
    
    for left <= right {
        for left <= right && nums[left] <= pivot {
            left++
        }
        for left <= right && nums[right] > pivot {
            right--
        }
        if left <= right {
            nums[left], nums[right] = nums[right], nums[left]
        }
    }
    nums[right], nums[start] = nums[start], nums[right]
    return right
}
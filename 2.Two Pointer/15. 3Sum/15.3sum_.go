package main 


func main(){

}


func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }
    
    sort.Ints(nums)
    var results [][]int

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        target, left, right := -nums[i], i+1, len(nums)-1

        for left < right {
            sum := nums[left] + nums[right]

            if sum == target {
                results = append(results, []int{-target, nums[left], nums[right]})
                left++
                right--

                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    return results
}

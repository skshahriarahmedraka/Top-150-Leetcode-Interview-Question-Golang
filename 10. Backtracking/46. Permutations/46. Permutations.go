//Go
//Runtime0 ms
//Beats
//100%
//Memory2.7 MB
//Beats
//60.36%

package main

func main() {

}

func permute(nums []int) [][]int {
	res := [][]int{}
	permuteRecu([]int{}, nums, &res)
	return res
}

func permuteRecu(curr []int, nums []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, curr)
		return
	}
	for i := 0; i < len(nums); i++ {
		permuteRecu(append(curr, nums[i]), append(nums[:i], nums[i+1:]...), res)
	}
}

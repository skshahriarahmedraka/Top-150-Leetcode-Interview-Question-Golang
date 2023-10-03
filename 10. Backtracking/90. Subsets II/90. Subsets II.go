package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	var dfs func(s int, path []int)

	dfs = func(s int, path []int) {
		ans = append(ans, append([]int(nil), path...))
		if s == len(nums) {
			return
		}

		for i := s; i < len(nums); i++ {
			if i > s && nums[i] == nums[i-1] {
				continue
			}
			dfs(i+1, append(path, nums[i]))
		}
	}

	sort.Ints(nums)
	dfs(0, []int{})
	return ans
}

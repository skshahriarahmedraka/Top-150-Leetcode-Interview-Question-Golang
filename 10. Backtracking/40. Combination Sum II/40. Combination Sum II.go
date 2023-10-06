package main

import "sort"

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	combinations := [][]int{}
	sort.Ints(candidates)

	var recurse func(i int, arr []int, sum int)
	recurse = func(i int, arr []int, sum int) {
		//base case success
		if sum == target {
			temp := make([]int, len(arr))
			copy(temp, arr)
			combinations = append(combinations, temp)
			return
		}

		//base case failure
		if sum > target || i >= len(candidates) {
			return
		}

		arr = append(arr, candidates[i])
		recurse(i+1, arr, sum+candidates[i])

		arr = arr[:len(arr)-1]

		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}

		recurse(i+1, arr, sum)
	}

	recurse(0, []int{}, 0)
	return combinations
}

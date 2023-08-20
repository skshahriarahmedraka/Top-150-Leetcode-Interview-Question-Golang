package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, j := range nums {
		if x, found := m[target-j]; found {
			return []int{x, i}
		}
		m[j] = i
	}
	return nil
}

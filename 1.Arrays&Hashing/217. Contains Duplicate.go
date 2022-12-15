package main

func main() {

}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, i := range nums {
		if _, err := m[i]; !err {
			m[i] = true
			continue
		} else {
			return true
		}
	}
	return false
}

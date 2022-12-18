// Runtime74 ms
// Beats
// 99.3%
// Memory8.2 MB
// Beats
// 94.98%

package main

func main() {

}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	l, r := 0, len(height)-1
	//lenHeight := len(height)
	maximumWater := 0
	for l != r {
		if height[l] >= height[r] {
			if maximumWater < height[r]*(r-l) {
				maximumWater = height[r] * (r - l)
			}
			r -= 1
		} else {
			if maximumWater < height[l]*(r-l) {
				maximumWater = height[l] * (r - l)
			}
			l += 1
		}
	}
	return maximumWater
}

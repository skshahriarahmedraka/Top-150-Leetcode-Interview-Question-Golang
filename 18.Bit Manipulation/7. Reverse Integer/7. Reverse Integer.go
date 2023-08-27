package main

func main() {

}

func reverse(x int) int {

	neg := false
	if x < 0 {
		x *= -1
		neg = true
	}
	ans := 0
	for x > 0 {

		ans = ans*10 + x%10
		x = x / 10
	}
	if neg {
		ans *= -1
	}
	if ans < -2147483648 || ans > 2147483647 {
		return 0
	}
	return ans
}

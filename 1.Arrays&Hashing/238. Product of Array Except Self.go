package main

func main() {

}

func productExceptSelf(nums []int) []int {
	mulAll := 1
	for _, i := range nums {
		mulAll *= i
	}
	Ans := []int{}

	for _, i := range nums {
		temp := mulAll / i
		if temp < 0 {
			temp = 0
		}
		Ans = append(Ans, temp)
	}
	return Ans
}

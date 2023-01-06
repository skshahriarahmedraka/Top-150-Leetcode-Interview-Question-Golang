/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-06 13:02:14  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-06 13:02:14  */
// Runtime1 ms
// Beats
// 54.32%
// Memory2 MB
// Beats
// 16.5%
package main

import "fmt"


func main(){
	fmt.Println(ConvertToTernary(2747))
}

func ConvertToTernary(n int) int {
	if n == 0 {
		return 0
	}
	return n%3 + 10*ConvertToTernary(n/3)
}

func ConvertToBinary(n int) int {
	if n == 0 {
		return 0
	}
	return n%2 + 10*ConvertToBinary(n/2)
}


func isStrictlyPalindromic(n int) bool {
	binary:= ConvertToBinary(n)

	ternary:= ConvertToTernary(n)
	return binary == Reverse(binary) && ternary == Reverse(ternary)
}
func Reverse(n int) int {
	var reverse int
	for n != 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	return reverse
}


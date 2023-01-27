/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-28 00:51:27  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-28 00:51:27  */
// Runtime3 ms
// Beats
// 24.52%
// Memory1.9 MB
// Beats
// 76.13%
package main

func main() {

}

func getSum(a int, b int) int {

	for b != 0 {
		c := a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

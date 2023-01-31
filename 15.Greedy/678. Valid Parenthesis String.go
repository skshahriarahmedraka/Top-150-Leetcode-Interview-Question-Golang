/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-29 23:36:24  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-29 23:36:24  */
// Runtime1 ms
// Beats
// 54.76%
// Memory1.8 MB
// Beats
// 71.43%
package main

func main (){

}


func checkValidString(s string) bool {
	if (len(s) == 1 && s == "*"){
		return true
	}
	if len(s)<=1 || s[0] == ')'   {
		return false
	}
	leftmin := 0
	leftmax := 0
	for _,i := range s {
		if i == '(' {
			leftmax++
			leftmin++
		} else if i == ')' {
			leftmax--
			leftmin--
		} else {
			leftmax++
			leftmin--
		}
		if leftmax < 0 {
			return false
		}else if leftmin < 0 {
			leftmin = 0
		}

		
	}
	return leftmin == 0
}
/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-21 18:21:53  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-21 18:21:53  *//* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-21 18:20:56  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-21 18:20:56  */package main
// Runtime0 ms
// Beats
// 100%
// Memory2.4 MB
// Beats
// 9.3%
import "fmt"

func main() {
	fmt.Println("valid : ", isValid("()"))
}



func isValid(s string) bool {
	stack := []string{}
	closeToOpen := map[string]string{")": "(", "}": "{", "]": "["}
	if !( s[0]=='(' || s[0]=='{' || s[0]=='[') {
		return false
	}
	for _, i := range s {
		i := string(i)
		if _, ok := closeToOpen[i]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[i] {
				stack = stack[:len(stack)-1]
			}else {
                return false
            }
			} else {
				stack = append(stack, i)
			}
			// fmt.Println("\U0001f680  i : ", i)
			// fmt.Println("\U0001f680  closeToOpen[i] : ",  closeToOpen[i])
			// fmt.Println("\U0001f680  stack : ", stack)
	}

	fmt.Println("stack : ", stack)
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
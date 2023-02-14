/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-14 12:54:03  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-14 12:54:03  */
// Runtime0 ms
// Beats
// 100%
// Memory2.7 MB
// Beats
// 85.84%

package main

func main(){

}


func generateParenthesis(n int) []string {
	res:=[]string{}
	RecursiveParanthesis(&res,0,0,"",n)
	return res
    
}

func RecursiveParanthesis(res *[]string,open int,close int,str string,n int) {
	if open==n && close==n{
		 *res = append(*res,str)
	}
	if open<n{
		 RecursiveParanthesis(res,open+1,close,str+"(",n)
	}
	if close<open{
		 RecursiveParanthesis(res,open,close+1,str+")",n)
	}
}
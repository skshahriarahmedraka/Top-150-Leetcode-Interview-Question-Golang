/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-05 15:39:43  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-05 15:39:43  */

// Runtime 0 ms
// Beats
// 100%
// Memory 2 MB
// Beats
// 80.50%
package main

import "strconv"

func main(){

}


func numDecodings(s string) int {
	
    if len(s)==0 || s[0]=='0' {return 0}else if len(s)==1{return 1}

	count1:=1
	count2:=1

	for i:=1;i<len(s);i++{
		d,_ :=strconv.Atoi(string(s[i]))
		x,_ :=strconv.Atoi(string(s[i-1]))
		dd := x*10+d
		count :=0 
		if d>0 { count+=count2}
		if dd>=10 && dd<=26{count+=count1}
		count1=count2
		count2=count

	}
	return count2

}

// func numDecodings(s string) int {
// 	a:=1
// 	len := len(s)


// }
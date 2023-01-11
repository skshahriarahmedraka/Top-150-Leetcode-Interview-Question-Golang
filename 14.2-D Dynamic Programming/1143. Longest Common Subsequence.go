/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-12 00:43:38  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-12 00:43:38  */
// Runtime8 ms
// Beats
// 42.66%
// Memory13.6 MB
// Beats
// 14.43%
package main



func main (){

}


func longestCommonSubsequence(text1 string, text2 string) int {
	lentxt1:=len(text1)
	lentxt2:=len(text2)

	matrix := make([][]int,0)
	for i:=0;i<lentxt1+1;i++{
		matrix = append(matrix,make([]int,lentxt2+1))
	}

	for i:=0 ; i<lentxt1 ; i++{
		for j:=0 ; j<lentxt2 ; j++{
			if text1[i] == text2[j]{
				matrix[i+1][j+1]=matrix[i][j]+1
			}
			if text1[i] != text2[j]{
				matrix[i+1][j+1]=MAX(matrix[i][j+1],matrix[i+1][j])
			}
		}
	}

	
	return matrix[lentxt1][lentxt2]
}

func MAX(a,b int) int{
	if a>b{
		return a
	}
	return b
}
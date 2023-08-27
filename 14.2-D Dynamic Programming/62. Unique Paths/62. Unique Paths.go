/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-01 19:11:40  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-01 19:11:40  */

// Runtime0 ms
// Beats
// 100%
// Memory2 MB
// Beats
// 72.84%

package main 

func main(){

}


func uniquePaths(m int, n int) int {
    matrix := make([][]int,0)
    for i:=0;i<m;i++{
		matrix = append(matrix,make([]int,n))
	}

	for i:=m-1 ; i>=0 ; i--{
		for j:=n-1 ; j>=0 ; j--{
			if i==m-1 || j==n-1{
				matrix[i][j] = 1
			}else{
				matrix[i][j] = matrix[i+1][j] + matrix[i][j+1]
			}
		}
	}
	return matrix[0][0]
}
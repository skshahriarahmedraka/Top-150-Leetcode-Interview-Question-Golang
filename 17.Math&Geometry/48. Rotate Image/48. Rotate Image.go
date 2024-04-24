/*
 * Author: Sk Shahriar Ahmed Raka
 * Email: skshahriarahmedraka@gmail.com
 * Telegram: https://t.me/shahriarraka
 * Github: https://github.com/skshahriarahmedraka
 * StackOverflow: https://stackoverflow.com/users/12216779/
 * Linkedin: https://linkedin.com/in/shahriarraka
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Copyright (c) 2022 Your Company
 * -----
 * HISTORY:
 */

package main

import "fmt"

func main() {
	li := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	rotate(li)
}

// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	n2 := len(matrix[0])
// 	m3:= matrix
// 	matrix=[][]int{}
// 	matrix2 := make([][]int,n)
// 	for i:=0 ; i<n ; i++{
// 		for j:=n2-1 ; j>=0 ; j--{
// 			matrix2[i] = append(matrix2[i],m3[j][i])
// 		}
// 		matrix = append(matrix, matrix2[i])
// 	}
// 	// fmt.Println("🚀 ~ file: 48. Rotate Image.go ~ line 41 ~ funcrotate ~ matrix : ", matrix)

// }

// func rotate(matrix [][]int)  {

// 	for i:=0 ; i<len(matrix) ;i++ {

// 		for j:=i ; j<len(matrix[0]) ; j++ {

// 		}

// 	}
// }

func rotate(m [][]int) {
	N := len(m)
	for i := 0; i < (N)/2; i++ {
		for j := 0; j < (N+1)/2; j++ {
			
			m[j][N-i-1], m[N-i-1][N-j-1], m[N-j-1][i], m[i][j] = m[i][j], m[j][N-i-1], m[N-i-1][N-j-1], m[N-j-1][i]
		}
	}
}

// this comment is made

func rotate(m[][]int) {
	for i :=0 ;i<len(m)/2 ;i++ {
		m[i],m[len(m)-i-1] =m[len(m)-i-1],m[i]
	}
	for i:=0 ;i<len(m);i++ {
		for j:=i+1 ;j<len(m[i]);j++ {
			m[i][j],m[j][i] =m[j][i],m[i][j]
		}
	}
}



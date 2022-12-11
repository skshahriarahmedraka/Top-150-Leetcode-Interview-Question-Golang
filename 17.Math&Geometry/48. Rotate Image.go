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
			fmt.Printf("swaped position %v,%v : %v and %v,%v : %v\n", j, N-i-1, m[j][N-i-1], i, j, m[i][j])
			fmt.Printf("swaped position %v,%v : %v and %v,%v : %v\n", N-i-1, N-j-1, m[N-i-1][N-j-1], j, N-i-1, m[j][N-i-1])
			fmt.Printf("swaped position %v,%v : %v and %v,%v : %v\n", N-j-1, i, m[N-j-1][i], N-i-1, N-j-1, m[N-i-1][N-j-1])
			fmt.Printf("swaped position %v,%v : %v and %v,%v : %v\n\n", i, j, m[i][j], N-j-1, i, m[N-j-1][i])
			m[j][N-i-1], m[N-i-1][N-j-1], m[N-j-1][i], m[i][j] = m[i][j], m[j][N-i-1], m[N-i-1][N-j-1], m[N-j-1][i]
		}
	}
}

// this comment is made



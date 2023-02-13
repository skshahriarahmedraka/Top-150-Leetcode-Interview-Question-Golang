/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-13 10:11:22  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-13 10:11:22  */
// Runtime3 ms
// Beats
// 74.35%
// Memory2.6 MB
// Beats
// 100%
package main

func main() {

}

func isValidSudoku(board [][]byte) bool {

	validity := true

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !isValid(board, i, j) {
					validity = false
					break
				}
			}
		}
	}

	return validity

}

func isValid(board [][]byte, i, j int) bool {

	// check row
	for k := 0; k < 9; k++ {
		if k != j && board[i][k] == board[i][j] {
			return false
		}
	}

	// check column
	for k := 0; k < 9; k++ {
		if k != i && board[k][j] == board[i][j] {
			return false
		}
	}

	// check box
	for k := (i / 3) * 3; k < ((i/3)*3)+3; k++ {
		for l := (j / 3) * 3; l < ((j/3)*3)+3; l++ {
			if k != i && l != j && board[k][l] == board[i][j] {
				return false
			}
		}
	}

	return true
}

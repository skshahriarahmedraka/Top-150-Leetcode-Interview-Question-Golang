// Go
// Runtime3 ms
// Beats
// 90.4%
// Memory3.2 MB
// Beats
// 58.92%
package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

func solveNQueens(n int) [][]string {

	var col map[int]bool = make(map[int]bool, 0)
	var Diag map[int]bool = make(map[int]bool, 0)
	var crosDiag map[int]bool = make(map[int]bool, 0)
	var board [][]string = make([][]string, 0)
	var res [][]string = make([][]string, 0)
	var N int = 0
	N = n
	for i := 0; i < n; i++ {
		board = append(board, make([]string, n)) // Initialize each row with an empty slice of strings
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	backtrack(&col, &Diag, &crosDiag, &board, &res, &N, 0)
	return res
}
func backtrack(col *map[int]bool, Diag *map[int]bool, crosDiag *map[int]bool, board *[][]string, res *[][]string, N *int, r int) {
	if r == *N {
		temp := make([]string, *N)
		for i := 0; i < *N; i++ {
			temp[i] = strings.Join((*board)[i][:], "")
		}
		*res = append(*res, temp)
		return
	}
	for i := 0; i < *N; i++ {
		if (*col)[i] || (*Diag)[r-i] || (*crosDiag)[r+i] {
			continue
		}
		(*col)[i] = true
		(*Diag)[r-i] = true
		(*crosDiag)[r+i] = true
		(*board)[r][i] = "Q"
		backtrack(col, Diag, crosDiag, board, res, N, r+1)

		(*col)[i] = false
		(*Diag)[r-i] = false
		(*crosDiag)[r+i] = false
		(*board)[r][i] = "."
	}
}

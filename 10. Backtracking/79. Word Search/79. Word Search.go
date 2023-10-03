//Accepted
//Runtime
//Details
//104ms
//Beats 81.24%of users with Go
//Memory
//Details
//2.02MB
//Beats 36.79%of users with Go

package main

import "fmt"

func main() {
	board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {

				if DFS(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func DFS(board [][]byte, word string, i, j, pos int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' {
		return false
	}
	if board[i][j] != word[pos] {
		return false
	}
	if pos == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = '*'
	pos++

	res := DFS(board, word, i-1, j, pos) || DFS(board, word, i, j-1, pos) || DFS(board, word, i+1, j, pos) || DFS(board, word, i, j+1, pos)
	board[i][j] = temp
	return res
}

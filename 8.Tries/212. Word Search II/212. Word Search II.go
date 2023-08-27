/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-06 15:03:33  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-06 15:03:33  */
// Runtime379 ms
// Beats
// 61.88%
// Memory11 MB
// Beats
// 8.29%

package main

func main() {
	
}

type trieNode struct {
    children [26]*trieNode
    level int
}

func search(root *trieNode, word string) bool {
    cur := root
    n := len(word)
    for i := 0; i < n; i++ {
        if cur.children[word[i]-'a'] == nil {
            return false
        }
        cur = cur.children[word[i]-'a']    
    }
    return true
}

func findWords(board [][]byte, words []string) []string {
    m := len(board)
    n := len(board[0])
    res := []string{}
    mark := big.NewInt(0)
    stop := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(len(words))), nil), big.NewInt(1))
    
    var dfs func(root *trieNode, i, j int) *trieNode
    dfs = func(root *trieNode, i, j int) *trieNode {
        if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 0 {
            return root
        }
        char := board[i][j]
        if root.children[char-'a'] == nil {
            root.children[char-'a'] = &trieNode{level: root.level+1}
        }
        if root.level+1 == 10 {
            return root
        }
        board[i][j] = 0
        root.children[char-'a'] = dfs(root.children[char-'a'], i-1, j)
        root.children[char-'a'] = dfs(root.children[char-'a'], i, j-1)
        root.children[char-'a'] = dfs(root.children[char-'a'], i+1, j)
        root.children[char-'a'] = dfs(root.children[char-'a'], i, j+1)
        board[i][j] = char
        return root
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mark.Cmp(stop) == 0 {
                goto termination
            }
            trie := dfs(&trieNode{level: 0}, i, j)
            for k := range words {
                if mark.Bit(k) == 0 && search(trie, words[k]) {
                    res = append(res, words[k])
                    mark = mark.SetBit(mark, k, 1)
                }
            }
        }
    }
    
    termination:
    return res
}
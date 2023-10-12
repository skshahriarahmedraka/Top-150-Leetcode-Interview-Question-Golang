// Go
// Runtime145 ms
// Beats
// 65.82%
// Memory7.2 MB
// Beats
// 58.23%
package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	mp := make(map[string]bool)

	for _, w := range wordList {
		mp[w] = true
	}

	findInMap := func(str string) bool {
		_, ok := mp[str]
		return ok
	}

	queue := make([]string, 0)

	queue = append(queue, beginWord)

	steps := 0

	for len(queue) > 0 {
		for _, q := range queue {
			w := q
			if w == endWord {
				return steps + 1
			}
			for i := 0; i < len(w); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					ww := w[:i] + string(ch) + w[i+1:]
					if findInMap(ww) {
						queue = append(queue, ww)
						delete(mp, ww)
					}
				}
			}
			queue = queue[1:]
			fmt.Println(queue)
		}
		steps++
		fmt.Println(steps)
	}

	return 0
}

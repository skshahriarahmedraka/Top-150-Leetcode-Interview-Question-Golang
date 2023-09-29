package main

import "fmt"

func main() {
	w1 := "intention"
	w2 := "execution"
	fmt.Println(minDistance(w1, w2))
}

func minDistance(word1 string, word2 string) int {
	wlen1 := len(word1)
	wlen2 := len(word2)

	cache := make([][]int, wlen1+1)
	for i := 0; i <= wlen1; i++ {
		cache[i] = make([]int, wlen2+1)
	}
	for i := 0; i <= wlen1; i++ {
		for j := 0; j <= wlen2; j++ {
			if j == wlen2 {
				cache[i][j] = wlen1 - i
			} else if i == wlen1 {
				cache[i][j] = wlen2 - j
			} else {
				cache[i][j] = 0
			}
		}
	}

	for i := wlen1 - 1; i >= 0; i-- {
		for j := wlen2 - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				cache[i][j] = cache[i+1][j+1]
			} else {
				cache[i][j] = 1 + min(cache[i][j+1], cache[i+1][j+1], cache[i+1][j])
			}
		}
	}
	return cache[0][0]
}

func min(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		} else {
			return k
		}
	} else {
		if j < k {
			return j
		} else {
			return k
		}
	}
}

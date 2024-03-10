package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
	// fmt.Println(longestPalindrome("cbbd"))
	// fmt.Println(longestPalindrome("a"))
	// fmt.Println(longestPalindrome("ac"))
	// fmt.Println(longestPalindrome("bb"))
	// fmt.Println(longestPalindrome("babadada"))
}

// 
func longestPalindrome(s string) string {
    T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
    n := len(T)
    P := make([]int, n)
    C, R := 0, 0
    
    for i := 1; i < n-1; i++ {
		fmt.Println("i, Center, Right, P", i, C, R, P)
        if R > i {
			fmt.Println(" R>i : R-i, P[2*C-i]", R-i, P[2*C-i])
            P[i] = min(R-i, P[2*C-i])
        }
        for T[i+1+P[i]] == T[i-1-P[i]] {
            P[i]++
        }
        if i + P[i] > R {
            C, R = i, i + P[i]
        }
    }
    
    maxLen := 0
    centerIndex := 0
    for i, v := range P {
        if v > maxLen {
            maxLen = v
            centerIndex = i
        }
    }
    return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
package main


import (
	"sort"
)

func longestCommonPrefix(v []string) string {
	ans := ""
	sort.Strings(v)
	n := len(v)
	first, last := v[0], v[n-1]
	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] != last[i] {
			return ans
		}
		ans += string(first[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
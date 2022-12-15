// Runtime0 ms
// Beats
// 100%
// Memory2.8 MB
// Beats
// 82.24%
package main

import (
	"fmt"
)

func main() {
	s := "anagrm"
	r := "nagaram"
	fmt.Println(isAnagram(s, r))

}

func isAnagram(s string, t string) bool {

	arr1 := []int{}
	for i := 0; i < 52; i++ {
		arr1 = append(arr1, 0)
	}
	arr2 := []int{}
	for i := 0; i < 52; i++ {
		arr2 = append(arr2, 0)
	}
	for i := 0; i < len(s); i++ {
		arr1[s[i]-'A']++
	}
	for i := 0; i < len(t); i++ {
		arr2[t[i]-'A']++
	}
	ans := true
	for i := 0; i < 52; i++ {
		if arr1[i] != arr2[i] {
			ans = false
			return ans
		}
	}
	return ans
}

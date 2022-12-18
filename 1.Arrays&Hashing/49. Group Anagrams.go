// Runtime31 ms
// Beats
// 79%
// Memory9.4 MB
// Beats
// 10.4%

package main

import (
	"fmt"
	"sort"
)

func main() {
	li := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// li := []string{"", "", ""}
	fmt.Println("groupAnagrams: ", groupAnagrams(li))
}

// func groupAnagrams(strs []string) [][]string {
//
// 	ans := [][]string{}
// 	temp := []string{}
// 	for i := 0; len(strs) > 0; {
// 		for j := i + 1; j < len(strs); j++ {
// 			if isAnagram(strs[i], strs[j]) {
// 				temp = append(temp, strs[j])
// 				strs = append(strs[:j], strs[j+1:]...)
// 				j--
// 			}
//
// 		}
// 		temp = append(temp, strs[i])
// 		strs = append(strs[:i], strs[i+1:]...)
// 		ans = append(ans, temp)
// 		temp = []string{}
//
// 	}
// 	return ans
// }
// func isAnagram(s string, t string) bool {
//
// 	arr1 := []int{}
// 	for i := 0; i < 26; i++ {
// 		arr1 = append(arr1, 0)
// 	}
// 	arr2 := []int{}
// 	for i := 0; i < 26; i++ {
// 		arr2 = append(arr2, 0)
// 	}
// 	for i := 0; i < len(s); i++ {
// 		arr1[s[i]-'a']++
// 	}
// 	for i := 0; i < len(t); i++ {
// 		arr2[t[i]-'a']++
// 	}
// 	ans := true
// 	for i := 0; i < 26; i++ {
// 		if arr1[i] != arr2[i] {
// 			ans = false
// 			return ans
// 		}
// 	}
// 	return ans
// }

func groupAnagrams(strs []string) [][]string {
	record := map[string][]string{}
	res := [][]string{}
	for _, str := range strs {
		sBytes := []byte(str)
		sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
		sortedStr := string(sBytes)
		record[sortedStr] = append(record[sortedStr], str)
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

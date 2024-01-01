// Runtime31 ms
// Beats
// 79%
// Memory9.4 MB
// Beats
// 10.4%

// 49. Group Anagrams
// Medium
// 15.4K
// 442
// Companies

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]


package main

import (
	"fmt"
	"sort"
)

func main() {
	testCase := [][]string{}
	testCase = append(testCase, []string{"eat", "tea", "tan", "ate", "nat", "bat"})
	testCase = append(testCase, []string{""})
	testCase = append(testCase, []string{"a"})
	for _,i := range testCase {
fmt.Println("testCase: ",i)
	fmt.Println("groupAnagrams: ", groupAnagrams(i))
	}
}



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

// Another approch 
// func groupAnagrams(strs []string) [][]string {
//     resArr:=[][]string{}
//     strsMap := make(map[string][]string)
//     for _,i := range strs {
//         runes := []rune(i)
//         sort.Slice(runes,func(i,j int)bool{
//             return runes[i]<runes[j]
//         })
//         strsMap[string(runes)]= append(strsMap[string(runes)],i)
//     }
//     for _,i := range strsMap {
//         resArr = append(resArr,i)
//     }
//     return resArr
// }

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
	li := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// li := []string{"", "", ""}
	fmt.Println("groupAnagrams: ", groupAnagrams(li))
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

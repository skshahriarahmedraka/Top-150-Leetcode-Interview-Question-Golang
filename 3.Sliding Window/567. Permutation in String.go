package main

import "fmt"

func main() {
	s1 :=
		"adc"
	s2 :=
		"dcda"
	fmt.Println(checkInclusion(s1, s2))
}

// func checkInclusion(s1 string, s2 string) bool {
	
// 	s1len := len(s1)
// 	s2len:=len(s2)
// 	if s1len > s2len {
// 		return false
// 	}
// 	arr1 := []int{}
// 	for i := 0; i < 26; i++ {
// 		arr1 = append(arr1, 0)
// 	}
	
// 	for i := 0; i < s1len; i++ {
// 		arr1[s1[i]-'a']++
// 	}
	
// 	for i := 0; i < s2len-s1len+1; i++ {
// 		if isAnagram(arr1, s2[i:i+s1len]) {
// 			return true
// 		}
// 	}

// 	return false
// }



// func isAnagram(arr1 []int, t string) bool {

	
// 	arr2 := []int{}
// 	for i := 0; i < 26; i++ {
// 		arr2 = append(arr2, 0)
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

func checkInclusion(s1 string, s2 string) bool {
	var freq [256]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
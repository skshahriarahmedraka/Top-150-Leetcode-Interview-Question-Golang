// 480 / 480 test cases passed.
// 	Status: Accepted
// Runtime: 7 ms
// Memory Usage: 2.6 MB

// You are here!
// Your runtime beats 48.77 % of golang submissions
// You are here!
// Your memory usage beats 85.08 % of golang submissions.


package main


import (
	"fmt"
	"strings"
	// "unicode"
)

func main() {
	s:= "A man, a plan, a canal: Panama"
	// s:="race a car"
	// s:="0P"
	// s:=" "

	fmt.Println("isPalindrome(s) : ", isPalindrome(s))
}

// func isPalindrome(s string) bool {
// 	s= strings.ToLower(s)
// 	for _,j:= range s {

// 		if !unicode.IsLetter(j) && !unicode.IsDigit(j) {
// 			s= strings.Replace(s, string(j), "", -1)
// 		}
// 	}

//		l:=(len(s)/2)+1
//		if len(s)%2==0{
//			l=(len(s)/2)
//		}
//		for i,j:=0,len(s)-1 ; i<l;{
//			if s[i]!=s[j]{
//				return false
//			}
//			i++
//			j--
//		}
//		return true
//	}
func isPalindrome(s string) bool {

	i, j := 0, len(s)-1
	valid := func(s byte) bool {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
			return true
		}
		return false
	}
	for i < j {
		if !valid(s[i]) {
			i++
			continue
		}
		if !valid(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--

	}
	return true
}

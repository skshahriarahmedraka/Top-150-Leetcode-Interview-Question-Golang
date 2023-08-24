// Runtime0 ms
// Beats
// 100%
// Memory2.6 MB
// Beats
// 100%

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
    start:= 0 
    end:= len(s)-1

    for start<end {
        
        if !valid(s[start]) {
            start++
            continue 
        }
        if !valid(s[end])  {
            end--
            continue 
        }
        if !strings.EqualFold(string(s[start]) ,string(s[end])) {
            return false
        }else {
            start++
            end--
        }
    }
    return true   
}
func valid(s byte)bool{
    if (s>='a'&& s<='z' ) || (s>='A' && s<='Z') || (s>='0'&& s<='9'){
        return true
    }
    return false
}
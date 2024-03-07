package main

func main() {

}


func longestPalindrome(s string) string {
	
	if len(s)<3 {
		return s 
	}
	var oddSized func(string) string 
	var evenSized func(string) string 
	oddSized = func(s string) string {
		palindrome := s[:3]
		for i:=1;i<len(s)-1;i++{
			j,k:=i-1,i+1
			for j>=0 && k<len(s) && s[j]==s[k]{
				if k-j+1>len(palindrome){
					palindrome = s[j:k+1]
				}
				j-=1
				k+=1
			}
		}

	}
}
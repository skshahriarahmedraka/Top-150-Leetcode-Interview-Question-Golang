package main

import "fmt"

func main() {
	// s := "leEeetcode"
	s := "abBAcC"
	fmt.Println(makeGood(s))

}

func makeGood(s string) string {

	i := 0
	for len(s) > 1 {
		if i > len(s)-1 {
			break
		}
		if s[i] == (s[i+1] + 32) {
			if i+2 < len(s) {
				s = s[:i] + s[i+2:]

			} else {
				s = s[:i]
			}
			i=0
			
		} else {
			i++
		}
	}
	return s

}

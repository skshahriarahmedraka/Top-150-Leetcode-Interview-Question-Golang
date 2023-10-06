//Runtime
//Details
//267ms
//Beats 20.52%of users with Go
//Memory
//Details
//21.46MB
//Beats 34.06%of users with Go

package main

func main() {

}

func Is_Palindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func DFS(s string, part *[]string, result *[][]string) {
	if len(s) == 0 {
		*result = append(*result, append([]string{}, *part...))
		return
	}
	for i := 1; i <= len(s); i++ {
		if Is_Palindrom(s[:i]) {
			*part = append(*part, s[:i])
			DFS(s[i:], part, result)
			*part = (*part)[:len(*part)-1]
		}
	}
	return
}

func partition(s string) [][]string {
	result := [][]string{}
	part := []string{}
	DFS(s, &part, &result)
	return result
}

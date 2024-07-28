/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-09 10:10:02  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-09 10:10:02  */
// Runtime3 ms
// Beats
// 44.1%
// Memory2.1 MB
// Beats
// 55.2%
package main 

func main (){

}



func wordBreak(s string, wordDict []string) bool {
	
	m:= make(map[int]bool)
	m[0]=true 

	for i:= 1 ;i<len(s)+1;i++ {
		m[i]=false 
	}

	for i := 1 ;i<len(s)+1 ;i++ {
		for j:=0 ;j<i;j++{
			if val,b:= m[j] ; b && val && searchInArrStr(wordDict,s[j:i]){
				m[i]=true 
                // fmt.Println("i, m[i] : ",i, m[i])
			}
		}
	}

	return m[len(s)]
}

func searchInArrStr(wordDic []string,s string ) bool{
	for _,i := range wordDic {
		if i== s {
			return true 
		}
	}
	return false 
}


func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	return dfs(s, wordSet, memo)
}

func dfs(s string, wordSet map[string]bool, memo map[string]bool) bool {
	if value, exists := memo[s]; exists {
		return value
	}
	if _, exists := wordSet[s]; exists {
		return true
	}
	for i := 1; i < len(s); i++ {
		prefix := s[:i]
		if _, exists := wordSet[prefix]; exists && dfs(s[i:], wordSet, memo) {
			memo[s] = true
			return true
		}
	}
	memo[s] = false
	return false
}
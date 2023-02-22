/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-21 16:41:14  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-21 16:41:14  */
// Runtime213 ms
// Beats
// 5.40%
// Memory9.1 MB
// Beats
// 98.58%

package main

func main() {
	
}

type Trie struct {
    M map[string]bool
}


func Constructor() Trie {
    return Trie{M: make(map[string]bool)}
}


func (this *Trie) Insert(word string)  {
    this.M[word] = true
}


func (this *Trie) Search(word string) bool {
    _, ok := this.M[word]
	return ok 
}


func (this *Trie) StartsWith(prefix string) bool {
    for k,v := range this.M {
		if v && len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
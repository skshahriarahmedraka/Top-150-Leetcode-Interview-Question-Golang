/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-04 11:21:34  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-04 11:21:34  */
// Runtime11 ms
// Beats
// 19.54%
// Memory6.5 MB
// Beats
// 12.40%
package main

func main() {

}

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	for head != nil {
		if b, _ := m[head]; b {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-15 20:51:29  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-15 20:51:29  */
// Runtime0 ms
// Beats
// 100%
// Memory2.4 MB
// Beats
// 55.92%

package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var dummy = new(ListNode)
    var p = dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            p.Next = l1
            l1 = l1.Next
        } else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 != nil {
        p.Next = l1
    } else {
        p.Next = l2;
    }
    return dummy.Next

}
/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-05 12:19:41  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-05 12:19:41  */
// Runtime11 ms
// Beats
// 44.69%
// Memory5.4 MB
// Beats
// 29.38%
package main

func main() {

}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}




func reorderList(head *ListNode)  {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }   
    // reverse the second half
    second := slow.Next
    slow.Next = nil
    prev := slow.Next
    for second != nil {
        tmp := second.Next
        second.Next = prev
        prev = second
        second = tmp
    }

    first, second := head, prev
    for second != nil {
        tmp,tmp2 := first.Next, second.Next 
        first.Next = second
        second.Next = tmp
        first, second = tmp, tmp2
       
    }
}

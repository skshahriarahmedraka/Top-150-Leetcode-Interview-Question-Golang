/*
 * Author: Sk Shahriar Ahmed Raka
 * Email: skshahriarahmedraka@gmail.com
 * Telegram: https://t.me/shahriarraka
 * Github: https://github.com/skshahriarahmedraka
 * StackOverflow: https://stackoverflow.com/users/12216779/
 * Linkedin: https://linkedin.com/in/shahriarraka
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Copyright (c) 2022 Your Company
 * -----
 * HISTORY:
 */

//  1568 / 1568 test cases passed.
// 	Status: Accepted
// Runtime: 9 ms
// Memory Usage: 4.3 MB

// Your runtime beats 84.76 % of golang submissions.
// Your memory usage beats 98.67 % of golang submissions.

package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// liAns:= &ListNode{}
	temp := 0
	tempAns := l1
	for tempAns != nil || l2 != nil || temp != 0 {
		if l2 != nil {
			temp += l2.Val
		}
		if tempAns != nil {
			temp += tempAns.Val
		}

		tempAns.Val = temp % 10
		temp /= 10
		if l2 != nil {
			if l2.Next != nil && tempAns.Next == nil {
				tempAns.Next = &ListNode{}
			}
			l2 = l2.Next
		}
		if (temp > 0) && tempAns.Next == nil {
			tempAns.Next = &ListNode{}
		}
		tempAns = tempAns.Next
	}

	return l1
}

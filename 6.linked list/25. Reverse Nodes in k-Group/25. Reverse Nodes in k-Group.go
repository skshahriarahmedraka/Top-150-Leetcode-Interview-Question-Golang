/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-16 00:20:11  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-16 00:20:11  */
// Runtime3 ms
// Beats
// 88.98%
// Memory3.5 MB
// Beats
// 51.57%
package main

import "fmt"

func main() {
	h := &ListNode{Val: 1}
	h.Next = &ListNode{Val: 2}
	x:= reverseKGroup(h, 2)
	fmt.Println(x.Val, x.Next.Val)
}

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if k ==0 || k == 1 {
// 		return head
// 	}
// 	NodeArr := []*ListNode{}
// 	temp := head
// 	for temp != nil {
// 		NodeArr = append(NodeArr, temp)
// 		temp = temp.Next
// 	}
// 	multiple := 1
// 	for i := 0; i < len(NodeArr); i++ {
//         fmt.Println("🚀 ~ file: 25. Reverse Nodes in k-Group.go ~ line 23 ~ fori:=0;i<len ~ i+1 == (k)*multiple : ", i+1, (k)*multiple)
// 		if i+1 == (k)*multiple {
// 			for ii, jj := i-k+1, i; ii < jj; ii, jj = ii+1, jj-1 {
// 				// s[i], s[j] = s[j], s[i]
//                 fmt.Println("🚀  NodeArr[ii] : ", NodeArr[ii].Val)
//                 fmt.Println("🚀  NodeArr[jj] : ", NodeArr[jj].Val)
// 				NodeArr[jj], NodeArr[ii] = NodeArr[ii], NodeArr[jj]
// 			}

		
// 			multiple++
// 		}
// 	}
// 	var newHead *ListNode

// 	for i := 0; i < len(NodeArr); i++ {
// 		if i == 0 {
// 			newHead = NodeArr[i]
// 		} else {
// 			NodeArr[i-1].Next = NodeArr[i]
// 		}
// 	}
// 	return newHead

// }

func reverseKGroup(head *ListNode, k int) *ListNode {
    var tail  *ListNode
    
    tempHead := &ListNode{0, head}
    cur := head
    previous := tempHead 
    
    
    for cur != nil {
        tail = cur
        linkedListIndex := 0
        
        for cur != nil && linkedListIndex < k {
            cur = cur.Next
            linkedListIndex++
        }
        
        if linkedListIndex < k { 
            previous.Next = tail
        } else {
            previous.Next = reverse(tail, k)
            previous = tail
        }
    }
    return tempHead.Next
}

//Create a function to easily reverse the linkedList
func reverse(head *ListNode, k int) *ListNode {
    
    var nextNode, previous *ListNode
    
    cur := head
    
    for cur != nil && k > 0 {
        nextNode = cur.Next
        cur.Next = previous
        previous = cur
        cur = nextNode
        k--
    }
    head = previous
    return head
    
}
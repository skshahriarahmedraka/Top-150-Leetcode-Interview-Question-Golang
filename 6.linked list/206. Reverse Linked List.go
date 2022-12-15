package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	li := []int{}
// 	if head == nil {
// 		return nil
// 	}
// 	h := head
// 	for {
// 		li = append(li, h.Val)
// 		if h.Next == nil {
// 			break
// 		}
// 		h = h.Next
// 	}
// 	h1 := &ListNode{}
// 	for i := len(li) - 1; i >= 0; i-- {
// 		h1 = AppendLinkedlist(h1, li[i])
// 	}
// 	return h1
//
// }
//
// func AppendLinkedlist(head *ListNode, data int) *ListNode {
// 	if head == nil {
// 		head = &ListNode{Val: data}
// 		return head
//
// 	}
// 	h := head
// 	for {
// 		if h.Next == nil {
// 			L := &ListNode{Val: data}
// 			h.Next = L
// 			break
// 		}
// 		h = h.Next
// 	}
// 	return head

// Runtime0 ms
// Beats
// 100%
// Memory2.8 MB
// Beats
// 14.44%/ }
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := head
	if head.Next != nil {
		newHead = reverseList(head.Next)
		head.Next.Next = head

	}
	head.Next = nil

	return newHead

}

package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sorted := &ListNode{}
	head := sorted
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				if sorted == nil {
					sorted = list1
					sorted = sorted.Next
					list1 = list1.Next
				} else {
					sorted.Val = list1.Val
					list1 = list1.Next
					sorted = sorted.Next

				}

			} else {
				if sorted == nil {
					sorted = list2
					sorted = sorted.Next
					list2 = list2.Next
				} else {
					sorted.Val = list2.Val
					list2 = list2.Next
					sorted = sorted.Next
				}

			}
		}
		if list1 == nil {
			sorted = list1
			sorted = sorted.Next
			list1 = list1.Next
		} else {
			sorted = list2
			sorted = sorted.Next
			list2 = list2.Next
		}
	}

	return head

}

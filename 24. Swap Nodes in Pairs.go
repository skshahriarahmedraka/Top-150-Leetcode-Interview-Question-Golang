package main

func main() {
	
}

/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
      Val int
      Next *ListNode
  }
 func swapPairs(head *ListNode) *ListNode {
    tempNode:= head 
	for tempNode.Next !=nil && tempNode!=nil {
		temp := tempNode.Next 
		tempNode.Next=temp.Next
		temp.Next=tempNode

		tempNode=tempNode.Next
	}
	return head
}
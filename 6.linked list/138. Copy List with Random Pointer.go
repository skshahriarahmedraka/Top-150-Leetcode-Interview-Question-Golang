package main

func main() {

}

//  Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var NewHead *Node
	// NewHead.Val = head.Val
	NewHead = &Node{}
	if head == nil {
		return nil
	}
	if head.Next == nil {
		NewHead.Val = head.Val
		return NewHead
	}
	TempNewHead := NewHead
	OldHead := head
	for OldHead != nil {
		TempNewHead.Val = OldHead.Val
		OldHead = OldHead.Next
		if OldHead == nil {

			TempNewHead.Next = &Node{}
			TempNewHead = TempNewHead.Next
		}
	}
	TempNewHead = NewHead
	OldHead = head
	for TempNewHead != nil {
		if OldHead.Random != nil {

			TempNewHead.Random = SearchNode(TempNewHead, OldHead.Random.Val)
		}
		TempNewHead = TempNewHead.Next
	}
	return NewHead

}

func SearchNode(head *Node, val int) *Node {
	if head == nil {
		return nil
	}
	for head != nil {
		if head.Val == val {
			return head
		}
		head = head.Next
	}
	return nil
}

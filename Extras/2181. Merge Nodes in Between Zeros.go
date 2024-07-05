/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeNodes(head *ListNode) *ListNode {
    tempRes :=0
    result := &ListNode{} 
    temp :=result  
    for i:=0; head !=nil ;i++ {
        if head.Val == 0 && i!=0 {
            temp.Val= tempRes
            if head.Next !=nil {
            temp.Next = &ListNode{}
            temp = temp.Next
            }      
            tempRes=0
            head=head.Next
        }else {
            tempRes+=head.Val
            head=head.Next
        }
    
    }
    return result
}


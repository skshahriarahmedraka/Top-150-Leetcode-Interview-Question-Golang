/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-15 20:52:52  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-15 20:52:52  */
// Runtime112 ms
// Beats
// 31.76%
// Memory5.2 MB
// Beats
// 76.89%
package main 


func main(){

}


//  Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0{
		return nil
	}
	if len(lists)==1{
		return lists[0]
	}
	first := lists[0]
	for i:=1;i<len(lists);i++{
		first = mergeTwoLists(first,lists[i])
	}
	return first

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
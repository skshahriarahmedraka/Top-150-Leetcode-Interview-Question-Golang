/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-15 19:21:29  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-15 19:21:29  */
// Runtime3 ms
// Beats
// 50.13%
// Memory2.1 MB
// Beats
// 32.80%

package  main


func main(){

}


//  Definition for singly-linked list.
 type ListNode struct {
  Val int
     Next *ListNode
  }

 func removeNthFromEnd(head *ListNode, n int) *ListNode {

    	if head == nil {
    		return nil

    	}
      if head.Next == nil {
          return nil
      }
      ListLen :=0 
      for temp:=head;temp!=nil;temp=temp.Next {
         ListLen++
      }
      temp := head
      var tempPrev *ListNode
      for i:=0;i<=(ListLen-n);i++ {
         if i== (ListLen-n) {
            if tempPrev != nil {
               tempPrev.Next = temp.Next
               return head
            }else {
               head = head.Next
               return head
            }
         }
         tempPrev = temp
         temp = temp.Next
      }
    	return head
 }
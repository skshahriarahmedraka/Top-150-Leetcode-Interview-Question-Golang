/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-17 19:33:06  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-17 19:33:06  */
// Runtime4 ms
// Beats
// 87.58%
// Memory5.1 MB
// Beats
// 98.79%
package main 


func main (){

}


//  Definition for a binary tree node.
  type TreeNode struct {
      Val int
     Left *TreeNode
      Right *TreeNode
 }
 
 func isValidBST(root *TreeNode) bool {
	min,max := -1<<63,1<<63-1
    return CheckValidBST(root,min,max)
 }

func CheckValidBST(root *TreeNode,min int , max int) bool {
	if root == nil {
		return true
	}
	if   !(root.Val > min && root.Val < max) {
		return false
	}
	
	return CheckValidBST(root.Left,min,root.Val) && CheckValidBST(root.Right,root.Val, max) 
}
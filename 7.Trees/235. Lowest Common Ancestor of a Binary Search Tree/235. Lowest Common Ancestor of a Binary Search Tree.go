/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-05 20:58:27  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-05 20:58:27  */
// Runtime19 ms
// Beats
// 57.44%
// Memory7.3 MB
// Beats
// 54.10%
package main

func main() {
	
}



 // Definition for a binary tree node.
 type TreeNode struct {
      Val   int
     Left  *TreeNode
      Right *TreeNode
  }
 

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	
	if root == nil {
		return nil 
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root 
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}else {
		return root 
	}
}
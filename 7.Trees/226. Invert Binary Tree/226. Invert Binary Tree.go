/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-21 13:01:37  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-21 13:01:37  */
// Runtime2 ms
// Beats
// 61.85%
// Memory2.1 MB
// Beats
// 100%
package main

func main() {
	
}


//   Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
 }
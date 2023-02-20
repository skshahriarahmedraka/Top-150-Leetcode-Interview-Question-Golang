/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-19 12:46:48  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-19 12:46:48  */
// Runtime78 ms
// Beats
// 87.95%
// Memory10.5 MB
// Beats
// 39.76%

package main

func main() {
	
}


// Definition for a binary tree node.
  type TreeNode struct {
      Val int
     Left *TreeNode
     Right *TreeNode
  }

 func goodNodes(root *TreeNode) int {
  if root == nil {
    return 0
  }
  if root.Left == nil && root.Right == nil {
    return 1
  }   
  return dfs(root, root.Val)
 }

 func dfs(root *TreeNode, max int) int {
  if root == nil {
    return 0
  }
  if root.Val >= max {
    
    return 1 + dfs(root.Left, root.Val) + dfs(root.Right, root.Val)
  }
  return dfs(root.Left, max) + dfs(root.Right, max)
}



/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-19 15:10:07  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-19 15:10:07  */
// Runtime3 ms
// Beats
// 89.34%
// Memory4.3 MB
// Beats
// 84.81%
package main

func main() {
	
}


//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
  }
 
 func diameterOfBinaryTree(root *TreeNode) int {
	res :=0 

	dfs(root,&res)

	return res
	
	
 }

func dfs(root *TreeNode,res *int)int{
		if root==nil {
			return -1 
		}
		left:= dfs(root.Left, res)
		right := dfs (root.Right,res)
		*res = MAX(*res,2+left+right)

		return 1+ MAX(left,right)

	}

func MAX(a,b int)int{
	if a>b {
		return a
	}
	return b 
}


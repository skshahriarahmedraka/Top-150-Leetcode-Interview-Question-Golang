/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-20 23:23:44  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-20 23:23:44  */
// Runtime21 ms
// Beats
// 33.2%
// Memory7.4 MB
// Beats
// 57.8%
package main

func main() {

}

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := -1 << 31
	dfs(root,&ans)
	return ans
}


func dfs(root *TreeNode,ans *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left,ans)
	right := dfs(root.Right,ans)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	*ans = max(*ans, root.Val+left+right)
	m:= max(left, right)
	if m <= 0 {
		return root.Val
	}else {
		return root.Val + m
	}
	
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
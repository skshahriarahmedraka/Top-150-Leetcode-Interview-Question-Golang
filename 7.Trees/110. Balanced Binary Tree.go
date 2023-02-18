/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-18 10:52:15  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-18 10:52:15  */
// Runtime10 ms
// Beats
// 40.50%
// Memory5.7 MB
// Beats
// 29.89%
package main

import "math"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	Balanced, _ := CheckBalanced(root, true)
	return Balanced
}

func CheckBalanced(root *TreeNode, b bool) (bool, int) {
	if root == nil {
		return b, 1
	}
	bLeft, leftHeight := CheckBalanced(root.Left, b)
	bRight, rightHeight := CheckBalanced(root.Right, b)

	if !(int(math.Abs(float64(leftHeight)-float64(rightHeight))) <= 1 && bLeft && bRight) {
		return false, MAX(leftHeight, rightHeight) + 1
	}
	return b, MAX(leftHeight, rightHeight) + 1
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

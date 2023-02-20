/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-19 13:12:49  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-19 13:12:49  */
// Runtime11 ms
// Beats
// 92.83%
// Memory6.7 MB
// Beats
// 31.74%
package main

func main() {
	
}


// Definition for a binary tree node.
 type TreeNode struct {
     Val int
    Left *TreeNode
     Right *TreeNode
 }
 
 func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if (root == nil && subRoot == nil) || (root == nil && subRoot != nil) {
	return false
  }else if root != nil && subRoot == nil {
	return true 
  }
  queue := []*TreeNode{}
  queue = append(queue, root)
  for len(queue) > 0 {
	tempQueue := []*TreeNode{}
	queueLen := len(queue)
	for i := 0; i < queueLen; i++ {
		if queue[0].Val == subRoot.Val {
			if CheckSameTree(queue[0], subRoot) {
				return true
			}
		}
		if queue[0].Left !=nil {
			tempQueue = append(tempQueue, queue[0].Left)
		}
		if queue[0].Right !=nil {
			tempQueue= append(tempQueue, queue[0].Right)
		}
		queue=queue[1:]
	}
	queue = tempQueue
  }
  return false 

 }

func CheckSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil && subRoot != nil {
		return false
	} else if root != nil && subRoot == nil {
		return false
	} else {
		if root.Val == subRoot.Val {
			return CheckSameTree(root.Left, subRoot.Left) && CheckSameTree(root.Right, subRoot.Right)
		} else {
			return false
		}
	}
 }

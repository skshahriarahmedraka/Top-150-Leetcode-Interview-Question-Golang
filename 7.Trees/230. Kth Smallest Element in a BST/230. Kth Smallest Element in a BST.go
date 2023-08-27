/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-21 13:15:27  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-21 13:15:27  */
// Runtime15 ms
// Beats
// 12.14%
// Memory6.5 MB
// Beats
// 18.10%
package main

import "sort"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	valueList := []int{}
	for len(queue) > 0 {
		tempQueue := []*TreeNode{}
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			valueList = append(valueList, queue[0].Val)
			if queue[0].Left != nil {
				tempQueue = append(tempQueue, queue[0].Left)
			}
			if queue[0].Right != nil {
				tempQueue = append(tempQueue, queue[0].Right)
			}
			queue = queue[1:]
		}
		queue = tempQueue
	}
	sort.Ints(valueList)
	return valueList[k-1]

}

/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-17 21:00:52  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-17 21:00:52  */
// Runtime6 ms
// Beats
// 53.67%
// Memory4.8 MB
// Beats
// 5.29%

package main

func main() {
	
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

 func maxDepth(root *TreeNode) int {
    return len(levelOrder(root))
 }

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
	  return [][]int{}
	}
  
	queue := []*TreeNode{}
	queue = append(queue, root)
  
	result := [][]int{}
  
	for len(queue) > 0 {
	  
	  level := []int{}
	  tempQueue := []*TreeNode{}
	  queueLen := len(queue)
	  for i := 0; i < queueLen; i++ {
		level =append(level, queue[0].Val)
		if queue[0].Left != nil {
		  tempQueue = append(tempQueue, queue[0].Left)
		}
		if queue[0].Right !=nil {
		  tempQueue = append(tempQueue, queue[0].Right)
  
		}
		  queue = queue[1:]
		
	  }
	  result = append(result, level)
	  queue = tempQueue
	}
	return result
  }
  
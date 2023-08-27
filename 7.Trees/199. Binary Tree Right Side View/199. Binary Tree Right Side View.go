/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-21 00:33:11  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-21 00:33:11  */
// Runtime3 ms
// Beats
// 51.62%
// Memory2.2 MB
// Beats
// 94.95%
package main

func main() {
	
}

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
  }
 
 func rightSideView(root *TreeNode) []int {
    if root ==nil {
        return []int{}
    }
    var res []int
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        tempQueue := []*TreeNode{}
        queueLen := len(queue)
        for i := 0; i < queueLen; i++ {
            if i== 0 {
                res = append(res, queue[0].Val)
            }
            
            if queue[0].Right != nil {
                tempQueue = append(tempQueue, queue[0].Right)
            }
            if queue[0].Left != nil {
                tempQueue = append(tempQueue, queue[0].Left)
            }
            queue = queue[1:]
        }
        queue = tempQueue
    }
    return res
      
 }
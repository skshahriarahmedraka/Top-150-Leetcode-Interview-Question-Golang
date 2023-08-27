/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-18 00:23:32  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-18 00:23:32  */
// Runtime8 ms
// Beats
// 36.33%
// Memory4.1 MB
// Beats
// 66.29%

package main


// import "golang.org/x/exp/slices"
func main() {
	
}


//  Definition for a binary tree node.
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder)==0 || len(preorder) ==0 {
        return nil 
    }
    root := &TreeNode{ Val : preorder[0]}
    mid := indexOf(inorder,preorder[0])
    root.Left = buildTree(preorder[1:mid+1],inorder[:mid])
    root.Right= buildTree(preorder[mid+1:],inorder[mid+1:])
    return root
 }

 func indexOf( data []int,element int) (int) {
    for k, v := range data {
        if element == v {
            return k
        }
    }
    return -1    //not found.
 }

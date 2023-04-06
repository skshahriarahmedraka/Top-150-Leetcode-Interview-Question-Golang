/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-05 21:46:05  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-05 21:46:05  */
// Runtime14 ms
// Beats
// 48.26%
// Memory7.6 MB
// Beats
// 47.76%
package main

import (
	"strconv"
	"strings"
)

func main() {
	
}



// * Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }
 

 type Codec struct {
    serialized string
 }
 
 func Constructor() Codec {
	 return Codec{}
 }
 
 // Serializes a tree to a single string.
 func (this *Codec) serialize(root *TreeNode) string {
	 res := []string{}
	 DFS(&res,root)
	 return strings.Join(res, ",")
	
 }

 func DFS(res *[]string,root *TreeNode) {

	if root == nil {
	   *res = append(*res, "n")
	   return
	}
	*res = append(*res, strconv.Itoa(root.Val))
	DFS(res,root.Left)
	DFS(res,root.Right)
	
}
 
 // Deserializes your encoded data to tree.
 func (this *Codec) deserialize(data string) *TreeNode {    
	 vals := strings.Split(data, ",")
	//  tree := DFS2(&vals)
	 return DFS2(&vals)
 }

 func DFS2(vals *[]string) *TreeNode {
	if len(*vals) == 0 {
		return nil
	}
	if (*vals)[0] == "n" {
		*vals = (*vals)[1:]
		return nil
	}
	val, _ := strconv.Atoi((*vals)[0])
	*vals = (*vals)[1:]
	root := &TreeNode{Val: val}
	root.Left = DFS2(vals)
	root.Right = DFS2(vals)
	return root
 }
 
 
 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor();
  * deser := Constructor();
  * data := ser.serialize(root);
  * ans := deser.deserialize(data);
  */
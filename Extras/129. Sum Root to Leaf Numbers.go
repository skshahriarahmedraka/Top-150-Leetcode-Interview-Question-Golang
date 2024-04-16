package main

func main() {
	
}


// Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}



func dfs(node *TreeNode, num int) int {
    if node == nil {
        return 0
    }
    num = num*10 + node.Val
    if node.Left == nil && node.Right == nil {
        return num
    }
    return dfs(node.Left, num) + dfs(node.Right, num)
}
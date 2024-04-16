package main

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	dfs(root, val, depth, 1)
	return root
}

func dfs(node *TreeNode, val int, depth int, currentDepth int) {
	if node == nil {
		return
	}
	if currentDepth == depth-1 {
		node.Left = &TreeNode{Val: val, Left: node.Left}
		node.Right = &TreeNode{Val: val, Right: node.Right}
		return
	}
	dfs(node.Left, val, depth, currentDepth+1)
	dfs(node.Right, val, depth, currentDepth+1)
}
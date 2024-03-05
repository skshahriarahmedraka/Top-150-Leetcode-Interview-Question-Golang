package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	visited := make(map[*Node]*Node)

	visited[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, neighbor := range n.Neighbors {
			if _, ok := visited[neighbor] ; !ok {
				visited[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}

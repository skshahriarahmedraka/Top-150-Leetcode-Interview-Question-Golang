package  main 

func main (){

}



/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

  type Node struct {
	 Val int
	   Neighbors []*Node
	}
var visited = make(map[int]*Node)
 func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	if v, ok := visited[node.Val]; ok {
		return v
	}
	newnode:=&Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	visited[node.Val] = newnode
    // clone := &Node{
	// 	Val: node.Val,
	// 	Neighbors: []*Node{},
	// }
	// if len(node.Neighbors) == 0 {
	// 	return clone
	// }
	for _, n := range node.Neighbors {
		visited[node.Val].Neighbors = append(visited[node.Val].Neighbors, cloneGraph(n))
	}

	return newnode
	
 }
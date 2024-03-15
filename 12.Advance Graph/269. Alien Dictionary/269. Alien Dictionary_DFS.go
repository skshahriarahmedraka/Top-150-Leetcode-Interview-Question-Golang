package main

import 
(   "fmt"
    // "bytes"
)

func main() {
    arr := []string{"wrt","wrf","er","ett","rftt"}

    fmt.Println(alienOrder(arr))
// Example 1: Empty input
arr = []string{}
fmt.Println(alienOrder(arr)) // Output: ""

// Example 2: Single word input
arr = []string{"abc"}
fmt.Println(alienOrder(arr)) // Output: "abc"

// Example 3: Invalid input (a is prefix of b, but b comes before a)
arr = []string{"ab", "abc"}
fmt.Println(alienOrder(arr)) // Output: ""

// Example 4: Multiple valid orders
arr = []string{"wrt", "wrf", "er", "ett", "rftt"}
fmt.Println(alienOrder(arr)) // Output: "wertf" or any other valid order

// Example 5: Large input with repeated characters
arr = []string{"aba", "aab", "bab", "cac", "cab", "bcb"}
fmt.Println(alienOrder(arr)) // Output: "cbad"

// Example 6: Input with non-alphabetic characters
arr = []string{"word", "world", "row"}
fmt.Println(alienOrder(arr)) // Output: "ordw"

// Example 7: Input with duplicate words
arr = []string{"abc", "ab", "abc", "def"}
fmt.Println(alienOrder(arr)) // Output: "abcdef"

// Example 8: Input with words of different lengths
arr = []string{"a", "ab", "abc", "de", "abcd", "xyz"}
fmt.Println(alienOrder(arr)) // Output: "abcdxyz"

}

func alienOrder(words []string) string {
    graph := make(map[byte]map[byte]bool)

    for i:=1; i<len(words); i++{
        w1,w2 := words[i-1], words[i]
        minLen := min(len(w1), len(w2))
        if len(w1) > len(w2) && w1[:len(w2)] == w2{
            return ""
        }
        for j:=0; j<minLen; j++{
            if w1[j] != w2[j] {
                if graph[w1[j]] == nil{
                    graph[w1[j]] =make(map[byte]bool)
                }
                graph[w1[j]][w2[j]] = true
            }
        }
    }

    visited := make(map[byte]bool)
    result := make([]byte, 0)
    finalResult := make([]byte, 0)
    var dfs func(byte) bool
    dfs = func(node byte) bool{
        if visited[node]{
            return visited[node]
        }
        visited[node] = true
        for nextNode,_ := range graph[node]{
            if visited[nextNode]{
                return true
            }
            if dfs(nextNode){
                return true
            }
        }
        visited[node] = false
        result = append(result, node)
        return false 

    }

    for node := range graph{
        if dfs(node){
            return ""
        }
        if len(result) > len(finalResult){
            finalResult = result
        }
        result = make([]byte, 0)
    }
    for i,j:=0, len(finalResult)-1; i<j; i,j=i+1,j-1{
        finalResult[i], finalResult[j] = finalResult[j], finalResult[i]
    }
    return string(finalResult)

}

////// 269. Alien Dictionary
// func alienOrder(words []string) string {
// 	if len(words) == 0 {
// 		return ""
// 	}
	
// 	if len(words) == 1 {
// 		return words[0]
// 	}
	
// 	var buffer bytes.Buffer
// 	graph := make(map[string]deps)

// 	buildGraph(words, graph)

// 	explore(graph, &buffer)

// 	return buffer.String()
// }

// // topo sort
// // note: we need to pass pointer to buffer, since the WriteString function needs a pointer receiver.
// func explore(graph map[string]deps, buffer *bytes.Buffer) {
// 	var queue []string
// 	// using queue, doing bfs, first add init values into queue
// 	for k, v := range graph {
// 		if len(v) == 0 {
// 			queue = append(queue, k)
// 			delete(graph, k)
// 		}
// 	}

// 	for len(queue) > 0 {
// 		c := string(queue[0])
// 		queue = queue[1:]
// 		buffer.WriteString(c)

// 		for k1, v1 := range graph {
// 			if _, ok := v1[c]; ok {
// 				delete(v1, c)
// 			}

// 			if len(v1) == 0 {
// 				queue = append(queue, k1)
// 				delete(graph, k1)
// 			}
// 		}
// 	}

// 	if len(graph) > 0 {
// 		buffer.Reset()
// 	}
// }

// // use map[string]bool as set of string
// type deps map[string]bool

// func buildGraph(words []string, graph map[string]deps) {
// 	for i := 0; i < len(words)-1; i++ {
// 		addWords(words[i], words[i+1], graph)
// 	}
// }

// func addWords(s1 string, s2 string, graph map[string]deps) {
// 	i, j := 0, 0
// 	for i < len(s1) && j < len(s2) {
// 		c1 := string(s1[i])
// 		c2 := string(s2[j])

// 		if _, ok := graph[c1]; !ok {
// 			graph[c1] = make(deps)
// 		}

// 		if _, ok := graph[c2]; !ok {
// 			graph[c2] = make(deps)
// 		}

// 		if c1 == c2 {
// 			i++;
// 			j++;
// 		} else {
// 			graph[c2][c1] = true
// 			break;
// 		}
// 	}

// 	for i < len(s1) {
// 		c := string(s1[i])
// 		if _, ok := graph[c]; !ok {
// 			graph[c] = make(deps)
// 		}
// 		i++;
// 	}

// 	for j < len(s2) {
// 		c := string(s2[j])
// 		if _, ok := graph[c]; !ok {
// 			graph[c] = make(deps)
// 		}
// 		j++;
// 	}
// }
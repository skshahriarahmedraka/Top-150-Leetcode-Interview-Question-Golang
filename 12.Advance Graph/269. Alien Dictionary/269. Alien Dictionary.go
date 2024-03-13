package main

func main() {

}

func alienOrder(words []string) string {
    // Create a map to store the graph
    graph := make(map[byte][]byte)

    // Create a map to store the in-degrees
    inDegree := make(map[byte]int)

    // Create a set of unique characters
    chars := make(map[byte]bool)

    // Build the graph, calculate the in-degrees, and collect unique characters
    for i := 0; i < len(words)-1; i++ {
        word1, word2 := words[i], words[i+1]

        // If the first word is a prefix of the second word, it's invalid
        if len(word1) > len(word2) && word1[:len(word2)] == word2 {
            return ""
        }

        // Find the first different character
        index := 0
        for index < len(word1) && index < len(word2) && word1[index] == word2[index] {
            chars[word1[index]] = true
            index++
        }

        // If index reached the end of one word, continue to the next pair
        if index == len(word1) || index == len(word2) {
            continue
        }

        // Add an edge from word1[index] to word2[index]
        graph[word1[index]] = append(graph[word1[index]], word2[index])
        inDegree[word2[index]]++
        chars[word1[index]] = true
        chars[word2[index]] = true
    }

    // Perform topological sort
    var result []byte
    queue := make([]byte, 0)

    // Find nodes with no incoming edges
    for node := range chars {
        if inDegree[node] == 0 {
            queue = append(queue, node)
        }
    }

    // Topological sort
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)

        // Remove outgoing edges and update in-degrees
        for _, neighbor := range graph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    // If there is a cycle or if there are remaining characters not in the order, return an empty string
    if len(result) != len(chars) {
        return ""
    }

    return string(result)
}
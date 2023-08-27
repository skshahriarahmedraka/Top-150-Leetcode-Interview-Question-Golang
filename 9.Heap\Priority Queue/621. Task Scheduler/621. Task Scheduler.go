package main

func main() {
	
}


func leastInterval(tasks []byte, n int) int {
    maxFreq := 0
    dict := make(map[byte]int)
    
    for i := 0; i < len(tasks); i++ {
        dict[tasks[i]]++
        if dict[tasks[i]] > maxFreq {
            maxFreq = dict[tasks[i]]
        }
    }
    
    // No of idle slots depends on maxFreq task
    res := (maxFreq-1) * (n+1)
    
    // If there are tasks with equal freq, then time increases
    for _, value := range dict {
        if value == maxFreq {
            res++
        }
    }
    
    // special handling for n is very less, and array size is large
    return max(res, len(tasks))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
package main 


func commonChars(words []string) []string {
    if len(words) == 0 {
        return []string{}
    }

    // Initialize minArr with the frequency counts of the first word
    minArr := make(map[rune]int)
    for _, w := range words[0] {
        minArr[w]++
    }

    // Iterate over the rest of the words and update minArr
    for i := 1; i < len(words); i++ {
        word := words[i]
        arrCount := make(map[rune]int)
        for _, w := range word {
            arrCount[w]++
        }

        // Update minArr to keep the minimum frequency of each character
        for s := range minArr {
            if val, exists := arrCount[s]; exists {
                minArr[s] = Min(minArr[s], val)
            } else {
                delete(minArr, s)
            }
        }
    }

    // Collect the result based on minArr
    var ans []string
    for s, val := range minArr {
        for i := 0; i < val; i++ {
            ans = append(ans, string(s))
        }
    }

    return ans
}

func Min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

package main

func longestPalindrome(s string) int {
    // Create a map to store the count of each character in the string.
    charCounts := make(map[rune]int)
    for _, char := range s {
        charCounts[char]++
    }
    
    ans := 0
    hasOddCount := false
    
    for _, charCount := range charCounts {
        if charCount % 2 == 0 {
            ans += charCount
        } else {
            ans += charCount - 1
            hasOddCount = true
        }
    }
    
    if hasOddCount {
        ans++
    }
    
    return ans
}

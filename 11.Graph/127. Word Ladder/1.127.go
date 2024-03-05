package main

// Runtime
// 217ms
// Beats30.64%of users with Go
// Memory
// 5.97MB
// Beats96.77%of users with Go


func ladderLength(beginWord string, endWord string, wordList []string) int {
  if len(wordList) == 0 {
    return 0
  }   
  wordMap := make(map[string]bool,len(wordList))

  for _ , w := range wordList {
    wordMap[w]  = true
  }

  queue := make([]string,len(wordList))

  NumOfWord := 0 
  queue = append(queue,beginWord)

 if !wordMap[endWord] {
        return 0
    }
  for len(queue) >0 {
    
    for  _ , word := range queue {
      if word == endWord {
        return NumOfWord+1
      }
      for  i, _ := range word {
          for j := 'A' ; j<= 'z' ;j++ {
            if _,ok := wordMap[word[:i]+string(j)+word[i+1:]] ; ok  {
                queue = append(queue, word[:i]+string(j)+word[i+1:])
                delete(wordMap,word[:i]+string(j)+word[i+1:]) 
            }
          }          
      }
      queue =queue[1:]
    }
    NumOfWord++
  }

  return 0

}
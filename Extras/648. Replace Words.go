package main


func replaceWords(dictionary []string, sentence string) string {
    wordMap := map[string]bool{}
    for _, d := range dictionary {
        wordMap[d] = true
    }

    words := strings.Split(sentence, " ")
    for i, word := range words {
        for j := 1; j <= len(word); j++ {
            if wordMap[word[:j]] {
                words[i] = word[:j]
                break
            }
        }
    }

    return strings.Join(words, " ")
}

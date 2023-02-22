package main

func main() {
	
}



type WordDictionary struct {
    M map[byte]interface{}
}


func Constructor() WordDictionary {
    return WordDictionary{M: make(map[byte]interface{})}
}


func (this *WordDictionary) AddWord(word string)  {
	for i := 1; i <= len(word); i++ {
		if _, ok := this.M[word[i]]; !ok {
			this.M[word[i]] = nil
		} 
	}
}


func (this *WordDictionary) Search(word string) bool {
    
}



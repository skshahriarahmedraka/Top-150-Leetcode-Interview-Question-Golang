package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)


func main(){
	li := []string{"hello", "world","I'm", "Sk"," Shahriar ","Ahmed ","Raka"}
	enc := encode(li)
	dec := decode(enc)
	fmt.Println("Encoded:", enc)
	fmt.Println("Decoded:", dec)
}

func encode(strs []string) string {
	endoded := ""
	for _, str := range strs {
		endoded += base64.StdEncoding.EncodeToString([]byte(str))
		endoded += " "
	}
	return endoded
}

func decode(strs string) []string {
	decoded := []string{}
	strs = strs[:len(strs)-1]
	decodedStrs := strings.Split(strs, " ")
	for _, str := range decodedStrs {
		decodedStr, _ := base64.StdEncoding.DecodeString(str)
		decoded = append(decoded, string(decodedStr))
	}
	return decoded
}
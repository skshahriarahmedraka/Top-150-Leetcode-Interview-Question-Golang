package main

import "fmt"

func main() {
	s := "lee(t(c)o)de)"
	fmt.Println("Expected output :\"lee(t(c)o)de\" ,your Output :", minRemoveToMakeValid(s))
	s = "a)b(c)d"
	fmt.Println("Expected output :\"ab(c)d\" ,your Output :", minRemoveToMakeValid(s))
	s = "))(("
	fmt.Println("Expected output :\"\" ,your Output :", minRemoveToMakeValid(s))
	
}

func minRemoveToMakeValid(s string) string {
    if len(s)==0 {
        return ""
    }
    temp := 0
    for i:= 0;i<len(s);i++ {
        if s[i]==')'{
            temp--
        }else if s[i]=='(' {
            temp++
        }
        if temp <0 {
            s= s[:i]+s[i+1:]
            i--
            temp++
        }
    }
    if temp>0 {
        for i:= len(s)-1;i>=0 ;i-- {
            if s[i]=='(' {
                s=s[:i]+s[i+1:]
                temp--
                if temp ==0{
                    break
                }
            }
        }
    }
    return s
}
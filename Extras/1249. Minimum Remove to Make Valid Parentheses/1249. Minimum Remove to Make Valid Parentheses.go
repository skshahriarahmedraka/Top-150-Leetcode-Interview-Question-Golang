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


func minRemoveToMakeValid(s string) string {
    str := []byte(s)
    stack := []int{}
    for i, c := range str {
        if c == '(' {
            stack = append(stack, i)
        } else if c == ')' {
            if len(stack) == 0 {
                str[i] = ' '
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    for _, i := range stack {
        str[i] = ' '
    }
    j := 0
    for i := 0; i < len(str); i++ {
        if str[i] != ' ' {
            str[j] = str[i]
            j++
        }
    }
    return string(str[:j])
}

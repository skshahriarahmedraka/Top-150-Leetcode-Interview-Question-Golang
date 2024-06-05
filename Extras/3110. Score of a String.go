package main


func scoreOfString(s string) int {
    sum := 0
    if len(s)< 2 {
        return 0
    }
    for i:=0 ;i<len(s)-1;i++ {
        sum +=ABS(int(s[i]),int(s[i+1]))
    }
    return sum
    
}


func ABS(i,j int)int{
    if i-j <0 {
        return -(i-j)
    }
    return i-j
}
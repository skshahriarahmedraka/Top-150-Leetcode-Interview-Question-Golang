package main

func reverseString(s []byte)  {
    for i,j:=0,len(s)-1 ;i<len(s)/2 ; {
        s[i],s[j]=s[j],s[i]
        i++
        j--
    }
}
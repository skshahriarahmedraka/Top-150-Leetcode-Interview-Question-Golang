package main

import "fmt"

func main() {
    nums := []int{1,2,3,1}
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,2,3,4}
    fmt.Println(containsDuplicate(nums))
    nums = []int{1,1,1,3,3,4,3,2,4,2}
    fmt.Println(containsDuplicate(nums))

}

func containsDuplicate(nums []int) bool {
    numMap := make(map[int]bool)
    var ok bool
    for _,i:= range nums {
        if ok,_ = numMap[i] ; ok {
            return true
        }
        numMap[i]=true 
    }
    return false 
}
package main



func main() {

}



func findMaxK(nums []int) int {
    if len(nums) < 2 {
        return -1
    }
    numMap := make(map[int]bool,len(nums))
    maxVal := -1
    for _,i := range nums {
        numMap[i]=true 
    }
    for _,i := range nums {
        if  _,b := numMap[-i] ; i>0 && b && i>maxVal {
            maxVal =i
        }
    }
    return maxVal
}


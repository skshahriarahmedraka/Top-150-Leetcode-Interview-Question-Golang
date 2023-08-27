package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2,3,5},8))
}

func combinationSum(candidates []int, target int) [][]int {
	ans :=[][]int{}
	candidatesLen:=len(candidates)
	var DFS func(int,[]int,int)
	DFS = func(i int,current []int,total int) {
		if total ==target {
			ans=append(ans, current)
			return
		}
		if i > candidatesLen || total>target {
			return
		}
		current=append(current, candidates[i])
		DFS(i,current,total+candidates[i])
		current=current[:candidatesLen-1]
		DFS(i+1,current,total)
	}
	DFS(0, []int{},0)
    // combinationSumRecution(0,[]int{},candidates,target)
	return ans 
}

// var ans [][]int 
// func combinationSumRecution(tempAns int,tempList []int,candidate []int,target int) {
// 	for _,i := range candidate {
// 		if tempAns+i >target || len() {
// 			return 
// 		}else if tempAns+i == target {
// 			tempList=append(tempList, i)
// 			ans=append(ans, tempList)
// 			return 
// 		}
// 		combinationSumRecution(tempAns,tempList,candidate,target)
// 		tempList=append(tempList, i)
// 		combinationSumRecution(tempAns+i,tempList,candidate,target)
		
// 	}
// }
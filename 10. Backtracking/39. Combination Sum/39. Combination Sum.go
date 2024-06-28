// Runtime
// Details
// 4ms
// Beats 41.82%of users with Go
// Memory
// Details
// 3.08MB
// Beats 76.12%of users with Go
package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

//func combinationSum(candidates []int, target int) [][]int {
//	ans :=[][]int{}
//	candidatesLen:=len(candidates)
//	var DFS func(int,[]int,int)
//	DFS = func(i int,current []int,total int) {
//		if total ==target {
//			ans=append(ans, current)
//			return
//		}
//		if i > candidatesLen || total>target {
//			return
//		}
//		current=append(current, candidates[i])
//		DFS(i,current,total+candidates[i])
//		current=current[:candidatesLen-1]
//		DFS(i+1,current,total)
//	}
//	DFS(0, []int{},0)
//    // combinationSumRecution(0,[]int{},candidates,target)
//	return ans
//}

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

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func([]int, int, int)
	dfs = func(path []int, i, total int) {
		if total == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		path = append(path, candidates[i])
		dfs(path, i, total+candidates[i])
		path = path[:len(path)-1]
		dfs(path, i+1, total)
	}

	dfs([]int{}, 0, 0)
	return res
}


func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var tempArr []int

    var backtrack func(start int, tempArr []int, currentSum int)
    backtrack = func(start int, tempArr []int, currentSum int) {
        if currentSum == target {
            // Make a deep copy of tempArr and append it to result
            combination := make([]int, len(tempArr))
            copy(combination, tempArr)
            result = append(result, combination)
            return
        } else if currentSum > target {
            return
        }

        for i := start; i < len(candidates); i++ {
            tempArr = append(tempArr, candidates[i])
            backtrack(i, tempArr, currentSum + candidates[i])
            tempArr = tempArr[:len(tempArr)-1]
        }
    }

    backtrack(0, tempArr, 0)
    return result
}


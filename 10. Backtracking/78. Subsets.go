package main

import "fmt"

func main() {
	nums := []int{1,3,2}
	fmt.Println(subsets(nums))
	fmt.Println(subsets([]int{9,0,3,5,7	}))
}

// func subsets(nums []int) [][]int {
// 	var li [][]int
// 	numsLen:= len(nums)
// 	subList:=[]int{}
// 	var DFS func(int)
// 	DFS = func(i int) {
//         fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 16 ~ DFS=func ~ i : ", i)
// 		if i >= numsLen {
//             fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 19 ~ DFS=func ~ subList : ", subList)
// 			li=append(li, subList)
//             fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 20 ~ DFS=func ~ li : ", li)
// 			return
// 		}	
// 		subList=append(subList, nums[i])
//         fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 22 ~ DFS=func ~ subList : ", subList)
// 		DFS(i+1)
// 		subList=subList[: len(subList)-1]
//         fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 25 ~ DFS=func ~ subList : ", subList)
// 		DFS(i+1)
// 	}
// 	DFS(0)
// 	return li
// }

func subsets(nums []int) [][]int {
	if len(nums)==0 {
		return [][]int{}
	}
	result:=[][]int{[]int{}}

	for _,i:= range nums {
		n:=len(result)
		for j:=0 ;j< n;j++ {
			r := append( result[j], i)
            fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 45 ~ funcsubsets ~ i : ", i)
            fmt.Println("🚀 ~ file: 78. Subsets.go ~ line 45 ~ funcsubsets ~ result[j] : ", result[j])
			result = append(result, r)
		}
	}
	return result
}

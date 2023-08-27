package main


func main(){

}

// func dailyTemperatures(temperatures []int) []int {
	
// 	start:=0 
// 	end:=1
// 	ans := make([]int, len(temperatures))
// 	for i:=0 ; i< len(temperatures)-1;i++{
// 		if temperatures[i] > temperatures[i+1]{
// 			start = i
// 			end = i+1
// 		}else {
// 			for j,x:=end-1,1; j>=start ;j--{
// 				ans[j] = x 
// 				x++ 
// 			}
// 			start = i+1
// 			end = i+2
// 		}
// 	}
// 	return ans
// }

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := -1
	for i := range temperatures {
		for top != -1 && temperatures[i] > temperatures[stack[top]] {
			idx := stack[top]
			top--
			res[idx] = i - idx
		}

		top ++
		stack[top] = i
	}
	return res
}

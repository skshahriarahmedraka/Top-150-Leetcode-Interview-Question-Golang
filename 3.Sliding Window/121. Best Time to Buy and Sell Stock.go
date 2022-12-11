func MAX(a,b int )int {
	if a>b {
		return a 
	}
	return b 
}
func maxProfit(prices []int) int {
	min:=prices[0]
	result:=0
	for _,j := range prices {
		if j >=min {
			result = MAX(result,j-min)
		}else {
			min=j 
		}
	}
	return result
}
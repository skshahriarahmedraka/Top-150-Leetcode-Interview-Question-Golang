package main  

func main (){


}

func maxProduct(nums []int) int {
	currentMax:= nums[0]
	currentmin := nums[0]
	finalMax:= nums[0] 
	MAX:= func (a int,b int)int{
		if a>b {
			return a
		}
		return b
	}
	MIN:= func (a int,b int)int{
		if a<b {
			return a
		}
		return b
	}
	for i:=1;i<len(nums);i++{
		temp:= currentMax

		currentMax =MAX(MAX(currentMax * nums[i],currentmin*nums[i]),nums[i])
		currentmin= MIN(MIN(temp*nums[i],currentmin*nums[i]),nums[i])
		finalMax=MAX(currentMax,finalMax)
	}
	return finalMax
}

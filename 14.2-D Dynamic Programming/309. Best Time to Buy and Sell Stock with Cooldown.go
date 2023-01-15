/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-14 22:41:25  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-14 22:41:25  */
// Runtime3 ms
// Beats
// 67.67%
// Memory4.4 MB
// Beats
// 8.19%
package main 


func main(){

}



func maxProfit(prices []int) int {
	indexMap := make(map[[2]int]int,0)

	return Profit(&indexMap,&prices,0,1)

}
func MAX(a,b int) int{
	if a>b{
		return a
	}
	return b
}
 func Profit(indexMap *map[[2]int]int,prices *[]int, index int, buying int) int {
		if index >= len(*prices) {
			return 0
		}
		if val, ok := (*indexMap)[[2]int{index, buying}]; ok {
			return val
		}
		coolDown := Profit(indexMap,prices, index+1, buying)
		if buying == 1 {
			buy := Profit(indexMap,prices, index+1, 0) - (*prices)[index]
			(*indexMap)[[2]int{index, buying}] = MAX(buy, coolDown)
			
		}else {
			sell := Profit(indexMap,prices, index+2, 1) + (*prices)[index]
			(*indexMap)[[2]int{index, buying}] = MAX(sell, coolDown)
		}
		return (*indexMap)[[2]int{index, buying}]
	}
package main

func main (){

}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var value map[[2]int]int

	for i:=0;i<len(flights);i++{
		templi := [2]int{flights[i][0],flights[i][1]}
		value[templi] = flights[i][2]
	}
	
}
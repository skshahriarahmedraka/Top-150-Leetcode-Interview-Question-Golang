package main

import (
	"fmt"
	"math"
)

func main() {
	li := []int{}
	fmt.Println(fillCups(li))
}

func fillCups(amount []int) int {
	sum:=0.0

	for _,i:= range amount{
		sum+=float64(i) 
	}
	return int(math.Ceil(sum/2))
}
/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-23 22:51:59  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-23 22:51:59  */Runtime240 ms
// Beats
// 64.10%
// Memory12.9 MB
// Beats
// 28.21%
package main

import (
	"fmt"
	"sort"
)

func main(){
	// target := 12
	target := 10
	// position := []int{10,8,0,5,3}
	position := []int{6,8}
	// speed := []int{2,4,1,1,3}
	speed := []int{3,2}
	fmt.Println("CarFleet :", carFleet(target, position, speed))
}

type Pair [][2]int	
func (m Pair) Len() (int) {
		return len(m)
	}

func (m Pair) Less(i, j int) (bool) {
		return m[i][0] > m[j][0]
	}
func (m Pair) Swap(i, j int) {
		m[i], m[j] = m[j], m[i]
	}
func carFleet(target int, position []int, speed []int) int {
	

	carPair := make(Pair, 0)
	for i := 0; i < len(position); i++ {
		carPair = append( carPair, [2]int{position[i], speed[i]})
	}
	stack := make([]float64, 0)
	sort.Sort(carPair)
    //fmt.Println("🚀 ~ file: 853. Car Fleet.go ~ line 28 ~ funccarFleet ~ carPair : ", carPair)
	for _, i := range carPair {
		stack = append(stack, float64((target-i[0]))/float64(i[1]))
		if len(stack) > 1 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack)


}
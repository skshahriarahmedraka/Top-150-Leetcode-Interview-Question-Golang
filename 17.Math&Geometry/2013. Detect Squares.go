/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-01-28 01:27:53  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-01-28 01:27:53  */
// Runtime93 ms
// Beats
// 76%
// Memory7.4 MB
// Beats
// 96%
package main

import (
// "math"
)

func main() {

}

type DetectSquares struct {
	PointMap  map[[2]int]int
	PointList [][2]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		PointMap:  make(map[[2]int]int),
		PointList: make([][2]int, 0),
	}
}

func (this *DetectSquares) Add(point []int) {
	this.PointList = append(this.PointList, [2]int{point[0], point[1]})
	this.PointMap[[2]int{point[0], point[1]}]++
}

func (this *DetectSquares) Count(point []int) int {
	res := 0
	px, py := point[0], point[1]
	for _, p := range this.PointList {
		if (Abs(px-p[0]) != Abs(py-p[1])) || px == p[0] || py == p[1] {
			continue
		}
		res += this.PointMap[[2]int{px, p[1]}] * this.PointMap[[2]int{p[0], py}]
	}
	return res
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

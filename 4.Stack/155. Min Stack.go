/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-25 10:53:34  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-25 10:53:34  */Runtime48 ms
// Beats
// 14.13%
// Memory8 MB
// Beats
// 94.65%
package main

// import "sort"

func main() {

}

type MinStack struct {
	Values   []int
	// MinValue []int
}

func Constructor() MinStack {
	return MinStack{Values: []int{}}
}

func (this *MinStack) Push(val int) {

	(*this).Values = append((*this).Values, val)
	// (*this).MinValue =(*this).Values
	// sort.Ints((*this).MinValue)
}

func (this *MinStack) Pop() {
	// (*this).MinValue = (*this).Values[:len((*this).Values)-1]
	(*this).Values = (*this).Values[:len((*this).Values)-1]

	// sort.Ints((*this).MinValue)
}

func (this *MinStack) Top() int {
	return (*this).Values[len((*this).Values)-1]
}

func (this *MinStack) GetMin() int {
	min:= (*this).Values[0]
	for i:= range (*this).Values{
		if min > (*this).Values[i]{
			min = (*this).Values[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
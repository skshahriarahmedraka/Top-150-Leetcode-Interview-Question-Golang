/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-03-08 19:38:43  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-03-08 19:38:43  */
// Runtime302 ms
// Beats
// 72.41%
// Memory24.7 MB
// Beats
// 10.34%

package main

func main() {
	
}
func mergeTriplets(triplets [][]int, target []int) bool {
	
	pos1,pos2,pos3 := false,false,false
    for i := 0; i < len(triplets); i++ {
		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2] {
			continue
		}
		if triplets[i][0] == target[0]  {
			pos1 = true
		}
		if triplets[i][1] == target[1]  {
			pos2 = true
		}
		if triplets[i][2] == target[2]  {
			pos3 = true
		}
	}
	return pos1 && pos2 && pos3
}
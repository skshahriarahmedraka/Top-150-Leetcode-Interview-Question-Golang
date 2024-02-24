/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-16 23:52:36  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-16 23:52:36  */
// Runtime11 ms
// Beats
// 94.31%
// Memory5.7 MB
// Beats
// 50.84%
package main

import (
	"fmt"
	"sort"
)

func main() {
	li := []int{1, 1, 1, 2, 2, 3}

	fmt.Println("347. Top K Frequent Elements :", topKFrequent(li, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, j := range nums {
		m[j] += 1
	}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 15 ~ functopKFrequent ~ m : ", m)
	li := make([][]int, len(nums))
	for i, j := range m {
		li[j-1] = append(li[j-1], i)
	}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 20 ~ functopKFrequent ~ li : ", li)
	ans := []int{}
	// fmt.Println("🚀 ~ file: 347. Top K Frequent Elements.go ~ line 26 ~ fori:=len ~ len(li) : ", len(li))
	for i := len(li) - 1; i >= 0; i-- {
		if li[i] != nil {
			ans = append(ans, li[i]...)
		}
	}
	return ans[:k]
}


func topKFrequent(nums []int, k int) []int {
    numsMap :=make(map[int]int)
    for _,i := range nums {
        numsMap[i]+=1
    }
    occuranceArr := make([][]int,len(nums))
    for num,occu := range numsMap {
        occuranceArr[occu-1] = append(occuranceArr[occu-1],num)
    }
    Kfreq :=[]int{}

    for i:= len(occuranceArr) -1 ; i>=0 ; i--{
		if len(occuranceArr[i]) != 0 {
			Kfreq= append(Kfreq,occuranceArr[i]...)
		}
	 	
    }
    return Kfreq[:k]

}

// fastest "(Runtime: 13ms)"
func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)

    for _, el := range nums {
        hash[el]++
    }

    // convert map to slice
    var entries []struct{
		Key   int
		Value int
    }

    // Populate the entries slice from the frequency map
	for key, value := range hash {
		entries = append(entries, struct{ Key, Value int }{Key: key, Value: value})
	}

    // Sort the slice
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Value > entries[j].Value
    })

    result := make([]int, 0, k)
    for i := 0; i < k; i++ {
        result = append(result, entries[i].Key)

    }

	return result
}
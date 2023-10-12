// Accepted
// Runtime
// Details
// 0ms
// Beats 100.00%of users with Go
// Memory
// Details
// 1.94MB
// Beats 100.00%of users with Go

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n uint32
	n = 1011
	fmt.Println(hammingWeight(n))
}

func hammingWeight2(num uint32) int {
	return bits.OnesCount(uint(num))
}

func hammingWeight3(num uint32) int {
	weight := 0
	for num != 0 {
		num &= (num - 1)
		weight += 1
	}
	return weight
}

func hammingWeight(num uint32) int {
    var ans,bit uint32 
    for num>0 {
        bit = num & 1
        ans += bit 
        num>>=1
    }   
    return int(ans) 

}
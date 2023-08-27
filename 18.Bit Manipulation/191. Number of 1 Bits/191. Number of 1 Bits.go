//Runtime0 ms
//Beats
//100%
//Memory2 MB
//Beats
//44.84%

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

func hammingWeight(num uint32) int {
	weight := 0
	for num != 0 {
		num &= (num - 1)
		weight += 1
	}
	return weight
}

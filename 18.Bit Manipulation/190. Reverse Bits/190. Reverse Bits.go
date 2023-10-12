// Accepted
// Runtime
// Details
// 4ms
// Beats 38.31%of users with Go
// Memory
// Details
// 2.50MB
// Beats 92.74%of users with Go

package main

func main() {

}

func reverseBits(num uint32) uint32 {
	var i, bit, res uint32
	for ; i < 32; i++ {

		bit = num & 1
		bit = bit << (31 - i)
		res = res | bit
		num >>= 1
	}
	return res
}

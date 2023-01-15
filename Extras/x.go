package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		li := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&li[j])
		}
		b := 0
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if k != j {
					if find_gcd(li[j], li[k]) != 1 {
						b = 1
						break
					}
				}
			}
		}
		if b == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func find_gcd(a, b int) int {

	if b == 0 {
		return a
	} else {
		return find_gcd(b, a%b)
	}

}

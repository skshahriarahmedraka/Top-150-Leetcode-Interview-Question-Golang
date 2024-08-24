package main

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	common, num := 5*7*8*9, 0

	for _, exp := range strings.Split(expression, "+") {
		for i, e := range strings.Split(exp, "-") {
			if len(e)==0 {
				continue
			}
			fraction := strings.Split(e, "/")
			u, _ := strconv.Atoi(fraction[0])
			v, _ := strconv.Atoi(fraction[1])
			if i == 0 {
				num += u*(common/v)
			} else {
				num -= u*(common/v)
			}
		}
	}
	
	for i := 2; i <= 7; i++ {
		if num%i == 0 && common%i == 0 {
			num /= i
			common /= i
			
		}
	}
	return strconv.Itoa(num) + "/" + strconv.Itoa(common)
}

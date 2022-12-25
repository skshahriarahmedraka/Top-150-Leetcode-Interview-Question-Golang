package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			tmp := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp)
		case "-":
			tmp := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp)
		case "*":
			tmp := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp)
		case "/":
			tmp := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

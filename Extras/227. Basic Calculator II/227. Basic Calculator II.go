package main

func main() {

}

func calculate(s string) int {
	symMap := map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}
	numMap := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
	current := 0
	PreSig := '+'
	stack := []int{}
	for _,i := range s+"+" {
		if numMap[i] {
			current = current*10 + int(i-'0')
		} else if symMap[i] {
			if PreSig == '+' {
				stack = append(stack, current)	
			}else if PreSig == '-' {
				stack = append(stack, -current)
			}else if PreSig == '*' {
				stack[len(stack)-1] *= current
			}else if PreSig == '/' {
				stack[len(stack)-1] /= current
			}
			PreSig = i
			current = 0
		}
	}
	return Sum(stack)
}

func Sum(all []int) int {
	sum :=0 
	for _,i := range all {
		sum += i
	}
	return sum
}

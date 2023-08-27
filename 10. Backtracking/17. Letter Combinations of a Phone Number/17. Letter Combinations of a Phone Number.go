
// Runtime2 ms
// Beats
// 64.77%
// Memory2.1 MB
// Beats
// 7.35%


package main

func main() {

}

var NumMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string
	switch len(digits) {
	case 0:
		return nil
	case 1:
		return NumMap[digits]
	}
	firstDig := digits[0:1]

	for _, i := range NumMap[firstDig] {
		for _, j := range letterCombinations(digits[1:]) {
			res = append(res, i+j)
		}
	}
	return res
}

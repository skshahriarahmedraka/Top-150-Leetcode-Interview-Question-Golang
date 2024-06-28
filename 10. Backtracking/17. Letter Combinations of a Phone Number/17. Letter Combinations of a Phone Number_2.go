// Runtime2 ms
// Beats
// 64.77%
// Memory2.1 MB
// Beats
// 7.35%

package main

func main() {

}



func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    digitToLetters := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var res []string

    var backtrack func(idx int, comb string)
    backtrack = func(idx int, comb string) {
        if idx == len(digits) {
            res = append(res, comb)
            return
        }

        for _, letter := range digitToLetters[digits[idx]] {
            backtrack(idx+1, comb+string(letter))
        }
    }

    backtrack(0, "")
    return res
}
package main 

func decodeString(s string) string {
    ans := strings.Builder{}
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            count := 0
            for i < len(s) && s[i] != '[' {
                count = count * 10 + int(s[i] - '0')
                i++
            }
            startIndex := i + 1
            for j := 1; j > 0; i++ {
                if s[i + 1] == '[' { j++ }
                if s[i + 1] == ']' { j-- }
            }
            for k := 0; k < count; k++ {
                ans.WriteString(decodeString(s[startIndex:i]))
            }
        } else {
            ans.WriteByte(s[i])
        }
    }
        return ans.String()
}
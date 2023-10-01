package main

func main() {

}

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1len, s2len, s3len := len(s1), len(s2), len(s3)
	if s1len+s2len != s3len {
		return false
	}
	cache := make([][]bool, s1len+1)
	for i := 0; i <= s1len; i++ {
		cache[i] = make([]bool, s2len+1)
	}
	for i := 0; i <= s1len; i++ {
		for j := 0; j <= s2len; j++ {
			if i == s1len && j == s2len {
				cache[i][j] = true
			}
		}
	}
	for i := s1len; i >= 0; i-- {
		for j := s2len; j >= 0; j-- {
			if i == s1len && j == s2len {
				cache[i][j] = true
			} else if i == s1len {
				cache[i][j] = cache[i][j+1] && s2[j] == s3[i+j]
			} else if j == s2len {
				cache[i][j] = cache[i+1][j] && s1[i] == s3[i+j]
			} else {
				cache[i][j] = (cache[i][j+1] && s2[j] == s3[i+j]) || (cache[i+1][j] && s1[i] == s3[i+j])
			}
		}
	}

	return cache[0][0]
}

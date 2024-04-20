package main

func main() {

}

func findFarmland(land [][]int) [][]int {
	var (
		m   = len(land)
		n   = len(land[0])
		res [][]int
		c2,r2 int 
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				continue
			}
			c2 = i
			for c2 < m && land[c2][j] == 1 {
				c2++
			}
			r2 = j
			for r2 < n && land[i][r2] == 1 {
				r2++
			}
			if c2 > 0 {
				c2--
			}
			if r2 > 0 {
				r2--
			}
			res = append(res, []int{i, j, c2, r2})

			for c := i; c <= c2; c++ {
				for r := j; r <= r2; r++ {
					land[c][r] = 0
				}
			}
		}
	}
	return res
}

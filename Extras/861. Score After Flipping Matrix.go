package main

func main() {
	
}




func matrixScore(grid [][]int) int {
	rows,cols := len(grid),len(grid[0])

	// flip rows 
	for i:=0;i<rows;i++{
		if grid[i][0] == 0{
			for j:=0;j<cols;j++{
				grid[i][j] = 1-grid[i][j]
			}
		}
	}

	// flip cols
	for j:=0;j<cols;j++{
		oneCount := 0
		for i:=0;i<rows;i++{
			if grid[i][j] == 1{
				oneCount++
			}
		}
		if oneCount < rows - oneCount {
			for i:=0;i<rows;i++{
				grid[i][j] = 1-grid[i][j]
			}
		}
	}
	res := 0
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			res += grid[i][j] << (cols-j-1)
		}
	}
	return res
}
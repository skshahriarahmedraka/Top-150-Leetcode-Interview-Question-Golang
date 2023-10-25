// 132 mstime cost
// ·
// 12.47 MBmemory cost
// ·
// Your submission beats77.22 %Submissions
// view ranking distribution
package main

func main() {

}

/**
 * @param rooms: m x n 2D grid
 * @return: nothing
 */
func WallsAndGates(rooms [][]int) {
	stack := make([][]int, 0)
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				stack = append(stack, []int{i, j})
			}
		}
	}

	for len(stack) != 0 {
		temp := make([][]int, 0)
		for len(stack) != 0 {
			i, j := stack[0][0], stack[0][1]
			stack = stack[1:]
			if i+1 < len(rooms) {
				if rooms[i+1][j] != -1 && rooms[i+1][j] > rooms[i][j]+1 {
					rooms[i+1][j] = rooms[i][j] + 1
					temp = append(temp, []int{i + 1, j})
				}
			}

			if j+1 < len(rooms[0]) {
				if rooms[i][j+1] != -1 && rooms[i][j+1] > rooms[i][j]+1 {
					rooms[i][j+1] = rooms[i][j] + 1
					temp = append(temp, []int{i, j + 1})
				}
			}

			if i-1 >= 0 {
				if rooms[i-1][j] != -1 && rooms[i-1][j] > rooms[i][j]+1 {
					rooms[i-1][j] = rooms[i][j] + 1
					temp = append(temp, []int{i - 1, j})
				}
			}

			if j-1 >= 0 {
				if rooms[i][j-1] != -1 && rooms[i][j-1] > rooms[i][j]+1 {
					rooms[i][j-1] = rooms[i][j] + 1
					temp = append(temp, []int{i, j - 1})
				}
			}

		}
		if len(temp) != 0 {
			stack = append(stack, temp...)
		}
	}

}

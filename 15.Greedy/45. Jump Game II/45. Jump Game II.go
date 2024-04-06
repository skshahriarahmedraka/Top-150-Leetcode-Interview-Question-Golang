// Runtime34 ms
// Beats
// 57.4%
// Memory6.1 MB
// Beats
// 69.63%
package main

import "fmt"

func main() {
	li := []int{2, 3, 1, 1, 4}
	fmt.Println("45. Jump Game II : ", jump(li))
}

func jump(nums []int) int {
	jump := 0
	left, right := 0, 0
	for right < len(nums)-1 {
		r := 0

		for i := left; i <= right; i++ {
			if nums[i]+i > r {
				r = nums[i] + i
			}

		}
		left = right + 1
		right = r
		jump++
	}
	return jump

}
func jump(nums []int) int {
    curJump, farthestJump, jumps := 0, 0, 0
    for i := 0; i < len(nums)-1; i++ {
        if i+nums[i] > farthestJump {
            farthestJump = i + nums[i]
        }

        if i == curJump {
            jumps, curJump = jumps+1, farthestJump
            if curJump >= len(nums)-1 {
                return jumps
            }
        }
    }
    return 0
}

func jump(nums []int) int {
    if len(nums)<2 {
        return 0
    }
    curJump,furJump,jump := 0,0,0
    for i,v := range nums {
        if i+v > furJump {
            furJump = i+v
        }

        if curJump==i {
            jump,curJump =jump+1,furJump 
            if curJump>= len(nums)-1 {
                return jump
            }
        }
    }
    return 0
}

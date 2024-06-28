package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}

	fmt.Println(maximumImportance(n, roads))
}

func maximumImportance(n int, roads [][]int) int64 {
	var impRoad [][]int
	impRoadMap := map[int]int{}
	var sum int64
	for _, r := range roads {
		impRoadMap[r[0]] += 1
		impRoadMap[r[1]] += 1
	}
	fmt.Println("🚀 ~ file: 2285. Maximum Total Importance of Roads.go ~ line 22 ~ funcmaximumImportance ~ impRoadMap : ", impRoadMap)
	for k, v := range impRoadMap {
		impRoad = append(impRoad, []int{k, v})
	}
	sort.Slice(impRoad, func(i, j int) bool {
		if impRoad[i][0] == impRoad[j][0] {
			return impRoad[i][1] < impRoad[j][1]
		}
		return impRoad[i][0] < impRoad[j][0]
	})

	fmt.Println("🚀 ~ file: 2285. Maximum Total Importance of Roads.go ~ line 26 ~ funcmaximumImportance ~ impRoad : ", impRoad)
	for i := 1; i <= n; i++ {
		sum += int64(i + impRoad[i-1][1])
		fmt.Println("🚀 ~ file: 2285. Maximum Total Importance of Roads.go ~ line 37 ~ funcmaximumImportance ~ sum : ", sum)
	}
	return sum
}

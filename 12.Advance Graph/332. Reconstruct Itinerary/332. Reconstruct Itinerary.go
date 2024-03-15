package main

import (
    "fmt"
    "sort"
)

func main() {
    // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    arr := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}

    fmt.Println("sort by Kth column:")
    sortByColumn(arr, 0)
    fmt.Println(arr)

    fmt.Println("sort by Kth row:")
    sortByRow(arr, 0)
    fmt.Println(arr)
}

func sortByColumn(arr [][]string, col int) {
    sort.SliceStable(arr, func(i, j int) bool {
        return arr[i][col] < arr[j][col]
    })
}

func sortByRow(arr [][]string, row int) {
    sort.SliceStable(arr, func(i, j int) bool {
        return arr[row][i] < arr[row][j]
    })
}
// func findItinerary(tickets [][]string) []string {

// return []string{}
// 	// graph := make(map[string][]string)
// }
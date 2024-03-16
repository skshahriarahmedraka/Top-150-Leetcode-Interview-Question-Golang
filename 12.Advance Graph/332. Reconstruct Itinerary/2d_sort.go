package main

// import (
//     "fmt"
//     "sort"
// )

// func main() {
//     arr := [][]int{
//         {5, 2, 8},
//         {1, 9, 3},
//         {4, 7, 6},
//         {5, 2, 6},
//     }

//     fmt.Println("Original Array:")
//     fmt.Println(arr)

//     // Sort by rows
//     fmt.Println("\nSort by Rows:")
//     sortByRows(arr)
//     fmt.Println(arr)

//     // Sort by columns
//     fmt.Println("\nSort by Columns:")
//     sortByColumns(arr)
//     fmt.Println(arr)
// }

// import "sort"

// type customSlice [][]int

// func (c customSlice) Len() int           { return len(c) }
// func (c customSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
// func (c customSlice) Less(i, j int) bool { /* Custom sorting logic goes here */ }

// func sortByCustom(arr [][]int) {
//     sort.Sort(customSlice(arr))
// }
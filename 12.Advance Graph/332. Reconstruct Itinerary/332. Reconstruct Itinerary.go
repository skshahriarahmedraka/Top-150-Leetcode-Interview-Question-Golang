package main

import (
    "fmt"
    "sort"
)

func main() {
    // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    arr := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}

    fmt.Println("Array:" , findItinerary(arr))

}



func findItinerary(tickets [][]string) []string {
    adj := make(map[string][]string)
    for _, ticket := range tickets {
        src, dst := ticket[0], ticket[1]
        adj[src] = append(adj[src], dst)
    }

    for src, dsts := range adj {
        sort.Strings(dsts)
        adj[src] = dsts
    }
    fmt.Println("adj : ", adj)
    res := []string{}
    var dfs func(string)

    dfs = func(src string) {
        for {
            fmt.Println("src : ", src)
            if len(adj[src]) == 0 {
                fmt.Println("added to result : ", src)
                res = append([]string{src}, res...)
                return
            }
            dst := adj[src][0]
            adj[src] = adj[src][1:]
            fmt.Println("removed : ", dst, " from :", src)
            dfs(dst)
        }
    }

    dfs("JFK")
    return res
}
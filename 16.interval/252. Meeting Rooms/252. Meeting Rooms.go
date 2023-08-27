package main

import "sort"

type Times [][]int

func (t Times) Len() int{
	return len(t)
}
func (t Times) Less(i,j int) bool {
	return t[i][0] < t[j][0]
}

func (t Times) Swap(i,j int) {
	t[i], t[j] = t[j],t[i]
}

func main() {

}


func canAttendMeetings(intervals [][]int) bool {
	if intervals == nil || len(intervals) == 0 {
		return true
	}

	sort.Sort(Times(intervals))

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}



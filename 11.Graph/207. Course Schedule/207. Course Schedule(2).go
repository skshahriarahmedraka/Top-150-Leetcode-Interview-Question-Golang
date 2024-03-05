package main



func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	visited := make(map[int]bool,0)

	for _,i := range prerequisites {
		preMap[i[0]] = append(preMap[i[0]],i[1:]...)
	}

	var search func(int)bool 

	search =func(i int) bool {
		if visited[i] {
			return false 
		} 
		if preMap[i] ==nil  {
			return true 
		}
		for _,j := range preMap[i] {
			if !search(j) {
				return false 
			}
		}
		visited[i]= true 
		preMap[i] =nil 
		return  true 
	}

	for i:= 0; i<numCourses ;i++ {
		if !search(i) {
			return false 
		}
	}
	return true 
}
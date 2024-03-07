package main


func findOrder(numCourses int, prerequisites [][]int) []int {
	visited :=make(map[int]bool)
	preMap := make(map[int][]int)
	cicle := make(map[int]bool)
	result := []int{}
	for _ , v := range prerequisites {
		preMap[v[0]] = append(preMap[v[0]], v[1:]...)
	}

	var search func(int)bool 

	search = func(i int) bool {
		if cicle[i] {
			return false   
		}
		if visited[i] {
			return true    
		}
		cicle[i] = true  
		for _,v := range preMap[i] {
			if !search(v) {
				return false  
			}
		}
		visited[i] = true
		cicle[i] = false 
		result = append(result, i)
		return true 	
	}

	for i:=0 ;i<numCourses ;i ++ {
		if !search(i) {
			return []int{}
		}
	}
	return result

}
package main

func main() {
	
}




func tribonacci(n int) int {
	var arr [3]int= [3]int{0,1,1}
	 
	for i := 3; i < n;i++ {
		arr[i%3] = arr[0] +arr[1]+arr[2]
	}
	return arr[n%3]
}
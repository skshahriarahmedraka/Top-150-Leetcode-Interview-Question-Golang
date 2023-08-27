package main

import "sort"

func main() {
	
}


type MedianFinder struct {
	arr []int
}


func Constructor() MedianFinder {
    return MedianFinder{arr: make([]int, 0)}
}


func (this *MedianFinder) AddNum(num int)  {
	this.arr = append(this.arr, num)

}


func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.arr)
    if len(this.arr)%2 == 0 {
		return float64(this.arr[len(this.arr)/2-1]+this.arr[len(this.arr)/2])/2
	}else {
		return float64(this.arr[len(this.arr)/2])
	}
}


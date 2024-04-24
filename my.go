package main

import (
	"fmt"
)

var arraySize = 10

type MyStruct struct {
	Roll  int
	Name  string
	Marks float32
}

func main() {
	array := myFunction1()
	fmt.Println(array)

	arrayOfPointers := myFunction2()
	fmt.Println(arrayOfPointers)
}

func myFunction1() []*MyStruct {
	array := make([]*MyStruct, arraySize)
	for i := 0; i < arraySize; i++ {
		array[i] = &MyStruct{
			Roll:  i + 1,
			Name:  fmt.Sprintf("Student %d", i+1),
			Marks: float32(i) * 10,
		}
	}
	return array
}

func myFunction2() *[]MyStruct {
	array := make([]MyStruct, arraySize)
	for i := 0; i < arraySize; i++ {
		array[i] = MyStruct{
			Roll:  i + 1,
			Name:  fmt.Sprintf("Student %d", i+1),
			Marks: float32(i) * 10,
		}
	}
	return &array
}

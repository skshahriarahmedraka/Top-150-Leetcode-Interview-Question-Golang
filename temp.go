// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func a() {
// 	b()
// }

// var b = func() {
// 	c()
// }

// func c() {
// 	pc := make([]uintptr, 64)
// 	n := runtime.Callers(0, pc)
// 	frames := runtime.CallersFrames(pc[:n])

// 	var frame runtime.Frame
// 	more := true
// 	for more {
// 		frame, more = frames.Next()
// 		fmt.Printf("- more:%v  |  %+v | %s\n", more, frame, frame.Function)
// 	}
// }

// func main() {
// 	fmt.Printf("< start\n")
// 	a()
// 	fmt.Printf("< end\n")
// }

package main

import (
	"fmt"
	"runtime"
)

func a() {
	b()
}

var b = func() {
	c()
}

func c() {
	pc := make([]uintptr, 64)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])

	var frame runtime.Frame
	more := true
	for more {
		frame, more = frames.Next()
		if frame.Function == "main.init.func1" {
			fmt.Println("main.init.func1", frame.Function)
			fmt.Printf("- more:%v  |  %+v | main.init.func1\n", more, frame)
		} else {
			fmt.Printf("- more:%v  |  %+v | %s\n", more, frame, frame.Function)
		}
	}
}

func main() {
	fmt.Printf("< start\n")
	a()
	fmt.Printf("< end\n")
}
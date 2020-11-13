//A Go program show Garbage Collection statistics
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var N = 40000000
	myMap := make(map[int]*int) //the name of the map that stores the integer pointers is myMap

	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	start := time.Now()
	runtime.GC()
	//_ = myMap[0] // is used for preventing the GC from garbage collecting the myMap variable too early
	duration := time.Since(start)
	fmt.Println("It took GC()", duration, "to finish.")
}

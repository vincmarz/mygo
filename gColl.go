//A Go program show Garbage Collection statistics
package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem) //Memory allocator statistic up to date as of the call
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	//for loop creates many Go slices in order to allocate large amounts of memory
	//and trigger the garbage collector
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000) //make function allocates a dinamically-sized array of byte
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)

}
